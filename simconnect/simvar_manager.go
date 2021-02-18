package simconnect

import (
	"bytes"
	"fmt"
	"sync"
)

type SimVarManager struct {
	Vars    []*SimVar
	nameMap map[string]*SimVar
	idMap   map[DWord]*SimVar
	mutex   sync.Mutex
}

func NewSimVarManager() *SimVarManager {
	return &SimVarManager{
		Vars:    make([]*SimVar, 0, 16),
		nameMap: make(map[string]*SimVar),
		idMap:   make(map[DWord]*SimVar),
	}
}

func (mgr *SimVarManager) Count() int {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()
	return len(mgr.nameMap)
}

func (mgr *SimVarManager) SimVars() []*SimVar {
	return mgr.Vars
}

func (mgr *SimVarManager) Add(name string, unit string, dataType DWord) DWord {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	if simVar, exists := mgr.simVarWithName(name); exists {
		return simVar.DefineID
	}

	defineID := NewDefineID()
	simVar := NewSimVar(defineID, name, unit, dataType)
	mgr.Vars = append(mgr.Vars, simVar)
	mgr.nameMap[name] = simVar
	mgr.idMap[defineID] = simVar
	return defineID
}

func (mgr *SimVarManager) Remove(defineID DWord) bool {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()
	simVar, exists := mgr.simVarWithID(defineID)
	if !exists {
		return false
	}
	delete(mgr.nameMap, simVar.Name)
	delete(mgr.idMap, simVar.DefineID)
	vars := mgr.Vars
	for i, simVar := range vars {
		if simVar.DefineID == defineID {
			vars[i] = vars[len(vars)-1]
			mgr.Vars = vars[:len(vars)-1]
			return true
		}
	}
	return false
}

func (mgr *SimVarManager) GetSimVar(defineID DWord) (*SimVar, bool) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()
	return mgr.simVarWithID(defineID)
}

func (mgr *SimVarManager) Update(requestID, defineID DWord, value interface{}) (*SimVar, bool) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()
	if simVar, ok := mgr.simVarWithID(defineID); ok {
		if simVar.Pending && simVar.RequestID == requestID {
			if !simVar.IsString {
				simVar.Value = value
			} else {
				simVar.Value = mgr.ToString(simVar.DataType, value)
			}
			simVar.UpdateCount++
			return simVar, true
		}
	}
	return nil, false
}

func (mgr *SimVarManager) ToString(dataType DWord, value interface{}) string {
	switch dataType {
	case DataTypeString8:
		b := value.([8]byte)
		return string(bytes.Trim(b[:], "\x00"))

	case DataTypeString32:
		b := value.([32]byte)
		return string(bytes.Trim(b[:], "\x00"))

	case DataTypeString64:
		b := value.([64]byte)
		return string(bytes.Trim(b[:], "\x00"))

	case DataTypeString128:
		b := value.([128]byte)
		return string(bytes.Trim(b[:], "\x00"))

	case DataTypeString256:
		b := value.([256]byte)
		return string(bytes.Trim(b[:], "\x00"))

	case DataTypeString260:
		b := value.([260]byte)
		return string(bytes.Trim(b[:], "\x00"))
	}
	return fmt.Sprintf("%v", value)
}

func (mgr *SimVarManager) SimVarDump(indent string) []string {
	dump := make([]string, 0)
	for i, simVar := range mgr.Vars {
		str := fmt.Sprintf("%s%02d: name: %s unit: %s value: %f type: %s updates: %d reqId: %d defid: %d registered: %v pending: %v",
			indent, i+1, simVar.Name, simVar.Unit, simVar.Value, DataTypeToString(simVar.DataType), simVar.UpdateCount,
			simVar.RequestID, simVar.DefineID, simVar.Registered, simVar.Pending)
		dump = append(dump, str)
	}
	return dump
}

func (mgr *SimVarManager) existsWithName(name string) bool {
	_, exists := mgr.nameMap[name]
	return exists
}

func (mgr *SimVarManager) existsWithID(defineID DWord) bool {
	_, exists := mgr.idMap[defineID]
	return exists
}

func (mgr *SimVarManager) simVarWithName(name string) (*SimVar, bool) {
	simVar, exists := mgr.nameMap[name]
	return simVar, exists
}

func (mgr *SimVarManager) simVarWithID(defineID DWord) (*SimVar, bool) {
	simVar, exists := mgr.idMap[defineID]
	return simVar, exists
}

// func (mgr *SimVarManager) AnyPending() bool {
// 	mgr.mutex.Lock()
// 	defer mgr.mutex.Unlock()
// 	for _, simVar := range mgr.Vars {
// 		if simVar.Pending {
// 			return true
// 		}
// 	}
// 	return false
// }
