package simconnect

import (
	"os"
	"time"

	"github.com/grumpypixel/msfs2020-simconnect-go/filepacker"
)

var (
	dataTypeMapper map[string]DWord
)

func init() {
	dataTypeMapper = stringToDataTypeMapping()
}

func UnpackDLL(path string) error {
	data := PackedSimConnectDLL()
	unpacked, err := filepacker.Unpack(string(data))
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if _, err := file.WriteString(string(unpacked)); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	time.Sleep(time.Second * 1)
	return nil
}

func StringToDataType(dataType string) DWord {
	if value, exists := dataTypeMapper[dataType]; exists {
		return value
	}
	return DataTypeInvalid
}

func DataTypeToString(dataType DWord) string {
	for name, value := range dataTypeMapper {
		if value == dataType {
			return name
		}
	}
	return "invalid"
}

func IsStringDataType(dataType DWord) bool {
	switch dataType {
	case DataTypeString8, DataTypeString32, DataTypeString64, DataTypeString128,
		DataTypeString256, DataTypeString260, DataTypeStringV:
		return true
	}
	return false
}

func ValueToInt32(value interface{}) int32 {
	return value.(int32)
}

func ValueToInt64(value interface{}) int64 {
	return value.(int64)
}

func ValueToFloat32(value interface{}) float32 {
	return value.(float32)
}

func ValueToFloat64(value interface{}) float64 {
	return value.(float64)
}

func ValueToString(value interface{}) string {
	return value.(string)
}

func stringToDataTypeMapping() map[string]DWord {
	return map[string]DWord{
		"invalid":      DataTypeInvalid,
		"int32":        DataTypeInt32,
		"int64":        DataTypeInt64,
		"float32":      DataTypeFloat32,
		"float64":      DataTypeFloat64,
		"string8":      DataTypeString8,
		"string32":     DataTypeString32,
		"string64":     DataTypeString64,
		"string128":    DataTypeString128,
		"string256":    DataTypeString256,
		"string260":    DataTypeString260,
		"stringv":      DataTypeStringV,
		"initposition": DataTypeInitPosition,
		"markerstate":  DataTypeMarkerState,
		"waypoint":     DataTypeWaypoint,
		"latlongalt":   DataTypeLatLonAlt,
		"xyz":          DataTypeXYZ,
	}
}
