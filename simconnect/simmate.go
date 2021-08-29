package simconnect

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

const (
	simVarRequestTimeout int64 = 10000
)

type OnOpenFunc func(applName, applVersion, applBuild, simConnectVersion, simConnectBuild string)
type OnQuitFunc func()
type OnSimObjectDataFunc func(data *RecvSimObjectData)
type OnSimObjectDataByTypeFunc func(data *RecvSimObjectDataByType)
type OnDataReadyFunc func()
type OnEventIDFunc func(eventID DWord)
type OnExceptionFunc func(exceptionCode DWord)

type EventListener struct {
	OnOpen                OnOpenFunc
	OnQuit                OnQuitFunc
	OnSimObjectData       OnSimObjectDataFunc
	OnSimObjectDataByType OnSimObjectDataByTypeFunc
	OnDataReady           OnDataReadyFunc
	OnEventID             OnEventIDFunc
	OnException           OnExceptionFunc
}

type SimMate struct {
	SimConnect
	simVarManager *SimVarManager
	mutex         sync.Mutex
	dirty         bool
}

func NewSimMate() *SimMate {
	if !initialized {
		// auto-initialize with default search paths
		Initialize("")
	}
	mate := &SimMate{
		simVarManager: NewSimVarManager(),
	}
	return mate
}

func (mate *SimMate) AddSimVar(name, unit string, dataType DWord) DWord {
	defineID := mate.simVarManager.Add(name, unit, dataType)
	mate.dirty = true
	return defineID
}

func (mate *SimMate) RemoveSimVar(defineID DWord) bool {
	if ok := mate.simVarManager.Remove(defineID); !ok {
		return false
	}
	return true
}

func (mate *SimMate) SimVarValueAndDataType(defineID DWord) (interface{}, DWord, bool) {
	mate.mutex.Lock()
	defer mate.mutex.Unlock()
	simVar, ok := mate.simVarManager.GetSimVar(defineID)
	if !ok {
		return nil, DataTypeInvalid, false
	}
	return simVar.Value, simVar.DataType, true
}

func (mate *SimMate) SimVar(defineID DWord) (SimVar, bool) {
	mate.mutex.Lock()
	defer mate.mutex.Unlock()
	simVar, ok := mate.simVarManager.GetSimVar(defineID)
	if !ok {
		return SimVar{}, false
	}
	return *simVar, true
}

func (mate *SimMate) SimVarDump(indent string) []string {
	mate.mutex.Lock()
	defer mate.mutex.Unlock()
	return mate.simVarManager.SimVarDump(indent)
}

func (mate *SimMate) SetSimObjectData(name, unit string, value interface{}, dataType DWord) error {
	defineID := NewDefineID()
	if err := mate.AddToDataDefinition(defineID, name, unit, DataTypeFloat64); err != nil {
		return err
	}
	switch dataType {
	case DataTypeFloat64:
		buffer := [1]float64{
			ValueToFloat64(value),
		}
		size := DWord(unsafe.Sizeof(buffer))
		mate.SetDataOnSimObject(defineID, ObjectIDUser, 0, 0, size, unsafe.Pointer(&buffer[0]))

	default:
		panic(fmt.Errorf("datatype not implemented"))
	}
	return nil
}

func (mate *SimMate) HandleEvents(requestDataInterval time.Duration, receiveDataInterval time.Duration, stop chan interface{}, listener *EventListener) {
	reqDataTicker := time.NewTicker(requestDataInterval)
	defer reqDataTicker.Stop()

	recvDataTicker := time.NewTicker(receiveDataInterval)
	defer recvDataTicker.Stop()

	requestCount := 0
	updateCount := 0

	for {
		select {
		case <-stop:
			return

		case <-reqDataTicker.C:
			if updateCount > 0 {
				if listener != nil && listener.OnDataReady != nil {
					listener.OnDataReady()
				}
			}
			mate.requestSimObjectData()
			requestCount++

		case <-recvDataTicker.C:
			ppData, r1, _ := mate.GetNextDispatch()
			if r1 < 0 {
				if uint32(r1) != EFail {
					// fmt.Printf("GetNextDispatch error: %d %s\n", r1, err)
					return
				}
				if ppData == nil {
					break
				}
			}

			recv := *(*Recv)(ppData)
			switch recv.ID {
			case RecvIDException:
				recvException := *(*RecvException)(ppData)
				if listener != nil && listener.OnException != nil {
					listener.OnException(recvException.Exception)
				}

			case RecvIDOpen:
				recvOpen := *(*RecvOpen)(ppData)
				applName := strings.Trim(string(recvOpen.ApplicationName[:256]), "\x00")
				applVersion := fmt.Sprintf("%d.%d", recvOpen.ApplicationVersionMajor, recvOpen.ApplicationVersionMinor)
				applBuild := fmt.Sprintf("%d.%d", recvOpen.ApplicationBuildMajor, recvOpen.ApplicationBuildMinor)
				simConnectVersion := fmt.Sprintf("%d.%d", recvOpen.SimConnectVersionMajor, recvOpen.SimConnectVersionMinor)
				simConnectBuild := fmt.Sprintf("%d.%d", recvOpen.SimConnectBuildMajor, recvOpen.SimConnectBuildMinor)
				if listener != nil && listener.OnOpen != nil {
					listener.OnOpen(applName, applVersion, applBuild, simConnectVersion, simConnectBuild)
				}

			case RecvIDQuit:
				if listener != nil && listener.OnQuit != nil {
					listener.OnQuit()
				}

			case RecvIDEvent:
				recvEvent := *(*RecvEvent)(ppData)
				if listener != nil && listener.OnEventID != nil {
					listener.OnEventID(recvEvent.EventID)
				}

			// case RecvIDEventObjectAddRemove:
			// case RecvIDEventFilename:
			// case RecvIDEventFrame:

			case RecvIDSimobjectData:
				recvData := *(*RecvSimObjectData)(ppData)
				if listener != nil && listener.OnSimObjectData != nil {
					listener.OnSimObjectData(&recvData)
				}

			case RecvIDSimObjectDataByType:
				recvData := *(*RecvSimObjectDataByType)(ppData)
				simVar, exists := mate.simVarManager.GetSimVar(recvData.DefineID)
				if !exists {
					continue
				}

				var value interface{}
				switch simVar.DataType {
				case DataTypeInt32:
					value = (*SimObjectData_int32)(ppData).Value

				case DataTypeInt64:
					value = (*SimObjectData_int64)(ppData).Value

				case DataTypeFloat32:
					value = (*SimObjectData_float32)(ppData).Value

				case DataTypeFloat64:
					value = (*SimObjectData_float64)(ppData).Value

				case DataTypeString8:
					value = (*SimObjectData_string8)(ppData).Value

				case DataTypeString32:
					value = (*SimObjectData_string32)(ppData).Value

				case DataTypeString64:
					value = (*SimObjectData_string64)(ppData).Value

				case DataTypeString128:
					value = (*SimObjectData_string128)(ppData).Value

				case DataTypeString256:
					value = (*SimObjectData_string256)(ppData).Value

				case DataTypeString260:
					value = (*SimObjectData_string260)(ppData).Value

				case DataTypeStringV:
					value = (*SimObjectData_stringv)(ppData).Value

				// case DataTypeInitPosition:
				// case DataTypeStringMarkerState:
				// case DataTypeWaypoint:
				// case DataTypeStringLatLonAlt:
				// case DataTypeStringXYZ:
				default:
				}
				if value != nil {
					mate.updateSimObjectData(recvData.RequestID, recvData.DefineID, value)
					updateCount++
				}
				if listener != nil && listener.OnSimObjectDataByType != nil {
					listener.OnSimObjectDataByType(&recvData)
				}

			// case RecvIDWeatherObservation:
			// case RecvIDCloudState:
			// case RecvIDAssignedObjectID:
			// case RecvIDReservedKey:
			// case RecvIDCustomAction:
			// case RecvIDSystemState:
			// case RecvIDClientData:
			// case RecvIDEventWeatherMode:
			// case RecvIDAirportList:
			// case RecvIDVORList:
			// case RecvIDNDBList:
			// case RecvIDWaypointList:
			// case RecvIDEventMultiplayerServerStarted:
			// case RecvIDEventMultiplayerClientStarted:
			// case RecvIDEventMultiplayerSessionEnded:
			// case RecvIDEventRaceEnd:
			// case RecvIDEventRaceLap:
			// case RecvIDPick:

			default:
				log.Debug("Unknown recvInfo ID ", recv.ID)
			}
		}
	}
}

func (mate *SimMate) registerSimVars() (int, error) {
	count := 0
	for _, simVar := range mate.simVarManager.Vars {
		if !simVar.Registered {
			err := mate.AddToDataDefinition(simVar.DefineID, simVar.Name, simVar.Unit, simVar.DataType)
			if err != nil {
				return count, err
			} else {
				simVar.Registered = true
				count++
			}
		}
	}
	return count, nil
}

func (mate *SimMate) requestSimObjectData() (bool, error) {
	mate.mutex.Lock()
	defer mate.mutex.Unlock()

	if mate.dirty {
		count, err := mate.registerSimVars()
		if err != nil {
			return false, err
		}
		if count > 0 {
			fmt.Printf("Registered %d simvars\n", count)
		}
		mate.dirty = false
	}

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	const radiusMeters = 0
	simObjectType := SimObjectTypeUser
	for _, simVar := range mate.simVarManager.Vars {
		if !simVar.Pending {
			simVar.RequestID = NewRequestID()
		} else {
			if timestamp-simVar.Timestamp < simVarRequestTimeout {
				continue
			}
		}
		mate.RequestDataOnSimObjectType(simVar.RequestID, simVar.DefineID, radiusMeters, simObjectType)
		simVar.Timestamp = timestamp
		simVar.Pending = true
	}
	return true, nil
}

func (mate *SimMate) updateSimObjectData(requestID, defineID DWord, value interface{}) {
	if simVar, updated := mate.simVarManager.Update(requestID, defineID, value); updated {
		simVar.Pending = false
	}
}

// Generics. Needed. Badly. Ugh.
type SimObjectData struct {
	RecvSimObjectDataByType
}

type SimObjectData_int32 struct {
	SimObjectData
	Value int32
}

type SimObjectData_int64 struct {
	SimObjectData
	Value int64
}

type SimObjectData_float32 struct {
	SimObjectData
	Value float32
}

type SimObjectData_float64 struct {
	SimObjectData
	Value float64
}

type SimObjectData_string8 struct {
	SimObjectData
	Value [8]byte
}

type SimObjectData_string32 struct {
	SimObjectData
	Value [32]byte
}

type SimObjectData_string64 struct {
	SimObjectData
	Value [64]byte
}

type SimObjectData_string128 struct {
	SimObjectData
	Value [128]byte
}

type SimObjectData_string256 struct {
	SimObjectData
	Value [256]byte
}

type SimObjectData_string260 struct {
	SimObjectData
	Value [260]byte
}

type SimObjectData_stringv struct {
	SimObjectData
	Value string // TODO: Not sure if this is correct
}
