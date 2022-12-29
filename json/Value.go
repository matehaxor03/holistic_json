package json

import (
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
	"time"
	"strings"
	"strconv"
)

type Value map[string]interface{}

func (v Value) GetMap() (*Map, []error) {
	if common.IsNil(v["value"]) {
		return nil, nil
	}

	var errors []error
	var result Map
	type_of := common.GetType(v["value"])
	if type_of == "json.Map" {
		result = (v["value"]).(Map)
	} else if type_of == "*json.Map" {
		result = *((v["value"]).(*Map))
	} else {
		errors = append(errors, fmt.Errorf("%s failed to unbox to json.Map", type_of))
		return nil, errors
	}

	return &result, nil
}

func (v Value) IsMap() (bool) {
	return common.IsMap(v["value"])
}

func (v Value) IsEmptyString() bool {
	if !common.IsString(v["value"]) {
		return false
	}

	string_value, string_value_errors := v.GetString()
	if string_value_errors != nil{
		return false
	} else if common.IsNil(string_value) {
		return false
	}

	return *string_value == ""
}

func (v Value) GetFunc() (func(Map) []error, []error) {
	if common.IsNil(v["value"]) {
		return nil, nil
	}

	var errors []error
	var result (func(Map) []error)
	rep := fmt.Sprintf("%T", v["value"])
	switch rep {
	case "func(json.Map) []error":
		result = v["value"].(func(Map) []error)
	case "*func(json.Map) []error":
		result = *(v["value"].(*func(Map) []error))
	default:
		errors = append(errors, fmt.Errorf("error: Map.Func: type %s is not supported please implement for field: %s", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v Value) GetErrors() ([]error, []error) {
	if common.IsNil(v["value"]) {
		return nil, nil
	}

	var errors []error
	var result []error
	
	rep := fmt.Sprintf("%T", v["value"])
	switch rep {
	case "*[]error":
		result = *(v["value"].(*[]error))
	case "[]error":
		result = v["value"].([]error)
	case "*[]string":
		string_array := v["value"].(*[]string)
		for _, string_array_value := range *string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "[]string":
		string_array := v["value"].([]string)
		for _, string_array_value := range string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "json.Array":
		string_array := v["value"].(Array)
		for _, string_array_value := range string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "*json.Array":
		string_array := v["value"].(*Array)
		for _, string_array_value := range *string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	default:
		errors = append(errors, fmt.Errorf("error: Value.GetErrors: type %s is not supported please implement for field: %s", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v Value) GetArray() (*Array, []error) {
	if common.IsNil(v["value"]) {
		return nil, nil
	}

	var errors []error
	var result *Array

	rep := fmt.Sprintf("%T", v["value"])
	switch rep {
	case "*json.Array":
		result = v["value"].(*Array)
	case "json.Array":
		temp := v["value"].(Array)
		result = &temp
	default:
		errors = append(errors, fmt.Errorf("error: Value.GetArray: type %s is not supported please implement for field: %s", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v Value) GetString() (*string, []error) {
	if common.IsNil(v["value"]){
		return nil, nil
	}

	var errors []error
	var result *string
	rep := fmt.Sprintf("%T", v["value"])
	switch rep {
	case "string":
		value := v["value"].(string)
		newValue := strings.Clone(value)
		result = &newValue
	case "*string":
		if fmt.Sprintf("%s", v["value"]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*((v["value"]).(*string)))
			result = &s
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetString: *string value is null for attribute: %s", rep))
		}
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetString: type %s is not supported please implement for attribute: %s", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v Value) IsFloat() (bool) {
	if common.IsNil(v["value"]) {
		return false
	}

	type_of := common.GetType(v["value"])
	if type_of == "float32" || 
	   type_of == "*float32" || 
	   type_of == "float64" || 
	   type_of == "*float64" {
		return true
	}

	return false
}

func (v Value) IsString() (bool) {
	if common.IsNil(v["value"]) {
		return false
	}

	return common.IsString(v["value"])
}

func (v Value) IsNil() bool {
	return common.IsNil(v["value"])
}

func (v Value) IsBool() bool {
	return common.IsBool(v["value"])
}

func (v Value) IsArray() bool {
	return common.IsArray(v["value"])
}

func (v Value) GetFloat64() (*float64, []error) {
	if common.IsNil(v["value"]) {
		return nil, nil
	}

	var errors []error
	var result *float64
	rep := fmt.Sprintf("%T",v["value"])
	switch rep {
	case "float64":
		value := v["value"].(float64)
		result = &value
	case "*float64":
		value := *(v["value"].(*float64))
		result = &value
	case "*string":
		value, value_error := strconv.ParseFloat((*(v["value"].(*string))),64)
		if value_error != nil {
			errors = append(errors, fmt.Errorf("error: Value.GetFloat64: cannot convert *string value to float64"))
		} else {
			result = &value
		}
	case "string":
		value, value_error := strconv.ParseFloat((v["value"].(string)), 64)
		if value_error != nil {
			errors = append(errors, fmt.Errorf("error: Value.GetFloat64: cannot convert *string value to float64"))
		} else {
			result = &value
		}
	case "float32":
		value := float64(v["value"].(float32))
		result = &value
	case "*float32":
		value := float64(*(v["value"].(*float32)))
		result = &value
	default:
		errors = append(errors, fmt.Errorf("error: Value.GetFloat64: type %s is not supported please implement for attribute: %s", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v Value) GetRunes() (*[]rune, []error) {
	if common.IsNil(v["value"]) {
		return nil, nil
	}

	var errors []error
	var result *string
	rep := fmt.Sprintf("%T", v["value"])
	switch rep {
	case "string":
		value := v["value"].(string)
		newValue := strings.Clone(value)
		result = &newValue
	case "*string":
		if fmt.Sprintf("%s", v["value"]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*((v["value"]).(*string)))
			result = &s
		} else {
			errors = append(errors, fmt.Errorf("error: Value.GetString: *string value is null for attribute: %s", rep))
		}
	default:
		errors = append(errors, fmt.Errorf("error: Value.GetString: type %s is not supported please implement for attribute: %s", rep))
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


func (v Value) GetBool() (*bool, []error) {
	if common.IsNil(v["value"]){
		return nil, nil
	}

	var result *bool
	var errors []error

	rep := fmt.Sprintf("%T", v["value"])
	switch rep {
	case "bool":
		value := v["value"].(bool)
		result = &value
		break
	case "*bool":
		if fmt.Sprintf("%s", v["value"]) != "%!s(*bool=<nil>)" {
			value := *((v["value"]).(*bool))
			result = &value
		} else {
			return nil, nil
		}
		break
	case "*string":
		if fmt.Sprintf("%s", v["value"]) != "%!s(*string=<nil>)" {
			value := *((v["value"]).(*string))
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
		value := ((v["value"]).(string))
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

func (v Value) GetInt64() (*int64, []error) {
	var errors []error
	var temp_value int64

	if common.IsNil(v["value"]) {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", v["value"])
	switch rep {
		case "*int64":
			x := v["value"].(*int64)
			temp_value = int64(*x)
		case "int64":
			x := v["value"].(int64)
			temp_value = x
		case "*int32":
			x :=v["value"].(*int32)
			temp_value = int64(*x)
		case "int32":
			x := v["value"].(int32)
			temp_value = int64(x)
		case "*int16":
			x := v["value"].(*int16)
			temp_value = int64(*x)
		case "int16":
			x := v["value"].(int16)
			temp_value = int64(x)
		case "*int8":
			x := v["value"].(*int8)
			temp_value = int64(*x)
		case "int8":
			x := v["value"].(int8)
			temp_value = int64(x)
		case "int":
			x := v["value"].(int)
			temp_value = int64(x)
		case "*int":
			x := v["value"].(*int)
			temp_value = int64(*x)
		case "*uint64":
			x := v["value"].(*uint64)
			if *x > 9223372036854775807 {
				errors = append(errors, fmt.Errorf("%d is greater than 9223372036854775807", *x))
			} else {
				temp_value = int64(*x)
			}
		case "uint64":
			x := v["value"].(uint64)
			if x > 9223372036854775807 {
				errors = append(errors, fmt.Errorf("%d is greater than 9223372036854775807", x))
			} else {
				temp_value = int64(x)
			}
		case "*uint32":
			x := v["value"].(*uint32)
			temp_value = int64(*x)
		case "uint32":
			x := v["value"].(uint32)
			temp_value = int64(x)
		case "*uint16":
			x := v["value"].(*uint16)
			temp_value = int64(*x)
		case "uint16":
			x := v["value"].(uint16)
			temp_value = int64(x)
		case "*uint8":
			x := v["value"].(*uint8)
			temp_value = int64(*x)
		case "uint8":
			x := v["value"].(uint8)
			temp_value = int64(x)
		case "*uint":
			x := v["value"].(*uint)
			temp_value = int64(*x)
		case "uint":
			x := v["value"].(uint)
			temp_value = int64(x)
		case "*string":
			value, value_error := strconv.ParseInt((*(v["value"].(*string))), 10, 64)
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

func (v Value) GetUInt64() (*uint64, []error) {
	var errors []error
	if common.IsNil(v["value"]){
		return nil, nil
	}

	var uint64_value uint64
	rep := fmt.Sprintf("%T",v["value"])
	switch rep {
	case "*int64":
		x := *(v["value"].(*int64))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int64":
		x := v["value"].(int64)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int32":
		x := *(v["value"].(*int32))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int32":
		x := v["value"].(int32)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int16":
		x := *(v["value"].(*int16))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int16":
		x := v["value"].(int16)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int8":
		x := *(v["value"].(*int8))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int8":
		x := v["value"].(int8)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int":
		x := (v["value"].(int))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int":
		x := *(v["value"].(*int))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "uint":
		uint_value := (v["value"].(uint))
		uint64_value = uint64(uint_value)
	case "*uint":
		uint_value := *((v["value"].(*uint)))
		uint64_value = uint64(uint_value)
	case "*uint64":
		uint64_value = *(v["value"].(*uint64))
	case "uint64":
		uint64_value = (v["value"].(uint64))
	case "*uint32":
		x := *(v["value"].(*uint32))
		uint64_value = uint64(x)
	case "uint32":
		x := (v["value"].(uint32))
		uint64_value = uint64(x)
	case "*uint16":
		x := *(v["value"].(*uint16))
		uint64_value = uint64(x)
	case "uint16":
		x := (v["value"].(uint16))
		uint64_value = uint64(x)
	case "*uint8":
		x := *(v["value"].(*uint8))
		uint64_value = uint64(x)
	case "uint8":
		x := (v["value"].(uint8))
		uint64_value = uint64(x)
	case "*string":
		string_value := (v["value"].(*string))
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

func (v Value) GetTime(decimal_places int) (*time.Time, []error) {
	if common.IsNil(v["value"]) {
		return nil, nil
	}

	return common.GetTime(v["value"], decimal_places)
}

func (v Value) GetType() string {
	if common.IsNil(v["value"]) {
		return "nil"
	}

	return fmt.Sprintf("%T", v["value"])
}

func (v Value) IsInteger() bool {
	return common.IsInteger(v["value"])
}

