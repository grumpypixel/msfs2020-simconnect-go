package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grumpypixel/msfs2020-simconnect-go/simconnect"
)

type Request struct {
	Name, Unit string
	DataType   simconnect.DWord
}

type Var struct {
	DefineID simconnect.DWord
	Name     string
}

type App struct {
	simconnect.EventListener
	mate    *simconnect.SimMate
	vars    []*Var
	done    chan bool
	counter uint32
}

var (
	requestDataInterval = time.Millisecond * 250
	receiveDataInterval = time.Millisecond * 1
	mate                *simconnect.SimMate
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

	app := &App{}
	app.run()
}

func (app *App) run() {
	app.done = make(chan bool, 1)
	defer close(app.done)

	app.mate = simconnect.NewSimMate()

	if err := app.mate.Open("Transpotato"); err != nil {
		panic(err)
	}

	// These are the sim vars we are looking for
	requests := []Request{
		{"AIRSPEED INDICATED", "knot", simconnect.DataTypeFloat64},
		{"PLANE LATITUDE", "degrees", simconnect.DataTypeFloat64},
		{"PLANE LONGITUDE", "degrees", simconnect.DataTypeFloat64},
		{"PLANE HEADING DEGREES MAGNETIC", "degrees", simconnect.DataTypeFloat64},
		{"TITLE", "", simconnect.DataTypeString256},
		{"ATC ID", "", simconnect.DataTypeString64},
	}
	app.vars = make([]*Var, 0)
	for _, request := range requests {
		defineID := app.mate.AddSimVar(request.Name, request.Unit, request.DataType)
		app.vars = append(app.vars, &Var{defineID, request.Name})
	}

	go app.handleTerminationSignal()
	go app.mate.HandleEvents(requestDataInterval, receiveDataInterval, app)

	<-app.done

	app.mate.Close()
}

func (app *App) handleTerminationSignal() {
	sigterm := make(chan os.Signal, 1)
	defer close(sigterm)

	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sigterm:
			app.done <- true
			return
		}
	}
}

func (app *App) OnOpen(applName, applVersion, applBuild, simConnectVersion, simConnectBuild string) {
	fmt.Println("\nConnected.")
	flightSimVersion := fmt.Sprintf(
		"Flight Simulator:\n Name: %s\n Version: %s (build %s)\n SimConnect: %s (build %s)",
		applName, applVersion, applBuild, simConnectVersion, simConnectBuild)
	fmt.Printf("\n%s\n\n", flightSimVersion)
	fmt.Printf("CLEAR PROP!\n\n")
}

func (app *App) OnQuit() {
	fmt.Println("Disconnected.")
	app.done <- true
}

func (app *App) OnEventID(eventID simconnect.DWord) {
	fmt.Println("Received event ID", eventID)
}

func (app *App) OnException(exceptionCode simconnect.DWord) {
	fmt.Printf("Exception (code: %d)\n", exceptionCode)
}

func (app *App) OnDataUpdate(defineID simconnect.DWord) {
	// ignore
}

func (app *App) OnDataReady() {
	fmt.Printf("\nUpdate %d...\n", app.counter)
	app.counter++
	for _, v := range app.vars {
		value, _, ok := app.mate.SimVarValueAndDataType(v.DefineID)
		if !ok || value == nil {
			continue
		}
		fmt.Printf("%s = %v\n", v.Name, value)
	}
}
