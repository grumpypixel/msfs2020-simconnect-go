package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grumpypixel/msfs2020-simconnect-go/simconnect"
)

type SimVar struct {
	DefineID   simconnect.DWord
	Name, Unit string
}

type SimObjectValue struct {
	simconnect.RecvSimObjectDataByType
	Value float64
}

var (
	requestDataInterval = time.Millisecond * 250
	receiveDataInterval = time.Millisecond * 1
	simConnect          *simconnect.SimConnect
	simVars             []*SimVar
)

func main() {
	additionalSearchPath := ""
	args := os.Args
	if len(args) > 1 {
		additionalSearchPath = args[1]
		fmt.Println("searchpath", additionalSearchPath)
	}

	if err := simconnect.Initialize(additionalSearchPath); err != nil {
		panic(err)
	}

	simConnect = simconnect.NewSimConnect()
	if err := simConnect.Open("Transpotato"); err != nil {
		panic(err)
	}

	simVars = make([]*SimVar, 0)
	nameUnitMapping := map[string]string{"AIRSPEED INDICATED": "knot", "INDICATED ALTITUDE": "feet", "PLANE LATITUDE": "degrees", "PLANE LONGITUDE": "degrees"}
	for name, unit := range nameUnitMapping {
		defineID := simconnect.NewDefineID()
		simConnect.AddToDataDefinition(defineID, name, unit, simconnect.DataTypeFloat64)
		simVars = append(simVars, &SimVar{defineID, name, unit})
	}

	done := make(chan bool, 1)
	defer close(done)
	go HandleTerminationSignal(done)
	go HandleEvents(done)

	<-done

	if err := simConnect.Close(); err != nil {
		panic(err)
	}
}

func HandleTerminationSignal(done chan bool) {
	sigterm := make(chan os.Signal, 1)
	defer close(sigterm)

	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-sigterm:
			done <- true
			return
		}
	}
}

func HandleEvents(done chan bool) {
	reqDataTicker := time.NewTicker(requestDataInterval)
	defer reqDataTicker.Stop()

	recvDataTicker := time.NewTicker(receiveDataInterval)
	defer recvDataTicker.Stop()

	var simObjectType = simconnect.SimObjectTypeUser
	var radius = simconnect.DWordZero

	for {
		select {
		case <-reqDataTicker.C:
			fmt.Println("\nRequesting data...")
			for _, simVar := range simVars {
				simConnect.RequestDataOnSimObjectType(simconnect.NewRequestID(), simVar.DefineID, radius, simObjectType)
			}

		case <-recvDataTicker.C:
			ppData, r1, err := simConnect.GetNextDispatch()
			if r1 < 0 {
				if uint32(r1) != simconnect.EFail {
					fmt.Printf("GetNextDispatch error: %d %s\n", r1, err)
					return
				}
				if ppData == nil {
					break
				}
			}

			recv := *(*simconnect.Recv)(ppData)
			switch recv.ID {
			case simconnect.RecvIDOpen:
				fmt.Println("Connected.")

			case simconnect.RecvIDQuit:
				fmt.Println("Disconnected.")
				done <- true

			case simconnect.RecvIDException:
				recvException := *(*simconnect.RecvException)(ppData)
				fmt.Println("Something exceptional happened.", recvException.Exception)

			case simconnect.RecvIDSimObjectDataByType:
				data := *(*SimObjectValue)(ppData)
				for _, simVar := range simVars {
					if simVar.DefineID == data.DefineID {
						fmt.Printf("[%d] %s %s %f\n", data.RequestID, simVar.Name, simVar.Unit, data.Value)
						break
					}
				}
			}
		}
	}
}
