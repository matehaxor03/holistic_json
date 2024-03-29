package json

import (
	"fmt"
	"strings"
	"time"
	common "github.com/matehaxor03/holistic_common/common"
)

type Map struct {
	GetKeys func() ([]string)
	HasKey func(s string) (bool)
	GetMap func(s string) (*Map, []error)
	GetMapValue func(s string) (Map, []error)
	Values func() *Array
	SetMap func(s string, zap *Map) 
	SetMapValue func(s string, zap Map)
	SetValue func(s string, value *Value) 
	GetValue func(s string) (*Value)

	//GetValue func(s string) (*Value, []error) 

	Clone func() (*Map, []error)


	RemoveKey func(key string) (bool, error)

	IsNil func() bool
	IsNull func(s string) bool
	IsBool func(s string) bool
	
	IsEmptyString func(s string) bool
	IsInteger func(s string) bool 
	IsInt func(s string) bool 
	IsInt8 func(s string) bool 
	IsInt16 func(s string) bool 
	IsInt32 func(s string) bool 
	IsInt64 func(s string) bool 
	IsUInt func(s string) bool 
	IsUInt8 func(s string) bool 
	IsUInt16 func(s string) bool 
	IsUInt32 func(s string) bool 
	IsUInt64 func(s string) bool 
	IsString func(s string) bool
	IsMap func(s string) bool
	IsBoolTrue func(s string) bool
	IsBoolFalse func(s string) bool
	ToJSONString func(json *strings.Builder) ([]error)
	
	SetErrors func(s string, errors []error)
	GetErrors func(s string) ([]error, []error)
	GetType func(s string) string
	GetFunc func(s string) (*func(string) []error, []error)
	SetFunc func(s string, function *func(string) []error)
	
	
	SetStringValue func(s string, v string)
	SetString func(s string, v *string)
	GetString func(s string) (*string, []error)
	GetStringValue func(s string) (string, []error)
	
	SetBoolValue func(s string, v bool)
	SetBool func(s string, v *bool)
	GetBool func(s string) (*bool, []error)
	GetBoolValue func(s string) (bool, []error)
	
	SetNil func(s string)
	
	IsFloat func(s string) bool
	IsFloat32 func(s string) bool
	IsFloat64 func(s string) bool
	SetFloat32Value func(s string, v float32)
	SetFloat32 func(s string, v *float32)
	GetFloat32 func(s string) (*float32, []error)
	GetFloat32Value func(s string) (float32, []error)
	SetFloat64Value func(s string, v float64)
	SetFloat64 func(s string, v *float64)
	GetFloat64 func(s string) (*float64, []error)
	GetFloat64Value func(s string) (float64, []error)
	
	GetIntValue func(s string) (int, []error)
	SetIntValue func(s string, v int)
	SetInt func(s string, v *int)
	GetInt func(s string) (*int, []error)
	GetInt8Value func(s string) (int8, []error)
	SetInt8Value func(s string, v int8)
	SetInt8 func(s string, v *int8)
	GetInt8 func(s string) (*int8, []error)
	GetInt16Value func(s string) (int16, []error)
	SetInt16Value func(s string, v int16)
	SetInt16 func(s string, v *int16)
	GetInt16 func(s string) (*int16, []error)
	GetInt32Value func(s string) (int32, []error)
	SetInt32Value func(s string, v int32)
	SetInt32 func(s string, v *int32)
	GetInt32 func(s string) (*int32, []error)
	GetInt64Value func(s string) (int64, []error)
	SetInt64Value func(s string, v int64)
	SetInt64 func(s string, v *int64)
	GetInt64 func(s string) (*int64, []error)

	GetUIntValue func(s string) (uint, []error)
	SetUIntValue func(s string, v uint)
	SetUInt func(s string, v *uint)
	GetUInt func(s string) (*uint, []error)
	GetUInt8Value func(s string) (uint8, []error)
	SetUInt8Value func(s string, v uint8)
	SetUInt8 func(s string, v *uint8)
	GetUInt8 func(s string) (*uint8, []error)
	GetUInt16Value func(s string) (uint16, []error)
	SetUInt16Value func(s string, v uint16)
	SetUInt16 func(s string, v *uint16)
	GetUInt16 func(s string) (*uint16, []error)
	GetUInt32Value func(s string) (uint32, []error)
	SetUInt32Value func(s string, v uint32)
	SetUInt32 func(s string, v *uint32)
	GetUInt32 func(s string) (*uint32, []error)
	GetUInt64Value func(s string) (uint64, []error)
	SetUInt64Value func(s string, v uint64)
	SetUInt64 func(s string, v *uint64)
	GetUInt64 func(s string) (*uint64, []error)

	IsArray func(s string) bool
	SetArray func(s string, array *Array)
	SetArrayValue func(s string, array Array)

	GetArray func(s string) ((*Array), []error)
	GetArrayValue func(s string) ((Array), []error)
	GetArrayOfInt func(s string) (*[]*int, []error)
	GetArrayOfInt8 func(s string) (*[]*int8, []error)
	GetArrayOfInt16 func(s string) (*[]*int16, []error)
	GetArrayOfInt32 func(s string) (*[]*int32, []error)
	GetArrayOfInt64 func(s string) (*[]*int64, []error)
	
	GetArrayOfUInt func(s string) (*[]*uint, []error)
	GetArrayOfUInt8 func(s string) (*[]*uint8, []error)
	GetArrayOfUInt16 func(s string) (*[]*uint16, []error)
	GetArrayOfUInt32 func(s string) (*[]*uint32, []error)
	GetArrayOfUInt64 func(s string) (*[]*uint64, []error)

	GetArrayOfFloat32 func(s string) (*[]*float32, []error)
	GetArrayOfFloat32Value func(s string) ([]float32, []error)
	GetArrayOfFloat64 func(s string) (*[]*float64, []error)
	GetArrayOfFloat64Value func(s string) ([]float64, []error)

	GetArrayOfBool func(s string) (*[]*bool, []error)
	GetArrayOfBoolValue func(s string) ([]bool, []error)

	GetArrayOfString func(s string) (*[]*string, []error)

	GetRunes func(s string) (*[]*rune, []error)

	GetObjectForMap func(s string) (interface{})
	SetObjectForMap func(s string, object interface{}) 

	SetTime func (s string, value *time.Time)
	GetTime func (s string) (*time.Time, []error)
	GetTimeWithDecimalPlaces func(s string, decimal_places int) (*time.Time, []error)

}

func NewMapValue() Map {
	var temp_map = NewMap()
	return *temp_map
}

func NewMap() *Map {
	return NewMapOfValues(nil)
}

func NewMapOfValues(m *map[string]interface{}) *Map {
	internal_map := make(map[string]*Value)
	internal_keys_obj := NewArray()
	internal_keys := []string{}

	findIndexOfKey := func(key string) int {
		for i, k := range internal_keys {
			if k == key {
				return i
			}
		}
		return -1
	}

	removeKey := func(key string) (bool, error) {
		i := findIndexOfKey(key)
		if i == -1 {
			return false, fmt.Errorf("key not found in array keys")
		}

		internal_keys = append(internal_keys[:i], internal_keys[i+1:]...)
		internal_keys_obj.RemoveValueAtIndex(i)
		return true, nil
	}

	getKeys := func() ([]string) {
		return internal_keys
	}

	hasKey := func(s string) (bool) {
		_, found := internal_map[s]
		return found
	}

	isValueNilForMap := func(s string) (bool) {
		if internal_map == nil {
			return true
		}

		if !hasKey(s) {
			return true
		}
		
		value, found := internal_map[s]
		if !found {
			return true
		}

		if common.IsNil(value) {
			return true
		}

		type_of := common.GetType(value)
		if type_of == "*json.Value" {
			temp_object := value.GetObject()
			if common.IsNil(temp_object) {
				return true
			}
		} else if type_of == "json.Value" {
			temp_object := value.GetObject()
			if common.IsNil(temp_object) {
				return true
			}
		}

		return false
	}

	setInternalValue := func(s string, v *Value) {
		if found := hasKey(s); !found || len(getKeys()) == 0 {
			internal_keys = append(internal_keys, s)
			internal_keys_obj.AppendValue(v)
		} else if index := findIndexOfKey(s); index > -1 {
			internal_keys_obj.SetValueAtIndex(index, v)
		}
		internal_map[s] = v
	}

	getInternalValue := func(s string) *Value {
		if !hasKey(s) {
			return nil
		}
		return internal_map[s]
	}

	getArray := func(s string) (*Array, []error) {
		if isValueNilForMap(s) {
			return  nil, nil
		}
		return getInternalValue(s).GetArray()
	}

	getArrayValue := func(s string) (Array, []error) {
		array, array_errors := getArray(s)
		if array_errors != nil {
			return NewArrayValue(), array_errors
		} else if common.IsNil(array) {
			var errors []error
			errors = append(errors, fmt.Errorf("json.Map.GetArrayValue array is nil"))
			return  NewArrayValue(), errors
		}
		return *array, nil
	}

	getMap := func(s string) (*Map, []error) {
		value := getInternalValue(s)
		if common.IsNil(value) {
			return nil, nil
		}
		return value.GetMap()
	}

	getMapValue := func(s string) (Map, []error) {
		value := getInternalValue(s)
		if common.IsNil(value) {
			var errors []error
			errors = append(errors, fmt.Errorf("Map.getMapValue value is nil"))
			return NewMapValue(), errors
		}
		return value.GetMapValue()
	}

	isMap := func(s string) bool {
		if isValueNilForMap(s) {
			return false
		}
		return getInternalValue(s).IsMap()
	}

	isArray := func(s string) bool {
		if isValueNilForMap(s) {
			return false
		}
		return getInternalValue(s).IsArray()
	}

	getValue := func(s string) (*Value) {
		return getInternalValue(s)
	}

	toJSONString := func(json *strings.Builder) ([]error) {
		if json == nil {
			var errors []error
			errors = append(errors, fmt.Errorf("error: *strings.Builder is nil"))
			return errors
		}
		
		temp_keys := getKeys()
		if common.IsNil(temp_keys) {
			json.WriteString("{}")
			return nil
		}
		
		length := len(temp_keys)
		if length == 0 {
			json.WriteString("{}")
			return nil
		}

		var errors []error
		json.WriteString("{")
		for i := 0; i < length; i++ {
			key := (temp_keys)[i]
			key_clone := common.CloneString(&key)
			escaped_key, escaped_key_error := common.EscapeString(*key_clone, "\"")
			if escaped_key_error != nil {
				errors = append(errors, escaped_key_error)
				return errors
			}
			json.WriteString("\"")
			json.WriteString(escaped_key)
			json.WriteString("\":")
			string_conversion_errors := ConvertInterfaceValueToJSONStringValue(json, internal_map[key])
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

	clone := func() (*Map, []error) {
		var errors []error
		cloned := NewMap()
		keys := getKeys()
		for _, key := range keys {
			if isMap(key) {
				temp_map, temp_map_errors := getMap(key)
				if temp_map_errors != nil {
					errors = append(errors, temp_map_errors...)
				} else if common.IsNil(temp_map) {
					errors = append(errors, fmt.Errorf("Map.Clone map is nil"))
				} else {
					cloned_map, cloned_map_errors := temp_map.Clone()
					if cloned_map_errors != nil {
						errors = append(errors, cloned_map_errors...)
					} else if common.IsNil(cloned_map) {
						errors = append(errors, fmt.Errorf("Map.Clone cloned map is nil"))
					} else {
						cloned.SetMap(key, cloned_map)
					}
				}
			} else if isArray(key) {
				temp_array, temp_array_errors := getArray(key)
				if temp_array_errors != nil {
					errors = append(errors, temp_array_errors...)
				} else if common.IsNil(temp_array) {
					errors = append(errors, fmt.Errorf("Map.Clone array is nil"))
				} else {
					cloned_array, cloned_array_errors := temp_array.Clone()
					if cloned_array_errors != nil {
						errors = append(errors, cloned_array_errors...)
					} else if common.IsNil(cloned_array) {
						errors = append(errors, fmt.Errorf("Map.Clone cloned array is nil"))
					} else {
						cloned.SetArray(key, cloned_array)
					}
				}
			} else {
				temp_value := getValue(key)
				if common.IsNil(temp_value) {
					cloned.SetNil(key)
				} else {
					cloned_value, cloned_value_errors := temp_value.Clone()
					if cloned_value_errors != nil {
						errors = append(errors, cloned_value_errors...)
					} else if common.IsNil(cloned_value) {
						errors = append(errors, fmt.Errorf("Map.Clone cloned value is nil"))
					} else {
						cloned.SetValue(key, cloned_value)
					}
				}
			}
		}

		if len(errors) > 0 {
			return nil, errors
		}

		return cloned, nil
	}

	if !common.IsNil(m) {
		current_type := ""
		for key, value := range *m {
			current_type = common.GetType(value)
			if current_type == "json.Value" {
				temp_value := ((value).(Value))
				setInternalValue(key, &temp_value)
			} else if current_type == "*json.Value" {
				setInternalValue(key,((value).(*Value)))
			} else {
				setInternalValue(key, NewValue(value))
			}
		}
	}

	return &Map{
		GetMap: func(s string) (*Map, []error) {
			return getMap(s)
		},
		GetMapValue: func(s string) (Map, []error) {
			return getMapValue(s)
		},
		GetValue: func(s string) (*Value) {
			return getValue(s)
		},
		SetMap: func(s string, zap *Map) {
			set_map_value := NewValue(zap)
			setInternalValue(s, set_map_value)
		},
		SetMapValue: func(s string, zap Map) {
			set_map_value := NewValue(&zap)
			setInternalValue(s, set_map_value)
		},
		IsNull: func(s string) bool {
			return isValueNilForMap(s)
		},
		IsNil: func() bool {			
			if common.IsNil(internal_map) {
				return true
			}
	
			return !(common.IsValue(internal_map) ||
					common.IsNumber(internal_map) || 
					common.IsBool(internal_map) || 
					common.IsString(internal_map) || 
					common.IsArray(internal_map) || 
					common.IsMap(internal_map) ||
					common.IsFunc(internal_map))
		},
		IsBool: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsBool()
		},
		IsArray: func(s string) bool {
			return isArray(s)
		},
		IsEmptyString: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsEmptyString()
		},
		IsInteger: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsInteger()
		},
		IsInt: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsInt()
		},
		IsInt8: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsInt8()
		},
		IsInt16: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsInt16()
		},
		IsInt32: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsInt32()
		},
		IsInt64: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsInt64()
		},
		IsUInt: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsUInt()
		},
		IsUInt8: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsUInt8()
		},
		IsUInt16: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsUInt16()
		},
		IsUInt32: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsUInt32()
		},
		IsUInt64: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsUInt64()
		},
		IsString: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsString()
		},
		IsMap: func(s string) bool {
			return isMap(s)
		},
		IsBoolTrue: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsBoolTrue()
		},
		IsBoolFalse: func(s string) bool {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsBoolFalse()
		},
		ToJSONString: func(json_payload_builder *strings.Builder) ([]error) {
			return toJSONString(json_payload_builder)
		},
		GetKeys: func() ([]string) {
			return getKeys()
		},
		SetArray: func(s string, array *Array) {
			set_map_value := NewValue(array)
			setInternalValue(s, set_map_value)
		},
		SetArrayValue: func(s string, array Array) {
			set_map_value := NewValue(&array)
			setInternalValue(s, set_map_value)
		},
		SetErrors: func(s string, errors []error) {
			set_map_value := NewValue(errors)
			setInternalValue(s, set_map_value)
		},
		GetErrors: func(s string) ([]error, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetErrors()
		},
		GetType: func(s string) string {
			if isValueNilForMap(s) {
				return "nil"
			}
			return getInternalValue(s).GetType()
		},
		GetFunc: func(s string) (*func(string) []error, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetFunc()
		},
		SetFunc: func(s string, function *func(string) []error) {
			set_map_value := NewValue(function)
			setInternalValue(s, set_map_value)
		},
		GetArray: func(s string) (*Array, []error) {
			return getArray(s)
		},
		GetArrayValue: func(s string) (Array, []error) {
			return getArrayValue(s)
		},
		GetArrayOfInt8: func(s string) (*[]*int8, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt8()
		},
		GetArrayOfInt16: func(s string) (*[]*int16, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt16()
		},
		GetArrayOfInt32: func (s string) (*[]*int32, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt32()
		},
		GetArrayOfInt64: func(s string) (*[]*int64, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfInt64()
		},
		GetArrayOfUInt8: func(s string) (*[]*uint8, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt8()
		},
		GetArrayOfUInt16: func(s string) (*[]*uint16, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt16()
		},
		GetArrayOfUInt32: func(s string) (*[]*uint32, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt32()
		},
		GetArrayOfUInt64: func(s string) (*[]*uint64, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfUInt64()
		},
		GetArrayOfFloat32: func(s string) (*[]*float32, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfFloat32()
		},
		GetArrayOfFloat32Value: func(s string) ([]float32, []error) {
			array, array_errors := getArrayValue(s)
			if array_errors != nil {
				return nil, array_errors
			} 
			return array.GetArrayOfFloat32Value()
		},
		GetArrayOfFloat64: func(s string) (*[]*float64, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfFloat64()
		},
		GetArrayOfFloat64Value: func(s string) ([]float64, []error) {
			array, array_errors := getArrayValue(s)
			if array_errors != nil {
				return nil, array_errors
			} 
			return array.GetArrayOfFloat64Value()
		},
		GetArrayOfString: func(s string) (*[]*string, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfString()
		},
		GetArrayOfBool: func(s string) (*[]*bool, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfBool()
		},
		GetArrayOfBoolValue: func(s string) ([]bool, []error) {
			array, array_errors := getArray(s)
			if array_errors != nil {
				return nil, array_errors
			} else if common.IsNil(array) {
				return nil, nil
			}
			return array.GetArrayOfBoolValue()
		},
		GetString: func(s string) (*string, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetString()
		},
		IsFloat: func(s string) (bool) {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsFloat()
		},
		IsFloat32: func(s string) (bool) {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsFloat32()
		},
		IsFloat64: func(s string) (bool) {
			if isValueNilForMap(s) {
				return false
			}
			return getInternalValue(s).IsFloat64()
		},
		GetStringValue: func(s string) (string, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetStringValue value is nil"))
				return "", errors
			}
			return getInternalValue(s).GetStringValue()
		},
		GetFloat32: func(s string) (*float32, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetFloat32()
		},
		GetFloat32Value: func(s string) (float32, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetFloat32Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetFloat32Value()
		},
		GetFloat64Value: func(s string) (float64, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetFloat64Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetFloat64Value()
		},
		GetFloat64: func (s string) (*float64, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetFloat64()
		},
		GetRunes: func(s string) (*[]*rune, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetRunes()
		},
		GetObjectForMap: func(s string) interface{} {
			if isValueNilForMap(s) {
				return nil
			}
			return getInternalValue(s).GetObject()
		},
		SetObjectForMap: func(s string, value interface{}) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		GetBool: func(s string) (*bool, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetBool()
		},
		SetBool: func(s string, value *bool) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetBoolValue: func(s string, value bool) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetString: func(s string, value *string) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetStringValue: func(s string, value string) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetNil: func(s string) {
			set_map_value := NewValue(nil)
			setInternalValue(s, set_map_value)
		},
		HasKey: func(s string) (bool) {
			return hasKey(s)
		},
		RemoveKey: func(s string) (bool, error) {
			if !hasKey(s) {
				return false, fmt.Errorf("key not found in map")
			} else {
				delete(internal_map, s)
			}
			_, remove_key_error := removeKey(s)
			return true, remove_key_error
		},
		GetInt64: func(s string) (*int64, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetInt64()
		},
		GetInt8: func(s string) (*int8, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetInt8()
		},
		GetInt8Value: func(s string) (int8, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("Map.GetInt8Value was nil"))
				return 0, errors
			}
			return getInternalValue(s).GetInt8Value()
		},
		GetUInt8: func(s string) (*uint8, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetUInt8()
		},
		GetUInt8Value: func(s string) (uint8, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetUInt8Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetUInt8Value()
		},
		GetInt16: func(s string) (*int16, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetInt16()
		},
		GetInt16Value: func(s string) (int16, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetInt16Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetInt16Value()
		},
		GetUInt16: func(s string) (*uint16, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetUInt16()
		},
		GetUInt16Value: func (s string) (uint16, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetUInt16Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetUInt16Value()
		},
		GetInt32: func(s string) (*int32, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetInt32()
		},	
		GetInt32Value: func(s string) (int32, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetInt32Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetInt32Value()
		},
		GetInt64Value: func(s string) (int64, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetInt64Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetInt64Value()
		},
		GetUInt32: func(s string) (*uint32, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetUInt32()
		},
		GetUInt32Value: func(s string) (uint32, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetUInt32Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetUInt32Value()
		},
		GetUInt64Value: func(s string) (uint64, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetUInt64Value value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetUInt64Value()
		},
		GetInt: func(s string) (*int, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}

			return getInternalValue(s).GetInt()
		},
		GetIntValue: func(s string) (int, []error) {
			if isValueNilForMap(s) {
				var errors []error
				errors = append(errors, fmt.Errorf("json.Map.GetIntValue value is nil"))
				return 0, errors
			}
			return getInternalValue(s).GetIntValue()
		},
		SetInt: func(s string, value *int) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetInt64: func(s string, value *int64) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetInt32: func(s string, value *int32) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetInt16: func(s string, value *int16) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetInt8: func(s string, value *int8) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetIntValue: func(s string, value int) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetInt64Value: func(s string, value int64) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetInt32Value: func(s string, value int32) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetInt16Value: func(s string, value int16) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetInt8Value: func(s string, value int8) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetUInt: func(s string, value *uint) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetUInt64: func(s string, value *uint64) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetUInt32: func(s string, value *uint32) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetUInt16: func(s string, value *uint16) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetUInt8: func(s string, value *uint8) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetUIntValue: func(s string, value uint) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetUInt64Value: func(s string, value uint64) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetUInt32Value: func(s string, value uint32) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetUInt16Value: func(s string, value uint16) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetUInt8Value: func(s string, value uint8) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetFloat64: func(s string, value *float64) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetFloat64Value: func(s string, value float64) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		SetFloat32: func(s string, value *float32) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		SetFloat32Value: func(s string, value float32) {
			set_map_value := NewValue(&value)
			setInternalValue(s, set_map_value)
		},
		GetUInt64: func(s string) (*uint64, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetUInt64()
		},
		SetTime: func (s string, value *time.Time) {
			set_map_value := NewValue(value)
			setInternalValue(s, set_map_value)
		},
		GetTime: func(s string) (*time.Time, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetTime()
		},
		GetTimeWithDecimalPlaces: func(s string, decimal_places int) (*time.Time, []error) {
			if isValueNilForMap(s) {
				return nil, nil
			}
			return getInternalValue(s).GetTimeWithDecimalPlaces(decimal_places)
		},
		Values: func() *Array {
			return internal_keys_obj
		},
		Clone: func() (*Map, []error) {
			return clone()
		},
		SetValue: func(s string, value *Value)  {
			setInternalValue(s, value)
		},
	}
}
