package simconnect

type SimVar struct {
	DefineID    DWord
	RequestID   DWord
	Name        string
	Unit        string
	DataType    DWord
	Value       interface{}
	UpdateCount int64
	IsString    bool
	Registered  bool
	Pending     bool
	Timestamp   int64
}

func NewSimVar(defineID DWord, name string, unit string, dataType DWord) *SimVar {
	return &SimVar{
		DefineID:  defineID,
		RequestID: 0,
		Name:      name,
		Unit:      unit,
		DataType:  dataType,
		IsString:  IsStringDataType(dataType),
	}
}

func (simVar *SimVar) ToInt32(defaultValue int32) int32 {
	if simVar.Value != nil {
		return ValueToInt32(simVar.Value)
	}
	return defaultValue
}

func (simVar *SimVar) ToInt64(defaultValue int64) int64 {
	if simVar.Value != nil {
		return ValueToInt64(simVar.Value)
	}
	return defaultValue
}

func (simVar *SimVar) ToFloat32(defaultValue float32) float32 {
	if simVar.Value != nil {
		return ValueToFloat32(simVar.Value)
	}
	return defaultValue
}

func (simVar *SimVar) ToFloat64(defaultValue float64) float64 {
	if simVar.Value != nil {
		return ValueToFloat64(simVar.Value)
	}
	return defaultValue
}

func (simVar *SimVar) ToString(defaultValue string) string {
	if simVar.Value != nil {
		return ValueToString(simVar.Value)
	}
	return defaultValue
}
