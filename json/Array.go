package json

import (
	"fmt"
	"strings"
	common "github.com/matehaxor03/holistic_common/common"

)

type Array [](*Value)

func (a *Array) ToJSONString(json *strings.Builder) ([]error) {
	var errors []error
	
	if json == nil {
		errors = append(errors, fmt.Errorf("error: *strings.Builder is nil"))
		return errors
	}
	
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
}

func (a *Array) AppendString(value *string) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendStringValue(value string) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt8(value *uint8) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt8Value(value uint8) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt16(value *uint16) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt16Value(value uint16) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt32(value *uint32) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt32Value(value uint32) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt64(value *uint64) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt64Value(value uint64) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt(value *int) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendIntValue(value int) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUInt(value *uint) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendUIntValue(value uint) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt8(value *int8) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt8Value(value int8) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt16(value *int16) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt16Value(value int16) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt32(value *int32) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt32Value(value int32) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt64(value *int64) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendInt64Value(value int64) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendFloat32(value *float32) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendFloat32Value(value float32) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendFloat64(value *float64) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendFloat64Value(value float64) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendBool(value *bool) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendBoolValue(value bool) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendNil() {
	appended_value := newValue(nil)
	*a = append(*a, appended_value)
}

func (a *Array) AppendMapValue(value Map) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendMap(value *Map) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendArrayValue(value Array) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) AppendArray(value *Array) {
	appended_value := newValue(value)
	*a = append(*a, appended_value)
}

func (a *Array) GetStringValue(index int) (string, []error) {
	var errors []error
	
	
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
}

func (a *Array) GetMap(index int) (*Map, []error) {
	var errors []error
	if index < 0 {
		errors = append(errors, fmt.Errorf("Array.GetStringValue index is less than 0"))
		return nil, errors
	}

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
}
