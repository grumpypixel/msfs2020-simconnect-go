package simconnect

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

const (
	SimConnectDLL = "SimConnect.dll"
)

var (
	library     *syscall.LazyDLL
	procs       map[string]*syscall.LazyProc
	lockID      sync.Mutex
	defineID    DWord
	eventID     DWord
	requestID   DWord
	initialized bool
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

type SimConnect struct {
	handle    unsafe.Pointer
	connected bool
}

func NewSimConnect() *SimConnect {
	if !initialized {
		panic("SimConnect not initialized.")
	}
	return &SimConnect{}
}

func LocateLibrary(additionalSearchPath string) error {
	paths, err := buildSearchPaths(additionalSearchPath)
	if err != nil {
		return err
	}
	_, err = findLibrary(paths)
	if err != nil {
		return err
	}
	return nil
}

func Initialize(additionalSearchPath string) error {
	if initialized {
		return nil
	}

	paths, err := buildSearchPaths(additionalSearchPath)
	if err != nil {
		return err
	}

	libPath, err := findLibrary(paths)
	if err != nil {
		return err
	}

	if err := loadLibrary(libPath); err != nil {
		return err
	}

	loadProcs()
	initialized = true
	return nil
}

func IsInitialized() bool {
	return initialized
}

func (simco *SimConnect) IsConnected() bool {
	return simco.connected
}

func NewDefineID() DWord {
	lockID.Lock()
	defer lockID.Unlock()
	defineID++
	return defineID
}

func NewRequestID() DWord {
	lockID.Lock()
	defer lockID.Unlock()
	requestID++
	return requestID
}

func NewEventID() DWord {
	lockID.Lock()
	defer lockID.Unlock()
	eventID++
	return eventID
}

func buildSearchPaths(searchPath string) ([]string, error) {
	paths := []string{}
	if len(searchPath) > 0 {
		paths = append(paths, searchPath)
	}

	execPath, err := os.Executable()
	if err != nil {
		return paths, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return paths, err
	}
	paths = append(paths, filepath.Dir(execPath), cwd)
	return append(paths, cwd), nil
}

func findLibrary(searchPaths []string) (string, error) {
	for _, path := range searchPaths {
		libPath := filepath.Join(path, SimConnectDLL)
		if _, err := os.Stat(libPath); os.IsNotExist(err) {
			continue
		}
		return libPath, nil
	}
	return "", fmt.Errorf("could not locate %s in search paths", SimConnectDLL)
}

func loadLibrary(path string) error {
	library = syscall.NewLazyDLL(path)
	if err := library.Load(); err != nil {
		return err
	}
	return nil
}

func loadProcs() {
	procs = make(map[string]*syscall.LazyProc)
	procNames := []string{
		scOpen,
		scClose,
		// scCallDispatch,
		scGetNextDispatch,
		scRequestSystemState,
		scMapClientEventToSimEvent,
		scSubscribeToSystemEvent,
		scSetSystemEventState,
		scUnsubscribeFromSystemEvent,
		scSetNotificationGroupPriority,
		scText,
		scRequestDataOnSimObject,
		scRequestDataOnSimObjectType,
		scAddClientEventToNotificationGroup,
		scRemoveClientEvent,
		scTransmitClientEvent,
		scMapClientDataNameToID,
		scRequestClientData,
		scCreateClientData,
		scAddToClientDataDefinition,
		scAddToDataDefinition,
		scSetClientData,
		scSetDataOnSimObject,
		scClearClientDataDefinition,
		scClearDataDefinition,
		scMapInputEventToClientEvent,
		scRequestNotificationGroup,
		scClearInputGroup,
		scClearNotificationGroup,
		// scRequestReservedKey,
		scSetInputGroupPriority,
		scSetInputGroupState,
		scRemoveInputEvent,
		scAICreateEnrouteATCAircraft,
		scAICreateNonATCAircraft,
		scAICreateParkedATCAircraft,
		scAICreateSimulatedObject,
		scAIReleaseControl,
		scAIRemoveObject,
		scAISetAircraftFlightPlan,
		scFlightLoad,
		scFlightSave,
		scFlightPlanLoad,
		scGetLastSentPacketID,
		// scRequestResponseTimes,
		// scInsertString,
		// scRetrieveString,
		scRequestFacilitiesList,
		scSubscribeToFacilities,
		scUnsubscribeToFacilities,
		// scCompleteCustomMissionAction,
		// scExecuteMissionAction,
		scMenuAddItem,
		scMenuAddSubItem,
		scMenuDeleteItem,
		scMenuDeleteSubItem,
		scCameraSetRelative6DOF,
		scSetSystemState,
	}
	for _, procName := range procNames {
		procs[procName] = library.NewProc(procName)
	}
}

func callProc(procName string, args ...uintptr) error {
	proc, ok := procs[procName]
	if !ok {
		return fmt.Errorf("proc %s not defined", procName)
	}
	r1, _, err := proc.Call(args...)
	if int32(r1) < 0 {
		return fmt.Errorf("%s error: %d %s", procName, r1, err)
	}
	return nil
}

func toNullTerminatedBytes(str string) []byte {
	return []byte(str + "\x00")
}

func toCharPtr(str string) uintptr {
	bytes := toNullTerminatedBytes(str)
	return uintptr(unsafe.Pointer(&bytes[0]))
}

func toBoolPtr(value bool) uintptr {
	v := 0
	if value {
		v = 1
	}
	return uintptr(v)
}
