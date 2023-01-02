package json

import (
	"fmt"
	"strings"
	common "github.com/matehaxor03/holistic_common/common"
)

type Array struct {
	ToJSONString func(json *strings.Builder) ([]error) 
	AppendString func (value *string)
	AppendStringValue func(value string)
	AppendInterface func(value *interface{})
	AppendInterfaceValue func(value interface{})
	AppendUInt func(value *uint)
	AppendUIntValue func(value uint)
	AppendUInt8 func(value *uint8)
	AppendUInt8Value func(value uint8)
	AppendUInt16 func(value *uint16)
	AppendUInt16Value func(value uint16)
	AppendUInt32 func(value *uint32)
	AppendUInt32Value func(value uint32)
	AppendUInt64 func(value *uint64)
	AppendUInt64Value func(value uint64)
	AppendInt func(value *int)
	AppendIntValue func(value int)
	AppendInt8 func(value *int8)
	AppendInt8Value func(value int8)
	AppendInt16 func(value *int16)
	AppendInt16Value func(value int16)
	AppendInt32 func(value *int32)
	AppendInt32Value func(value int32)
	AppendInt64 func(value *int64)
	AppendInt64Value func(value int64)
	AppendFloat32 func(value *float32)
	AppendFloat32Value func(value float32)
	AppendFloat64 func(value *float64)
	AppendFloat64Value func(value float64)
	AppendBool func(value *bool)
	AppendBoolValue func(value bool)
	AppendNil func()
	AppendValue func(value *Value)
	AppendValueValue func(value Value)
	AppendMap func(value *Map)
	AppendMapValue func(value Map)
	AppendArray func(value *Array)
	AppendArrayValue func(value Array)
	GetStringValue func(index int) (string, []error)
	GetMap func(index int) (*Map, []error)
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
	GetArrayOfString func() (*[](*string), []error)
	GetArrayOfStringValue func() ([](string), []error)
	GetArrayOfFloat32 func() (*[]*float32, []error)
	GetArrayOfFloat32Value func() ([]float32, []error)
	GetArrayOfFloat64 func() (*[]*float64, []error)
	GetArrayOfFloat64Value func() ([]float64, []error)
	GetArrayOfBool func() (*[]*bool, []error)
	GetArrayOfBoolValue func() ([]bool, []error)

	GetObject func() (*[](*interface{}))
	SetObject func(object *[](*interface{})) 
	Values func() *[](*Value)
	GetValue func() *Value
}

func NewArrayValue() (Array) {
	return *(NewArray())
}

func NewArray() (*Array) {
	array, _ :=  NewArrayOfValues(nil)
	return array
}

func NewArrayOfValues(a *[]*interface{}) (*Array, []error) {
	var errors []error
	var this_value *Value
	var this_array *Array
	interface_values := a
	var internal_values *[](*Value)

	if !common.IsNil(a) {
		temp_array := make([](*Value), len(*a))
		internal_values = &temp_array
	} else {
		temp_array := make([](*Value), 0)
		internal_values = &temp_array
	}
	
	if !common.IsNil(a) {
		for index, value := range *a {
			if common.IsMap(value) ||
			   common.IsArray(value) ||
			   common.IsValue(value) {
				errors = append(errors, fmt.Errorf("cannot create array with non-primitive type %s", common.GetType(value)))
			} else {
				// todo: for now just assume it's a value... to do map array and maps etc
				converted_value := NewValue(value)
				(*internal_values)[index] = converted_value
			}
		}
	} 

	set_this := func(array *Array) {
		this_array = array
		this_value = this_array.GetValue()
	}
	
	this := func() *Array {
		return this_array
	}

	values := func() *[](*Value) {
		return internal_values
	}

	setObject := func(a *[]*interface{}) {
		interface_values = a
	}

	getObject :=  func() *[]*interface{} {
		return interface_values
	}

	getValue := func() *Value {
		return this_value
	}

	created_array := Array{
		ToJSONString: func (json *strings.Builder) ([]error) {
			var errors []error
			
			if json == nil {
				errors = append(errors, fmt.Errorf("error: *strings.Builder is nil"))
				return errors
			}
			
			a := (*this()).Values()
			length := len(*a)
		
			if length == 0 {
				json.WriteString("[]")
				return nil
			}
		
			json.WriteString("[")
			for i, value := range *a {
				string_conversion_error := ConvertInterfaceValueToJSONStringValue(json, value)
				if string_conversion_error != nil {
					errors = append(errors, string_conversion_error...)
				} 
		
				if i < length - 1 {
					json.WriteString(",")
				}
			}
			json.WriteString("]")
		
			if len(errors) > 0 {
				return errors
			}
		
			return nil
		},
		AppendString: func(value *string) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendStringValue: func(value string) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt: func(value *uint) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUIntValue: func(value uint) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt8: func(value *uint8) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt8Value: func(value uint8) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt16: func(value *uint16) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt16Value: func(value uint16) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt32: func(value *uint32) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt32Value: func(value uint32) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt64: func(value *uint64) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt64Value: func(value uint64) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt: func(value *int) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendIntValue: func(value int) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt8: func(value *int8) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt8Value: func(value int8) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt16: func(value *int16) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt16Value: func(value int16) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt32: func(value *int32) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt32Value: func(value int32) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt64: func(value *int64) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt64Value: func(value int64) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendFloat32: func(value *float32) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendFloat32Value: func(value float32) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendFloat64: func(value *float64) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendFloat64Value: func(value float64) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendBool: func(value *bool) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendBoolValue: func(value bool) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInterface: func(value *interface{}) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendInterfaceValue: func(value interface{}) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendNil: func() {
			a := this().Values()
			*a = append(*a, nil)
		},
		AppendMap: func(value *Map) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendMapValue: func(value Map) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendArray: func(value *Array) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendArrayValue: func(value Array) {
			a := this().Values()
			appended_value := NewValue(value)
			*a = append(*a, appended_value)
		},
		AppendValue: func(value *Value) {
			a := this().Values()
			*a = append(*a, value)
		},
		AppendValueValue: func(value Value) {
			a := this().Values()
			*a = append(*a, &value)
		},
		Values: func() *[](*Value) {
			return values()
		},
		GetValue: func() *Value {
			return getValue()
		},
		GetStringValue: func(index int) (string, []error) {
			var errors []error
			a := this().Values()
			
			if index < 0 {
				errors = append(errors, fmt.Errorf("Array.GetStringValue index is less than 0"))
				return "", errors
			}
		
			if index > (len(*a) - 1) {
				errors = append(errors, fmt.Errorf("Array.GetStringValue index is out of range"))
				return "", errors
			}
		
			value, value_errors := ((*a)[index]).GetStringValue()
			if value_errors != nil {
				errors = append(errors, fmt.Errorf("Array.GetStringValue has errors: %s", fmt.Sprintf("%s", value_errors)))
			} else if common.IsNil(value) {
				errors = append(errors, fmt.Errorf("Array.GetStringValue value is nil"))
			}
		
			if len(errors) > 0 {
				return "", errors
			}
		
			return value, nil
		},
		GetMap: func(index int) (*Map, []error) {
			var errors []error
			if index < 0 {
				errors = append(errors, fmt.Errorf("Array.GetStringValue index is less than 0"))
				return nil, errors
			}
		
			a := this().Values()
			if index > (len(*a) - 1) {
				errors = append(errors, fmt.Errorf("Array.GetStringValue index is out of range"))
				return nil, errors
			}
		
			value, value_errors := ((*a)[index]).GetMap()
			if value_errors != nil {
				errors = append(errors, fmt.Errorf("Array.GetStringValue has errors: %s", fmt.Sprintf("%s", value_errors)))
			} else if common.IsNil(value) {
				errors = append(errors, fmt.Errorf("Array.GetStringValue value is nil"))
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
		
			return value, nil
		},
		GetArrayOfInt8: func() (*[](*int8), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*int8))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					int8_value, int8_value_errors := array_value.GetInt8()
					if int8_value_errors != nil {
						errors = append(errors, int8_value_errors...)
					} else {
						result = append(result, int8_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfInt16: func() (*[](*int16), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*int16))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					int16_value, int16_value_errors := array_value.GetInt16()
					if int16_value_errors != nil {
						errors = append(errors, int16_value_errors...)
					} else {
						result = append(result, int16_value)			 
					}
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfInt32: func() (*[](*int32), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*int32))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					int32_value, int32_value_errors := array_value.GetInt32()
					if int32_value_errors != nil {
						errors = append(errors, int32_value_errors...)
					} else {
						result = append(result, int32_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfInt64: func() (*[](*int64), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*int64))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					int64_value, int64_value_errors := array_value.GetInt64()
					if int64_value_errors != nil {
						errors = append(errors, int64_value_errors...)
					} else {
						result = append(result, int64_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfUInt8: func() (*[](*uint8), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*uint8))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					uint8_value, uint8_value_errors := array_value.GetUInt8()
					if uint8_value_errors != nil {
						errors = append(errors, uint8_value_errors...)
					} else {
						result = append(result, uint8_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfUInt16: func() (*[](*uint16), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*uint16))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					uint16_value, uint16_value_errors := array_value.GetUInt16()
					if uint16_value_errors != nil {
						errors = append(errors, uint16_value_errors...)
					} else {
						result = append(result, uint16_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfUInt32: func() (*[](*uint32), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*uint32))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					uint32_value, uint32_value_errors := array_value.GetUInt32()
					if uint32_value_errors != nil {
						errors = append(errors, uint32_value_errors...)
					} else {
						result = append(result, uint32_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfUInt64: func() (*[](*uint64), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*uint64))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					uint64_value, uint64_value_errors := array_value.GetUInt64()
					if uint64_value_errors != nil {
						errors = append(errors, uint64_value_errors...)
					} else {
						result = append(result, uint64_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfInt8Value: func() ([](int8), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](int8))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					int8_value, int8_value_errors := array_value.GetInt8Value()
					if int8_value_errors != nil {
						errors = append(errors, int8_value_errors...)
					} else {
						result = append(result, int8_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfInt16Value: func() ([](int16), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](int16))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					int16_value, int16_value_errors := array_value.GetInt16Value()
					if int16_value_errors != nil {
						errors = append(errors, int16_value_errors...)
					} else {
						result = append(result, int16_value)			 
					}
				}
			}

			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfInt32Value: func() ([](int32), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](int32))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					int32_value, int32_value_errors := array_value.GetInt32Value()
					if int32_value_errors != nil {
						errors = append(errors, int32_value_errors...)
					} else {
						result = append(result, int32_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfInt64Value: func() ([](int64), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](int64))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					int64_value, int64_value_errors := array_value.GetInt64Value()
					if int64_value_errors != nil {
						errors = append(errors, int64_value_errors...)
					} else {
						result = append(result, int64_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfUInt8Value: func() ([]uint8, []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](uint8))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					uint8_value, uint8_value_errors := array_value.GetUInt8Value()
					if uint8_value_errors != nil {
						errors = append(errors, uint8_value_errors...)
					} else {
						result = append(result, uint8_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfUInt16Value: func() ([](uint16), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](uint16))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					uint16_value, uint16_value_errors := array_value.GetUInt16Value()
					if uint16_value_errors != nil {
						errors = append(errors, uint16_value_errors...)
					} else {
						result = append(result, uint16_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfUInt32Value: func() ([](uint32), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](uint32))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					uint32_value, uint32_value_errors := array_value.GetUInt32Value()
					if uint32_value_errors != nil {
						errors = append(errors, uint32_value_errors...)
					} else {
						result = append(result, uint32_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfUInt64Value: func() ([](uint64), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](uint64))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					uint64_value, uint64_value_errors := array_value.GetUInt64Value()
					if uint64_value_errors != nil {
						errors = append(errors, uint64_value_errors...)
					} else {
						result = append(result, uint64_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfBool: func() (*[](*bool), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*bool))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					string_value, string_value_errors := array_value.GetBool()
					if string_value_errors != nil {
						errors = append(errors, string_value_errors...)
					} else {
						result = append(result, string_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfBoolValue: func() ([](bool), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](bool))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained nil value"))		 
				} else {
					string_value, string_value_errors := array_value.GetBoolValue()
					if string_value_errors != nil {
						errors = append(errors, string_value_errors...)
					} else {
						result = append(result, string_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfString: func() (*[](*string), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*string))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					string_value, string_value_errors := array_value.GetString()
					if string_value_errors != nil {
						errors = append(errors, string_value_errors...)
					} else {
						result = append(result, string_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfStringValue: func() ([](string), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](string))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					string_value, string_value_errors := array_value.GetStringValue()
					if string_value_errors != nil {
						errors = append(errors, string_value_errors...)
					} else {
						result = append(result, string_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil
		},
		GetArrayOfFloat32: func() (*[](*float32), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](*float32))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					result = append(result, nil)			 
				} else {
					string_value, string_value_errors := array_value.GetFloat32()
					if string_value_errors != nil {
						errors = append(errors, string_value_errors...)
					} else {
						result = append(result, string_value)			 
					}
				}
			}
		
			if len(errors) > 0 {
				return nil, errors
			}
			
			return &result, nil
		},
		GetArrayOfFloat32Value: func() ([](float32), []error) {
			array_values := this().Values()
			if common.IsNil(array_values) {
				return nil, nil
			}
		
			var errors []error
			var result ([](float32))
			for _, array_value := range *array_values {
				if common.IsNil(array_value) {
					errors = append(errors, fmt.Errorf("array contained a null value")) 
				} else {
					string_value, string_value_errors := array_value.GetFloat32Value()
					if string_value_errors != nil {
						errors = append(errors, string_value_errors...)
					} else {
						result = append(result, string_value)			 
					}
				}
			}


			if len(errors) > 0 {
				return nil, errors
			}
			
			return result, nil

			
			},
			GetArrayOfFloat64: func() (*[](*float64), []error) {
				array_values := this().Values()
				if common.IsNil(array_values) {
					return nil, nil
				}
			
				var errors []error
				var result ([](*float64))
				for _, array_value := range *array_values {
					if common.IsNil(array_value) {
						result = append(result, nil)			 
					} else {
						string_value, string_value_errors := array_value.GetFloat64()
						if string_value_errors != nil {
							errors = append(errors, string_value_errors...)
						} else {
							result = append(result, string_value)			 
						}
					}
				}
			
				if len(errors) > 0 {
					return nil, errors
				}
				
				return &result, nil
			},
			GetArrayOfFloat64Value: func() ([](float64), []error) {
			
				array_values := this().Values()
				if common.IsNil(array_values) {
					return nil, nil
				}
			
				var errors []error
				var result ([](float64))
				for _, array_value := range *array_values {
					if common.IsNil(array_value) {
						errors = append(errors, fmt.Errorf("array contained a null value")) 
					} else {
						string_value, string_value_errors := array_value.GetFloat64Value()
						if string_value_errors != nil {
							errors = append(errors, string_value_errors...)
						} else {
							result = append(result, string_value)			 
						}
					}
				}
			
		
				if len(errors) > 0 {
					return nil, errors
				}
				
				return result, nil
		},
		SetObject: func(value *[](*interface{})) {
			setObject(value)
		},
		GetObject: func() (*[](*interface{})) {
			return getObject()
		},
	}
	set_this(&created_array)
	if len(errors) > 0 {
		return nil, errors
	}
	return &created_array, nil
}





