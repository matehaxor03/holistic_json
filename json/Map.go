package json

import (
	"fmt"
	"strings"
	"time"
	common "github.com/matehaxor03/holistic_common/common"
)

type Map map[string](Value)

func (m Map) GetMap(s string) (*Map, []error) {
	if m.IsNil(s) {
		return nil, nil
	}

	return m[s].GetMap()
}

func (m Map) SetMap(s string, zap *Map) {
	m[s] = Value{"value":zap}
}

func (m Map) SetMapValue(s string, zap Map) {
	m[s] = Value{"value":zap}
}

func (m Map) IsNil(s string) bool {
	if common.IsNil(m[s]) {
		return true
	}

	return m[s].IsNil()
}

func (m Map) IsBool(s string) bool {
	if common.IsNil(s) {
		return false
	}
	return m[s].IsBool()
}

func (m Map) IsArray(s string) bool {
	if common.IsNil(s) {
		return false
	}
	return m[s].IsArray()
}

func (m Map) IsEmptyString(s string) bool {
	if m.IsNil(s) {
		return false
	}
	return m[s].IsEmptyString()
}

func (m Map) IsInteger(s string) bool {
	if common.IsNil(m[s]) {
		return false
	}
	return m[s].IsInteger()
}

func (m Map) IsString(s string) bool {
	if m.IsNil(s) {
		return false
	}

	return m[s].IsString()
}

func (m Map) IsMap(s string) bool {
	if m.IsNil(s) {
		return false
	}

	return m[s].IsMap()
}

func (m Map) IsBoolTrue(s string) bool {
	if m.IsNil(s) {
		return false
	}
	return m[s].IsBoolTrue()
}

func (m Map) IsBoolFalse(s string) bool {
	if m.IsNil(s) {
		return false
	}
	return m[s].IsBoolFalse()
}

func (m Map) ToJSONString(json *strings.Builder) ([]error) {
	var errors []error
	if json == nil {
		errors = append(errors, fmt.Errorf("error: *strings.Builder is nil"))
		return errors
	}
	
	length := len(m)
	
	if length == 0 {
		json.WriteString("{}")
		return nil
	}

	keys := m.Keys()
	

	json.WriteString("{")
	for i := 0; i < length; i++ {
		key := keys[i]
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll(key, "\"", "\\\""))
		json.WriteString("\":")
		string_conversion_errors := ConvertInterfaceValueToJSONStringValue(json, m[key])
		if string_conversion_errors != nil {
			errors = append(errors, string_conversion_errors...)
		}

		if i < length - 1 {
			json.WriteString(",")
		}
	}
	json.WriteString("}")

	if len(errors) > 0 {
		return errors
	}
	return nil
}

func (m Map) SetArray(s string, array *Array) {
	m[s] = Value{"value":array}
}

func (m Map) SetErrors(s string, errors []error) {
	m[s] = Value{"value":errors}
}


func (m Map) GetErrors(s string) ([]error, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetErrors()
}

func (m Map) GetType(s string) string {
	if common.IsNil(m[s]) {
		return "nil"
	}
	
	return m[s].GetType()
}

func (m Map) Func(s string) (func(Map) []error, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetFunc()
}

func (m Map) SetFunc(s string, function func(Map) []error) {
	m[s] = Value{"value":function}
}

func (m Map) GetArray(s string) (*Array, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArray()
}

func (m Map) GetArrayOfInt8(s string) (*[]*int8, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfInt8()
}

func (m Map) GetArrayOfInt16(s string) (*[]*int16, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfInt16()
}

func (m Map) GetArrayOfInt32(s string) (*[]*int32, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfInt32()
}

func (m Map) GetArrayOfInt64(s string) (*[]*int64, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfInt64()
}

func (m Map) GetArrayOfUInt8(s string) (*[]*uint8, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfUInt8()
}

func (m Map) GetArrayOfUInt16(s string) (*[]*uint16, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfUInt16()
}

func (m Map) GetArrayOfUInt32(s string) (*[]*uint32, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfUInt32()
}

func (m Map) GetArrayOfUInt64(s string) (*[]*uint64, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfUInt64()
}

func (m Map) GetArrayOfString(s string) (*[]*string, []error) {
	if common.IsNil(m[s]){
		return nil, nil
	}

	return m[s].GetArrayOfString()
}

func (m Map) GetString(s string) (*string, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetString()
}

func (m Map) IsFloat(s string) (bool) {
	if common.IsNil(m[s]) {
		return false
	}

	return m[s].IsFloat()
}

func (m Map) GetStringValue(s string) (string, []error) {
	string_value, string_value_errors := m.GetString(s)
	if string_value_errors != nil {
		return "", string_value_errors
	} else if string_value == nil {
		return "", nil
	}
	return *string_value, nil
}

func (m Map) GetFloat32(s string) (*float32, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetFloat32()
}

func (m Map) GetFloat32Value(s string) (float32, []error) {
	var errors []error
	if common.IsNil(m[s]) {
		errors = append(errors, fmt.Errorf("error: Map.GetFloat32Value returned a nil value"))
		return 0, errors
	}
	return m[s].GetFloat32Value()
}

func (m Map) GetFloat64Value(s string) (float64, []error) {
	var errors []error
	if common.IsNil(m[s]) {
		errors = append(errors, fmt.Errorf("error: Map.GetFloat64Value returned a nil value"))
		return 0, errors
	}
	return m[s].GetFloat64Value()
}

func (m Map) GetFloat64(s string) (*float64, []error) {
	if m.IsNil(s) {
		return nil, nil
	}

	return m[s].GetFloat64()
}

func (m Map) GetRunes(s string) (*[]rune, []error) {
	if m[s] == nil {
		return nil, nil
	}

	return m[s].GetRunes()
}

func (m Map) GetObject(s string) interface{} {
	if !m.HasKey(s) {
		return nil
	}

	if m.IsNil(s) {
		return nil
	}
	return m[s]
}

func (m Map) SetObject(s string, object interface{}) {
	m[s] = Value{"value":object}
}

func (m Map) GetBool(s string) (*bool, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}

	return m[s].GetBool()
}

func (m Map) SetBool(s string, value *bool) {
	m[s] = Value{"value":value}
}

func (m Map) SetString(s string, value *string) {
	m[s] = Value{"value":value}
}

func (m Map) SetStringValue(s string, value string) {
	m[s] = Value{"value":value}
}

func (m Map) SetNil(s string) {
	m[s] = Value{"value":nil}
}

func (m Map) Keys() []string {
	keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
	return keys
}

func (m Map) HasKey(key string) bool {
	if _, found := m[key]; found {
		return true
	} else {
		return false
	}
}

func (m Map) RemoveKey(key string) (*bool, []error) {
	var errors []error
	var result bool
	if !m.HasKey(key) {
		result = false
		errors = append(errors, fmt.Errorf("error: key %s not found", key))
		return &result, errors
	}	

	result = true
	delete(m, key)
	return &result, nil 
}

func (m Map) GetInt64(s string) (*int64, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}

	return m[s].GetInt64()
}

func (m Map) GetInt8(s string) (*int8, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetInt8()
}

func (m Map) GetInt8Value(s string) (int8, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetInt8Value was nil"))
		return 0, errors
	}
	return m[s].GetInt8Value()
}

func (m Map) GetUInt8(s string) (*uint8, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetUInt8()
}

func (m Map) GetUInt8Value(s string) (uint8, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetUInt8Value was nil"))
		return 0, errors
	}
	return m[s].GetUInt8Value()
}

func (m Map) GetInt16(s string) (*int16, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetInt16()
}

func (m Map) GetInt16Value(s string) (int16, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetInt16Value was nil"))
		return 0, errors
	}
	return m[s].GetInt16Value()
}

func (m Map) GetUInt16(s string) (*uint16, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetUInt16()
}

func (m Map) GetUInt16Value(s string) (uint16, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetUInt16Value was nil"))
		return 0, errors
	}
	return m[s].GetUInt16Value()
}


func (m Map) GetInt32(s string) (*int32, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetInt32()
}

func (m Map) GetInt32Value(s string) (int32, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetInt32Value was nil"))
		return 0, errors
	}
	return m[s].GetInt32Value()
}

func (m Map) GetInt64Value(s string) (int64, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetInt64Value was nil"))
		return 0, errors
	}
	return m[s].GetInt64Value()
}

func (m Map) GetUInt32(s string) (*uint32, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetUInt32()
}

func (m Map) GetUInt32Value(s string) (uint32, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetUInt32Value was nil"))
		return 0, errors
	}
	return m[s].GetUInt32Value()
}

func (m Map) GetUInt64Value(s string) (uint64, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetUInt64Value was nil"))
		return 0, errors
	}
	return m[s].GetUInt64Value()
}

func (m Map) GetInt(s string) (*int, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}
	return m[s].GetInt()
}

func (m Map) GetIntValue(s string) (int, []error) {
	if common.IsNil(m[s]) {
		var errors []error
		errors = append(errors, fmt.Errorf("Map.GetIntValue was nil"))
		return 0, errors
	}
	return m[s].GetIntValue()
}

func (m Map) SetInt(s string, v *int) {
	m[s] = Value{"value":v}
}

func (m Map) SetInt64(s string, v *int64) {
	m[s] = Value{"value":v}
}

func (m Map) SetInt32(s string, v *int32) {
	m[s] = Value{"value":v}
}

func (m Map) SetInt16(s string, v *int16) {
	m[s] = Value{"value":v}
}

func (m Map) SetInt8(s string, v *int8) {
	m[s] = Value{"value":v}
}

func (m Map) SetIntValue(s string, v int) {
	m[s] = Value{"value":v}
}

func (m Map) SetInt64Value(s string, v int64) {
	m[s] = Value{"value":v}
}

func (m Map) SetInt32Value(s string, v int32) {
	m[s] = Value{"value":v}
}

func (m Map) SetInt16Value(s string, v int16) {
	m[s] = Value{"value":v}
}

func (m Map) SetInt8Value(s string, v int8) {
	m[s] = Value{"value":v}
}

///

func (m Map) SetUInt(s string, v *uint) {
	m[s] = Value{"value":v}
}

func (m Map) SetUInt64(s string, v *uint64) {
	m[s] = Value{"value":v}
}

func (m Map) SetUInt32(s string, v *uint32) {
	m[s] = Value{"value":v}
}

func (m Map) SetUInt16(s string, v *uint16) {
	m[s] = Value{"value":v}
}

func (m Map) SetUInt8(s string, v *uint8) {
	m[s] = Value{"value":v}
}

func (m Map) SetUIntValue(s string, v uint) {
	m[s] = Value{"value":v}
}

func (m Map) SetUInt64Value(s string, v uint64) {
	m[s] = Value{"value":v}
}

func (m Map) SetUInt32Value(s string, v uint32) {
	m[s] = Value{"value":v}
}

func (m Map) SetUInt16Value(s string, v uint16) {
	m[s] = Value{"value":v}
}

func (m Map) SetUInt8Value(s string, v uint8) {
	m[s] = Value{"value":v}
}

func (m Map) SetFloat64(s string, v *float64) {
	m[s] = Value{"value":v}
}

func (m Map) SetFloat64Value(s string, v float64) {
	m[s] = Value{"value":v}
}

func (m Map) SetFloat32(s string, v *float32) {
	m[s] = Value{"value":v}
}

func (m Map) SetFloat32Value(s string, v float32) {
	m[s] = Value{"value":v}
}

func (m Map) GetUInt64(s string) (*uint64, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}

	return m[s].GetUInt64()
}
	
func (m Map) SetTime(s string, value *time.Time) {
	m[s] = Value{"value":value}
}

func (m Map) GetTime(s string, decimal_places int) (*time.Time, []error) {
	if common.IsNil(m[s]) {
		return nil, nil
	}

	return m[s].GetTime(decimal_places)
}

func (m Map) Values() Array {
	array := Array{}
	for _, f := range m {
		array = append(array, f)
	}
	return array
}

func (m Map) Clone() (*Map, []error) {
	var json_payload_builder strings.Builder
	request_payload_as_string_errors := m.ToJSONString(&json_payload_builder)
	if request_payload_as_string_errors != nil {
		return nil, request_payload_as_string_errors
	}

	return Parse(json_payload_builder.String())
}
