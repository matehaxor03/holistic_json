package json

import (
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
	"time"
	"strings"
	"strconv"
	"net/url"
	"os/exec"
)

type Value struct {

	GetMap func() (*Map, []error) 
	SetMap func(*Map)
	SetMapValue func(Map)
	GetMapValue func() (Map, []error) 
	IsMap func() (bool) 
	IsEmptyString func() bool
	GetFunc func() (*func(string) []error, []error)
	GetErrors func() ([]error, []error)
	SetErrors func(*[]error)
	SetErrorsValue func([]error)
	SetError func(*error)
	SetErrorValue func(error)
	GetArray func() (*Array, []error)
	GetArrayValue func() (Array, []error)
	SetArray func(*Array)
	SetArrayValue func(Array)
	
	
	GetBool func() (*bool, []error) 
	GetBoolValue func() (bool, []error) 
	SetBool func(*bool)
	SetBoolValue func(bool)
	IsArray func() bool
	GetType func() string
	GetObject func() (interface{})
	SetObject func(value interface{})

	IsString func() bool
	IsValueEqualToStringValue func(s string) bool
	IsValueEqualToString func(s *string) bool
	GetStringValue func() (string, []error) 
	GetString func() (*string, []error) 
	SetStringValue func(string)
	SetString func(*string)

	IsNumber func() bool
	IsFloat func() bool
	IsFloat32 func() bool
	IsFloat64 func() bool
	IsBool func() bool
	IsBoolTrue func() bool
	IsBoolFalse func() bool
	IsInteger func() bool
	IsInt func() bool
	IsInt8 func() bool
	IsInt16 func() bool
	IsInt32 func() bool
	IsInt64 func() bool
	IsUInt func() bool
	IsUInt8 func() bool
	IsUInt16 func() bool
	IsUInt32 func() bool
	IsUInt64 func() bool
	
	SetNil func()
	GetFloat32 func() (*float32, []error) 
	GetFloat32Value func() (float32, []error) 
	SetFloat32 func(*float32)
	SetFloat32Value func(float32)
	GetFloat64 func() (*float64, []error) 
	GetFloat64Value func() (float64, []error) 
	SetFloat64 func(*float64)
	SetFloat64Value func(float64)
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
	//
	SetInt func(*int)
	SetIntValue func(int)
	SetInt8 func(*int8)
	SetInt8Value func(int8)
	SetInt16 func(*int16)
	SetInt16Value func(int16)
	SetInt32 func(*int32)
	SetInt32Value func(int32)
	SetInt64 func(*int64)
	SetInt64Value func(int64)
	SetUInt func(*uint)
	SetUIntValue func(uint)
	SetUInt8 func(*uint8)
	SetUInt8Value func(uint8)
	SetUInt16 func(*uint16)
	SetUInt16Value func(uint16)
	SetUInt32 func(*uint32)
	SetUInt32Value func(uint32)
	SetUInt64 func(*uint64)
	SetUInt64Value func(uint64)

	GetRunes func() (*[]*rune, []error)
	GetTimeWithDecimalPlaces func(decimal_places int) (*time.Time, []error)
	GetTime func() (*time.Time, []error)
	GetTimeValue func() (time.Time, []error)
	SetTime func(*time.Time)
	SetTimeValue func(time.Time)
	IsNil func() bool
	Clone func() (*Value, []error)
}

func NewValueValue(v interface{}) (Value) {
	return *NewValue(v)
}


func NewValue(v interface{}) (*Value) {
	var this_value *Value
	internal_value := v

	set_this := func(value *Value) {
		this_value = value
	}
	
	this := func() *Value {
		return this_value
	}

	getObject := func() interface{} {
		return internal_value
	}

	setObject := func(value interface{}) {
		internal_value = value
	}

	clone := func() (*Value, []error) {
		var errors []error
		temp_value := this()
		cloned_value := NewValue(nil)
		if temp_value.IsMap() {
			temp_map, temp_map_errors := temp_value.GetMap()
			if temp_map_errors != nil {
				errors = append(errors, temp_map_errors...)
			} else if common.IsNil(temp_map) {
				errors = append(errors, fmt.Errorf("Value.Clone map is nil"))
			} else {
				cloned_map, cloned_map_errors := temp_map.Clone()
				if cloned_map_errors != nil {
					errors = append(errors, cloned_map_errors...)
				} else if common.IsNil(cloned_map) {
					errors = append(errors, fmt.Errorf("Value.Clone cloned map is nil"))
				} else {
					cloned_value.SetMap(cloned_map)
				}
			}
		} else if temp_value.IsArray() {
			temp_array, temp_array_errors := temp_value.GetArray()
			if temp_array_errors != nil {
				errors = append(errors, temp_array_errors...)
			} else if common.IsNil(temp_array) {
				errors = append(errors, fmt.Errorf("Value.Clone array is nil"))
			} else {
				cloned_array, cloned_array_errors := temp_array.Clone()
				if cloned_array_errors != nil {
					errors = append(errors, cloned_array_errors...)
				} else if common.IsNil(cloned_array) {
					errors = append(errors, fmt.Errorf("Value.Clone cloned array is nil"))
				} else {
					cloned_value.SetArray(cloned_array)
				}
			}
		} else {
			rep := temp_value.GetType()
			temp_object := temp_value.GetObject()
			switch rep {
				case "nil":
					cloned_value.SetNil()
				case "string":
					result := temp_object.(string)
					cloned_value.SetStringValue(result)
				case "*string":
					result := *(temp_object.(*string))
					cloned_value.SetString(&result)
				case "error":
					result := fmt.Sprintf("%s",  temp_object.(error).Error())
					cloned_value.SetErrorValue(fmt.Errorf(result))
				case "*error":
					result := fmt.Sprintf("%s", (*(temp_object.(*error))).Error())
					cloned_error := fmt.Errorf(result)
					cloned_value.SetError(&cloned_error)
				case "*url.Error":
					result := fmt.Sprintf("%s", (*(temp_object.(*url.Error))).Error()) 
					cloned_error := fmt.Errorf(result)
					cloned_value.SetError(&cloned_error)
				case "exec.ExitError":
					result := fmt.Sprintf("%s", temp_object.(exec.ExitError))
					cloned_error := fmt.Errorf(result)
					cloned_value.SetErrorValue(cloned_error)
				case "*exec.ExitError":
					result := fmt.Sprintf("%s", *(temp_object.(*exec.ExitError)))
					cloned_error := fmt.Errorf(result)
					cloned_value.SetError(&cloned_error)
				case "errors.errorString":
					result := fmt.Sprintf("%s", temp_object)
					cloned_error := fmt.Errorf(result)
					cloned_value.SetErrorValue(cloned_error)
				case "*errors.errorString":
					result := fmt.Sprintf("%s", temp_object)
					cloned_error := fmt.Errorf(result)
					cloned_value.SetError(&cloned_error)
				case "bool":
					if temp_object.(bool) {
						cloned_value.SetBoolValue(true)
					} else {
						cloned_value.SetBoolValue(false)
					}
				case "*bool":
					if *(temp_object.(*bool)) {
						cloned_bool := true
						cloned_value.SetBool(&cloned_bool)
					} else {
						cloned_bool := false
						cloned_value.SetBool(&cloned_bool)
					}
				case "*time.Time":
					result := (*(temp_object.(*time.Time)))
					cloned_time, cloned_time_errors := common.GetTime(result)
					if cloned_time_errors != nil {
						errors = append(errors, cloned_time_errors...)
					} else if common.IsNil(cloned_time) {
						errors = append(errors, fmt.Errorf("Value.Clone time is nil"))
					} else {
						cloned_value.SetTime(cloned_time)
					}
				case "time.Time":
					result := ((temp_object.(time.Time)))
					cloned_time, cloned_time_errors := common.GetTime(result)
					if cloned_time_errors != nil {
						errors = append(errors, cloned_time_errors...)
					} else if common.IsNil(cloned_time) {
						errors = append(errors, fmt.Errorf("Value.Clone time is nil"))
					} else {
						cloned_value.SetTime(cloned_time)
					}
				case "map[string]map[string][][]string":
					result := "map[string]map[string][][]string"
					cloned_value.SetStringValue(result)
				case "*uint64":
					result := *(temp_object.(*uint64))
					cloned_value.SetUInt64(&result)
				case "uint64":
					result := (temp_object.(uint64))
					cloned_value.SetUInt64Value(result)
				case "*uint32":
					result := *(temp_object.(*uint32))
					cloned_value.SetUInt32(&result)
				case "uint32":
					result := (temp_object.(uint32))
					cloned_value.SetUInt32Value(result)
				case "*uint16":
					result := *(temp_object.(*uint16))
					cloned_value.SetUInt16(&result)
				case "uint16":
					result := (temp_object.(uint16))
					cloned_value.SetUInt16Value(result)
				case "*uint8":
					result := *(temp_object.(*uint8))
					cloned_value.SetUInt8(&result)
				case "uint8":
					result := (temp_object.(uint8))
					cloned_value.SetUInt8Value(result)
				case "*uint":
					result := *(temp_object.(*uint))
					cloned_value.SetUInt(&result)
				case "uint":
					result := (temp_object.(uint))
					cloned_value.SetUIntValue(result)
				case "*int64":
					result := *(temp_object.(*int64))
					cloned_value.SetInt64(&result)
				case "int64":
					result := (temp_object.(int64))
					cloned_value.SetInt64Value(result)
				case "*int32":
					result := *(temp_object.(*int32))
					cloned_value.SetInt32(&result)
				case "int32":
					result := (temp_object.(int32))
					cloned_value.SetInt32Value(result)
				case "*int16":
					result := *(temp_object.(*int16))
					cloned_value.SetInt16(&result)
				case "int16":
					result := (temp_object.(int16))
					cloned_value.SetInt16Value(result)
				case "*int8":
					result := *(temp_object.(*int8))
					cloned_value.SetInt8(&result)
				case "int8":
					result := (temp_object.(int8))
					cloned_value.SetInt8Value(result)
				case "*int":
					result := *(temp_object.(*int))
					cloned_value.SetInt(&result)
				case "int":
					result := (temp_object.(int))
					cloned_value.SetIntValue(result)
				case "*float64":
					result := *(temp_object.(*float64))
					cloned_value.SetFloat64(&result)
				case "float64":
					result := (temp_object.(float64))
					cloned_value.SetFloat64Value(result)
				case "*float32":
					result := *(temp_object.(*float32))
					cloned_value.SetFloat32(&result)
				case "float32":
					result := (temp_object.(float32))
					cloned_value.SetFloat32Value(result)
				default:
					errors = append(errors, fmt.Errorf("error: Value.Clone: type %s is not supported please implement", rep))
			}
			
		}
		

		if len(errors) > 0 {
			return nil, errors
		}

		return cloned_value, nil
	}
	
	created_value := Value{
		SetMap: func(value *Map) {
			internal_value = value
		},
		SetMapValue: func(value Map) {
			internal_value = value
		},
		SetArray: func(value *Array) {
			internal_value = value
		},
		SetArrayValue: func(value Array) {
			internal_value = value
		},
		IsMap: func() (bool) {
			return common.IsMap(this().GetObject())
		},
		IsEmptyString: func() bool {
			if this().IsNil() {
				return false
			}

			if !common.IsString(this().GetObject()) {
				return false
			}
		
			string_value, string_value_errors := this().GetString()
			if string_value_errors != nil{
				return false
			} else if this().IsNil() {
				return false
			}
		
			return *string_value == ""
		},
		GetMap: func() (*Map, []error) {
			if this().IsNil() { 
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
		GetMapValue: func() (Map, []error) {
			var errors []error
			if this().IsNil() { 
				errors = append(errors, fmt.Errorf("map is nil"))
				return  NewMapValue(), errors
			}
	
			var result Map
			type_of := this().GetType()
			if type_of == "*json.Map" {
				result = (*this().GetObject().(*Map))
			} else if type_of == "json.Map" {
				result = this().GetObject().(Map)
			} else {
				errors = append(errors, fmt.Errorf("%s failed to unbox to json.Map", type_of))
				return NewMapValue(), errors
			}
		
			return result, nil
		},
		GetFunc: func() (*func(string) []error, []error) {
			if this().IsNil() {
				return nil, nil
			}
		
			var errors []error
			var result (func(string) []error)
			rep := common.GetType(getObject())
			switch rep {
			case "func(string) []error":
				result = (this().GetObject()).(func(string) []error)
			case "*func(string) []error":
				result = *(this().GetObject().(*func(string) []error))
			default:
				errors = append(errors, fmt.Errorf("error: Map.Func: type %s is not supported please implement", rep))
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
		
			return &result, nil
		},
		GetErrors: func() ([]error, []error) {
			if this().IsNil() {
				return nil, nil
			}
		
			var errors []error
			var result []error
			
			rep := common.GetType(getObject())
			switch rep {
			case "*[]error":
				result = *(this().GetObject().(*[]error))
			case "[]error":
				result = this().GetObject().([]error)
			case "*[]string":
				string_array := this().GetObject().(*[]string)
				for _, string_array_value := range *string_array {
					result = append(result, fmt.Errorf("%s", string_array_value))
				}
			case "[]string":
				string_array := this().GetObject().([]string)
				for _, string_array_value := range string_array {
					result = append(result, fmt.Errorf("%s", string_array_value))
				}
			case "json.Array":
				string_array := this().GetObject().(Array)
				for _, string_array_value := range *(string_array.GetValues()) {
					converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
					if converted_errors != nil {
						errors = append(errors, converted_errors...)
					} else {
						result = append(result, fmt.Errorf("%s", converted))
					}
				}
			case "*json.Array":
				string_array := this().GetObject().(*Array)
				for _, string_array_value := range *(string_array.GetValues()) {
					converted, converted_errors := ConvertInterfaceValueToStringValue(string_array_value)
					if converted_errors != nil {
						errors = append(errors, converted_errors...)
					} else {
						result = append(result, fmt.Errorf("%s", converted))
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
			if this().IsNil() {
				return nil, nil
			}
		
			var errors []error
			var result *Array
		
			rep := common.GetType(getObject())
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
				return NewArrayValue(), array_errors
			} else if this().IsNil() {
				var errors []error
				errors = append(errors, fmt.Errorf("Value.GetArrayValue array is nil"))
				return NewArrayValue(), errors
			}

			return *array, nil
		},
		GetString: func() (*string, []error) {
			if this().IsNil() {
				return nil, nil
			}
		
			var errors []error
			var result *string
			rep := common.GetType(getObject())
			switch rep {
			case "string":
				value := this().GetObject().(string)
				result = &value
			case "*string":
				result = this().GetObject().(*string)
			default:
				panic("debug")
				errors = append(errors, fmt.Errorf("error Value.GetString: type %s is not supported please implement for attribute", rep))
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
		
			return result, nil
		},
		SetString: func(value *string) {
			setObject(value)
		},
		SetStringValue: func(value string) {
			setObject(value)
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
			return common.IsFloat(this().GetObject())
		},
		IsFloat32: func() (bool) {
			return common.IsFloat32(this().GetObject())
		},
		IsFloat64: func() (bool) {
			return common.IsFloat64(this().GetObject())
		},
		IsNumber: func() (bool) {
			return common.IsNumber(this().GetObject())
		},
		IsString: func() (bool) {
			return common.IsString(this().GetObject())
		},
		IsValueEqualToString: func(s *string) (bool) {
			this_nil := this().IsNil()
			compare_nil := common.IsNil(s)
			
			if this_nil && compare_nil {
				return true
			}
			
			if this_nil != compare_nil {
				return false
			}

			value, value_errors := this().GetStringValue()
			if value_errors != nil {
				return false
			}

			return value == *s			
		},
		IsValueEqualToStringValue: func(s string) (bool) {
			this_nil := this().IsNil()
			compare_nil := common.IsNil(s)
			
			if this_nil && compare_nil {
				return true
			}
			
			if this_nil != compare_nil {
				return false
			}

			value, value_errors := this().GetStringValue()
			if value_errors != nil {
				return false
			}

			return value == s			
		},
		IsNil: func() bool {
			temp_value := this().GetObject()
			if common.IsNil(temp_value) {
				return true
			}
			
			return !(common.IsValue(temp_value) ||
					common.IsNumber(temp_value) || 
					common.IsBool(temp_value) || 
					common.IsString(temp_value) || 
					common.IsArray(temp_value) || 
					common.IsMap(temp_value) ||
					common.IsFunc(temp_value))
		},
		IsBool: func() bool {
			return common.IsBool(this().GetObject())
		},
		IsArray: func() bool {
			return common.IsArray(this().GetObject())
		},
		GetFloat64: func() (*float64, []error) {
			if this().IsNil() {
				return nil, nil
			}
		
			var errors []error
			var result *float64
			rep := common.GetType(getObject())
			switch rep {
			case "float64":
				value := this().GetObject().(float64)
				result = &value
			case "*float64":
				value := *(this().GetObject().(*float64))
				result = &value
			case "float32":
				value := float64(this().GetObject().(float32))
				result = &value
			case "*float32":
				value := float64(*(this().GetObject().(*float32)))
				result = &value
			case "int":
				int_value := this().GetObject().(int)
				float_value := float64(int_value)
				result = &float_value
			case "*int":
				int_value := *(this().GetObject().(*int))
				float_value := float64(int_value)
				result = &float_value
			case "int8":
				int_value := this().GetObject().(int8)
				float_value := float64(int_value)
				result = &float_value
			case "*int8":
				int_value := *(this().GetObject().(*int8))
				float_value := float64(int_value)
				result = &float_value
			case "int16":
				int_value := this().GetObject().(int16)
				float_value := float64(int_value)
				result = &float_value
			case "*int16":
				int_value := *(this().GetObject().(*int16))
				float_value := float64(int_value)
				result = &float_value
			case "int32":
				int_value := this().GetObject().(int32)
				float_value := float64(int_value)
				result = &float_value
			case "*int32":
				int_value := *(this().GetObject().(*int32))
				float_value := float64(int_value)
				result = &float_value
			case "int64":
				int_value := this().GetObject().(int64)
				float_value := float64(int_value)
				result = &float_value
			case "*int64":
				int_value := *(this().GetObject().(*int64))
				float_value := float64(int_value)
				result = &float_value
			case "uint":
				int_value := this().GetObject().(uint)
				float_value := float64(int_value)
				result = &float_value
			case "*uint":
				int_value := *(this().GetObject().(*uint))
				float_value := float64(int_value)
				result = &float_value
			case "uint8":
				int_value := this().GetObject().(uint8)
				float_value := float64(int_value)
				result = &float_value
			case "*uint8":
				int_value := *(this().GetObject().(*uint8))
				float_value := float64(int_value)
				result = &float_value
			case "uint16":
				int_value := this().GetObject().(uint16)
				float_value := float64(int_value)
				result = &float_value
			case "*uint16":
				int_value := *(this().GetObject().(*uint16))
				float_value := float64(int_value)
				result = &float_value
			case "uint32":
				int_value := this().GetObject().(uint32)
				float_value := float64(int_value)
				result = &float_value
			case "*uint32":
				int_value := *(this().GetObject().(*uint32))
				float_value := float64(int_value)
				result = &float_value
			case "uint64":
				int_value := this().GetObject().(uint64)
				float_value := float64(int_value)
				result = &float_value
			case "*uint64":
				int_value := *(this().GetObject().(*uint64))
				float_value := float64(int_value)
				result = &float_value
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
			default:
				errors = append(errors, fmt.Errorf("error: Value.GetFloat64: type %s is not supported please implement for attribute: %s", rep))
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
		
			return result, nil
		},
		SetFloat64: func(value *float64) {
			setObject(value)
		},
		SetFloat64Value: func(value float64) {
			setObject(value)
		},
		GetFloat32: func() (*float32, []error) {
			if this().IsNil() {
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
		SetFloat32: func(value *float32) {
			setObject(value)
		},
		SetFloat32Value: func(value float32) {
			setObject(value)
		},
		SetError: func(value *error) {
			setObject(value)
		},
		SetErrorValue: func(value error) {
			setObject(value)
		},
		SetErrors: func(value *[]error) {
			setObject(value)
		},
		SetErrorsValue: func(value []error) {
			setObject(value)
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
			if this().IsNil() {
				return nil, nil
			}
		
			var errors []error
			var result *string
			rep := common.GetType(getObject())
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
		SetBool: func(value *bool) {
			setObject(value)
		},
		SetBoolValue: func(value bool) {
			setObject(value)
		},
		GetBool: func() (*bool, []error) {
			if this().IsNil() {
				return nil, nil
			}
		
			var result *bool
			var errors []error
		
			rep := common.GetType(getObject())
			switch rep {
			case "bool":
				value := this().GetObject().(bool)
				result = &value
				break
			case "*bool":
				result = this().GetObject().(*bool)
			case "*string":
				value := *(this().GetObject().(*string))
				if value == "1" || value == "true"{
					boolean_result := true
					result = &boolean_result
				} else if value == "0" || value == "false" {
					boolean_result := false
					result = &boolean_result
				} else {
					errors = append(errors, fmt.Errorf("error: Map.GetBool: unknown value for *string: %s", value))
					result = nil
				}
			case "string":
				value := (this().GetObject().(string))
				if value == "1" || value == "true" {
					boolean_result := true
					result = &boolean_result
				} else if value == "0" || value == "false" {
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
		GetBoolValue: func() (bool, []error) {
			var errors []error
			if this().IsNil() {
				errors = append(errors, fmt.Errorf("Value.GetBoolValue value is nil"))
				return false, errors
			}
		
			var result bool
			rep := common.GetType(getObject())
			switch rep {
			case "bool":
				result = this().GetObject().(bool)
				break
			case "*bool":
				result = *(this().GetObject().(*bool))
				break
			case "*string":
				value := *(this().GetObject().(*string))
				if value == "1" || value == "true" {
					result = true
				} else if value == "0" || value == "false" {
					result = false
				} else {
					errors = append(errors, fmt.Errorf("error: Value.GetBoolValue: unknown value"))
					result = false
				}
			case "string":
				value := (this().GetObject().(string))
				if value == "1" || value == "true" {
					result = true
				} else if value == "0" || value == "false" {
					result = false
				} else {
					errors = append(errors, fmt.Errorf("error: Value.GetBoolValue: unknown value"))
					result = false
				}
			default:
				errors = append(errors, fmt.Errorf("error: Map.GetBool: type %s is not supported please implement", rep))
			}
		
			if len(errors) > 0 {
				return false, errors
			}
		
			return result, nil
		},
		SetNil: func() {
			setObject(nil)
		},
		GetInt64: func() (*int64, []error) {
			var errors []error
			var temp_value int64
		
			if this().IsNil() {
				return nil, nil
			}
		
			rep := common.GetType(getObject())
			switch rep {
				case "json.Value":
					string_value, convert_errors := ConvertInterfaceValueToStringValue(this().GetObject().(Value))
					if convert_errors != nil {
						errors = append(errors, convert_errors...)
					} else if common.IsNil(string_value) {
						return nil, nil
					} else {
						value, value_error := strconv.ParseInt(string_value, 10, 64)
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
						value, value_error := strconv.ParseInt(string_value, 10, 64)
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
		SetInt64: func(value *int64) {
			setObject(value)
		},
		SetInt64Value: func(value int64) {
			setObject(value)
		},
		SetInt32: func(value *int32) {
			setObject(value)
		},
		SetInt32Value: func(value int32) {
			setObject(value)
		},
		SetInt16: func(value *int16) {
			setObject(value)
		},
		SetInt16Value: func(value int16) {
			setObject(value)
		},
		SetInt8: func(value *int8) {
			setObject(value)
		},
		SetInt8Value: func(value int8) {
			setObject(value)
		},
		SetInt: func(value *int) {
			setObject(value)
		},
		SetIntValue: func(value int) {
			setObject(value)
		},
		SetUInt64: func(value *uint64) {
			setObject(value)
		},
		SetUInt64Value: func(value uint64) {
			setObject(value)
		},
		SetUInt32: func(value *uint32) {
			setObject(value)
		},
		SetUInt32Value: func(value uint32) {
			setObject(value)
		},
		SetUInt16: func(value *uint16) {
			setObject(value)
		},
		SetUInt16Value: func(value uint16) {
			setObject(value)
		},
		SetUInt8: func(value *uint8) {
			setObject(value)
		},
		SetUInt8Value: func(value uint8) {
			setObject(value)
		},
		SetUInt: func(value *uint) {
			setObject(value)
		},
		SetUIntValue: func(value uint) {
			setObject(value)
		},
		GetUInt64: func() (*uint64, []error) {
			var errors []error
			if this().IsNil() {
				return nil, nil
			}

			var uint64_value uint64
			rep := common.GetType(getObject())
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
			if this().IsNil() {
				return nil, nil
			}
		
			return common.GetTimeWithDecimalPlaces(this().GetObject(), decimal_places)
		},
		GetTime: func() (*time.Time, []error) {
			if this().IsNil() {
				return nil, nil
			}
		
			return common.GetTime(this().GetObject())
		},
		GetTimeValue: func() (time.Time, []error) {
			temp_time, temp_time_errors := common.GetTime(this().GetObject())
			if temp_time_errors != nil {
				return (*common.GetTimeZero()), temp_time_errors
			} else if common.IsNil(temp_time) {
				return (*common.GetTimeZero()), temp_time_errors
			}
			return *temp_time, nil
		},
		SetTime: func(value *time.Time) {
			setObject(value)
		},
		SetTimeValue: func(value time.Time) {
			setObject(value)
		},
		GetType: func() string {
			return common.GetType(this().GetObject())
		},
		IsInteger: func() bool {
			return common.IsInteger(this().GetObject())
		},
		IsInt: func() bool {
			return common.IsInt(this().GetObject())
		},
		IsInt8: func() bool {
			return common.IsInt8(this().GetObject())
		},
		IsInt16: func() bool {
			return common.IsInt16(this().GetObject())
		},
		IsInt32: func() bool {
			return common.IsInt32(this().GetObject())
		},
		IsInt64: func() bool {
			return common.IsInt64(this().GetObject())
		},
		IsUInt: func() bool {
			return common.IsUInt(this().GetObject())
		},
		IsUInt8: func() bool {
			return common.IsUInt8(this().GetObject())
		},
		IsUInt16: func() bool {
			return common.IsUInt16(this().GetObject())
		},
		IsUInt32: func() bool {
			return common.IsUInt32(this().GetObject())
		},
		IsUInt64: func() bool {
			return common.IsUInt64(this().GetObject())
		},
		IsBoolTrue: func() bool {
			return common.IsBoolTrue(this().GetObject())
		},
		IsBoolFalse: func() bool {
			return common.IsBoolFalse(this().GetObject())
		},
		GetInt8: func() (*int8, []error) {
			if this().IsNil() {
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
			if this().IsNil() {	
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
		SetObject: func(value interface{}) () {
			setObject(value)
		},
		GetObject: func() (interface{}) {
			return getObject()
		},
		Clone: func() (*Value, []error) {
			return clone()
		},
	}
	set_this( &created_value)
	return &created_value
}













