package json

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	common "github.com/matehaxor03/holistic_common/common"
)

type Map map[string]interface{}

func (m Map) GetMap(s string) (*Map, []error) {
	var errors []error
	if m.IsNil(s) {
		return nil, nil
	}

	var result *Map
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "json.Map":
		value := m[s].(Map)
		result = &value
	case "*json.Map":
		value := *(m[s].(*Map))
		result = &value
	default:
		errors = append(errors, fmt.Errorf("error: Map.M: type %s is not supported please implement for key %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) SetMap(s string, zap *Map) {
	m[s] = zap
}

func (m Map) SetMapValue(s string, zap Map) {
	m[s] = &zap
}

func (m Map) IsNil(s string) bool {
	if m[s] == nil {
		return true
	}

	string_value := fmt.Sprintf("%s", m[s])

	if string_value == "<nil>" {
		return true
	}

	rep := fmt.Sprintf("%T", m[s])

	if string_value == "%!s("+rep+"=<nil>)" {
		return true
	}

	return false
}

func (m Map) IsBool(s string) bool {
	type_of := m.GetType(s)
	if type_of == "bool" || type_of == "*bool" {
		return true
	}

	return false
}

func (m Map) IsArray(s string) bool {
	type_of := m.GetType(s)
	if type_of == "json.Array" || type_of == "*json.Array" {
		return true
	}

	return false
}

func (m Map) IsEmptyString(s string) bool {
	if m.IsNil(s) {
		return false
	}

	type_of := m.GetType(s)
	if type_of == "string" || type_of == "*string" {
		string_value, _ := m.GetString(s)
		return *string_value == ""
	}

	return false
}

func (m Map) IsNumber(s string) bool {
	type_of := m.GetType(s)
	switch type_of {
	case "*int", 
		  "int",
		  "*uint", 
		  "uint",
		  "*int64",
		  "int64",
		  "*uint64",
		  "uint64",
		  "*int32",
		  "int32",
		  "*uint32",
		  "uint32",
		  "*int16",
		  "int16",
		  "*uint16",
		  "uint16",
		  "*int8",
		  "int8",
		  "*uint8",
		  "uint8":
		return true
	default: 
		return false
	}
}

func (m Map) IsString(s string) bool {
	if m.IsNil(s) {
		return false
	}

	type_of := m.GetType(s)
	if type_of == "string" || type_of == "*string" {
		return true
	}

	return false
}

func (m Map) IsMap(s string) bool {
	if m.IsNil(s) {
		return false
	}

	type_of := m.GetType(s)
	if type_of == "*json.Map" || type_of == "json.Map" {
		return true
	}

	return false
}

func (m Map) IsBoolTrue(s string) bool {
	if m.IsNil(s) {
		return false
	}

	if !m.IsBool(s) {
		return false
	}

	value, _ := m.GetBool(s)
	return *value == true
}

func (m Map) IsBoolFalse(s string) bool {
	if m.IsNil(s) {
		return true
	}

	if !m.IsBool(s) {
		return true
	}

	value, _ := m.GetBool(s)
	return *value == false
}

func (m Map) ToJSONString(json *strings.Builder) ([]error) {
	var errors []error
	if json == nil {
		errors = append(errors, fmt.Errorf("error: *strings.Builder is nil"))
		return errors
	}
	/*
	var b strings.Builder
	for i := 0; i < len(keys); i++ {
	  b.WriteString(keys[i])
	}
	return b.String()*/
	length := len(m)
	
	if length == 0 {
		json.WriteString("{}")
		return nil
	}

	keys := m.Keys()
	

	json.WriteString("{\n")
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
		json.WriteString("\n")
	}
	json.WriteString("}")

	if len(errors) > 0 {
		return errors
	}
	return nil
}

func (m Map) SetArray(s string, array *Array) {
	m[s] = array
}

func (m Map) SetErrors(s string, errors *[]error) {
	if errors == nil {
		m[s] = nil
		return
	}

	m[s] = errors
}


func (m Map) GetErrors(s string) ([]error, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var errors []error
	var result []error
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*[]error":
		result = *(m[s].(*[]error))
	case "[]error":
		result = m[s].([]error)
	case "*[]string":
		string_array := m[s].(*[]string)
		for _, string_array_value := range *string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "[]string":
		string_array := m[s].([]string)
		for _, string_array_value := range string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "json.Array":
		string_array := m[s].(Array)
		for _, string_array_value := range string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "*json.Array":
		string_array := m[s].(*Array)
		for _, string_array_value := range *string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetErrors: type %s is not supported please implement for field: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) GetType(s string) string {
	if m.IsNil(s) {
		return "nil"
	}
	
	return fmt.Sprintf("%T", m[s])
}

func (m Map) Func(s string) func(Map) []error {
	if m[s] == nil {
		return nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "func(json.Map) []error":
		return m[s].(func(Map) []error)
	case "*func(json.Map) []error":
		value := m[s].(*func(Map) []error)
		return *value
	default:
		panic(fmt.Errorf("error: Map.Func: type %s is not supported please implement for field: %s", rep, s))
	}

	return nil
}

func (m Map) SetFunc(s string, function func(Map) []error) {
	m[s] = function
}

func (m Map) GetArray(s string) (*Array, []error) {
	if m[s] == nil || m.IsNil(s) {
		return nil, nil
	}

	var errors []error
	var result *Array

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*json.Array":
		result = m[s].(*Array)
	case "json.Array":
		temp := m[s].(Array)
		result = &temp
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetArray: type %s is not supported please implement for field: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) GetString(s string) (*string, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var errors []error
	var result *string
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "string":
		value := m[s].(string)
		newValue := strings.Clone(value)
		result = &newValue
	case "*string":
		if fmt.Sprintf("%s", m[s]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*((m[s]).(*string)))
			result = &s
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetString: *string value is null for attribute: %s", rep, s))
		}
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetString: type %s is not supported please implement for attribute: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) IsFloat(s string) (bool) {
	if m.IsNil(s) {
		return false
	}

	type_of := m.GetType(s)
	if type_of == "float32" || 
	   type_of == "*float32" || 
	   type_of == "float64" || 
	   type_of == "*float64" {
		return true
	}

	return false
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
	var errors []error
	float64_value, float64_value_errors := m.GetFloat64(s)
	if float64_value_errors != nil {
		return nil, float64_value_errors
	} else if float64_value == nil {
		return nil, nil
	}

	value, value_error := strconv.ParseFloat(fmt.Sprintf("%f", *float64_value), 32)
	if value_error != nil {
		errors = append(errors, fmt.Errorf("error: strconv.ParseFloat returned error for converting to float32 number may be out of range"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	float_32_value := float32(value)
	return &float_32_value, nil
}

func (m Map) GetFloat32Value(s string) (float32, []error) {
	var errors []error
	float64_value, float64_value_errors := m.GetFloat64(s)
	if float64_value_errors != nil {
		return 0, float64_value_errors
	} else if float64_value == nil {
		errors = append(errors, fmt.Errorf("error: Map.GetFloat64 returned a nil value"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	value, value_error := strconv.ParseFloat(fmt.Sprintf("%f", *float64_value), 32)
	if value_error != nil {
		errors = append(errors, fmt.Errorf("error: strconv.ParseFloat returned error for converting to float32 number may be out of range"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	float_32_value := float32(value)
	return float_32_value, nil
}

func (m Map) GetFloat64Value(s string) (float64, []error) {
	var errors []error
	float64_value, float64_value_errors := m.GetFloat64(s)
	if float64_value_errors != nil {
		return 0, float64_value_errors
	} else if float64_value == nil {
		errors = append(errors, fmt.Errorf("error: Map.GetFloat64 returned a nil value"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	return *float64_value, nil
}

func (m Map) GetFloat64(s string) (*float64, []error) {
	if m.IsNil(s) {
		return nil, nil
	}

	var errors []error
	var result *float64
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "float64":
		value := m[s].(float64)
		result = &value
	case "*float64":
		value := *(m[s].(*float64))
		result = &value
	case "*string":
		value, value_error := strconv.ParseFloat((*(m[s].(*string))),64)
		if value_error != nil {
			errors = append(errors, fmt.Errorf("error: Map.GetFloat64: cannot convert *string value to float64"))
		} else {
			result = &value
		}
	case "string":
		value, value_error := strconv.ParseFloat((m[s].(string)), 64)
		if value_error != nil {
			errors = append(errors, fmt.Errorf("error: Map.GetFloat64: cannot convert *string value to float64"))
		} else {
			result = &value
		}
	case "float32":
		value := float64(m[s].(float32))
		result = &value
	case "*float32":
		value := float64(*(m[s].(*float32)))
		result = &value
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetFloat64: type %s is not supported please implement for attribute: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) GetRunes(s string) (*[]rune, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var errors []error
	var result *string
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "string":
		value := m[s].(string)
		newValue := strings.Clone(value)
		result = &newValue
	case "*string":
		if fmt.Sprintf("%s", m[s]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*((m[s]).(*string)))
			result = &s
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetString: *string value is null for attribute: %s", rep, s))
		}
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetString: type %s is not supported please implement for attribute: %s", rep, s))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	var runes []rune
	for _, runeValue := range *result {
		runes = append(runes, runeValue)
	}

	return &runes, nil
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
	m[s] = object
}

func (m Map) GetBool(s string) (*bool, []error) {
	if m[s] == nil {
		return nil, nil
	}

	var result *bool
	var errors []error

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "bool":
		value := m[s].(bool)
		result = &value
		break
	case "*bool":
		if fmt.Sprintf("%s", m[s]) != "%!s(*bool=<nil>)" {
			value := *((m[s]).(*bool))
			result = &value
		} else {
			return nil, nil
		}
		break
	case "*string":
		if fmt.Sprintf("%s", m[s]) != "%!s(*string=<nil>)" {
			value := *((m[s]).(*string))
			if value == "1" {
				boolean_result := true
				result = &boolean_result
			} else if value == "0" {
				boolean_result := false
				result = &boolean_result
			} else {
				errors = append(errors, fmt.Errorf("error: Map.GetBool: unknown value for *string: %s", value))
				result = nil
			}
		} else {
			return nil, nil
		}
		break
	case "string":
		value := ((m[s]).(string))
		if value == "1" {
			boolean_result := true
			result = &boolean_result
		} else if value == "0" {
			boolean_result := false
			result = &boolean_result
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetBool: unknown value for string: %s", value))
			result = nil
		}
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetBool: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (m Map) SetBool(s string, value *bool) {
	m[s] = value
}

func (m Map) SetString(s string, value *string) {
	if value == nil {
		m[s] = nil
		return 
	}

	clone_string := cloneString(value)
	m[s] = clone_string
}

func (m Map) SetStringValue(s string, value string) {
	clone_string := cloneString(&value)
	m[s] = clone_string
}

func (m Map) SetNil(s string) {
	m[s] = nil
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
	var errors []error
	var temp_value int64

	if m[s] == nil || m.IsNil(s) {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", m[s])
	switch rep {
		case "*int64":
			x := m[s].(*int64)
			temp_value = int64(*x)
		case "int64":
			x := m[s].(int64)
			temp_value = x
		case "*int32":
			x := m[s].(*int32)
			temp_value = int64(*x)
		case "int32":
			x := m[s].(int32)
			temp_value = int64(x)
		case "*int16":
			x := m[s].(*int16)
			temp_value = int64(*x)
		case "int16":
			x := m[s].(int16)
			temp_value = int64(x)
		case "*int8":
			x := m[s].(*int8)
			temp_value = int64(*x)
		case "int8":
			x := m[s].(int8)
			temp_value = int64(x)
		case "int":
			x := m[s].(int)
			temp_value = int64(x)
		case "*int":
			x := m[s].(*int)
			temp_value = int64(*x)
		case "*uint64":
			x := m[s].(*uint64)
			if *x > 9223372036854775807 {
				errors = append(errors, fmt.Errorf("%d is greater than 9223372036854775807", *x))
			} else {
				temp_value = int64(*x)
			}
		case "uint64":
			x := m[s].(uint64)
			if x > 9223372036854775807 {
				errors = append(errors, fmt.Errorf("%d is greater than 9223372036854775807", x))
			} else {
				temp_value = int64(x)
			}
		case "*uint32":
			x := m[s].(*uint32)
			temp_value = int64(*x)
		case "uint32":
			x := m[s].(uint32)
			temp_value = int64(x)
		case "*uint16":
			x := m[s].(*uint16)
			temp_value = int64(*x)
		case "uint16":
			x := m[s].(uint16)
			temp_value = int64(x)
		case "*uint8":
			x := m[s].(*uint8)
			temp_value = int64(*x)
		case "uint8":
			x := m[s].(uint8)
			temp_value = int64(x)
		case "*uint":
			x := m[s].(*uint)
			temp_value = int64(*x)
		case "uint":
			x := m[s].(uint)
			temp_value = int64(x)
		case "*string":
			value, value_error := strconv.ParseInt((*(m[s].(*string))), 10, 64)
			if value_error != nil {
				errors = append(errors, fmt.Errorf("error: Map.GetInt64: cannot convert *string value to int64"))
			} else {
				temp_value = value
			}
		default:
			errors = append(errors, fmt.Errorf("error: Map.GetInt64: type %s is not supported please implement", rep))
	}
	
	if len(errors) > 0 {
		return nil, errors
	}

	result := &temp_value

	return result, nil
}

func (m Map) GetInt8(s string) (*int8, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		return nil, nil
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < -128 || *int64_value > 127 {
		errors = append(errors, fmt.Errorf("error: value is not in range [-128, 127]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int8_conv := int8(*int64_value)
	result := &int8_conv

	return result, nil
}

func (m Map) GetInt8Value(s string) (int8, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: m.GetInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	if *int64_value < -128 || *int64_value > 127 {
		errors = append(errors, fmt.Errorf("error: value is not in range [-128, 127]"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int8_conv := int8(*int64_value)
	result := int8_conv

	return result, nil
}

func (m Map) GetUInt8(s string) (*uint8, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		return nil, nil
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < 0 || *int64_value > 255 {
		errors = append(errors, fmt.Errorf("error: value is not in range [0, 255]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int8_conv := uint8(*int64_value)
	result := &int8_conv

	return result, nil
}

func (m Map) GetUInt8Value(s string) (uint8, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: field: %s m.GetUInt64(s) returned nil", s))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	if *int64_value < 0 || *int64_value > 255 {
		errors = append(errors, fmt.Errorf("error: value is not in range [0, 255]"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int8_conv := uint8(*int64_value)
	result := int8_conv

	return result, nil
}

func (m Map) GetInt16(s string) (*int16, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		return nil, nil
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < -32768 || *int64_value > 32767 {
		errors = append(errors, fmt.Errorf("error: value is not in range [-32768, 32767]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int16_conv := int16(*int64_value)
	result := &int16_conv

	return result, nil
}

func (m Map) GetInt16Value(s string) (int16, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error:  m.GetInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	if *int64_value < -32768 || *int64_value > 32767 {
		errors = append(errors, fmt.Errorf("error: value is not in range [-32768, 32767]"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int16_conv := int16(*int64_value)
	result := int16_conv

	return result, nil
}

func (m Map) GetUInt16(s string) (*uint16, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		return nil, nil
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < 0 || *int64_value > 65535 {
		errors = append(errors, fmt.Errorf("error: value is not in range [0, 65535]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int16_conv := uint16(*int64_value)
	result := &int16_conv

	return result, nil
}

func (m Map) GetUInt16Value(s string) (uint16, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: field: %s m.GetUInt64(s) returned nil", s))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	if *int64_value < 0 || *int64_value > 65535 {
		errors = append(errors, fmt.Errorf("error: value is not in range [0, 65535]"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int16_conv := uint16(*int64_value)
	result := int16_conv

	return result, nil
}


func (m Map) GetInt32(s string) (*int32, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		return nil, nil
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < -2147483648 || *int64_value > 2147483647 {
		errors = append(errors, fmt.Errorf("error: value is not in range [-2147483648, 2147483647]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int32_conv := int32(*int64_value)
	result := &int32_conv

	return result, nil
}

func (m Map) GetInt32Value(s string) (int32, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error:  m.GetInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	if *int64_value < -2147483648 || *int64_value > 2147483647 {
		errors = append(errors, fmt.Errorf("error: value is not in range [-2147483648, 2147483647]"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int32_conv := int32(*int64_value)
	result := int32_conv

	return result, nil
}

func (m Map) GetInt64Value(s string) (int64, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error:  m.GetInt64(s) returned nil"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int64_conv := int64(*int64_value)
	result := int64_conv

	return result, nil
}

func (m Map) GetUInt32(s string) (*uint32, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		return nil, nil
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if *int64_value < 0 || *int64_value > 4294967295 {
		errors = append(errors, fmt.Errorf("error: value is not in range [0, 4294967295]"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	int32_conv := uint32(*int64_value)
	result := &int32_conv

	return result, nil
}

func (m Map) GetUInt32Value(s string) (uint32, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: field: %s m.GetUInt64(s) returned nil", s))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	if *int64_value < 0 || *int64_value > 4294967295 {
		errors = append(errors, fmt.Errorf("error: value is not in range [0, 4294967295]"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int32_conv := uint32(*int64_value)
	result := int32_conv

	return result, nil
}

func (m Map) GetUInt64Value(s string) (uint64, []error) {
	var errors []error
	int64_value, int64_value_errors := m.GetUInt64(s)
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: field: %s m.GetUInt64(s) returned nil", s))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	if *int64_value < 0 {
		errors = append(errors, fmt.Errorf("error: value is not in range [0, 18446744073709551615]"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int64_conv := uint64(*int64_value)
	result := int64_conv

	return result, nil
}

func (m Map) GetInt(s string) (*int, []error) {
	var errors []error
	var result int

	if m[s] == nil {
		return nil, nil
	}

	bit_size := strconv.IntSize
	if bit_size == 32 {
		temp_value, temp_value_errors := m.GetInt32(s)
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}

		result = int(*temp_value)
	} else if bit_size == 64 {
		temp_value, temp_value_errors := m.GetInt64(s)
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}

		result = int(*temp_value)
	} else {
		errors = append(errors, fmt.Errorf("Mao.GetInt bit size is not supported: %d", bit_size))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &result, nil
}

func (m Map) GetIntValue(s string) (int, []error) {
	var errors []error
	int_value, int_value_errors := m.GetInt(s)
	if int_value_errors != nil {
		errors = append(errors, int_value_errors...)
	} else if int_value == nil {
		errors = append(errors, fmt.Errorf("error:  m.GetInt(s) returned nil"))
	}

	if len(errors) > 0 {
		return 0, errors
	}

	int_conv := int(*int_value)
	result := int_conv

	return result, nil
}

func (m Map) SetInt(s string, v *int) {
	m[s] = v
}

func (m Map) SetInt64(s string, v *int64) {
	m[s] = v
}

func (m Map) SetInt32(s string, v *int32) {
	m[s] = v
}

func (m Map) SetInt16(s string, v *int16) {
	m[s] = v
}

func (m Map) SetInt8(s string, v *int8) {
	m[s] = v
}

///

func (m Map) SetIntValue(s string, v int) {
	m[s] = v
}

func (m Map) SetInt64Value(s string, v int64) {
	m[s] = v
}

func (m Map) SetInt32Value(s string, v int32) {
	m[s] = v
}

func (m Map) SetInt16Value(s string, v int16) {
	m[s] = v
}

func (m Map) SetInt8Value(s string, v int8) {
	m[s] = v
}

///

func (m Map) SetUInt(s string, v *uint) {
	m[s] = v
}

func (m Map) SetUInt64(s string, v *uint64) {
	m[s] = v
}

func (m Map) SetUInt32(s string, v *uint32) {
	m[s] = v
}

func (m Map) SetUInt16(s string, v *uint16) {
	m[s] = v
}

func (m Map) SetUInt8(s string, v *uint8) {
	m[s] = v
}

///

func (m Map) SetUIntValue(s string, v uint) {
	m[s] = v
}

func (m Map) SetUInt64Value(s string, v uint64) {
	m[s] = v
}

func (m Map) SetUInt32Value(s string, v uint32) {
	m[s] = v
}

func (m Map) SetUInt16Value(s string, v uint16) {
	m[s] = v
}

func (m Map) SetUInt8Value(s string, v uint8) {
	m[s] = v
}

///


func (m Map) SetFloat64(s string, v *float64) {
	m[s] = v
}

func (m Map) SetFloat64Value(s string, v float64) {
	m[s] = v
}

func (m Map) SetFloat32(s string, v *float32) {
	m[s] = v
}

func (m Map) SetFloat32Value(s string, v float32) {
	m[s] = v
}

func (m Map) GetUInt64(s string) (*uint64, []error) {
	var errors []error
	if m[s] == nil || m.IsNil(s) {
		return nil, nil
	}

	var uint64_value uint64
	rep := fmt.Sprintf("%T", m[s])
	switch rep {
	case "*int64":
		x := *(m[s].(*int64))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int64":
		x := m[s].(int64)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int32":
		x := *(m[s].(*int32))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int32":
		x := m[s].(int32)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int16":
		x := *(m[s].(*int16))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int16":
		x := m[s].(int16)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int8":
		x := *(m[s].(*int8))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int8":
		x := m[s].(int8)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int":
		x := (m[s].(int))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int":
		x := *(m[s].(*int))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "uint":
		uint_value := (m[s].(uint))
		uint64_value = uint64(uint_value)
	case "*uint":
		uint_value := *((m[s].(*uint)))
		uint64_value = uint64(uint_value)
	case "*uint64":
		uint64_value = *(m[s].(*uint64))
	case "uint64":
		uint64_value = (m[s].(uint64))
	case "*uint32":
		x := *(m[s].(*uint32))
		uint64_value = uint64(x)
	case "uint32":
		x := (m[s].(uint32))
		uint64_value = uint64(x)
	case "*uint16":
		x := *(m[s].(*uint16))
		uint64_value = uint64(x)
	case "uint16":
		x := (m[s].(uint16))
		uint64_value = uint64(x)
	case "*uint8":
		x := *(m[s].(*uint8))
		uint64_value = uint64(x)
	case "uint8":
		x := (m[s].(uint8))
		uint64_value = uint64(x)
	case "*string":
		string_value := (m[s].(*string))
		if *string_value == "NULL" {
			return nil, nil
		} else {
			value, value_error := strconv.ParseUint(*string_value, 10, 64)
			if value_error != nil {
				errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert *string value to uint64"))
			} else {
				uint64_value = value
			}
		}
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetUInt64: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &uint64_value, nil
}

func (m Map) SetTime(s string, value *time.Time) {
	m[s] = value
}

func (m Map) GetTime(s string, decimal_places int) (*time.Time, []error) {
	if m[s] == nil {
		return nil, nil
	}

	return common.GetTime(m[s], decimal_places)
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

	return ParseJSON(json_payload_builder.String())
}
