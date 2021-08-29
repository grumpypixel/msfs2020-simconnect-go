package simconnect

import (
	"syscall"
	"unsafe"
)

// SimConnect_Open: Used to send a request to the Flight Simulator server to open up communications with a new client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_Open.htm
func (simco *SimConnect) Open(name string) error {
	// SimConnect_Open(
	// 	HANDLE * phSimConnect,
	// 	LPCSTR szName,
	// 	HWND hWnd,
	// 	DWORD UserEventWin32,
	// 	HANDLE hEventHandle,
	// 	DWORD ConfigIndex)

	const hwnd DWord = 0
	const userEventWin32 = WmUserSimConnect
	const eventHandle DWord = 0
	const configIndex DWord = 0 // TODO: make this a function parameter

	var namePtr *uint16
	namePtr, namePtrErr := syscall.UTF16PtrFromString(name)
	if namePtrErr != nil {
		return namePtrErr
	}

	args := []uintptr{
		uintptr(unsafe.Pointer(&simco.handle)),
		uintptr(unsafe.Pointer(namePtr)),
		uintptr(hwnd),
		uintptr(userEventWin32),
		uintptr(eventHandle),
		uintptr(configIndex),
	}
	err := callProc(scOpen, args...)
	if err == nil {
		simco.connected = true
	}
	return err
}

// SimConnect_Close: Used to request that the communication with the server is ended.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_Close.htm
func (simco *SimConnect) Close() error {
	// SimConnect_Close(
	//  HANDLE hSimConnect)

	args := []uintptr{
		uintptr(simco.handle),
	}
	err := callProc(scClose, args...)
	if err == nil {
		simco.connected = false
	}
	return err
}

// SimConnect_CallDispatch: Used to process the next SimConnect message received through the specified callback function.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_CallDispatch.htm
// TODO: SimConnect_CallDispatch(HANDLE hSimConnect, DispatchProc pfcnDispatch, void * pContext)

// SimConnect_GetNextDispatch: Used to process the next SimConnect message received, without the use of a callback function.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_GetNextDispatch.htm
func (simco *SimConnect) GetNextDispatch() (unsafe.Pointer, int32, error) {
	// SimConnect_GetNextDispatch(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_RECV ** ppData,
	// 	DWORD * pcbData)

	var ppData unsafe.Pointer
	var ppDataLength DWord
	r1, _, err := procs[scGetNextDispatch].Call(
		uintptr(simco.handle),
		uintptr(unsafe.Pointer(&ppData)),
		uintptr(unsafe.Pointer(&ppDataLength)),
	)
	return ppData, int32(r1), err
}

// SimConnect_RequestSystemState: Used to request information from a number of Flight Simulator system components.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_RequestSystemState.htm
func (simco *SimConnect) RequestSystemState(requestID DWord, state string) error {
	// SimConnect_RequestSystemState(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID,
	//  const char * szState)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(requestID),
		uintptr(toCharPtr(state)),
	}
	return callProc(scRequestSystemState, args...)
}

// SimConnect_MapClientEventToSimEvent: Used to associate a client defined event ID with a Flight Simulator event name.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_MapClientEventToSimEvent.htm
func (simco *SimConnect) MapClientEventToSimEvent(eventID DWord, eventName string) error {
	// SimConnect_MapClientEventToSimEvent(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_EVENT_ID EventID,
	//  const char * EventName = "")

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(eventID),
		toCharPtr(eventName),
	}
	return callProc(scMapClientEventToSimEvent, args...)
}

// SimConnect_SubscribeToSystemEvent: Used to request that a specific system event is notified to the client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_SubscribeToSystemEvent.htm
func (simco *SimConnect) SubscribeToSystemEvent(eventID DWord, systemEventName string) error {
	// SimConnect_SubscribeToSystemEvent(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_EVENT_ID EventID,
	//  const char * SystemEventName)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(eventID),
		toCharPtr(systemEventName),
	}
	return callProc(scSubscribeToSystemEvent, args...)
}

// SimConnect_SetSystemEventState: Used to turn requests for event information from the server on and off.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_SetSystemEventState.htm
func (simco *SimConnect) SetSystemEventState(eventID, state DWord) error {
	// SIMCONNECTAPI SimConnect_SetSystemEventState(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_EVENT_ID EventID,
	//  SIMCONNECT_STATE dwState)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(eventID),
		uintptr(state),
	}
	return callProc(scSetSystemEventState, args...)
}

// SimConnect_UnsubscribeFromSystemEvent: Used to request that notifications are no longer received for the specified system event.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_UnsubscribeFromSystemEvent.htm
func (simco *SimConnect) UnsubscribeFromSystemEvent(eventID DWord) error {
	// SimConnect_UnsubscribeFromSystemEvent(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_EVENT_ID EventID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(eventID),
	}
	return callProc(scUnsubscribeFromSystemEvent, args...)
}

// SimConnect_SetNotificationGroupPriority: Used to set the priority of a notification group.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_SetNotificationGroupPriority.htm
func (simco *SimConnect) SetNotificationGroupPriority(groupID, priority DWord) error {
	// SimConnect_SetNotificationGroupPriority(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_NOTIFICATION_GROUP_ID GroupID,
	//  DWORD uPriority)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
		uintptr(priority),
	}
	return callProc(scSetNotificationGroupPriority, args...)
}

// SimConnect_Text: Displays text to the user. (This function is not currently available for use.)
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/General/SimConnect_Text.htm
func (simco *SimConnect) Text(text string, textType DWord, timeSeconds float32, eventID DWord) error {
	// SimConnect_Text(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_TEXT_TYPE type,
	// 	float fTimeSeconds,
	// 	SIMCONNECT_CLIENT_EVENT_ID EventID,
	// 	DWORD cbUnitSize,
	// 	void * pDataSet)

	size := len(text)
	args := []uintptr{
		uintptr(simco.handle),
		uintptr(textType),
		uintptr(timeSeconds),
		uintptr(eventID),
		uintptr(DWord(size)),
		toCharPtr(text),
	}
	return callProc(scText, args...)
}

// Event And Data functions:

// SimConnect_RequestDataOnSimObject: Used to request when the SimConnect client is to receive data values for a specific object.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_RequestDataOnSimObject.htm
func (simco *SimConnect) RequestDataOnSimObject(requestID, defineID, objectID, period, flags DWord) error {
	// SimConnect_RequestDataOnSimObject(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID,
	//  SIMCONNECT_DATA_DEFINITION_ID DefineID,
	//  SIMCONNECT_OBJECT_ID ObjectID,
	//  SIMCONNECT_PERIOD Period,
	//  SIMCONNECT_DATA_REQUEST_FLAG Flags = 0,
	//  DWORD origin = 0,
	//  DWORD interval = 0,
	//  DWORD limit = 0)

	const origin DWord = 0
	const interval DWord = 0
	const limit DWord = 0

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(requestID),
		uintptr(defineID),
		uintptr(objectID),
		uintptr(period),
		uintptr(flags),
		uintptr(origin),
		uintptr(interval),
		uintptr(limit),
	}
	return callProc(scRequestDataOnSimObject, args...)
}

// SimConnect_RequestDataOnSimObjectType: Used to retrieve information about simulation objects of a given type that are within a specified radius of the user's aircraft.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_RequestDataOnSimObjectType.htm
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
	return callProc(scRequestDataOnSimObjectType, args...)
}

// SimConnect_AddClientEventToNotificationGroup: Used to add an individual client defined event to a notification group.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_AddClientEventToNotificationGroup.htm
func (simco *SimConnect) AddClientEventToNotificationGroup(groupID, eventID DWord, maskable bool) error {
	// SimConnect_AddClientEventToNotificationGroup(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_NOTIFICATION_GROUP_ID GroupID,
	//  SIMCONNECT_CLIENT_EVENT_ID EventID,
	//  BOOL bMaskable = FALSE)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
		uintptr(eventID),
		uintptr(toBoolPtr(maskable)),
	}
	return callProc(scAddClientEventToNotificationGroup, args...)
}

// SimConnect_RemoveClientEvent: Used to remove a client defined event from a notification group.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_RemoveClientEvent.htm
func (simco *SimConnect) RemoveClientEvent(groupID, eventID DWord) error {
	// SimConnect_RemoveClientEvent(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_NOTIFICATION_GROUP_ID GroupID,
	//  SIMCONNECT_CLIENT_EVENT_ID EventID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
		uintptr(eventID),
	}
	return callProc(scRemoveClientEvent, args...)
}

// SimConnect_TransmitClientEvent: Used to request that the Flight Simulator server transmit to all SimConnect clients the specified client event.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_TransmitClientEvent.htm
func (simco *SimConnect) TransmitClientEvent(objectID uint32, eventID uint32, data DWord, groupID DWord, flags DWord) error {
	// SimConnect_TransmitClientEvent(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_OBJECT_ID ObjectID,
	//  SIMCONNECT_CLIENT_EVENT_ID EventID,
	//  DWORD dwData,
	//  SIMCONNECT_NOTIFICATION_GROUP_ID GroupID,
	//  SIMCONNECT_EVENT_FLAG Flags)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(objectID),
		uintptr(eventID),
		uintptr(data),
		uintptr(groupID),
		uintptr(flags),
	}
	return callProc(scTransmitClientEvent, args...)
}

// SimConnect_MapClientDataNameToID: Used to associate an ID with a named client date area.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_MapClientDataNameToID.htm
func (simco *SimConnect) MapClientDataNameToID(clientDataName string, clientDataID DWord) error {
	// SimConnect_MapClientDataNameToID(
	//  HANDLE hSimConnect,
	//  const char * szClientDataName,
	//  SIMCONNECT_CLIENT_DATA_ID ClientDataID)

	args := []uintptr{
		uintptr(simco.handle),
		toCharPtr(clientDataName),
		uintptr(clientDataID),
	}
	return callProc(scMapClientDataNameToID, args...)
}

// SimConnect_RequestClientData: Used to request that the data in an area created by another client be sent to this client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_RequestClientData.htm
func (simco *SimConnect) RequestClientData(clientDataID, requestID, defineID, period, flags DWord) error {
	// SimConnect_RequestClientData(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_DATA_ID ClientDataID,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID,
	//  SIMCONNECT_CLIENT_DATA_DEFINITION_ID DefineID,
	//  SIMCONNECT_CLIENT_DATA_PERIOD Period = SIMCONNECT_CLIENT_DATA_PERIOD_ONCE,
	//  SIMCONNECT_CLIENT_DATA_REQUEST_FLAG Flags = 0,
	//  DWORD origin = 0,
	//  DWORD interval = 0,
	//  DWORD limit = 0)

	const origin DWord = 0
	const interval DWord = 0
	const limit DWord = 0

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(clientDataID),
		uintptr(requestID),
		uintptr(defineID),
		uintptr(period),
		uintptr(flags),
		uintptr(origin),
		uintptr(interval),
		uintptr(limit),
	}
	return callProc(scRequestClientData, args...)
}

// SimConnect_CreateClientData: Used to request the creation of a reserved data area for this client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_CreateClientData.htm
func (simco *SimConnect) CreateClientData(clientDataID, size, flags DWord) error {
	// SimConnect_CreateClientData(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_DATA_ID ClientDataID,
	//  DWORD dwSize,
	//  SIMCONNECT_CREATE_CLIENT_DATA_FLAG Flags)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(clientDataID),
		uintptr(size),
		uintptr(flags),
	}
	return callProc(scCreateClientData, args...)
}

// SimConnect_AddToClientDataDefinition: Used to add an offset and a size in bytes, or a type, to a client data definition.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_AddToClientDataDefinition.htm
func (simco *SimConnect) AddToClientDataDefinition(defineID, offset, sizeOrType DWord) error {
	// SimConnect_AddToClientDataDefinition(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_DATA_DEFINITION_ID DefineID,
	//  DWORD dwOffset,
	//  DWORD dwSizeOrType,
	//  float fEpsilon = 0,
	//  DWORD DatumID = SIMCONNECT_UNUSED)

	const epsilon float32 = 0
	const datumID = Unused

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(defineID),
		uintptr(offset),
		uintptr(sizeOrType),
		uintptr(epsilon),
		uintptr(datumID),
	}
	return callProc(scAddToClientDataDefinition, args...)
}

// SimConnect_AddToDataDefinition: Used to add a Flight Simulator simulation variable name to a client defined object definition.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_AddToDataDefinition.htm
func (simco *SimConnect) AddToDataDefinition(defineID DWord, datumName string, unitName string, datumType DWord) error {
	// SimConnect_AddToDataDefinition(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_DATA_DEFINITION_ID DefineID,
	// 	const char * DatumName,
	// 	const char * UnitsName,
	// 	SIMCONNECT_DATATYPE DatumType = SIMCONNECT_DATATYPE_FLOAT64,
	// 	float fEpsilon = 0,
	// 	DWORD DatumID = SIMCONNECT_UNUSED)

	var unitArg uintptr = 0
	if len(unitName) > 0 {
		unitArg = toCharPtr(unitName)
	}

	const epsilon float32 = 0
	const datumID = Unused

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(defineID),
		toCharPtr(datumName),
		unitArg,
		uintptr(datumType),
		uintptr(epsilon),
		uintptr(datumID),
	}
	return callProc(scAddToDataDefinition, args...)
}

// SimConnect_SetClientData: Used to write one or more units of data to a client data area.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_SetClientData.htm
func (simco *SimConnect) SetClientData(clientDataID, defineID, flags DWord, unitSize DWord, buf unsafe.Pointer) error {
	// SimConnect_SetClientData(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_DATA_ID ClientDataID,
	//  SIMCONNECT_CLIENT_DATA_DEFINITION_ID DefineID,
	//  SIMCONNECT_CLIENT_DATA_SET_FLAG Flags,
	//  DWORD dwReserved,
	//  DWORD cbUnitSize,
	//  void * pDataSet)

	const reserved DWord = 0
	args := []uintptr{
		uintptr(simco.handle),
		uintptr(clientDataID),
		uintptr(defineID),
		uintptr(flags),
		uintptr(reserved),
		uintptr(unitSize),
		uintptr(buf),
	}
	return callProc(scSetClientData, args...)

}

// SimConnect_SetDataOnSimObject: Used to make changes to the data properties of an object.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_SetDataOnSimObject.htm
func (simco *SimConnect) SetDataOnSimObject(defineID, objectID, flags, arrayCount, unitSize DWord, buf unsafe.Pointer) error {
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
		uintptr(unitSize),
		uintptr(buf),
	}
	return callProc(scSetDataOnSimObject, args...)
}

// SimConnect_ClearClientDataDefinition: Used to clear the definition of the specified client data.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_ClearClientDataDefinition.htm
func (simco *SimConnect) ClearClientDataDefinition(defineID DWord) error {
	// SimConnect_ClearClientDataDefinition(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_DATA_DEFINITION_ID DefineID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(defineID),
	}
	return callProc(scClearClientDataDefinition, args...)
}

// SimConnect_ClearDataDefinition: Used to remove all simulation variables from a client defined object.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_ClearDataDefinition.htm
func (simco *SimConnect) ClearDataDefinition(defineID DWord) error {
	// SIMCONNECTAPI SimConnect_ClearDataDefinition(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_DATA_DEFINITION_ID DefineID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(defineID),
	}
	return callProc(scClearDataDefinition, args...)
}

// SimConnect_MapInputEventToClientEvent: Used to connect input events (such as keystrokes, joystick or mouse movements) with the sending of appropriate event notifications.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_MapInputEventToClientEvent.htm
func (simco *SimConnect) MapInputEventToClientEvent(groupID DWord, inputDefinition string, downEventID DWord) error {
	// SimConnect_MapInputEventToClientEvent(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_INPUT_GROUP_ID GroupID,
	//  const char * szInputDefinition,
	//  SIMCONNECT_CLIENT_EVENT_ID DownEventID,
	//  DWORD DownValue = 0,
	//  SIMCONNECT_CLIENT_EVENT_ID UpEventID = (SIMCONNECT_CLIENT_EVENT_ID)SIMCONNECT_UNUSED,
	//  DWORD UpValue = 0,
	//  BOOL bMaskable = FALSE)

	const downValue DWord = 0
	const upEventID = Unused
	const upValue DWord = 0
	const maskable = false

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
		toCharPtr(inputDefinition),
		uintptr(downEventID),
		uintptr(downValue),
		uintptr(upEventID),
		uintptr(upValue),
		toBoolPtr(maskable),
	}
	return callProc(scMapInputEventToClientEvent, args...)
}

// SimConnect_RequestNotificationGroup: Used to request events from a notification group when the simulation is in Dialog Mode.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_RequestNotificationGroup.htm
func (simco *SimConnect) RequestNotificationGroup(groupID DWord) error {
	// SimConnect_RequestNotificationGroup(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_NOTIFICATION_GROUP_ID GroupID,
	//  DWORD dwReserved = 0,
	//  DWORD Flags = 0)

	const reserved DWord = 0
	const flags DWord = 0

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
		uintptr(reserved),
		uintptr(flags),
	}
	return callProc(scRequestNotificationGroup, args...)
}

// SimConnect_ClearInputGroup: Used to remove all the input events from a specified input group object.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_ClearInputGroup.htm
func (simco *SimConnect) ClearInputGroup(groupID DWord) error {
	// SimConnect_ClearInputGroup(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_INPUT_GROUP_ID GroupID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
	}
	return callProc(scClearInputGroup, args...)
}

// SimConnect_ClearNotificationGroup: Used to remove all the client defined events from a notification group.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_ClearNotificationGroup.htm
func (simco *SimConnect) ClearNotificationGroup(groupID DWord) error {
	// SimConnect_ClearNotificationGroup(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_NOTIFICATION_GROUP_ID GroupID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
	}
	return callProc(scClearNotificationGroup, args...)
}

// SimConnect_RequestReservedKey: Used to request a specific keyboard TAB-key combination applies only to this client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_RequestReservedKey.htm
// func (simco *SimConnect) RequestReservedKey(eventID uint32, keyChoice1, keyChoice2, keyChoice3 string) error {
// 	// SimConnect_RequestReservedKey(
// 	//  HANDLE hSimConnect,
// 	//  SIMCONNECT_CLIENT_EVENT_ID EventID,
// 	//  const char * szKeyChoice1 = "",
// 	//  const char * szKeyChoice2 = "",
// 	//  const char * szKeyChoice3 = "")
// 	return errors.New("Not implemented")
// }

// SimConnect_SetInputGroupPriority: Used to set the priority for a specified input group object.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_SetInputGroupPriority.htm
func (simco *SimConnect) SetInputGroupPriority(groupID, priority DWord) error {
	// SimConnect_SetInputGroupPriority(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_INPUT_GROUP_ID GroupID,
	//  DWORD uPriority)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
		uintptr(priority),
	}
	return callProc(scSetInputGroupPriority, args...)
}

// SimConnect_SetInputGroupState: Used to turn requests for input event information from the server on and off.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_SetInputGroupState.htm
func (simco *SimConnect) SetInputGroupState(groupID, state DWord) error {
	// SimConnect_SetInputGroupState(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_INPUT_GROUP_ID GroupID,
	//  DWORD dwState)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
		uintptr(state),
	}
	return callProc(scSetInputGroupState, args...)
}

// SimConnect_RemoveInputEvent: Used to remove an input event from a specified input group object.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Events_And_Data/SimConnect_RemoveInputEvent.htm
func (simco *SimConnect) RemoveInputEvent(groupID DWord, inputDefinition string) error {
	// SimConnect_RemoveInputEvent(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_INPUT_GROUP_ID GroupID,
	//  const char * szInputDefinition)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(groupID),
		toCharPtr(inputDefinition),
	}
	return callProc(scRemoveInputEvent, args...)
}

// AI Object functions:

// SimConnect_AICreateEnrouteATCAircraft: Used to create an AI controlled aircraft that is about to start or is already underway on its flight plan.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/AI_Object/SimConnect_AICreateEnrouteATCAircraft.htm
func (simco *SimConnect) AICreateEnrouteATCAircraft(containerTitle, tailNumber string, flightNumber int, flightPlanPath string, flightPlanPosition float64, touchAndGo bool, requestID uint32) error {
	// SimConnect_AICreateEnrouteATCAircraft(
	//  HANDLE hSimConnect,
	//  const char * szContainerTitle,
	//  const char * szTailNumber,
	//  int iFlightNumber,
	//  const char * szFlightPlanPath,
	//  double dFlightPlanPosition,
	//  BOOL bTouchAndGo,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(containerTitle)),
		uintptr(toCharPtr(tailNumber)),
		uintptr(flightNumber),
		uintptr(toCharPtr(flightPlanPath)),
		uintptr(flightPlanPosition),
		uintptr(toBoolPtr(touchAndGo)),
		uintptr(requestID),
	}
	return callProc(scAICreateEnrouteATCAircraft, args...)
}

// SimConnect_AICreateNonATCAircraft: Used to create an aircraft that is not flying under ATC control (so is typically flying under VFR rules).
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/AI_Object/SimConnect_AICreateNonATCAircraft.htm
func (simco *SimConnect) AICreateNonATCAircraft(containerTitle, tailNumber string, initPos InitPosition, requestID DWord) error {
	// SimConnect_AICreateNonATCAircraft(
	//  HANDLE hSimConnect,
	//  const char * szContainerTitle,
	//  const char * szTailNumber,
	//  SIMCONNECT_DATA_INITPOSITION InitPos,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(containerTitle)),
		uintptr(toCharPtr(tailNumber)),
		uintptr(unsafe.Pointer(&initPos)),
		uintptr(requestID),
	}
	return callProc(scAICreateNonATCAircraft, args...)
}

// SimConnect_AICreateParkedATCAircraft: Used to create an AI controlled aircraft that is currently parked and does not have a flight plan.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/AI_Object/SimConnect_AICreateParkedATCAircraft.htm
func (simco *SimConnect) AICreateParkedATCAircraft(containerTitle, tailNumber, airportID string, requestID DWord) error {
	// TODO: SimConnect_AICreateParkedATCAircraft(
	//  HANDLE hSimConnect,
	//  const char * szContainerTitle,
	//  const char * szTailNumber,
	//  const char * szAirportID,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(containerTitle)),
		uintptr(toCharPtr(tailNumber)),
		uintptr(toCharPtr(airportID)),
		uintptr(requestID),
	}
	return callProc(scAICreateParkedATCAircraft, args...)
}

// SimConnect_AICreateSimulatedObject: Used to create AI controlled objects other than aircraft.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/AI_Object/SimConnect_AICreateSimulatedObject.htm
func (simco *SimConnect) AICreateSimulatedObject(containerTitle string, initPos InitPosition, requestID DWord) error {
	// SimConnect_AICreateSimulatedObject(
	//  HANDLE hSimConnect,
	//  const char * szContainerTitle,
	//  SIMCONNECT_DATA_INITPOSITION InitPos,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(containerTitle)),
		uintptr(unsafe.Pointer(&initPos)),
		uintptr(requestID),
	}
	return callProc(scAICreateSimulatedObject, args...)
}

// SimConnect_AIReleaseControl: Used to clear the AI control of a simulated object, typically an aircraft, in order for it to be controlled by a SimConnect client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/AI_Object/SimConnect_AIReleaseControl.htm
func (simco *SimConnect) AIReleaseControl(objectID, requestID DWord) error {
	// SimConnect_AIReleaseControl(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_OBJECT_ID ObjectID,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(objectID),
		uintptr(requestID),
	}
	return callProc(scAIReleaseControl, args...)
}

// SimConnect_AIRemoveObject: Used to remove any object created by the client using one of the AI creation functions.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/AI_Object/SimConnect_AIRemoveObject.htm
func (simco *SimConnect) AIRemoveObject(objectID, requestID DWord) error {
	// SimConnect_AIRemoveObject(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_OBJECT_ID ObjectID,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(objectID),
		uintptr(requestID),
	}
	return callProc(scAIRemoveObject, args...)
}

// SimConnect_AISetAircraftFlightPlan: Used to set or change the flight plan of an AI controlled aircraft.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/AI_Object/SimConnect_AISetAircraftFlightPlan.htm
func (simco *SimConnect) AISetAircraftFlightPlan(objectID, requestID DWord, flightPlanPath string) error {
	// SimConnect_AISetAircraftFlightPlan(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_OBJECT_ID ObjectID,
	//  const char * szFlightPlanPath,
	//  SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(objectID),
		uintptr(toCharPtr(flightPlanPath)),
		uintptr(requestID),
	}
	return callProc(scAISetAircraftFlightPlan, args...)
}

// Flights functions:

// SimConnect_FlightLoad: Used to load an existing flight file.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Flights/SimConnect_FlightLoad.htm
func (simco *SimConnect) FlightLoad(fileName string) error {
	// SimConnect_FlightLoad(
	//  HANDLE hSimConnect,
	//  const char * szFileName)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(fileName)),
	}
	return callProc(scFlightLoad, args...)
}

// SimConnect_FlightSave: Used to save the current state of a flight to a flight file.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Flights/SimConnect_FlightSave.htm
func (simco *SimConnect) FlightSave(fileName, title, description string, flags DWord) error {
	// SimConnect_FlightSave(
	//  HANDLE hSimConnect,
	//  const char * szFileName,
	//  const char * szTitle,
	//  const char * szDescription,
	//  DWORD Flags)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(fileName)),
		uintptr(toCharPtr(title)),
		uintptr(toCharPtr(description)),
		uintptr(flags),
	}
	return callProc(scFlightSave, args...)
}

// SimConnect_FlightPlanLoad: Used to load an existing flight plan.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Flights/SimConnect_FlightPlanLoad.htm
// (fileName: .PLN file format, no extension)
func (simco *SimConnect) FlightPlanLoad(fileName string) error {
	// SimConnect_FlightPlanLoad(
	// HANDLE hSimConnect,
	// const char * szFileName)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(fileName)),
	}
	return callProc(scFlightPlanLoad, args...)
}

// Debug functions:

// SimConnect_GetLastSentPacketID: Returns the ID of the last packet sent to the SimConnect server.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Debug/SimConnect_GetLastSentPacketID.htm
func (simco *SimConnect) GetLastSentPacketID(pdwError *DWord) error {
	// SimConnect_GetLastSentPacketID(
	//  HANDLE hSimConnect,
	//  DWORD * pdwError);

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(unsafe.Pointer(pdwError)),
	}
	return callProc(scGetLastSentPacketID, args...)
}

// SimConnect_RequestResponseTimes: Used to provide some data on the performance of the client-server connection.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Debug/SimConnect_RequestResponseTimes.htm
// TODO: SimConnect_RequestResponseTimes(HANDLE hSimConnect, DWORD nCount, float * fElapsedSeconds)

// SimConnect_InsertString: Used to assist in adding variable length strings to a structure.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Debug/SimConnect_InsertString.htm
// TODO: SimConnect_InsertString(char * pDest, DWORD cbDest, void ** ppEnd, DWORD * pcbStringV, const char * pSource)

// SimConnect_RetrieveString: Used to assist in retrieving variable length strings from a structure.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Debug/SimConnect_RetrieveString.htm
// TODO: SimConnect_RetrieveString(SIMCONNECT_RECV * pData, DWORD cbData, void * pStringV, char ** pszString, DWORD * pcbString)

// Facilities functions:

// SimConnect_RequestFacilitesList: Request a list of all the facilities of a given type currently held in the facilities cache.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Facilities/SimConnect_RequestFacilitesList.htm
func (simco *SimConnect) RequestFacilitiesList(facilityListType, requestID DWord) error {
	// SimConnect_RequestFacilitiesList(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_FACILITY_LIST_TYPE type,
	// 	SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(facilityListType),
		uintptr(requestID),
	}
	return callProc(scRequestFacilitiesList, args...)
}

// SimConnect_SubscribeToFacilities: Used to request notifications when a facility of a certain type is added to the facilities cache.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Facilities/SimConnect_SubscribeToFacilities.htm
func (simco *SimConnect) SubscribeToFacilities(facilityListType, requestID DWord) error {
	// SimConnect_SubscribeToFacilities(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_FACILITY_LIST_TYPE type,
	// 	SIMCONNECT_DATA_REQUEST_ID RequestID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(facilityListType),
		uintptr(requestID),
	}
	return callProc(scSubscribeToFacilities, args...)
}

// SimConnect_UnsubscribeToFacilities: Used to request that notifications of additions to the facilities cache are not longer sent.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Facilities/SimConnect_UnsubscribeToFacilities.htm
func (simco *SimConnect) UnsubscribeToFacilities(facilityListType DWord) error {
	// SimConnect_UnsubscribeToFacilities(
	// 	HANDLE hSimConnect,
	// 	SIMCONNECT_FACILITY_LIST_TYPE type)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(facilityListType),
	}
	return callProc(scUnsubscribeToFacilities, args...)
}

// Mission functions:
// see https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/SimConnect_API_Reference.htm

// TODO: SimConnect_CompleteCustomMissionAction(HANDLE hSimConnect, const GUID guidInstanceId)
// TODO: SimConnect_ExecuteMissionAction(HANDLE hSimConnect, const GUID guidInstanceId)

// Menu functions:

// SimConnect_MenuAddItem is mentioned in the docs but there is no further description
// see https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/SimConnect_API_Reference.htm
func (simco *SimConnect) MenuAddItem(menuItem string, menuEventID, data DWord) error {
	// SimConnect_MenuAddItem(
	//  HANDLE hSimConnect,
	//  const char * szMenuItem,
	//  SIMCONNECT_CLIENT_EVENT_ID MenuEventID,
	//  DWORD dwData)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(menuItem)),
		uintptr(eventID),
		uintptr(data),
	}
	return callProc(scMenuAddItem, args...)
}

// SimConnect_MenuAddSubItem is mentioned in the docs but there is no further description
// see https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/SimConnect_API_Reference.htm
func (simco *SimConnect) MenuAddSubItem(menuEventID DWord, menuItem string, subMenuEventID, data DWord) error {
	// SimConnect_MenuAddSubItem(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_EVENT_ID MenuEventID,
	//  const char * szMenuItem,
	//  SIMCONNECT_CLIENT_EVENT_ID SubMenuEventID,
	//  DWORD dwData)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(menuEventID),
		uintptr(toCharPtr(menuItem)),
		uintptr(subMenuEventID),
		uintptr(data),
	}
	return callProc(scMenuAddSubItem, args...)
}

// SimConnect_MenuDeleteItem is mentioned in the docs but there is no further description
// see https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/SimConnect_API_Reference.htm
func (simco *SimConnect) MenuDeleteItem(menuEventID DWord) error {
	// SimConnect_MenuDeleteItem(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_EVENT_ID MenuEventID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(menuEventID),
	}
	return callProc(scMenuDeleteItem, args...)
}

// SimConnect_MenuDeleteSubItem is mentioned in the docs but there is no further description
// see https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/SimConnect_API_Reference.htm
func (simco *SimConnect) MenuDeleteSubItem(menuEventID, subMenuEventID DWord) error {
	// SimConnect_MenuDeleteSubItem(
	//  HANDLE hSimConnect,
	//  SIMCONNECT_CLIENT_EVENT_ID MenuEventID,
	//  const SIMCONNECT_CLIENT_EVENT_ID SubMenuEventID)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(menuEventID),
		uintptr(subMenuEventID),
	}
	return callProc(scMenuDeleteSubItem, args...)
}

// SimConnect_CameraSetRelative6DOF is not documented (see SimConnect.h)
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
	return callProc(scCameraSetRelative6DOF, args...)
}

// SimConnect_SetSystemState is not documented (see SimConnect.h)
func (simco *SimConnect) SetSystemState(state string, integerValue DWord, floatValue float32, stringValue string) error {
	// SimConnect_SetSystemState(
	//  HANDLE hSimConnect,
	//  const char * szState,
	//  DWORD dwInteger,
	//  float fFloat,
	//  const char * szString)

	args := []uintptr{
		uintptr(simco.handle),
		uintptr(toCharPtr(state)),
		uintptr(integerValue),
		uintptr(floatValue),
		uintptr(toCharPtr(stringValue)),
	}
	return callProc(scSetSystemState, args...)
}
