package json

import (
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
	"time"
	"strings"
	"strconv"
)

//type Value map[string](interface{})

type Value struct {

	GetMap func() (*Map, []error) 
	SetMap func(s string, value *Map) []error
	SetValue func(s string, value *Value) []error
	SetArray func(s string, value *Array) []error
	SetArrayValue func(s string, a Array) []error
	IsMap func() (bool) 
	IsEmptyString func() bool
	GetFunc func() (func(Map) []error, []error)
	GetErrors func() ([]error, []error)
	GetArray func() (*Array, []error)
	GetArrayValue func() (Array, []error)
	GetArrayOfInt func() (*[](*int), []error)
	GetArrayOfInt8 func() (*[](*int8), []error)
	GetArrayOfInt16 func() (*[](*int16), []error)
	GetArrayOfInt32 func() (*[](*int32), []error) 
	GetArrayOfInt64 func() (*[](*int64), []error) 
	GetArrayOfUInt func() (*[](*uint), []error)
	GetArrayOfUInt8 func() (*[](*uint8), []error)
	GetArrayOfUInt16 func() (*[](*uint16), []error)
	GetArrayOfUInt32 func() (*[](*uint32), []error) 
	GetArrayOfUInt64 func() (*[](*uint64), []error) 
	GetArrayOfIntValue func() ([](int), []error)
	GetArrayOfInt8Value func() ([](int8), []error)
	GetArrayOfInt16Value func() ([](int16), []error)
	GetArrayOfInt32Value func() ([](int32), []error) 
	GetArrayOfInt64Value func() ([](int64), []error) 
	GetArrayOfUIntValue func() ([](uint), []error)
	GetArrayOfUInt8Value func() ([](uint8), []error)
	GetArrayOfUInt16Value func() ([](uint16), []error)
	GetArrayOfUInt32Value func() ([](uint32), []error) 
	GetArrayOfUInt64Value func() ([](uint64), []error) 
	GetStringValue func() (string, []error) 
	GetString func() (*string, []error) 
	GetBool func() (*bool, []error) 
	GetBoolValue func() (bool, []error) 
	AppendValue func(add *Value) []error
	IsArray func() bool
	GetType func() string
	GetObject func() (interface{})
	SetObject func(value interface{})

	IsFloat func() bool
	IsBool func() bool
	IsBoolTrue func() bool
	IsBoolFalse func() bool
	IsInteger func() bool
	IsString func() bool
	GetArrayOfString func() (*[](*string), []error)
	GetArrayOfStringValue func() ([](string), []error)
	GetFloat32 func() (*float32, []error) 
	GetFloat32Value func() (float32, []error) 
	GetFloat64 func() (*float64, []error) 
	GetFloat64Value func() (float64, []error) 
	GetInt func() (*int, []error) 
	GetIntValue func() (int, []error) 
	GetInt8 func() (*int8, []error) 
	GetInt8Value func() (int8, []error) 
	GetInt16 func() (*int16, []error) 
	GetInt16Value func() (int16, []error) 
	GetInt32 func() (*int32, []error) 
	GetInt32Value func() (int32, []error) 
	GetInt64 func() (*int64, []error) 
	GetInt64Value func() (int64, []error) 
	GetUInt func() (*uint, []error) 
	GetUIntValue func() (uint, []error) 
	GetUInt8 func() (*uint8, []error) 
	GetUInt8Value func() (uint8, []error) 
	GetUInt16 func() (*uint16, []error) 
	GetUInt16Value func() (uint16, []error) 
	GetUInt32 func() (*uint32, []error) 
	GetUInt32Value func() (uint32, []error) 
	GetUInt64 func() (*uint64, []error) 
	GetUInt64Value func() (uint64, []error) 
	GetRunes func() (*[]*rune, []error)
	GetTimeWithDecimalPlaces func(decimal_places int) (*time.Time, []error)
	GetTime func() (*time.Time, []error)
	IsNil func() bool
	GetIntFromMap func(key string) (*int, []error)
	SetIntToMap func(key string, int_value *int) ([]error)
	SetIntValueToMap func(key string, int_value int) ([]error)
	AppendValueValue func(add Value) []error
	GetArrayOfFloat32 func() (*[]*float32, []error)
	GetArrayOfFloat32Value func() ([]float32, []error)
	GetArrayOfFloat64 func() (*[]*float64, []error)
	GetArrayOfFloat64Value func() ([]float64, []error)



}

func newValue(v interface{}) (*Value) {
	var this_value *Value
	internal_value := v

	set_this_value := func(value *Value) {
		this_value = value
	}
	
	this := func() *Value {
		return this_value
	}

	get_internal_value := func() interface{} {
		return internal_value
	}

	getObject := func() interface{} {
		if common.IsNil(get_internal_value()) {
			return nil
		}
	
		return get_internal_value()
	}

	setObject := func(value interface{}) {
		internal_value = value
	}
	
	created_value := Value{
		SetValue: func(s string, value *Value) []error {
			var errors []error
			if common.IsMap(this().GetObject()) {
				temp_map, temp_map_errors := this().GetMap()
				if temp_map_errors != nil {
					return temp_map_errors
				} else if common.IsNil(temp_map) {
					errors = append(errors, fmt.Errorf("Value.SetMap map is nil"))
				} else {
					set_map_value := value
					(*temp_map).SetValue(s, set_map_value)
				}
			} else {
				errors = append(errors, fmt.Errorf("Value.SetMap cannot set map on type %s", this().GetType()))
			}
		
			if len(errors) > 0 {
				return errors
			}
		
			return nil
		},
		SetArray: func(s string, value *Array) []error {
			var errors []error
			if common.IsMap(this().GetObject()) {
				temp_map, temp_map_errors := this().GetMap()
				if temp_map_errors != nil {
					return temp_map_errors
				} else if common.IsNil(temp_map) {
					errors = append(errors, fmt.Errorf("Value.SetMap map is nil"))
				} else {
					set_value := newValue(value)
					(*temp_map).SetValue(s, set_value)
				}
			} else {
				errors = append(errors, fmt.Errorf("Value.SetArray cannot set arrary on type %s", this().GetType()))
			}
		
			if len(errors) > 0 {
				return errors
			}
		
			return nil
		},
		SetArrayValue: func(s string, a Array) []error {
			return this().SetArray(s, &a)
		},
		IsMap: func() (bool) {
			if common.IsNil(this().GetObject()) {
				return false
			}
			return common.IsMap(this().GetObject())
		},
		SetMap: func(s string, value *Map) []error {
			var errors []error
			if common.IsMap(this().GetObject()) {
				temp_map, temp_map_errors := this().GetMap()
				if temp_map_errors != nil {
					return temp_map_errors
				} else if common.IsNil(temp_map) {
					errors = append(errors, fmt.Errorf("Value.SetMap map is nil"))
				} else {
					set_map_value := newValue(value)
					(*temp_map).SetValue(s, set_map_value)
				}
			} else {
				errors = append(errors, fmt.Errorf("Value.SetMap cannot set map on type %s", this().GetType()))
			}
		
			if len(errors) > 0 {
				return errors
			}
		
			return nil
		},
		IsEmptyString: func() bool {
			if common.IsNil(this().GetObject()) {
				return false
			}

			if !common.IsString(this().GetObject()) {
				return false
			}
		
			string_value, string_value_errors := this().GetString()
			if string_value_errors != nil{
				return false
			} else if common.IsNil(string_value) {
				return false
			}
		
			return *string_value == ""
		},
		GetMap: func() (*Map, []error) {
			if common.IsNil(this().GetObject()) { 
				return nil, nil
			}
		
			var errors []error
			var result *Map
			type_of := this().GetType()
			if type_of == "*json.Map" {
				result = this().GetObject().(*Map)
			} else if type_of == "json.Map" {
				temp := this().GetObject().(Map)
				result = &temp
			} else {
				errors = append(errors, fmt.Errorf("%s failed to unbox to json.Map", type_of))
				return nil, errors
			}
		
			return result, nil
		},
		GetFunc: func() (func(Map) []error, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			var errors []error
			var result (func(Map) []error)
			rep := common.GetType(get_internal_value())
			switch rep {
			case "func(json.Map) []error":
				result = (this().GetObject()).(func(Map) []error)
			case "*func(json.Map) []error":
				result = *(this().GetObject().(*func(Map) []error))
			default:
				errors = append(errors, fmt.Errorf("error: Map.Func: type %s is not supported please implement", rep))
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
		
			return result, nil
		},
		GetErrors: func() ([]error, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			var errors []error
			var result []error
			
			rep := common.GetType(get_internal_value())
			switch rep {
			case "*[]error":
				result = *(this().GetObject().(*[]error))
			case "[]error":
				result = this().GetObject().([]error)
			case "*[]string":
				string_array := this().GetObject().(*[]string)
				for _, string_array_value := range *string_array {
					converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
					if converted_errors != nil {
						errors = append(errors, converted_errors...)
					} else {
						result = append(result, fmt.Errorf("%s", *converted))
					}
				}
			case "[]string":
				string_array := this().GetObject().([]string)
				for _, string_array_value := range string_array {
					converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
					if converted_errors != nil {
						errors = append(errors, converted_errors...)
					} else {
						result = append(result, fmt.Errorf("%s", *converted))
					}
				}
			case "json.Array":
				string_array := this().GetObject().(Array)
				for _, string_array_value := range *(string_array.Values()) {
					converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
					if converted_errors != nil {
						errors = append(errors, converted_errors...)
					} else {
						result = append(result, fmt.Errorf("%s", *converted))
					}
				}
			case "*json.Array":
				string_array := this().GetObject().(*Array)
				for _, string_array_value := range *(string_array.Values()) {
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
		},
		GetArray: func() (*Array, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			var errors []error
			var result *Array
		
			rep := common.GetType(get_internal_value())
			switch rep {
			case "*json.Array":
				result = this().GetObject().(*Array)
			case "json.Array":
				temp := this().GetObject().(Array)
				result = &temp
			default:
				errors = append(errors, fmt.Errorf("error: Value.GetArray: type %s is not supported please implement", rep))
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
		
			return result, nil
		},
		GetArrayValue: func() (Array, []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return Array{}, array_errors
			} else if common.IsNil(array) {
				var errors []error
				errors = append(errors, fmt.Errorf("Value.GetArrayValue array is nil"))
				return Array{}, errors
			}

			return *array, nil
		},
		GetArrayOfFloat32: func() (*[](*float32), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfFloat32()
		},
		GetArrayOfFloat32Value: func() ([](float32), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfFloat32Value()
		},
		GetArrayOfFloat64: func() (*[](*float64), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfFloat64()
		},
		GetArrayOfFloat64Value: func() ([](float64), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfFloat64Value()
		},
		GetArrayOfInt: func() (*[](*int), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt()
		},
		GetArrayOfInt8: func() (*[](*int8), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt8()
		},
		GetArrayOfInt16: func() (*[](*int16), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt16()
		},
		GetArrayOfInt32: func() (*[](*int32), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt32()
		},
		GetArrayOfInt64: func() (*[](*int64), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt64()
		},
		GetArrayOfIntValue: func() ([](int), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfIntValue()
		},
		GetArrayOfInt8Value: func() ([](int8), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt8Value()
		},
		GetArrayOfInt16Value: func() ([](int16), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt16Value()
		},
		GetArrayOfInt32Value: func() ([](int32), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt32Value()
		},
		GetArrayOfInt64Value: func() ([](int64), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt64Value()
		},
		GetArrayOfUInt: func() (*[](*uint), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt()
		},
		GetArrayOfUInt8: func() (*[](*uint8), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt8()
		},
		GetArrayOfUInt16: func() (*[](*uint16), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt16()
		},
		GetArrayOfUInt32: func() (*[](*uint32), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt32()
		},
		GetArrayOfUInt64: func() (*[](*uint64), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt64()
		},
		GetArrayOfUIntValue: func() ([](uint), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUIntValue()
		},
		GetArrayOfUInt8Value: func() ([](uint8), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt8Value()
		},
		GetArrayOfUInt16Value: func() ([](uint16), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt16Value()
		},
		GetArrayOfUInt32Value: func() ([](uint32), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt32Value()
		},
		GetArrayOfUInt64Value: func() ([](uint64), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt64Value()
		},
		GetArrayOfString: func() (*[](*string), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfString()
		},
		GetArrayOfStringValue: func() ([](string), []error) {
			array, array_errors := this().GetArray()
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfStringValue()
		},
		GetString: func() (*string, []error) {
			if common.IsNil(this().GetObject()){
				return nil, nil
			}
		
			var errors []error
			var result *string
			rep := common.GetType(get_internal_value())
			switch rep {
			case "string":
				value := this().GetObject().(string)
				result = &value
			case "*string":
				result = this().GetObject().(*string)
			default:
				errors = append(errors, fmt.Errorf("error: Map.GetString: type %s is not supported please implement for attribute: %s", rep))
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
		
			return result, nil
		},
		GetStringValue: func() (string, []error) {
			var errors []error
			result, result_errors := this().GetString()
			if result_errors != nil {
				errors = append(errors, result_errors...)
			} else if common.IsNil(result) {
				errors = append(errors, fmt.Errorf("Value.GetStringValue is nil"))
			}
			
			if len(errors) > 0 {
				return "", errors
			}
		
			return *result, nil
		},
		IsFloat: func() (bool) {
			if common.IsNil(this().GetObject()) {
				return false
			}
		
			type_of := common.GetType(this().GetObject())
			if type_of == "float32" || 
			   type_of == "*float32" || 
			   type_of == "float64" || 
			   type_of == "*float64" {
				return true
			}
		
			return false
		},
		IsString: func() (bool) {
			if common.IsNil(this().GetObject()) {
				return false
			}
		
			return common.IsString(this().GetObject())
		},
		IsNil: func() bool {
			return common.IsNil(this().GetObject())
		},
		IsBool: func() bool {
			return common.IsBool(this().GetObject())
		},
		IsArray: func() bool {
			return common.IsArray(this().GetObject())
		},
		GetFloat64: func() (*float64, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			var errors []error
			var result *float64
			rep := common.GetType(get_internal_value())
			switch rep {
			case "float64":
				value := this().GetObject().(float64)
				result = &value
			case "*float64":
				value := *(this().GetObject().(*float64))
				result = &value
			case "*string":
				value, value_error := strconv.ParseFloat(*(this().GetObject().(*string)),64)
				if value_error != nil {
					errors = append(errors, fmt.Errorf("error: Value.GetFloat64: cannot convert *string value to float64"))
				} else {
					result = &value
				}
			case "string":
				value, value_error := strconv.ParseFloat((this().GetObject().(string)), 64)
				if value_error != nil {
					errors = append(errors, fmt.Errorf("error: Value.GetFloat64: cannot convert *string value to float64"))
				} else {
					result = &value
				}
			case "float32":
				value := float64(this().GetObject().(float32))
				result = &value
			case "*float32":
				value := float64(*(this().GetObject().(*float32)))
				result = &value
			default:
				errors = append(errors, fmt.Errorf("error: Value.GetFloat64: type %s is not supported please implement for attribute: %s", rep))
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
		
			return result, nil
		},
		GetFloat32: func() (*float32, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			var errors []error
			float64_value, float64_value_errors := this().GetFloat64()
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
		},
		GetFloat32Value: func() (float32, []error) {
			var errors []error
			float64_value, float64_value_errors := this().GetFloat64()
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
		},
		GetFloat64Value: func() (float64, []error) {
			var errors []error
			float64_value, float64_value_errors := this().GetFloat64()
			if float64_value_errors != nil {
				return 0, float64_value_errors
			} else if float64_value == nil {
				errors = append(errors, fmt.Errorf("error: Map.GetFloat64 returned a nil value"))
			}
		
			if len(errors) > 0 {
				return 0, errors
			}
		
			return *float64_value, nil
		},
		GetRunes: func() (*[]*rune, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			var errors []error
			var result *string
			rep := common.GetType(get_internal_value())
			switch rep {
			case "string":
				value := this().GetObject().(string)
				cloned_value := strings.Clone(value)
				result = &cloned_value
			case "*string":
				if fmt.Sprintf("%s", this().GetObject()) != "%!s(*string=<nil>)" {
					s := strings.Clone(*(this().GetObject().(*string)))
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
		
			var runes [](*rune)
			for _, runeValue := range *result {
				runes = append(runes, &runeValue)
			}
		
			return &runes, nil
		},
		GetBool: func() (*bool, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			var result *bool
			var errors []error
		
			rep := common.GetType(get_internal_value())
			switch rep {
			case "bool":
				value := this().GetObject().(bool)
				result = &value
				break
			case "*bool":
				if fmt.Sprintf("%s", this().GetObject()) != "%!s(*bool=<nil>)" {
					value := *(this().GetObject().(*bool))
					result = &value
				} else {
					return nil, nil
				}
				break
			case "*string":
				if fmt.Sprintf("%s", this().GetObject()) != "%!s(*string=<nil>)" {
					value := *(this().GetObject().(*string))
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
				value := (this().GetObject().(string))
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
		},
		GetInt64: func() (*int64, []error) {
			var errors []error
			var temp_value int64
		
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			rep := common.GetType(get_internal_value())
			switch rep {
				case "json.Value":
					string_value, convert_errors := ConvertInterfaceValueToStringValue(this().GetObject().(Value))
					if convert_errors != nil {
						errors = append(errors, convert_errors...)
					} else if common.IsNil(string_value) {
						return nil, nil
					} else {
						value, value_error := strconv.ParseInt(*string_value, 10, 64)
						if value_error != nil {
							errors = append(errors, fmt.Errorf("error: Map.GetInt64: cannot convert json.Value to int64 %s", value_error))
						} else {
							temp_value = value
						}
					}
				case "*json.Value":
					string_value, convert_errors := ConvertInterfaceValueToStringValue(*(this().GetObject().(*Value)))
					if convert_errors != nil {
						errors = append(errors, convert_errors...)
					} else if common.IsNil(string_value) {
						return nil, nil
					} else {
						value, value_error := strconv.ParseInt(*string_value, 10, 64)
						if value_error != nil {
							errors = append(errors, fmt.Errorf("error: Map.GetInt64: cannot convert json.Value to int64 %s", value_error))
						} else {
							temp_value = value
						}
					}
				case "*int64":
					x := this().GetObject().(*int64)
					temp_value = int64(*x)
				case "int64":
					x := this().GetObject().(int64)
					temp_value = x
				case "*int32":
					x := this().GetObject().(*int32)
					temp_value = int64(*x)
				case "int32":
					x := this().GetObject().(int32)
					temp_value = int64(x)
				case "*int16":
					x := this().GetObject().(*int16)
					temp_value = int64(*x)
				case "int16":
					x := this().GetObject().(int16)
					temp_value = int64(x)
				case "*int8":
					x := this().GetObject().(*int8)
					temp_value = int64(*x)
				case "int8":
					x := this().GetObject().(int8)
					temp_value = int64(x)
				case "int":
					x := this().GetObject().(int)
					temp_value = int64(x)
				case "*int":
					x := this().GetObject().(*int)
					temp_value = int64(*x)
				case "*uint64":
					x := this().GetObject().(*uint64)
					if *x > 9223372036854775807 {
						errors = append(errors, fmt.Errorf("%d is greater than 9223372036854775807", *x))
					} else {
						temp_value = int64(*x)
					}
				case "uint64":
					x := this().GetObject().(uint64)
					if x > 9223372036854775807 {
						errors = append(errors, fmt.Errorf("%d is greater than 9223372036854775807", x))
					} else {
						temp_value = int64(x)
					}
				case "*uint32":
					x := this().GetObject().(*uint32)
					temp_value = int64(*x)
				case "uint32":
					x := this().GetObject().(uint32)
					temp_value = int64(x)
				case "*uint16":
					x := this().GetObject().(*uint16)
					temp_value = int64(*x)
				case "uint16":
					x := this().GetObject().(uint16)
					temp_value = int64(x)
				case "*uint8":
					x := this().GetObject().(*uint8)
					temp_value = int64(*x)
				case "uint8":
					x := this().GetObject().(uint8)
					temp_value = int64(x)
				case "*uint":
					x := this().GetObject().(*uint)
					temp_value = int64(*x)
				case "uint":
					x := this().GetObject().(uint)
					temp_value = int64(x)
				case "*string":
					value, value_error := strconv.ParseInt(*(this().GetObject().(*string)), 10, 64)
					if value_error != nil {
						errors = append(errors, fmt.Errorf("error: Map.GetInt64: cannot convert *string value to int64"))
					} else {
						temp_value = value
					}
				case "string":
					value, value_error := strconv.ParseInt((this().GetObject().(string)), 10, 64)
					if value_error != nil {
						errors = append(errors, fmt.Errorf("error: Map.GetInt64: cannot convert string value to int64"))
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
		},
		GetUInt64: func() (*uint64, []error) {
			var errors []error
			if common.IsNil(this().GetObject()){
				return nil, nil
			}

			var uint64_value uint64
			rep := common.GetType(get_internal_value())
			switch rep {
			case "*int64":
				x := *(this().GetObject().(*int64))
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "int64":
				x := this().GetObject().(int64)
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "*int32":
				x := *(this().GetObject().(*int32))
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "int32":
				x := this().GetObject().(int32)
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "*int16":
				x := *(this().GetObject().(*int16))
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "int16":
				x := (this().GetObject().(int16))
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "*int8":
				x := *(this().GetObject().(*int8))
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "int8":
				x := this().GetObject().(int8)
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "int":
				x := this().GetObject().(int)
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "*int":
				x := *(this().GetObject().(*int))
				if x >= 0 {
					uint64_value = uint64(x)
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetUInt64: cannot convert negative numbers for uint64"))
				}
			case "uint":
				uint_value := (this().GetObject().(uint))
				uint64_value = uint64(uint_value)
			case "*uint":
				uint_value := *(this().GetObject().(*uint))
				uint64_value = uint64(uint_value)
			case "*uint64":
				uint64_value = *(this().GetObject().(*uint64))
			case "uint64":
				uint64_value = (this().GetObject().(uint64))
			case "*uint32":
				x := *(this().GetObject().(*uint32))
				uint64_value = uint64(x)
			case "uint32":
				x := (this().GetObject().(uint32))
				uint64_value = uint64(x)
			case "*uint16":
				x := *(this().GetObject().(*uint16))
				uint64_value = uint64(x)
			case "uint16":
				x := (this().GetObject().(uint16))
				uint64_value = uint64(x)
			case "*uint8":
				x := *(this().GetObject().(*uint8))
				uint64_value = uint64(x)
			case "uint8":
				x := (this().GetObject().(uint8))
				uint64_value = uint64(x)
			case "*string":
				string_value := (this().GetObject().(*string))
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
			case "string":
				string_value := (this().GetObject().(string))
				if string_value == "NULL" {
					return nil, nil
				} else {
					value, value_error := strconv.ParseUint(string_value, 10, 64)
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
		},
		GetTimeWithDecimalPlaces: func(decimal_places int) (*time.Time, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			return common.GetTimeWithDecimalPlaces(this().GetObject(), decimal_places)
		},
		GetTime: func() (*time.Time, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
		
			return common.GetTime(this().GetObject())
		},
		GetType: func() string {
			return common.GetType(this().GetObject())
		},
		IsInteger: func() bool {
			return common.IsInteger(this().GetObject())
		},
		IsBoolTrue: func() bool {
			return common.IsBoolTrue(this().GetObject())
		},
		IsBoolFalse: func() bool {
			return common.IsBoolFalse(this().GetObject())
		},
		GetInt8: func() (*int8, []error) {
			if common.IsNil(this().GetObject()) {
				return nil, nil
			}
			
			var errors []error
			int64_value, int64_value_errors := this().GetInt64()
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
		},
		GetInt8Value: func() (int8, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetInt64()
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
		},
		GetUInt8: func() (*uint8, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetUInt64()
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
		},
		GetUInt8Value: func() (uint8, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetUInt64()
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
		},
		GetInt16: func() (*int16, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetInt64()
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
		},
		GetInt16Value: func() (int16, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetInt64()
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
		},
		GetUInt16: func() (*uint16, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetUInt64()
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
		},
		GetUInt16Value: func() (uint16, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetUInt64()
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
		},
		GetInt32: func() (*int32, []error) {
			if common.IsNil(this().GetObject()) {	
				return nil, nil
			}
			
			var errors []error
			int64_value, int64_value_errors := this().GetInt64()
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
		},
		GetInt32Value: func() (int32, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetInt64()
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
		},
		GetInt64Value: func() (int64, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetInt64()
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
		},
		GetUInt32: func() (*uint32, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetUInt64()
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
		},
		GetUInt32Value: func() (uint32, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetUInt64()
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
		},
		GetUInt64Value: func() (uint64, []error) {
			var errors []error
			int64_value, int64_value_errors := this().GetUInt64()
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
		},
		GetInt: func() (*int, []error) {
			var errors []error
			var result int
			bit_size := strconv.IntSize
			if bit_size == 32 {
				temp_value, temp_value_errors := this().GetInt32()
				if temp_value_errors != nil {
					return nil, temp_value_errors
				} else if temp_value == nil {
					return nil, nil
				}
		
				result = int(*temp_value)
			} else if bit_size == 64 {
				temp_value, temp_value_errors := this().GetInt64()
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
		},
		GetIntFromMap: func(key string) (*int, []error) {
			var errors []error
			map_value, map_value_errors := this().GetMap()
			if map_value_errors != nil {
				errors = append(errors, map_value_errors...)
			} else if common.IsNil(map_value) {
				errors = append(errors, fmt.Errorf("Value.GetIntFromMap map is nil"))
			}
			if len(errors) > 0 {
				return nil, errors
			}
			return map_value.GetInt(key)
		},
		SetIntToMap: func(key string, int_value *int) ([]error) {
			var errors []error
			map_value, map_value_errors := this().GetMap()
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
		},
		SetIntValueToMap: func(key string, int_value int) ([]error) {
			var errors []error
			map_value, map_value_errors := this().GetMap()
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
		},
		GetIntValue: func() (int, []error) {
			var errors []error
			int_value, int_value_errors := this().GetInt()
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
		},
		AppendValue: func(add *Value) []error {
			var errors []error
			if common.IsArray(this().GetObject()) {	
				array, array_errors := this().GetArray()
				if array_errors != nil {
					errors = append(errors, array_errors...)
				} else if common.IsNil(array) {
					errors = append(errors, fmt.Errorf("Value.AppendValue array is nil"))
				} else {
					(*array).AppendValue(add)
				}
			} else {
				errors = append(errors, fmt.Errorf("Value.AppendValue not supported for type %s", this().GetType()))
			}

			if len(errors) > 0 {
				return errors
			}

			return nil
		},
		AppendValueValue: func(add Value) []error {
			var errors []error
			if common.IsArray(this().GetObject()) {	
				array, array_errors := this().GetArray()
				if array_errors != nil {
					errors = append(errors, array_errors...)
				} else if common.IsNil(array) {
					errors = append(errors, fmt.Errorf("Value.AppendValue array is nil"))
				} else {
					(*array).AppendValueValue(add)
				}
			} else {
				errors = append(errors, fmt.Errorf("Value.AppendValue not supported for type %s", this().GetType()))
			}
		
			if len(errors) > 0 {
				return errors
			}
		
			return nil
		},
		SetObject: func(value interface{}) () {
			setObject(value)
		},
		GetObject: func() (interface{}) {
			return getObject()
		},
	}
	set_this_value( &created_value)
	return &created_value
}













