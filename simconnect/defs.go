package simconnect

const simConnectLibraryName string = "SimConnect.dll"

const (
	simConnectAddToDataDefinitionProcName        = "SimConnect_AddToDataDefinition"
	simConnectClearDataDefinitionProcName        = "SimConnect_ClearDataDefinition"
	simConnectRequestDataOnSimObjectTypeProcName = "SimConnect_RequestDataOnSimObjectType"
	simConnectSetDataOnSimObjectProcName         = "SimConnect_SetDataOnSimObject"
	simConnectCloseProcName                      = "SimConnect_Close"
	simConnectOpenProcName                       = "SimConnect_Open"
	simConnectGetNextDispatchProcName            = "SimConnect_GetNextDispatch"
	simConnectCameraSetRelative6DOFProcName      = "SimConnect_CameraSetRelative6DOF"
	simConnectTextProcName                       = "SimConnect_Text"
	simConnectSubscribeToFacilitiesProcName      = "SimConnect_SubscribeToFacilities"
	simConnectUnsubscribeToFacilitiesProcName    = "SimConnect_UnsubscribeToFacilities"
	simConnectRequestFacilitiesListProcName      = "SimConnect_RequestFacilitiesList"
)

type DWord uint32 // DWORD

const (
	DWordMax         DWord  = 0xffffffff // DWORD_MAX
	EFail            uint32 = 0x80004005 // E_FAIL
	ObjectIDUser     DWord  = 0          // OBJECT_ID_USER
	Unused           DWord  = DWordMax   // UNUSED
	WmUserSimConnect DWord  = 0x0402     // WM_USER_SIMCONNECT
	DWordZero        DWord  = 0
)

// SIMCONNECT_RECV_ID
const (
	RecvIDNull                          = iota // RECV_ID_NULL
	RecvIDException                            // RECV_ID_EXCEPTION
	RecvIDOpen                                 // RECV_ID_OPEN
	RecvIDQuit                                 // RECV_ID_QUIT
	RecvIDEvent                                // RECV_ID_EVENT
	RecvIDEventObjectAddRemove                 // RECV_ID_EVENT_OBJECT_ADDREMOVE
	RecvIDEventFilename                        // RECV_ID_EVENT_FILENAME
	RecvIDEventFrame                           // RECV_ID_EVENT_FRAME
	RecvIDSimobjectData                        // RECV_ID_SIMOBJECT_DATA
	RecvIDSimObjectDataByType                  // RECV_ID_SIMOBJECT_DATA_BYTYPE
	RecvIDWeatherObservation                   // RECV_ID_WEATHER_OBSERVATION
	RecvIDCloudState                           // RECV_ID_CLOUD_STATE
	RecvIDAssignedObjectID                     // RECV_ID_ASSIGNED_OBJECT_ID
	RecvIDReservedKey                          // RECV_ID_RESERVED_KEY
	RecvIDCustomAction                         // RECV_ID_CUSTOM_ACTION
	RecvIDSystemState                          // RECV_ID_SYSTEM_STATE
	RecvIDClientData                           // RECV_ID_CLIENT_DATA
	RecvIDEventWeatherMode                     // RECV_ID_EVENT_WEATHER_MODE
	RecvIDAirportList                          // RECV_ID_AIRPORT_LIST
	RecvIDVORList                              // RECV_ID_VOR_LIST
	RecvIDNDBList                              // RECV_ID_NDB_LIST
	RecvIDWaypointList                         // RECV_ID_WAYPOINT_LIST
	RecvIDEventMultiplayerServerStarted        // RECV_ID_EVENT_MULTIPLAYER_SERVER_STARTED
	RecvIDEventMultiplayerClientStarted        // RECV_ID_EVENT_MULTIPLAYER_CLIENT_STARTED
	RecvIDEventMultiplayerSessionEnded         // RECV_ID_EVENT_MULTIPLAYER_SESSION_ENDED
	RecvIDEventRaceEnd                         // RECV_ID_EVENT_RACE_END
	RecvIDEventRaceLap                         // RECV_ID_EVENT_RACE_LAP
	RecvIDPick                                 // RECV_ID_PICK
)

// SIMCONNECT_DATATYPE
const (
	DataTypeInvalid      = iota // DATATYPE_INVALID
	DataTypeInt32               // DATATYPE_INT32
	DataTypeInt64               // DATATYPE_INT64
	DataTypeFloat32             // DATATYPE_FLOAT32
	DataTypeFloat64             // DATATYPE_FLOAT64
	DataTypeString8             // DATATYPE_STRING8
	DataTypeString32            // DATATYPE_STRING32
	DataTypeString64            // DATATYPE_STRING64
	DataTypeString128           // DATATYPE_STRING128
	DataTypeString256           // DATATYPE_STRING256
	DataTypeString260           // DATATYPE_STRING260
	DataTypeStringV             // DATATYPE_STRINGV
	DataTypeInitPosition        // DATATYPE_INITPOSITION
	DataTypeMarkerState         // DATATYPE_MARKERSTATE
	DataTypeWaypoint            // DATATYPE_WAYPOINT
	DataTypeLatLonAlt           // DATATYPE_LATLONALT
	DataTypeXYZ                 // DATATYPE_XYZ
)

// SIMCONNECT_EXCEPTION
const (
	ExceptionNone                          = iota // EXCEPTION_NONE
	ExceptionError                                // EXCEPTION_ERROR
	ExceptionSizeMismatch                         // EXCEPTION_SIZE_MISMATCH
	ExceptionUnrecognizedID                       // EXCEPTION_UNRECOGNIZED_ID
	ExceptionUnopened                             // EXCEPTION_UNOPENED
	ExceptionVersionMismatch                      // EXCEPTION_VERSION_MISMATCH
	ExceptionTooManyGroups                        // EXCEPTION_TOO_MANY_GROUPS
	ExceptionNameUnrecognized                     // EXCEPTION_NAME_UNRECOGNIZED
	ExceptionTooManyEventNames                    // EXCEPTION_TOO_MANY_EVENT_NAMES
	ExceptionEventIDDuplicate                     // EXCEPTION_EVENT_ID_DUPLICATE
	ExceptionTooManyMaps                          // EXCEPTION_TOO_MANY_MAPS
	ExceptionTooManyObjects                       // EXCEPTION_TOO_MANY_OBJECTS
	ExceptionTooManyRequests                      // EXCEPTION_TOO_MANY_REQUESTS
	ExceptionWeatherInvalidPort                   // EXCEPTION_WEATHER_INVALID_PORT
	ExceptionWeatherInvalidMetar                  // EXCEPTION_WEATHER_INVALID_METAR
	ExceptionWeatherUnableToGetObservation        // EXCEPTION_WEATHER_UNABLE_TO_GET_OBSERVATION
	ExceptionWeatherUnableToCreateStation         // EXCEPTION_WEATHER_UNABLE_TO_CREATE_STATION
	ExceptionWeatherUnableToRemoveStation         // EXCEPTION_WEATHER_UNABLE_TO_REMOVE_STATION
	ExceptionInvalidDataType                      // EXCEPTION_INVALID_DATA_TYPE
	ExceptionInvalidDataSize                      // EXCEPTION_INVALID_DATA_SIZE
	ExceptionDataError                            // EXCEPTION_DATA_ERROR
	ExceptionInvalidArray                         // EXCEPTION_INVALID_ARRAY
	ExceptionCreateObjectFailed                   // EXCEPTION_CREATE_OBJECT_FAILED
	ExceptionLoadFlightplanFailed                 // EXCEPTION_LOAD_FLIGHTPLAN_FAILED
	ExceptionOperationInvalidForObjectType        // EXCEPTION_OPERATION_INVALID_FOR_OBJECT_TYPE
	ExceptionIllegalOperation                     // EXCEPTION_ILLEGAL_OPERATION
	ExceptionAlreadySubscribed                    // EXCEPTION_ALREADY_SUBSCRIBED
	ExceptionInvalidEnum                          // EXCEPTION_INVALID_ENUM
	ExceptionDefinitionError                      // EXCEPTION_DEFINITION_ERROR
	ExceptionDuplicateID                          // EXCEPTION_DUPLICATE_ID
	ExceptionDatumID                              // EXCEPTION_DATUM_ID
	ExceptionOutOfBounds                          // EXCEPTION_OUT_OF_BOUNDS
	ExceptionAlreadyCreated                       // EXCEPTION_ALREADY_CREATED
	ExceptionObjectOutsideRealityBubble           // EXCEPTION_OBJECT_OUTSIDE_REALITY_BUBBLE
	ExceptionObjectContainer                      // EXCEPTION_OBJECT_CONTAINER
	ExceptionObjectAt                             // EXCEPTION_OBJECT_AI
	ExceptionObjectATC                            // EXCEPTION_OBJECT_ATC
	ExceptionObjectSchedule                       // EXCEPTION_OBJECT_SCHEDULE
)

// SIMCONNECT_SIMOBJECT_TYPE
const (
	SimObjectTypeUser       DWord = iota // SIMOBJECT_TYPE_USER
	SimObjectTypeAll                     // SIMOBJECT_TYPE_ALL
	SimObjectTypeAircraft                // SIMOBJECT_TYPE_AIRCRAFT
	SimObjectTypeHelicopter              // SIMOBJECT_TYPE_HELICOPTER
	SimObjectTypeBoat                    // SIMOBJECT_TYPE_BOAT
	SimObjectTypeGround                  // SIMOBJECT_TYPE_GROUND
)

// SIMCONNECT_TEXT_TYPE
const (
	TextTypeScrollBlack   DWord = iota          // TEXT_TYPE_SCROLL_BLACK DWord
	TextTypeScrollWhite                         // TEXT_TYPE_SCROLL_WHITE
	TextTypeScrollRed                           // TEXT_TYPE_SCROLL_RED
	TextTypeScrollGreen                         // TEXT_TYPE_SCROLL_GREEN
	TextTypeScrollBlue                          // TEXT_TYPE_SCROLL_BLUE
	TextTypeScrollYellow                        // TEXT_TYPE_SCROLL_YELLOW
	TextTypeScrollMagenta                       // TEXT_TYPE_SCROLL_MAGENTA
	TextTypeScrollCyan                          // TEXT_TYPE_SCROLL_CYAN
	TextTypePrintBlack    DWord = iota + 0x0100 // TEXT_TYPE_PRINT_BLACK
	TextTypePrintWhite                          // TEXT_TYPE_PRINT_WHITE
	TextTypePrintRed                            // TEXT_TYPE_PRINT_RED
	TextTypePrintGreen                          // TEXT_TYPE_PRINT_GREEN
	TextTypePrintBlue                           // TEXT_TYPE_PRINT_BLUE
	TextTypePrintYellow                         // TEXT_TYPE_PRINT_YELLOW
	TextTypePrintMagenta                        // TEXT_TYPE_PRINT_MAGENTA
	TextTypePrintCyan                           // TEXT_TYPE_PRINT_CYAN
	TextTypeMenu          DWord = iota + 0x0200 // TEXT_TYPE_MENU
)

// SIMCONNECT_RECV
type Recv struct {
	Size    DWord
	Version DWord
	ID      DWord
}

// SIMCONNECT_RECV_EXCEPTION
type RecvException struct {
	Recv
	Exception DWord // SIMCONNECT_EXCEPTION
	SendID    DWord // SimConnect_GetLastSentPacketID
	Index     DWord // index of parameter that was source of error
}

// SIMCONNECT_RECV_OPEN
type RecvOpen struct {
	Recv
	ApplicationName         [256]byte
	ApplicationVersionMajor DWord
	ApplicationVersionMinor DWord
	ApplicationBuildMajor   DWord
	ApplicationBuildMinor   DWord
	SimConnectVersionMajor  DWord
	SimConnectVersionMinor  DWord
	SimConnectBuildMajor    DWord
	SimConnectBuildMinor    DWord
	Reserved1               DWord
	Reserved2               DWord
}

// SIMCONNECT_RECV_QUIT
type RecvQuit struct {
	Recv
}

// SIMCONNECT_RECV_EVENT
type RecvEvent struct {
	Recv
	GroupID DWord
	EventID DWord
	Data    DWord
}

// SIMCONNECT_RECV_SIMOBJECT_DATA
type RecvSimObjectData struct {
	Recv
	RequestID   DWord
	ObjectID    DWord
	DefineID    DWord
	Flags       DWord
	EntryNumber DWord
	OutOf       DWord
	DefineCount DWord
}

// SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE
type RecvSimObjectDataByType struct {
	RecvSimObjectData
}

// SIMCONNECT_DATATYPE_INITPOSITION
type InitPositionDataType struct {
	Latitude  float64 // degrees
	Longitude float64 // degrees
	Altitude  float64 // feet
	Pitch     float64 // degrees
	Bank      float64 // degrees
	Heading   float64 // degrees
	OnGround  DWord   // 1=force to be on the ground
	Airspeed  DWord   // knots
}

// SIMCONNECT_DATATYPE_MARKERSTATE
type MarkerStateDataType struct {
	MarkerName  [64]byte
	MarkerState DWord
}

// SIMCONNECT_DATATYPE_WAYPOINT
type WaypointDataType struct {
	Latitude        float64 // degrees
	Longitude       float64 // degrees
	Altitude        float64 // feet
	Flags           uint32  // unsigned long
	ktsSpeed        float64 // knots
	percentThrottle float64
}

// SIMCONNECT_DATA_LATLONALT
type LatLogAltDataType struct {
	Latitude, Longitude, Altitude float64
}

// // SIMCONNECT_DATA_XYZ
type XYZDataType struct {
	x, y, z float64
}
