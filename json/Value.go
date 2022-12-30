package json

import (
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
	"time"
	"strings"
	"strconv"
)

type Value map[string](interface{})

func (v *Value) GetMap() (*Map, []error) {
	if v.IsNil() {
		return nil, nil
	}

	var errors []error
	var result *Map
	type_of := v.GetType()
	if type_of == "*json.Map" {
		result = ((*v)["value"]).(*Map)
	} else if type_of == "json.Map" {
		temp := ((*v)["value"]).(Map)
		result = &temp
	} else {
		errors = append(errors, fmt.Errorf("%s failed to unbox to json.Map", type_of))
		return nil, errors
	}

	return result, nil
}

func (v *Value) SetMap(s string, value *Map) []error {
	var errors []error
	if v.IsMap() {
		temp_map, temp_map_errors := v.GetMap()
		if temp_map_errors != nil {
			return temp_map_errors
		} else if common.IsNil(temp_map) {
			errors = append(errors, fmt.Errorf("Value.SetMap map is nil"))
		} else {
			set_map_value := Value{"value":value}
			(*temp_map)[s] = &set_map_value
		}
	} else {
		errors = append(errors, fmt.Errorf("Value.SetMap cannot set map on type %s", v.GetType()))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (v *Value) SetValue(s string, value *Value) []error {
	var errors []error
	if v.IsMap() {
		temp_map, temp_map_errors := v.GetMap()
		if temp_map_errors != nil {
			return temp_map_errors
		} else if common.IsNil(temp_map) {
			errors = append(errors, fmt.Errorf("Value.SetMap map is nil"))
		} else {
			set_map_value := value
			(*temp_map)[s] = set_map_value
		}
	} else {
		errors = append(errors, fmt.Errorf("Value.SetMap cannot set map on type %s", v.GetType()))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (v *Value) SetArray(s string, value *Array) []error {
	var errors []error
	if v.IsMap() {
		temp_map, temp_map_errors := v.GetMap()
		if temp_map_errors != nil {
			return temp_map_errors
		} else if common.IsNil(temp_map) {
			errors = append(errors, fmt.Errorf("Value.SetMap map is nil"))
		} else {
			set_value := Value{"value":value}
			(*temp_map)[s] = &set_value
		}
	} else {
		errors = append(errors, fmt.Errorf("Value.SetArray cannot set arrary on type %s", v.GetType()))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (v *Value) SetArrayValue(s string, a Array) []error {
	return v.SetArray(s, &a)
}

func (v *Value) IsMap() (bool) {
	return common.IsMap((*v)["value"])
}

func (v *Value) IsEmptyString() bool {
	if !common.IsString((*v)["value"]) {
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

func (v *Value) GetFunc() (func(Map) []error, []error) {
	if common.IsNil((*v)["value"]) {
		return nil, nil
	}

	var errors []error
	var result (func(Map) []error)
	rep := v.GetType()
	switch rep {
	case "func(json.Map) []error":
		result = (*v)["value"].(func(Map) []error)
	case "*func(json.Map) []error":
		result = *((*v)["value"].(*func(Map) []error))
	default:
		errors = append(errors, fmt.Errorf("error: Map.Func: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v *Value) GetErrors() ([]error, []error) {
	if v.IsNil() {
		return nil, nil
	}

	var errors []error
	var result []error
	
	rep := v.GetType()
	switch rep {
	case "*[]error":
		result = *((*v)["value"].(*[]error))
	case "[]error":
		result = (*v)["value"].([]error)
	case "*[]string":
		string_array := (*v)["value"].(*[]string)
		for _, string_array_value := range *string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "[]string":
		string_array := (*v)["value"].([]string)
		for _, string_array_value := range string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "json.Array":
		string_array := (*v)["value"].(Array)
		for _, string_array_value := range string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	case "*json.Array":
		string_array := (*v)["value"].(*Array)
		for _, string_array_value := range *string_array {
			converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
			if converted_errors != nil {
				errors = append(errors, converted_errors...)
			} else {
				result = append(result, fmt.Errorf("%s", *converted))
			}
		}
	default:
		errors = append(errors, fmt.Errorf("error: Value.GetErrors: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v *Value) GetArray() (*Array, []error) {
	if v.IsNil() {
		return nil, nil
	}

	var errors []error
	var result *Array

	rep := v.GetType()
	switch rep {
	case "*json.Array":
		result = (*v)["value"].(*Array)
	case "json.Array":
		temp := (*v)["value"].(Array)
		result = &temp
	default:
		errors = append(errors, fmt.Errorf("error: Value.GetArray: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v *Value) GetArrayOfInt8() (*[](*int8), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*int8))
	for _, array_value := range *array {
		int8_value, int8_value_errors := array_value.GetInt8()
		if int8_value_errors != nil {
			errors = append(errors, int8_value_errors...)
		} else {
			result = append(result, int8_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

func (v *Value) GetArrayOfInt16() (*[](*int16), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*int16))
	for _, array_value := range *array {
		int16_value, int16_value_errors := array_value.GetInt16()
		if int16_value_errors != nil {
			errors = append(errors, int16_value_errors...)
		} else {
			result = append(result, int16_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

func (v *Value) GetArrayOfInt32() (*[](*int32), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*int32))
	for _, array_value := range *array {
		int32_value, int32_value_errors := array_value.GetInt32()
		if int32_value_errors != nil {
			errors = append(errors, int32_value_errors...)
		} else {
			result = append(result, int32_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

func (v *Value) GetArrayOfInt64() (*[](*int64), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*int64))
	for _, array_value := range *array {
		int64_value, int64_value_errors := array_value.GetInt64()
		if int64_value_errors != nil {
			errors = append(errors, int64_value_errors...)
		} else {
			result = append(result, int64_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

///
func (v *Value) GetArrayOfUInt8() (*[](*uint8), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*uint8))
	for _, array_value := range *array {
		uint8_value, uint8_value_errors := array_value.GetUInt8()
		if uint8_value_errors != nil {
			errors = append(errors, uint8_value_errors...)
		} else {
			result = append(result, uint8_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

func (v *Value) GetArrayOfUInt16() (*[](*uint16), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*uint16))
	for _, array_value := range *array {
		uint16_value, uint16_value_errors := array_value.GetUInt16()
		if uint16_value_errors != nil {
			errors = append(errors, uint16_value_errors...)
		} else {
			result = append(result, uint16_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

func (v *Value) GetArrayOfUInt32() (*[](*uint32), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*uint32))
	for _, array_value := range *array {
		uint32_value, uint32_value_errors := array_value.GetUInt32()
		if uint32_value_errors != nil {
			errors = append(errors, uint32_value_errors...)
		} else {
			result = append(result, uint32_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

func (v *Value) GetArrayOfUInt64() (*[](*uint64), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*uint64))
	for _, array_value := range *array {
		uint64_value, uint64_value_errors := array_value.GetUInt64()
		if uint64_value_errors != nil {
			errors = append(errors, uint64_value_errors...)
		} else {
			result = append(result, uint64_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

func (v *Value) GetArrayOfString() (*[](*string), []error) {
	array, array_errors := v.GetArray()
	if array_errors != nil {
		return nil, array_errors
	} else if common.IsNil(array) {
		return nil, nil
	}

	var errors []error
	var result ([](*string))
	for _, array_value := range *array {
		string_value, string_value_errors := array_value.GetString()
		if string_value_errors != nil {
			errors = append(errors, string_value_errors...)
		} else {
			result = append(result, string_value)			 
		}
	}

	if len(errors) > 0 {
		return nil, errors
	}
	
	return &result, nil
}

func (v *Value) GetString() (*string, []error) {
	if common.IsNil((*v)["value"]){
		return nil, nil
	}

	var errors []error
	var result *string
	rep := fmt.Sprintf("%T", (*v)["value"])
	switch rep {
	case "string":
		value := (*v)["value"].(string)
		result = &value
	case "*string":
		result = (*v)["value"].(*string)
	default:
		errors = append(errors, fmt.Errorf("error: Map.GetString: type %s is not supported please implement for attribute: %s", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v *Value) GetStringValue() (string, []error) {
	var errors []error
	result, result_errors := v.GetString()
	if result_errors != nil {
		errors = append(errors, result_errors...)
	} else if common.IsNil(result) {
		errors = append(errors, fmt.Errorf("Value.GetStringValue is nil"))
	}
	
	if len(errors) > 0 {
		return "", errors
	}

	return *result, nil
}

func (v *Value) IsFloat() (bool) {
	if common.IsNil((*v)["value"]) {
		return false
	}

	type_of := common.GetType((*v)["value"])
	if type_of == "float32" || 
	   type_of == "*float32" || 
	   type_of == "float64" || 
	   type_of == "*float64" {
		return true
	}

	return false
}

func (v *Value) IsString() (bool) {
	if common.IsNil((*v)["value"]) {
		return false
	}

	return common.IsString((*v)["value"])
}

func (v *Value) IsNil() bool {
	return common.IsNil((*v)["value"])
}

func (v Value) IsNilForValue() bool {
	return common.IsNil((v)["value"])
}

func (v *Value) IsBool() bool {
	return common.IsBool((*v)["value"])
}

func (v *Value) IsArray() bool {
	return common.IsArray((*v)["value"])
}

func (v *Value) GetFloat64() (*float64, []error) {
	if v.IsNil() {
		return nil, nil
	}

	var errors []error
	var result *float64
	rep := fmt.Sprintf("%T", (*v)["value"])
	switch rep {
	case "float64":
		value := (*v)["value"].(float64)
		result = &value
	case "*float64":
		value := *((*v)["value"].(*float64))
		result = &value
	case "*string":
		value, value_error := strconv.ParseFloat((*((*v)["value"].(*string))),64)
		if value_error != nil {
			errors = append(errors, fmt.Errorf("error: Value.GetFloat64: cannot convert *string value to float64"))
		} else {
			result = &value
		}
	case "string":
		value, value_error := strconv.ParseFloat(((*v)["value"].(string)), 64)
		if value_error != nil {
			errors = append(errors, fmt.Errorf("error: Value.GetFloat64: cannot convert *string value to float64"))
		} else {
			result = &value
		}
	case "float32":
		value := float64((*v)["value"].(float32))
		result = &value
	case "*float32":
		value := float64(*((*v)["value"].(*float32)))
		result = &value
	default:
		errors = append(errors, fmt.Errorf("error: Value.GetFloat64: type %s is not supported please implement for attribute: %s", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return result, nil
}

func (v *Value) GetFloat32() (*float32, []error) {
	if v.IsNil() {
		return nil, nil
	}

	var errors []error
	float64_value, float64_value_errors := v.GetFloat64()
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

func (v *Value) GetFloat32Value() (float32, []error) {
	var errors []error
	float64_value, float64_value_errors := v.GetFloat64()
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

func (v *Value) GetFloat64Value() (float64, []error) {
	var errors []error
	float64_value, float64_value_errors := v.GetFloat64()
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

func (v *Value) GetRunes() (*[]rune, []error) {
	if v.IsNil() {
		return nil, nil
	}

	var errors []error
	var result *string
	rep := fmt.Sprintf("%T", (*v)["value"])
	switch rep {
	case "string":
		value := (*v)["value"].(string)
		newValue := strings.Clone(value)
		result = &newValue
	case "*string":
		if fmt.Sprintf("%s", (*v)["value"]) != "%!s(*string=<nil>)" {
			s := strings.Clone(*(((*v)["value"]).(*string)))
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

func (v *Value) GetBool() (*bool, []error) {
	if v.IsNil() {
		return nil, nil
	}

	var result *bool
	var errors []error

	rep := fmt.Sprintf("%T", (*v)["value"])
	switch rep {
	case "bool":
		value := (*v)["value"].(bool)
		result = &value
		break
	case "*bool":
		if fmt.Sprintf("%s", (*v)["value"]) != "%!s(*bool=<nil>)" {
			value := *(((*v)["value"]).(*bool))
			result = &value
		} else {
			return nil, nil
		}
		break
	case "*string":
		if fmt.Sprintf("%s", (*v)["value"]) != "%!s(*string=<nil>)" {
			value := *(((*v)["value"]).(*string))
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
		value := (((*v)["value"]).(string))
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

func (v *Value) GetInt64() (*int64, []error) {
	var errors []error
	var temp_value int64

	if v.IsNil() {
		return nil, nil
	}

	rep := fmt.Sprintf("%T", (*v)["value"])
	switch rep {
		case "*int64":
			x := (*v)["value"].(*int64)
			temp_value = int64(*x)
		case "int64":
			x := (*v)["value"].(int64)
			temp_value = x
		case "*int32":
			x :=(*v)["value"].(*int32)
			temp_value = int64(*x)
		case "int32":
			x := (*v)["value"].(int32)
			temp_value = int64(x)
		case "*int16":
			x := (*v)["value"].(*int16)
			temp_value = int64(*x)
		case "int16":
			x := (*v)["value"].(int16)
			temp_value = int64(x)
		case "*int8":
			x := (*v)["value"].(*int8)
			temp_value = int64(*x)
		case "int8":
			x := (*v)["value"].(int8)
			temp_value = int64(x)
		case "int":
			x := (*v)["value"].(int)
			temp_value = int64(x)
		case "*int":
			x := (*v)["value"].(*int)
			temp_value = int64(*x)
		case "*uint64":
			x := (*v)["value"].(*uint64)
			if *x > 9223372036854775807 {
				errors = append(errors, fmt.Errorf("%d is greater than 9223372036854775807", *x))
			} else {
				temp_value = int64(*x)
			}
		case "uint64":
			x := (*v)["value"].(uint64)
			if x > 9223372036854775807 {
				errors = append(errors, fmt.Errorf("%d is greater than 9223372036854775807", x))
			} else {
				temp_value = int64(x)
			}
		case "*uint32":
			x := (*v)["value"].(*uint32)
			temp_value = int64(*x)
		case "uint32":
			x := (*v)["value"].(uint32)
			temp_value = int64(x)
		case "*uint16":
			x := (*v)["value"].(*uint16)
			temp_value = int64(*x)
		case "uint16":
			x := (*v)["value"].(uint16)
			temp_value = int64(x)
		case "*uint8":
			x := (*v)["value"].(*uint8)
			temp_value = int64(*x)
		case "uint8":
			x := (*v)["value"].(uint8)
			temp_value = int64(x)
		case "*uint":
			x := (*v)["value"].(*uint)
			temp_value = int64(*x)
		case "uint":
			x := (*v)["value"].(uint)
			temp_value = int64(x)
		case "*string":
			value, value_error := strconv.ParseInt((*((*v)["value"].(*string))), 10, 64)
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

func (v *Value) GetUInt64() (*uint64, []error) {
	var errors []error
	if common.IsNil((*v)["value"]){
		return nil, nil
	}

	var uint64_value uint64
	rep := fmt.Sprintf("%T",(*v)["value"])
	switch rep {
	case "*int64":
		x := *((*v)["value"].(*int64))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int64":
		x := (*v)["value"].(int64)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int32":
		x := *((*v)["value"].(*int32))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int32":
		x := (*v)["value"].(int32)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int16":
		x := *((*v)["value"].(*int16))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int16":
		x := (*v)["value"].(int16)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int8":
		x := *((*v)["value"].(*int8))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int8":
		x := (*v)["value"].(int8)
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "int":
		x := ((*v)["value"].(int))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "*int":
		x := *((*v)["value"].(*int))
		if x >= 0 {
			uint64_value = uint64(x)
		} else {
			errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
		}
	case "uint":
		uint_value := ((*v)["value"].(uint))
		uint64_value = uint64(uint_value)
	case "*uint":
		uint_value := *(((*v)["value"].(*uint)))
		uint64_value = uint64(uint_value)
	case "*uint64":
		uint64_value = *((*v)["value"].(*uint64))
	case "uint64":
		uint64_value = ((*v)["value"].(uint64))
	case "*uint32":
		x := *((*v)["value"].(*uint32))
		uint64_value = uint64(x)
	case "uint32":
		x := ((*v)["value"].(uint32))
		uint64_value = uint64(x)
	case "*uint16":
		x := *((*v)["value"].(*uint16))
		uint64_value = uint64(x)
	case "uint16":
		x := ((*v)["value"].(uint16))
		uint64_value = uint64(x)
	case "*uint8":
		x := *((*v)["value"].(*uint8))
		uint64_value = uint64(x)
	case "uint8":
		x := ((*v)["value"].(uint8))
		uint64_value = uint64(x)
	case "*string":
		string_value := ((*v)["value"].(*string))
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

func (v *Value) GetTime(decimal_places int) (*time.Time, []error) {
	if v.IsNil() {
		return nil, nil
	}

	return common.GetTime((*v)["value"], decimal_places)
}

func (v *Value) GetType() string {
	return common.GetType((*v)["value"])
}

func (v Value) GetTypeForValue() string {
	return common.GetType((v)["value"])
}

func (v *Value) IsInteger() bool {
	return common.IsInteger((*v)["value"])
}

func (v *Value) IsBoolTrue() bool {
	return common.IsBoolTrue((*v)["value"])
}

func (v *Value) IsBoolFalse() bool {
	return common.IsBoolFalse((*v)["value"])
}

func (v *Value) GetInt8() (*int8, []error) {
	if common.IsNil((*v)["value"]) {
		return nil, nil
	}
	
	var errors []error
	int64_value, int64_value_errors := v.GetInt64()
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

func (v *Value) GetInt8Value() (int8, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetInt64()
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

func (v *Value) GetUInt8() (*uint8, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetUInt64()
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

func (v *Value) GetUInt8Value() (uint8, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetUInt64()
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: Value.GetUInt64(s) returned nil"))
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

func (v *Value) GetInt16() (*int16, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetInt64()
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

func (v *Value) GetInt16Value() (int16, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetInt64()
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

func (v *Value) GetUInt16() (*uint16, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetUInt64()
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

func (v *Value) GetUInt16Value() (uint16, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetUInt64()
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: field: Value.GetUInt64(s) returned nil"))
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


func (v *Value) GetInt32() (*int32, []error) {
	if v.IsNil() {
		return nil, nil
	}
	
	var errors []error
	int64_value, int64_value_errors := v.GetInt64()
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

func (v *Value) GetInt32Value() (int32, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetInt64()
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

func (v *Value) GetInt64Value() (int64, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetInt64()
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

func (v *Value) GetUInt32() (*uint32, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetUInt64()
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

func (v *Value) GetUInt32Value() (uint32, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetUInt64()
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: Value.GetUInt64(s) returned nil"))
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

func (v *Value) GetUInt64Value() (uint64, []error) {
	var errors []error
	int64_value, int64_value_errors := v.GetUInt64()
	if int64_value_errors != nil {
		errors = append(errors, int64_value_errors...)
	} else if int64_value == nil {
		errors = append(errors, fmt.Errorf("error: Value.GetUInt64(s) returned nil"))
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

func (v *Value) GetInt() (*int, []error) {
	var errors []error
	var result int
	bit_size := strconv.IntSize
	if bit_size == 32 {
		temp_value, temp_value_errors := v.GetInt32()
		if temp_value_errors != nil {
			return nil, temp_value_errors
		} else if temp_value == nil {
			return nil, nil
		}

		result = int(*temp_value)
	} else if bit_size == 64 {
		temp_value, temp_value_errors := v.GetInt64()
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

func (v *Value) GetIntFromMap(key string) (*int, []error) {
	var errors []error
	map_value, map_value_errors := v.GetMap()
	if map_value_errors != nil {
		errors = append(errors, map_value_errors...)
	} else if common.IsNil(map_value) {
		errors = append(errors, fmt.Errorf("Value.GetIntFromMap map is nil"))
	}
	if len(errors) > 0 {
		return nil, errors
	}
	return map_value.GetInt(key)
}

func (v *Value) SetIntToMap(key string, int_value *int) ([]error) {
	var errors []error
	map_value, map_value_errors := v.GetMap()
	if map_value_errors != nil {
		errors = append(errors, map_value_errors...)
	} else if common.IsNil(map_value) {
		errors = append(errors, fmt.Errorf("Value.SetIntToMap map is nil"))
	}
	if len(errors) > 0 {
		return errors
	}
	map_value.SetInt(key, int_value)
	return nil
}

func (v *Value) SetIntValueToMap(key string, int_value int) ([]error) {
	var errors []error
	map_value, map_value_errors := v.GetMap()
	if map_value_errors != nil {
		errors = append(errors, map_value_errors...)
	} else if common.IsNil(map_value) {
		errors = append(errors, fmt.Errorf("Value.SetIntToMap map is nil"))
	}
	if len(errors) > 0 {
		return errors
	}
	map_value.SetIntValue(key, int_value)
	return nil
}

func (v *Value) GetIntValue() (int, []error) {
	var errors []error
	int_value, int_value_errors := v.GetInt()
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

func (v *Value) AppendValue(add *Value) []error {
	var errors []error
	if v.IsArray() {
		array, array_errors := v.GetArray()
		if array_errors != nil {
			errors = append(errors, array_errors...)
		} else if common.IsNil(array) {
			errors = append(errors, fmt.Errorf("Value.AppendValue array is nil"))
		} else {
			add_value := add
			*array = append(*array, add_value)
		}
	} else {
		errors = append(errors, fmt.Errorf("Value.AppendValue not supported for type %s", v.GetType()))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (v *Value) AppendValueValue(add Value) []error {
	var errors []error
	if v.IsArray() {
		array, array_errors := v.GetArray()
		if array_errors != nil {
			errors = append(errors, array_errors...)
		} else if common.IsNil(array) {
			errors = append(errors, fmt.Errorf("Value.AppendValue array is nil"))
		} else {
			add_value := add
			*array = append(*array, &add_value)
		}
	} else {
		errors = append(errors, fmt.Errorf("Value.AppendValue not supported for type %s", v.GetType()))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (v *Value) GetObject() (interface{}, []error) {
	if v.IsNil() {
		return nil, nil
	}

	return (*v)["value"], nil
}

func (v Value) GetObjectForValue() (interface{}, []error) {
	if v.IsNilForValue() {
		return nil, nil
	}

	return (v)["value"], nil
}

