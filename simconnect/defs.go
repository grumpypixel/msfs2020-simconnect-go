package simconnect

import (
	"math"
)

/*
	Transfused code from: SimConnect.h
	Location: /MSFS SDK/SimConnect SDK/include
	Copyright (c) Microsoft Corporation. All Rights Reserved.

	Please also check out the official documentation:
	https://docs.flightsimulator.com/html/index.htm
*/

const (
	// General
	scOpen                         = "SimConnect_Open"
	scClose                        = "SimConnect_Close"
	scCallDispatch                 = "SimConnect_CallDispatch" // Not implemented
	scGetNextDispatch              = "SimConnect_GetNextDispatch"
	scRequestSystemState           = "SimConnect_RequestSystemState"
	scMapClientEventToSimEvent     = "SimConnect_MapClientEventToSimEvent"
	scSubscribeToSystemEvent       = "SimConnect_SubscribeToSystemEvent"
	scSetSystemEventState          = "SimConnect_SetSystemEventState"
	scUnsubscribeFromSystemEvent   = "SimConnect_UnsubscribeFromSystemEvent"
	scSetNotificationGroupPriority = "SimConnect_SetNotificationGroupPriority"
	scText                         = "SimConnect_Text"
	// Events and Data
	scRequestDataOnSimObject            = "SimConnect_RequestDataOnSimObject"
	scRequestDataOnSimObjectType        = "SimConnect_RequestDataOnSimObjectType"
	scAddClientEventToNotificationGroup = "SimConnect_AddClientEventToNotificationGroup"
	scRemoveClientEvent                 = "SimConnect_RemoveClientEvent"
	scTransmitClientEvent               = "SimConnect_TransmitClientEvent"
	scMapClientDataNameToID             = "SimConnect_MapClientDataNameToID"
	scRequestClientData                 = "SimConnect_RequestClientData"
	scCreateClientData                  = "SimConnect_CreateClientData"
	scAddToClientDataDefinition         = "SimConnect_AddToClientDataDefinition"
	scAddToDataDefinition               = "SimConnect_AddToDataDefinition"
	scSetClientData                     = "SimConnect_SetClientData"
	scSetDataOnSimObject                = "SimConnect_SetDataOnSimObject"
	scClearClientDataDefinition         = "SimConnect_ClearClientDataDefinition"
	scClearDataDefinition               = "SimConnect_ClearDataDefinition"
	scMapInputEventToClientEvent        = "SimConnect_MapInputEventToClientEvent"
	scRequestNotificationGroup          = "SimConnect_RequestNotificationGroup"
	scClearInputGroup                   = "SimConnect_ClearInputGroup"
	scClearNotificationGroup            = "SimConnect_ClearNotificationGroup"
	scRequestReservedKey                = "SimConnect_RequestReservedKey" // Not implemented
	scSetInputGroupPriority             = "SimConnect_SetInputGroupPriority"
	scSetInputGroupState                = "SimConnect_SetInputGroupState"
	scRemoveInputEvent                  = "SimConnect_RemoveInputEvent"
	// AI Objects
	scAICreateEnrouteATCAircraft = "SimConnect_AICreateEnrouteATCAircraft"
	scAICreateNonATCAircraft     = "SimConnect_AICreateNonATCAircraft"
	scAICreateParkedATCAircraft  = "SimConnect_AICreateParkedATCAircraft"
	scAICreateSimulatedObject    = "SimConnect_AICreateSimulatedObject"
	scAIReleaseControl           = "SimConnect_AIReleaseControl"
	scAIRemoveObject             = "SimConnect_AIRemoveObject"
	scAISetAircraftFlightPlan    = "SimConnect_AISetAircraftFlightPlan"
	// Flights
	scFlightLoad     = "SimConnect_FlightLoad"
	scFlightSave     = "SimConnect_FlightSave"
	scFlightPlanLoad = "SimConnect_FlightPlanLoad"
	// Debug
	scGetLastSentPacketID  = "SimConnect_GetLastSentPacketID"
	scRequestResponseTimes = "SimConnect_RequestResponseTimes" // Not implemented
	scInsertString         = "SimConnect_InsertString"         // Not implemented
	scRetrieveString       = "SimConnect_RetrieveString"       // Not implemented
	// Facilities
	scRequestFacilitiesList   = "SimConnect_RequestFacilitiesList"
	scSubscribeToFacilities   = "SimConnect_SubscribeToFacilities"
	scUnsubscribeToFacilities = "SimConnect_UnsubscribeToFacilities"
	// Missions
	// scCompleteCustomMissionAction = "SimConnect_CompleteCustomMissionAction" // Not implemented
	// scExecuteMissionAction        = "SimConnect_ExecuteMissionAction"        // Not implemented
	// Menu
	scMenuAddItem       = "SimConnect_MenuAddItem"
	scMenuAddSubItem    = "SimConnect_MenuAddSubItem"
	scMenuDeleteItem    = "SimConnect_MenuDeleteItem"
	scMenuDeleteSubItem = "SimConnect_MenuDeleteSubItem"
	// Undocumented
	scCameraSetRelative6DOF = "SimConnect_CameraSetRelative6DOF"
	scSetSystemState        = "SimConnect_SetSystemState"
)

type DWord uint32 // DWORD

const (
	DWordMax          DWord   = 0xffffffff      // DWORD_MAX
	EFail             uint32  = 0x80004005      // E_FAIL
	Unused            DWord   = DWordMax        // SIMCONNECT_UNUSED
	ObjectIDUser      DWord   = 0               // SIMCONNECT_OBJECT_ID_USER
	CameraIgnoreField float32 = math.MaxFloat32 // SIMCONNECT_CAMERA_IGNORE_FIELD: Used to tell the Camera API to NOT modify the value in this part of the argument.
	ClientDataMaxSize DWord   = 8192            // SIMCONNECT_CLIENTDATA_MAX_SIZE: maximum value for SimConnect_CreateClientData dwSize parameter
	WmUserSimConnect  DWord   = 0x0402          // WM_USER_SIMCONNECT
	DWordZero         DWord   = 0
)

// Notification Group priority values
const (
	GroupPriorityHighest         DWord = 1          // SIMCONNECT_GROUP_PRIORITY_HIGHEST: highest priority
	GroupPriorityHighestMaskable DWord = 10000000   // SIMCONNECT_GROUP_PRIORITY_HIGHEST_MASKABLE: highest priority that allows events to be masked
	GroupPriorityStandard        DWord = 1900000000 // SIMCONNECT_GROUP_PRIORITY_STANDARD: standard priority
	GroupPriorityDefault         DWord = 2000000000 // SIMCONNECT_GROUP_PRIORITY_DEFAULT: default priority
	GroupPriorityLowest          DWord = 4000000000 // SIMCONNECT_GROUP_PRIORITY_LOWEST: priorities lower than this will be ignored
)

// SIMCONNECT_RECV_ID: Receive data types
const (
	RecvIDNull                          = iota // SIMCONNECT_RECV_ID_NULL
	RecvIDException                            // SIMCONNECT_RECV_ID_EXCEPTION
	RecvIDOpen                                 // SIMCONNECT_RECV_ID_OPEN
	RecvIDQuit                                 // SIMCONNECT_RECV_ID_QUIT
	RecvIDEvent                                // SIMCONNECT_RECV_ID_EVENT
	RecvIDEventObjectAddRemove                 // SIMCONNECT_RECV_ID_EVENT_OBJECT_ADDREMOVE
	RecvIDEventFilename                        // SIMCONNECT_RECV_ID_EVENT_FILENAME
	RecvIDEventFrame                           // SIMCONNECT_RECV_ID_EVENT_FRAME
	RecvIDSimobjectData                        // SIMCONNECT_RECV_ID_SIMOBJECT_DATA
	RecvIDSimObjectDataByType                  // SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE
	RecvIDWeatherObservation                   // SIMCONNECT_RECV_ID_WEATHER_OBSERVATION
	RecvIDCloudState                           // SIMCONNECT_RECV_ID_CLOUD_STATE
	RecvIDAssignedObjectID                     // SIMCONNECT_RECV_ID_ASSIGNED_OBJECT_ID
	RecvIDReservedKey                          // SIMCONNECT_RECV_ID_RESERVED_KEY
	RecvIDCustomAction                         // SIMCONNECT_RECV_ID_CUSTOM_ACTION
	RecvIDSystemState                          // SIMCONNECT_RECV_ID_SYSTEM_STATE
	RecvIDClientData                           // SIMCONNECT_RECV_ID_CLIENT_DATA
	RecvIDEventWeatherMode                     // SIMCONNECT_RECV_ID_EVENT_WEATHER_MODE
	RecvIDAirportList                          // SIMCONNECT_RECV_ID_AIRPORT_LIST
	RecvIDVORList                              // SIMCONNECT_RECV_ID_VOR_LIST
	RecvIDNDBList                              // SIMCONNECT_RECV_ID_NDB_LIST
	RecvIDWaypointList                         // SIMCONNECT_RECV_ID_WAYPOINT_LIST
	RecvIDEventMultiplayerServerStarted        // SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_SERVER_STARTED
	RecvIDEventMultiplayerClientStarted        // SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_CLIENT_STARTED
	RecvIDEventMultiplayerSessionEnded         // SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_SESSION_ENDED
	RecvIDEventRaceEnd                         // SIMCONNECT_RECV_ID_EVENT_RACE_END
	RecvIDEventRaceLap                         // SIMCONNECT_RECV_ID_EVENT_RACE_LAP
	RecvIDPick                                 // SIMCONNECT_RECV_ID_PICK
)

// SIMCONNECT_DATATYPE: Data data types
const (
	DataTypeInvalid      = iota // SIMCONNECT_DATATYPE_INVALID
	DataTypeInt32               // SIMCONNECT_DATATYPE_INT32
	DataTypeInt64               // SIMCONNECT_DATATYPE_INT64
	DataTypeFloat32             // SIMCONNECT_DATATYPE_FLOAT32
	DataTypeFloat64             // SIMCONNECT_DATATYPE_FLOAT64
	DataTypeString8             // SIMCONNECT_DATATYPE_STRING8
	DataTypeString32            // SIMCONNECT_DATATYPE_STRING32
	DataTypeString64            // SIMCONNECT_DATATYPE_STRING64
	DataTypeString128           // SIMCONNECT_DATATYPE_STRING128
	DataTypeString256           // SIMCONNECT_DATATYPE_STRING256
	DataTypeString260           // SIMCONNECT_DATATYPE_STRING260
	DataTypeStringV             // SIMCONNECT_DATATYPE_STRINGV
	DataTypeInitPosition        // SIMCONNECT_DATATYPE_INITPOSITION
	DataTypeMarkerState         // SIMCONNECT_DATATYPE_MARKERSTATE
	DataTypeWaypoint            // SIMCONNECT_DATATYPE_WAYPOINT
	DataTypeLatLonAlt           // SIMCONNECT_DATATYPE_LATLONALT
	DataTypeXYZ                 // SIMCONNECT_DATATYPE_XYZ
)

// SIMCONNECT_EXCEPTION: Exception error types
const (
	ExceptionNone                          = iota // SIMCONNECT_EXCEPTION_NONE
	ExceptionError                                // SIMCONNECT_EXCEPTION_ERROR
	ExceptionSizeMismatch                         // SIMCONNECT_EXCEPTION_SIZE_MISMATCH
	ExceptionUnrecognizedID                       // SIMCONNECT_EXCEPTION_UNRECOGNIZED_ID
	ExceptionUnopened                             // SIMCONNECT_EXCEPTION_UNOPENED
	ExceptionVersionMismatch                      // SIMCONNECT_EXCEPTION_VERSION_MISMATCH
	ExceptionTooManyGroups                        // SIMCONNECT_EXCEPTION_TOO_MANY_GROUPS
	ExceptionNameUnrecognized                     // SIMCONNECT_EXCEPTION_NAME_UNRECOGNIZED
	ExceptionTooManyEventNames                    // SIMCONNECT_EXCEPTION_TOO_MANY_EVENT_NAMES
	ExceptionEventIDDuplicate                     // SIMCONNECT_EXCEPTION_EVENT_ID_DUPLICATE
	ExceptionTooManyMaps                          // SIMCONNECT_EXCEPTION_TOO_MANY_MAPS
	ExceptionTooManyObjects                       // SIMCONNECT_EXCEPTION_TOO_MANY_OBJECTS
	ExceptionTooManyRequests                      // SIMCONNECT_EXCEPTION_TOO_MANY_REQUESTS
	ExceptionWeatherInvalidPort                   // SIMCONNECT_EXCEPTION_WEATHER_INVALID_PORT
	ExceptionWeatherInvalidMetar                  // SIMCONNECT_EXCEPTION_WEATHER_INVALID_METAR
	ExceptionWeatherUnableToGetObservation        // SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_GET_OBSERVATION
	ExceptionWeatherUnableToCreateStation         // SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_CREATE_STATION
	ExceptionWeatherUnableToRemoveStation         // SIMCONNECT_EXCEPTION_WEATHER_UNABLE_TO_REMOVE_STATION
	ExceptionInvalidDataType                      // SIMCONNECT_EXCEPTION_INVALID_DATA_TYPE
	ExceptionInvalidDataSize                      // SIMCONNECT_EXCEPTION_INVALID_DATA_SIZE
	ExceptionDataError                            // SIMCONNECT_EXCEPTION_DATA_ERROR
	ExceptionInvalidArray                         // SIMCONNECT_EXCEPTION_INVALID_ARRAY
	ExceptionCreateObjectFailed                   // SIMCONNECT_EXCEPTION_CREATE_OBJECT_FAILED
	ExceptionLoadFlightplanFailed                 // SIMCONNECT_EXCEPTION_LOAD_FLIGHTPLAN_FAILED
	ExceptionOperationInvalidForObjectType        // SIMCONNECT_EXCEPTION_OPERATION_INVALID_FOR_OBJECT_TYPE
	ExceptionIllegalOperation                     // SIMCONNECT_EXCEPTION_ILLEGAL_OPERATION
	ExceptionAlreadySubscribed                    // SIMCONNECT_EXCEPTION_ALREADY_SUBSCRIBED
	ExceptionInvalidEnum                          // SIMCONNECT_EXCEPTION_INVALID_ENUM
	ExceptionDefinitionError                      // SIMCONNECT_EXCEPTION_DEFINITION_ERROR
	ExceptionDuplicateID                          // SIMCONNECT_EXCEPTION_DUPLICATE_ID
	ExceptionDatumID                              // SIMCONNECT_EXCEPTION_DATUM_ID
	ExceptionOutOfBounds                          // SIMCONNECT_EXCEPTION_OUT_OF_BOUNDS
	ExceptionAlreadyCreated                       // SIMCONNECT_EXCEPTION_ALREADY_CREATED
	ExceptionObjectOutsideRealityBubble           // SIMCONNECT_EXCEPTION_OBJECT_OUTSIDE_REALITY_BUBBLE
	ExceptionObjectContainer                      // SIMCONNECT_EXCEPTION_OBJECT_CONTAINER
	ExceptionObjectAt                             // SIMCONNECT_EXCEPTION_OBJECT_AI
	ExceptionObjectATC                            // SIMCONNECT_EXCEPTION_OBJECT_ATC
	ExceptionObjectSchedule                       // SIMCONNECT_EXCEPTION_OBJECT_SCHEDULE
)

// SIMCONNECT_SIMOBJECT_TYPE: Object types
const (
	SimObjectTypeUser       DWord = iota // SIMCONNECT_SIMOBJECT_TYPE_USER
	SimObjectTypeAll                     // SIMCONNECT_SIMOBJECT_TYPE_ALL
	SimObjectTypeAircraft                // SIMCONNECT_SIMOBJECT_TYPE_AIRCRAFT
	SimObjectTypeHelicopter              // SIMCONNECT_SIMOBJECT_TYPE_HELICOPTER
	SimObjectTypeBoat                    // SIMCONNECT_SIMOBJECT_TYPE_BOAT
	SimObjectTypeGround                  // SIMCONNECT_SIMOBJECT_TYPE_GROUND
)

// SIMCONNECT_STATE: EventState values
const (
	StateOff DWord = iota // SIMCONNECT_STATE_OFF
	StateOn               // SIMCONNECT_STATE_ON
)

// SIMCONNECT_PERIOD: Object Data Request Period values
const (
	PeriodNever       DWord = iota // SIMCONNECT_PERIOD_NEVER
	PeriodOnce                     // SIMCONNECT_PERIOD_ONCE
	PeriodVisualFrame              // SIMCONNECT_PERIOD_VISUAL_FRAME
	PeriodSimFrame                 // SIMCONNECT_PERIOD_SIM_FRAME
	PeriodSecond                   // SIMCONNECT_PERIOD_SECOND
)

// SIMCONNECT_MISSION_END
const (
	MissionFailed    DWord = iota // SIMCONNECT_MISSION_FAILED
	MissionCrashed                // SIMCONNECT_MISSION_CRASHED
	MissionSucceeded              // SIMCONNECT_MISSION_SUCCEEDED
)

// SIMCONNECT_CLIENT_DATA_PERIOD: ClientData Request Period values
// Used with the SimConnect_RequestClientData call to specify how often data is to be sent to the client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_CLIENT_DATA_PERIOD.htm
const (
	ClientDataPeriodNever       DWord = iota // SIMCONNECT_CLIENT_DATA_PERIOD_NEVER
	ClientDataPeriodOnce                     // SIMCONNECT_CLIENT_DATA_PERIOD_ONCE
	ClientDataPeriodVisualFrame              // SIMCONNECT_CLIENT_DATA_PERIOD_VISUAL_FRAME
	ClientDataPeriodOnSet                    // SIMCONNECT_CLIENT_DATA_PERIOD_ON_SET
	ClientDataPeriodSecond                   // SIMCONNECT_CLIENT_DATA_PERIOD_SECOND
)

// SIMCONNECT_TEXT_TYPE
// Used to specify which type of text is to be displayed by the SimConnect_Text function
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_TEXT_TYPE.htm
const (
	TextTypeScrollBlack   DWord = iota          // SIMCONNECT_TEXT_TYPE_SCROLL_BLACK DWord
	TextTypeScrollWhite                         // SIMCONNECT_TEXT_TYPE_SCROLL_WHITE
	TextTypeScrollRed                           // SIMCONNECT_TEXT_TYPE_SCROLL_RED
	TextTypeScrollGreen                         // SIMCONNECT_TEXT_TYPE_SCROLL_GREEN
	TextTypeScrollBlue                          // SIMCONNECT_TEXT_TYPE_SCROLL_BLUE
	TextTypeScrollYellow                        // SIMCONNECT_TEXT_TYPE_SCROLL_YELLOW
	TextTypeScrollMagenta                       // SIMCONNECT_TEXT_TYPE_SCROLL_MAGENTA
	TextTypeScrollCyan                          // SIMCONNECT_TEXT_TYPE_SCROLL_CYAN
	TextTypePrintBlack    DWord = iota + 0x0100 // SIMCONNECT_TEXT_TYPE_PRINT_BLACK
	TextTypePrintWhite                          // SIMCONNECT_TEXT_TYPE_PRINT_WHITE
	TextTypePrintRed                            // SIMCONNECT_TEXT_TYPE_PRINT_RED
	TextTypePrintGreen                          // SIMCONNECT_TEXT_TYPE_PRINT_GREEN
	TextTypePrintBlue                           // SIMCONNECT_TEXT_TYPE_PRINT_BLUE
	TextTypePrintYellow                         // SIMCONNECT_TEXT_TYPE_PRINT_YELLOW
	TextTypePrintMagenta                        // SIMCONNECT_TEXT_TYPE_PRINT_MAGENTA
	TextTypePrintCyan                           // SIMCONNECT_TEXT_TYPE_PRINT_CYAN
	TextTypeMenu          DWord = iota + 0x0200 // SIMCONNECT_TEXT_TYPE_MENU
)

// SIMCONNECT_TEXT_RESULT
// Used to specify which event has occurred as a result of a call to SimConnect_Text.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_TEXT_RESULT.htm
const (
	TextResultMenuSelect1  DWord = iota       // SIMCONNECT_TEXT_RESULT_MENU_SELECT_1
	TextResultMenuSelect2                     // SIMCONNECT_TEXT_RESULT_MENU_SELECT_2
	TextResultMenuSelect3                     // SIMCONNECT_TEXT_RESULT_MENU_SELECT_3
	TextResultMenuSelect4                     // SIMCONNECT_TEXT_RESULT_MENU_SELECT_4
	TextResultMenuSelect5                     // SIMCONNECT_TEXT_RESULT_MENU_SELECT_5
	TextResultMenuSelect6                     // SIMCONNECT_TEXT_RESULT_MENU_SELECT_6
	TextResultMenuSelect7                     // SIMCONNECT_TEXT_RESULT_MENU_SELECT_7
	TextResultMenuSelect8                     // SIMCONNECT_TEXT_RESULT_MENU_SELECT_8
	TextResultMenuSelect9                     // SIMCONNECT_TEXT_RESULT_MENU_SELECT_9
	TextResultMenuSelect10                    // SIMCONNECT_TEXT_RESULT_MENU_SELECT_10
	TextResultDisplayed    DWord = 0x00010000 // SIMCONNECT_TEXT_RESULT_DISPLAYED = 0x00010000
	TextResultQueued                          // SIMCONNECT_TEXT_RESULT_QUEUED
	TextResultRemoved                         // SIMCONNECT_TEXT_RESULT_REMOVED
	TextResultReplaced                        // SIMCONNECT_TEXT_RESULT_REPLACED
	TextResultTimeout                         // SIMCONNECT_TEXT_RESULT_TIMEOUT
)

// // SIMCONNECT_WEATHER_MODE
// const (
// 	WeatherModeTheme  DWord = iota // SIMCONNECT_WEATHER_MODE_THEME
// 	WeatherModeRWW                 // SIMCONNECT_WEATHER_MODE_RWW
// 	WeatherModeCustom              // SIMCONNECT_WEATHER_MODE_CUSTOM
// 	WeatherModeGlobal              // SIMCONNECT_WEATHER_MODE_GLOBAL
// )

// SIMCONNECT_FACILITY_LIST_TYPE
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_FACILITY_LIST_TYPE.htm
// Used to determine which type of facilities data is being requested or returned.
const (
	FacilityListTypeAirport  DWord = iota // SIMCONNECT_FACILITY_LIST_TYPE_AIRPORT
	FacilityListTypeWaypoint              // SIMCONNECT_FACILITY_LIST_TYPE_WAYPOINT
	FacilityListTypeNDB                   // SIMCONNECT_FACILITY_LIST_TYPE_NDB
	FacilityListTypeVOR                   // SIMCONNECT_FACILITY_LIST_TYPE_VOR
	FacilityListTypeCount                 // SIMCONNECT_FACILITY_LIST_TYPE_COUNT: invalid
)

// SIMCONNECT_VOR_FLAGS: flags for SIMCONNECT_RECV_ID_VOR_LIST
const (
	RecvIDVORListHasNAVSignal  DWord = 0x00000001 // SIMCONNECT_RECV_ID_VOR_LIST_HAS_NAV_SIGNAL: Has Nav signal
	RecvIDVORListHasLocalizer  DWord = 0x00000002 // SIMCONNECT_RECV_ID_VOR_LIST_HAS_LOCALIZER: Has localizer
	RecvIDVORListHasGlideScope DWord = 0x00000004 // SIMCONNECT_RECV_ID_VOR_LIST_HAS_GLIDE_SLOPE: Has Nav signal
	RecvIDVORListHasDME        DWord = 0x00000008 // SIMCONNECT_RECV_ID_VOR_LIST_HAS_DME: Station has DME
)

// SIMCONNECT_WAYPOINT_FLAGS: bits for the Waypoint Flags field: may be combined
const (
	WaypointNone                 DWord = 0x00       // SIMCONNECT_WAYPOINT_NONE
	WaypointSpeedRequested       DWord = 0x04       // SIMCONNECT_WAYPOINT_SPEED_REQUESTED: requested speed at waypoint is valid
	WaypointThrottleRequested    DWord = 0x08       // SIMCONNECT_WAYPOINT_THROTTLE_REQUESTED: request a specific throttle percentage
	WaypointComputeVerticalSpeed DWord = 0x10       // SIMCONNECT_WAYPOINT_COMPUTE_VERTICAL_SPEED: compute vertical to speed to reach waypoint altitude when crossing the waypoint
	WaypointAltitudeIsAGL        DWord = 0x20       // SIMCONNECT_WAYPOINT_ALTITUDE_IS_AGL: AltitudeIsAGL
	WaypointOnGround             DWord = 0x00100000 // SIMCONNECT_WAYPOINT_ON_GROUND: place this waypoint on the ground
	WaypointReverse              DWord = 0x00200000 // SIMCONNECT_WAYPOINT_REVERSE: Back up to this waypoint. Only valid on first waypoint
	WaypointWrapFirst            DWord = 0x00400000 // SIMCONNECT_WAYPOINT_WRAP_TO_FIRST: Wrap around back to first waypoint. Only valid on last waypoint
)

// SIMCONNECT_EVENT_FLAG
const (
	EventFlagDefault           DWord = 0x00000000 // SIMCONNECT_EVENT_FLAG_DEFAULT
	EventFlagFastRepeatTimer   DWord = 0x00000001 // SIMCONNECT_EVENT_FLAG_FAST_REPEAT_TIMER: set event repeat timer to simulate fast repeat
	EventFlagSlowRepeatTimer   DWord = 0x00000002 // DWORD SIMCONNECT_EVENT_FLAG_SLOW_REPEAT_TIMER: set event repeat timer to simulate slow repeat
	EventFlagGroupIDIsPriority DWord = 0x00000010 // SIMCONNECT_EVENT_FLAG_GROUPID_IS_PRIORITY: interpret GroupID parameter as priority value
)

// SIMCONNECT_DATA_REQUEST_FLAG
const (
	DataRequestFlagDefault DWord = 0x00000000 // SIMCONNECT_DATA_REQUEST_FLAG_DEFAULT
	DataRequestFlagChanged DWord = 0x00000001 // SIMCONNECT_DATA_REQUEST_FLAG_CHANGED: send requested data when value(s) change
	DataRequestFlagTagged  DWord = 0x00000002 // SIMCONNECT_DATA_REQUEST_FLAG_TAGGED: send requested data in tagged format
)

// SIMCONNECT_DATA_SET_FLAG
const (
	DataSetFlagDefault DWord = 0x00000000 // SIMCONNECT_DATA_SET_FLAG_DEFAULT
	DataSetFlagTagged  DWord = 0x00000001 // SIMCONNECT_DATA_SET_FLAG_TAGGED: data is in tagged format
)

// SIMCONNECT_CREATE_CLIENT_DATA_FLAG
const (
	CreateClientDataFlagDefault  DWord = 0x00000000 // SIMCONNECT_CREATE_CLIENT_DATA_FLAG_DEFAULT
	CreateClientDataFlagReadOnly DWord = 0x00000001 // SIMCONNECT_CREATE_CLIENT_DATA_FLAG_READ_ONLY: permit only ClientData creator to write into ClientData
)

// SIMCONNECT_CLIENT_DATA_REQUEST_FLAG
const (
	ClientDataRequestFlagDefault DWord = 0x00000000 // SIMCONNECT_CLIENT_DATA_REQUEST_FLAG_DEFAULT
	ClientDataRequestFlagChanged DWord = 0x00000001 // SIMCONNECT_CLIENT_DATA_REQUEST_FLAG_CHANGED: send requested ClientData when value(s) change
	ClientDataRequestFlagTagged  DWord = 0x00000002 // SIMCONNECT_CLIENT_DATA_REQUEST_FLAG_TAGGED: send requested ClientData in tagged format
)

// SIMCONNECT_CLIENT_DATA_SET_FLAG
const (
	ClientDataSetFlagDefault DWord = 0x00000000 // SIMCONNECT_CLIENT_DATA_SET_FLAG_DEFAULT
	ClientDataSetFlagTagged  DWord = 0x00000001 // SIMCONNECT_CLIENT_DATA_SET_FLAG_TAGGED: data is in tagged format
)

// SIMCONNECT_VIEW_SYSTEM_EVENT_DATA: dwData contains these flags for the "View" System Event
const (
	ViewSystemEventDataCockpit2D      DWord = 0x00000001 // SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_2D: 2D Panels in cockpit view
	ViewSystemEventDataCockpitVirtual DWord = 0x00000002 // SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_VIRTUAL: Virtual (3D) panels in cockpit view
	ViewSystemEventDataOrthogonal     DWord = 0x00000004 // SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_ORTHOGONAL: Orthogonal (Map) view
)

// SIMCONNECT_SOUND_SYSTEM_EVENT_DATA: dwData contains these flags for the "Sound" System Event
const (
	SoundSystemEventDataMaster DWord = 0x00000001 // SIMCONNECT_SOUND_SYSTEM_EVENT_DATA_MASTER: Sound Master
)

// SIMCONNECT_USER_ENUM SIMCONNECT_NOTIFICATION_GROUP_ID;     //client-defined notification group ID
// SIMCONNECT_USER_ENUM SIMCONNECT_INPUT_GROUP_ID;            //client-defined input group ID
// SIMCONNECT_USER_ENUM SIMCONNECT_DATA_DEFINITION_ID;        //client-defined data definition ID
// SIMCONNECT_USER_ENUM SIMCONNECT_DATA_REQUEST_ID;           //client-defined request data ID

// SIMCONNECT_USER_ENUM SIMCONNECT_CLIENT_EVENT_ID;           //client-defined client event ID
// SIMCONNECT_USER_ENUM SIMCONNECT_CLIENT_DATA_ID;            //client-defined client data ID
// SIMCONNECT_USER_ENUM SIMCONNECT_CLIENT_DATA_DEFINITION_ID; //client-defined client data definition ID

// SIMCONNECT_RECV
// Used with the SIMCONNECT_RECV_ID enumeration to indicate which type of structure has been returned.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV.htm
type Recv struct {
	Size    DWord // record size
	Version DWord // interface version
	ID      DWord // see SIMCONNECT_RECV_ID
}

// SIMCONNECT_RECV_EXCEPTION: when dwID == SIMCONNECT_RECV_ID_EXCEPTION
// Used with the SIMCONNECT_EXCEPTION enumeration type to return information on an error that has occurred.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_EXCEPTION.htm
type RecvException struct {
	Recv
	Exception DWord // SIMCONNECT_EXCEPTION
	SendID    DWord // SimConnect_GetLastSentPacketID
	Index     DWord // index of parameter that was source of error
}

// SIMCONNECT_RECV_OPEN: when dwID == SIMCONNECT_RECV_ID_OPEN
// Used to return information to the client, after a successful call to SimConnect_Open.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_OPEN.htm
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

// SIMCONNECT_RECV_QUIT: when dwID == SIMCONNECT_RECV_ID_QUIT
// This is an identical structure to the SIMCONNECT_RECV structure.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_QUIT.htm
type RecvQuit struct {
	Recv
}

// SIMCONNECT_RECV_EVENT: when dwID == SIMCONNECT_RECV_ID_EVENT
// Used to return an event ID to the client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_EVENT.htm
type RecvEvent struct {
	Recv
	GroupID DWord
	EventID DWord
	Data    DWord // uEventID-dependent context
}

// SIMCONNECT_RECV_EVENT_FILENAME: when dwID == SIMCONNECT_RECV_ID_EVENT_FILENAME
// Used with the SimConnect_SubscribeToSystemEvent to return a filename and an event ID to the client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_FILENAME.htm
type RecvEventFilename struct {
	RecvEvent
	FileName [260]byte // uEventID-dependent context
	Flags    DWord
}

// SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE: when dwID == SIMCONNECT_RECV_ID_EVENT_OBJECT_ADDREMOVE
type RecvEventObjectAddRemove struct {
	RecvEvent
	ObjType DWord // SIMCONNECT_SIMOBJECT_TYPE
}

// SIMCONNECT_RECV_EVENT_FRAME: when dwID == SIMCONNECT_RECV_ID_EVENT_FRAME
// Used with the SimConnect_SubscribeToSystemEvent to return the frame rate and simulation speed to the client.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_EVENT_FRAME.htm
type RecvEventFrame struct {
	RecvEvent
	FrameRate float32
	SimSpeed  float32
}

// SIMCONNECT_RECV_EVENT_MULTIPLAYER_SERVER_STARTED: when dwID == SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_SERVER_STARTED
// type RecvEventMultiplayerServerStarted struct {
// 	RecvEvent
// 	// No event specific data, for now
// }

// SIMCONNECT_RECV_EVENT_MULTIPLAYER_CLIENT_STARTED: when dwID == SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_CLIENT_STARTED
// type RecvEventMultiplayerClientStarted struct {
// 	RecvEvent
// 	// No event specific data, for now
// }

// SIMCONNECT_RECV_EVENT_MULTIPLAYER_SESSION_ENDED: when dwID == SIMCONNECT_RECV_ID_EVENT_MULTIPLAYER_SESSION_ENDED
// type RecvEventMultiplayerSessionEnded struct {
// 	RecvEvent
// 	// No event specific data, for now
// }

// GUID structure from guiddef.h
// typedef struct _GUID {
// 	unsigned long  Data1;
// 	unsigned short Data2;
// 	unsigned short Data3;
// 	unsigned char  Data4[8];
// } GUID;

// SIMCONNECT_DATA_RACE_RESULT
// type DataRaceResult struct {
// 	NumberOfRacers DWord        // The total number of racers
// 	MissionGUID    windows.GUID // The name of the mission to execute, NULL if no mission
// 	PlayerName     [260]byte    // The name of the player
// 	SessionType    [260]byte    // The type of the multiplayer session: "LAN", "GAMESPY")
// 	Aircraft       [260]byte    // The aircraft type
// 	PlayerRole     [260]byte    // The player role in the mission
// 	TotalTime      float64      // Total time in seconds, 0 means DNF
// 	PenaltyTime    float64      // Total penalty time in seconds
// 	IsDisqualified DWord        // non 0 - disqualified, 0 - not disqualified
// }

// SIMCONNECT_RECV_EVENT_RACE_END: when dwID == SIMCONNECT_RECV_ID_EVENT_RACE_END
// type RecvEventRaceEnd struct {
// 	RecvEvent
// 	RacerNumber DWord // The index of the racer the results are for
// 	RacerData   DataRaceResult
// }

// SIMCONNECT_RECV_EVENT_RACE_LAP: when dwID == SIMCONNECT_RECV_ID_EVENT_RACE_LAP
// type RecvEventRaceLap struct {
// 	LapIndex  DWord // The index of the lap the results are for
// 	RacerData DataRaceResult
// }

// SIMCONNECT_RECV_SIMOBJECT_DATA: when dwID == SIMCONNECT_RECV_ID_SIMOBJECT_DATA
// Will be received by the client after a successful call to SimConnect_RequestDataOnSimObject or SimConnect_RequestDataOnSimObjectType.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_SIMOBJECT_DATA.htm
type RecvSimObjectData struct {
	Recv
	RequestID   DWord
	ObjectID    DWord
	DefineID    DWord
	Flags       DWord // SIMCONNECT_DATA_REQUEST_FLAG
	EntryNumber DWord // if multiple objects returned, this is number <entrynumber> out of <outof>
	OutOf       DWord // note: starts with 1, not 0
	DefineCount DWord // data count (number of datums, *not* byte count)
	// SIMCONNECT_DATAV(dwData, dwDefineID); // data begins here, dwDefineCount data items
}

// SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE: when dwID == SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE
// Will be received by the client after a successful call to SimConnect_RequestDataOnSimObjectType. The structure is identical to SIMCONNECT_RECV_SIMOBJECT_DATA.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE.htm
type RecvSimObjectDataByType struct {
	RecvSimObjectData
}

// SIMCONNECT_RECV_CLIENT_DATA: when dwID == SIMCONNECT_RECV_ID_CLIENT_DATA
// Will be received by the client after a successful call to SimConnect_RequestClientData. The structure is identical to SIMCONNECT_RECV_SIMOBJECT_DATA.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_CLIENT_DATA.htm
type RecvClientData struct {
	RecvSimObjectData
}

// SIMCONNECT_RECV_WEATHER_OBSERVATION: when dwID == SIMCONNECT_RECV_ID_WEATHER_OBSERVATION
// type RecvWeatherObservation struct {
// 	Recv
// 	requestID DWord
// 	metar     [1]byte // SIMCONNECT_STRINGV(szMetar): Variable length string whose maximum size is MAX_METAR_LENGTH
// }

// const (
// 	CloudStateArrayWidth int = 64                                          // SIMCONNECT_CLOUD_STATE_ARRAY_WIDTH
// 	CloudStateArraySize  int = CloudStateArrayWidth * CloudStateArrayWidth // SIMCONNECT_CLOUD_STATE_ARRAY_SIZE
// )

// SIMCONNECT_RECV_CLOUD_STATE: when dwID == SIMCONNECT_RECV_ID_CLOUD_STATE
// type RecvCloudState struct {
// 	Recv
// 	RequestID DWord
// 	ArraySize DWord
// 	// SIMCONNECT_FIXEDTYPE_DATAV(BYTE, rgbData, dwArraySize, U1 /*member of UnmanagedType enum*/ , System::Byte /*cli type*/);
// }

// SIMCONNECT_RECV_ASSIGNED_OBJECT_ID: when dwID == SIMCONNECT_RECV_ID_ASSIGNED_OBJECT_ID
// Used to return an object ID that matches a request ID.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_ASSIGNED_OBJECT_ID.htm
type RecvAssignedObjectID struct {
	Recv
	RequestID DWord
	ObjectID  DWord
}

// SIMCONNECT_RECV_RESERVED_KEY: when dwID == SIMCONNECT_RECV_ID_RESERVED_KEY
// Used with the SimConnect_RequestReservedKey function to return the reserved key combination.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_RESERVED_KEY.htm
type RecvReservedKey struct {
	Recv
	ChoiceReserved [30]byte
	ReservedKey    [50]byte
}

// SIMCONNECT_RECV_SYSTEM_STATE : public SIMCONNECT_RECV // when dwID == SIMCONNECT_RECV_ID_SYSTEM_STATE
// Used with the SimConnect_RequestSystemState function to retrieve specific Flight Simulator systems states and information.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_SYSTEM_STATE.htm
type RecvSystemState struct {
	Recv
	RequestID DWord
	Integer   DWord
	Float     float32
	String    [260]byte
}

// SIMCONNECT_RECV_CUSTOM_ACTION : public SIMCONNECT_RECV_EVENT
// type RecvCustomAction struct {
// 	RecvEvent
// 	InstanceID        windows.GUID // Instance id of the action that executed
// 	WaitForCompletion DWord        // Wait for completion flag on the action
// 	Payload           [1]byte      // SIMCONNECT_STRINGV(szPayLoad): Variable length string payload associated with the mission action
// }

// SIMCONNECT_RECV_EVENT_WEATHER_MODE : public SIMCONNECT_RECV_EVENT
// type RecvEventWeatherMode struct {
// 	// No event specific data - the new weather mode is in the base structure dwData member
// }

// SIMCONNECT_RECV_FACILITIES_LIST
// Used to provide information on the number of elements in a list of facilities returned to the client, and the number of packets that were used to transmit the data.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_FACILITIES_LIST.htm
type RecvFacilitiesList struct {
	Recv
	RequestID   DWord
	ArraySize   DWord
	EntryNumber DWord // when the array of items is too big for one send, which send this is (0..dwOutOf-1)
	OutOf       DWord // total number of transmissions the list is chopped into
}

// SIMCONNECT_DATA_FACILITY_AIRPORT
// Used to return information on a single airport in the facilities cache.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_AIRPORT.htm
type DataFacilityAirport struct {
	Icao      [9]byte // ICAO of the object
	Latitude  float64 // degrees
	Longitude float64 // degrees
	Altitude  float64 // meters
}

// SIMCONNECT_RECV_AIRPORT_LIST
// Used to return a list of SIMCONNECT_DATA_FACILITY_AIRPORT structures.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_AIRPORT_LIST.htm
type RecvAirportList struct {
	RecvFacilitiesList
	// SIMCONNECT_FIXEDTYPE_DATAV(SIMCONNECT_DATA_FACILITY_AIRPORT, rgData, dwArraySize, U1 /*member of UnmanagedType enum*/, SIMCONNECT_DATA_FACILITY_AIRPORT /*cli type*/)
}

// SIMCONNECT_DATA_FACILITY_WAYPOINT
// Used to return information on a single waypoint in the facilities cache.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_WAYPOINT.htm
type DataFacilityWaypoint struct {
	DataFacilityAirport
	MagVar float32 // Magvar in degrees
}

// SIMCONNECT_RECV_WAYPOINT_LIST
// Used to return a list of SIMCONNECT_DATA_FACILITY_WAYPOINT structures.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_WAYPOINT_LIST.htm
type RecvWaypointList struct {
	RecvFacilitiesList
	// SIMCONNECT_FIXEDTYPE_DATAV(SIMCONNECT_DATA_FACILITY_WAYPOINT, rgData, dwArraySize, U1 /*member of UnmanagedType enum*/, SIMCONNECT_DATA_FACILITY_WAYPOINT /*cli type*/)
}

// SIMCONNECT_DATA_FACILITY_NDB
// Used to return information on a single NDB station in the facilities cache.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_NDB.htm
type DataFacilityNDB struct {
	DataFacilityWaypoint
	Frequency DWord // frequency in Hz
}

// SIMCONNECT_RECV_NDB_LIST
// Used to return a list of SIMCONNECT_DATA_FACILITY_NDB structures.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_NDB_LIST.htm
type RecvNDBList struct {
	RecvFacilitiesList
	// SIMCONNECT_FIXEDTYPE_DATAV(SIMCONNECT_DATA_FACILITY_NDB, rgData, dwArraySize, U1 /*member of UnmanagedType enum*/, SIMCONNECT_DATA_FACILITY_NDB /*cli type*/)
}

// SIMCONNECT_DATA_FACILITY_VOR
// Used to return information on a single VOR station in the facilities cache.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_DATA_FACILITY_VOR.htm
type DataFacilityVOR struct {
	DataFacilityNDB
	Flags           DWord   // SIMCONNECT_VOR_FLAGS
	Localizer       float32 // Localizer in degrees
	GlideLat        float64 // Glide Slope Location (deg, deg, meters)
	GlideLon        float64
	GlideAlt        float64
	GlideSlopeAngle float32 // Glide Slope in degrees
}

// SIMCONNECT_RECV_VOR_LIST
// Used to return a list of SIMCONNECT_DATA_FACILITY_VOR structures.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_RECV_VOR_LIST.htm
type RecvVORList struct {
	RecvFacilitiesList
	// SIMCONNECT_FIXEDTYPE_DATAV(SIMCONNECT_DATA_FACILITY_VOR, rgData, dwArraySize, U1 /*member of UnmanagedType enum*/, SIMCONNECT_DATA_FACILITY_VOR /*cli type*/)
}

// SIMCONNECT_DATATYPE_INITPOSITION
type InitPosition struct {
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
// type MarkerState struct {
// 	MarkerName  [64]byte
// 	MarkerState DWord
// }

// SIMCONNECT_DATATYPE_WAYPOINT
// type Waypoint struct {
// 	Latitude        float64 // degrees
// 	Longitude       float64 // degrees
// 	Altitude        float64 // feet
// 	Flags           uint32
// 	ktsSpeed        float64 // knots
// 	percentThrottle float64
// }

// SIMCONNECT_DATA_LATLONALT
// Used to hold a world position.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_DATA_LATLONALT.htm
type LatLogAlt struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

// SIMCONNECT_DATA_XYZ
// Used to hold a 3D co-ordinate.
// https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/API_Reference/Structures_And_Enumerations/SIMCONNECT_DATA_XYZ.htm
type XYZ struct {
	X float64
	Y float64
	Z float64
}
