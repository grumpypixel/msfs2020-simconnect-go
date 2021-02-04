package simconnect

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"unsafe"
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

type SimConnect struct {
	handle    unsafe.Pointer
	close     chan bool
	connected bool
}

func NewSimConnect() *SimConnect {
	if !initialized {
		panic(fmt.Errorf("simconnect not initialized. Initialize() is your friend."))
	}
	simco := &SimConnect{}
	return simco
}

func LocateLibrary(additionalSearchPath string) bool {
	paths := buildSearchPaths(additionalSearchPath)
	_, err := findLibrary(paths)
	if err != nil {
		return false
	}
	return true
}

func Initialize(additionalSearchPath string) error {
	if initialized {
		return nil
	}

	paths := buildSearchPaths(additionalSearchPath)
	libPath, err := findLibrary(paths)
	if err != nil {
		return err
	}

	fmt.Println("Loading library:", libPath)
	if err := loadLibrary(libPath); err != nil {
		return err
	}

	loadProcedures()
	initialized = true
	return nil
}

func IsInitialized() bool {
	return initialized
}

func (simco *SimConnect) IsConnected() bool {
	return simco.connected
}

func (simco *SimConnect) AddToDataDefinition(defineID DWord, datumName string, unitName string, datumType DWord) error {
	// SimConnect_AddToDataDefinition(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_DATA_DEFINITION_ID DefineID,
	// 	const char * DatumName,
	// 	const char * UnitsName,
	// 	SIMCONNECT_DATATYPE DatumType = SIMCONNECT_DATATYPE_FLOAT64,
	// 	float fEpsilon = 0,
	// 	DWORD DatumID = SIMCONNECT_UNUSED)

	name := []byte(datumName + "\x00")
	unit := []byte(unitName + "\x00")

	var unitArg uintptr = 0
	if len(unit) > 0 {
		unitArg = uintptr(unsafe.Pointer(&unit[0]))
	}

	const epsilon float32 = 0
	args := []uintptr{
		uintptr(simco.handle),
		uintptr(defineID),
		uintptr(unsafe.Pointer(&name[0])),
		unitArg,
		uintptr(datumType),
		uintptr(epsilon),
		uintptr(Unused),
	}
	procName := simConnectAddToDataDefinitionProcName
	r1, _, err := procs[procName].Call(args...)
	if int32(r1) < 0 {
		return fmt.Errorf("%s for data %s error: %d %s", procName, name, r1, err)
	}
	return nil
}

func (simco *SimConnect) ClearDataDefinition(defineID DWord) error {
	// SIMCONNECTAPI SimConnect_ClearDataDefinition(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_DATA_DEFINITION_ID DefineID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(defineID),
	}
	procName := simConnectClearDataDefinitionProcName
	r1, _, err := procs[procName].Call(args...)
	if int32(r1) < 0 {
		return fmt.Errorf("%s error: %d %s", procName, r1, err)
	}
	return nil
}

func (simco *SimConnect) Close() error {
	// SimConnect_Close(HANDLE hSimConnect)

	procName := simConnectCloseProcName
	r1, _, err := procs[procName].Call(uintptr(simco.handle))
	if int32(r1) < 0 {
		return fmt.Errorf("%s error: %d %s", procName, r1, err)
	}
	simco.connected = false
	return nil
}

func (simco *SimConnect) GetNextDispatch() (unsafe.Pointer, int32, error) {
	// SimConnect_GetNextDispatch(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_RECV ** ppData,
	// 	DWORD * pcbData)

	var ppData unsafe.Pointer
	var ppDataLength DWord
	procName := simConnectGetNextDispatchProcName
	r1, _, err := procs[procName].Call(
		uintptr(simco.handle),
		uintptr(unsafe.Pointer(&ppData)),
		uintptr(unsafe.Pointer(&ppDataLength)),
	)
	return ppData, int32(r1), err
}

func (simco *SimConnect) Open(name string) error {
	// SimConnect_Open(
	// 	HANDLE * phSimConnect,
	// 	LPCSTR szName,
	// 	HWND hWnd,
	// 	DWORD UserEventWin32,
	// 	HANDLE hEventHandle,
	// 	DWORD ConfigIndex)

	fsxCompatible := 0
	args := []uintptr{
		uintptr(unsafe.Pointer(&simco.handle)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		0,
		uintptr(WmUserSimConnect),
		0,
		uintptr(fsxCompatible),
	}
	procName := simConnectOpenProcName
	r1, _, err := procs[procName].Call(args...)
	if int32(r1) < 0 {
		return fmt.Errorf("%s error: %d %s", procName, r1, err)
	}
	simco.connected = true
	return nil
}

func (simco *SimConnect) SetDataOnSimObject(defineID, objectID, flags, arrayCount, size DWord, buf unsafe.Pointer) error {
	// SimConnect_SetDataOnSimObject(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_DATA_DEFINITION_ID DefineID,
	// 	SIMCONNECT_OBJECT_ID ObjectID,
	// 	SIMCONNECT_DATA_SET_FLAG Flags,
	// 	DWORD ArrayCount,
	// 	DWORD cbUnitSize,
	// 	void * pDataSet)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(defineID),
		uintptr(objectID),
		uintptr(flags),
		uintptr(arrayCount),
		uintptr(size),
		uintptr(buf),
	}
	procName := simConnectSetDataOnSimObjectProcName
	r1, _, err := procs[procName].Call(args...)
	if int32(r1) < 0 {
		return fmt.Errorf("%s error: %d %s", procName, r1, err)
	}
	return nil
}

func (simco *SimConnect) RequestDataOnSimObjectType(requestID, defineID, radius, simobjectType DWord) error {
	// SimConnect_RequestDataOnSimObjectType(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_DATA_REQUEST_ID RequestID,
	// 	SIMCONNECT_DATA_DEFINITION_ID DefineID,
	// 	DWORD dwRadiusMeters,
	// 	SIMCONNECT_SIMOBJECT_TYPE type)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(requestID),
		uintptr(defineID),
		uintptr(radius),
		uintptr(simobjectType),
	}
	procName := simConnectRequestDataOnSimObjectTypeProcName
	r1, _, err := procs[procName].Call(args...)
	if int32(r1) < 0 {
		return fmt.Errorf("%s error: %d %s", procName, r1, err)
	}
	return nil
}

func (simco *SimConnect) CameraSetRelative6DOF(deltaX, deltaY, deltaZ, pitchDeg, bankDeg, headingDeg float64) error {
	// SimConnect_CameraSetRelative6DOF(
	// 	HANDLE hSimConnect,
	// 	float fDeltaX,
	// 	float fDeltaY,
	// 	float fDeltaZ,
	// 	float fPitchDeg,
	// 	float fBankDeg,
	// 	float fHeadingDeg)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(deltaX),
		uintptr(deltaY),
		uintptr(deltaZ),
		uintptr(pitchDeg),
		uintptr(bankDeg),
		uintptr(headingDeg),
	}
	procName := simConnectCameraSetRelative6DOFProcName
	r1, _, err := procs[procName].Call(args...)
	if int32(r1) < 0 {
		return fmt.Errorf("%s error: %d %s", procName, r1, err)
	}
	return nil
}

func (simco *SimConnect) Text(text string, textType DWord, duration float64, eventID DWord) error {
	// SimConnect_Text(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_TEXT_TYPE type,
	// 	float fTimeSeconds,
	// 	SIMCONNECT_CLIENT_EVENT_ID EventID,
	// 	DWORD cbUnitSize,
	// 	void * pDataSet)

	txt := []byte(text + "\x00")
	args := []uintptr{
		uintptr(simco.handle),
		uintptr(textType),
		uintptr(duration),
		uintptr(eventID),
		uintptr(DWord(len(text))),
		uintptr(unsafe.Pointer(&txt[0])),
	}
	procName := simConnectTextProcName
	r1, _, err := procs[procName].Call(args...)
	if int32(r1) < 0 {
		return fmt.Errorf("%s error: %d %s", procName, r1, err)
	}
	return nil
}

func (simco *SimConnect) SubscribeToFacilities() error {
	// SimConnect_SubscribeToFacilities(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_FACILITY_LIST_TYPE type,
	// 	SIMCONNECT_DATA_REQUEST_ID RequestID)

	panic(fmt.Errorf("Not implemented."))
	// return nil
}

func (simco *SimConnect) UnsubscribeToFacilities() error {
	// SimConnect_UnsubscribeToFacilities(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_FACILITY_LIST_TYPE type)

	panic(fmt.Errorf("Not implemented."))
	// return nil
}

func (simco *SimConnect) RequestFacilitiesList() error {
	// SimConnect_RequestFacilitiesList(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_FACILITY_LIST_TYPE type,
	// 	SIMCONNECT_DATA_REQUEST_ID RequestID)

	panic(fmt.Errorf("Not implemented."))
	// return nil
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

func buildSearchPaths(searchPath string) []string {
	paths := []string{}
	if len(searchPath) > 0 {
		paths = append(paths, searchPath)
	}

	execPath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	paths = append(paths, filepath.Dir(execPath))

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return append(paths, cwd)
}

func findLibrary(searchPaths []string) (string, error) {
	for _, path := range searchPaths {
		libPath := filepath.Join(path, simConnectLibraryName)
		if _, err := os.Stat(libPath); os.IsNotExist(err) {
			continue
		}
		return libPath, nil
	}
	return "", fmt.Errorf("Could not locate %s in search paths", simConnectLibraryName)
}

func loadLibrary(path string) error {
	library = syscall.NewLazyDLL(path)
	if err := library.Load(); err != nil {
		return err
	}
	return nil
}

func loadProcedures() {
	procs = make(map[string]*syscall.LazyProc)
	procNames := []string{
		simConnectAddToDataDefinitionProcName,
		simConnectClearDataDefinitionProcName,
		simConnectRequestDataOnSimObjectTypeProcName,
		simConnectSetDataOnSimObjectProcName,
		simConnectCloseProcName,
		simConnectOpenProcName,
		simConnectGetNextDispatchProcName,
		simConnectCameraSetRelative6DOFProcName,
		simConnectTextProcName,
		simConnectSubscribeToFacilitiesProcName,
		simConnectUnsubscribeToFacilitiesProcName,
		simConnectRequestFacilitiesListProcName,
	}
	for _, procName := range procNames {
		procs[procName] = library.NewProc(procName)
	}
}
