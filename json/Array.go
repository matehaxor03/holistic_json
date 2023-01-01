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
	//GetObject func() (*[](*interface{}))
	//SetObject func(object *[](*interface{})) 
	Values func() *[](*Value)
}

func newArray() *Array {
	var this_value *Array
	temp_array := make([](*Value), 0)
	array := &temp_array

	set_this_value := func(value *Array) {
		this_value = value
	}
	
	this := func() *Array {
		return this_value
	}

	values := func() *[](*Value) {
		return array
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
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendStringValue: func(value string) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt: func(value *uint) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUIntValue: func(value uint) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt8: func(value *uint8) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt8Value: func(value uint8) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt16: func(value *uint16) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt16Value: func(value uint16) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt32: func(value *uint32) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt32Value: func(value uint32) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt64: func(value *uint64) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendUInt64Value: func(value uint64) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt: func(value *int) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendIntValue: func(value int) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt8: func(value *int8) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt8Value: func(value int8) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt16: func(value *int16) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt16Value: func(value int16) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt32: func(value *int32) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt32Value: func(value int32) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt64: func(value *int64) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendInt64Value: func(value int64) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendFloat32: func(value *float32) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendFloat32Value: func(value float32) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendFloat64: func(value *float64) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendFloat64Value: func(value float64) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendBool: func(value *bool) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendBoolValue: func(value bool) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendNil: func() {
			a := this().Values()
			*a = append(*a, nil)
		},
		AppendMap: func(value *Map) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendMapValue: func(value Map) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendArray: func(value *Array) {
			a := this().Values()
			appended_value := newValue(value)
			*a = append(*a, appended_value)
		},
		AppendArrayValue: func(value Array) {
			a := this().Values()
			appended_value := newValue(value)
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
		/*
		SetObject: func(value *[](*interface{})) {
			setObject(value)
		},
		GetObject: func() (*[](*interface{})) {
			return getObject()
		},*/
	}
	set_this_value(&created_array)
	return &created_array
}





