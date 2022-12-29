package json

import (
	"fmt"
	"strings"
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

func (a *Array) AppendString(s *string) {
	fmt.Println("appending " + *s)
	appended_value := Value{"value":*s}
	*a = append(*a, &appended_value)
	fmt.Println(*a)
	fmt.Println(len(*a))
}

func (a *Array) AppendStringValue(s string) {
	appended_value := Value{"value":s}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt8(value *uint8) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt8Value(value uint8) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt16(value *uint16) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt16Value(value uint16) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt32(value *uint32) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt32Value(value uint32) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt64(value *uint64) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt64Value(value uint64) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt(value *int) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendIntValue(value int) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUInt(value *uint) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendUIntValue(value uint) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt8(value *int8) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt8Value(value int8) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt16(value *int16) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt16Value(value int16) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt32(value *int32) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt32Value(value int32) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt64(value *int64) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendInt64Value(value int64) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendFloat32(value *float32) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendFloat32Value(value float32) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendFloat64(value *float64) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendFloat64Value(value float64) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendBool(b *bool) {
	appended_value := Value{"value":b}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendBoolValue(b bool) {
	appended_value := Value{"value":b}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendNil() {
	appended_value := Value{"value":nil}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendMapValue(value Map) {
	appended_value := Value{"value":&value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendMap(value *Map) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendArrayValue(value Array) {
	appended_value := Value{"value":&value}
	*a = append(*a, &appended_value)
}

func (a *Array) AppendArray(value *Array) {
	appended_value := Value{"value":value}
	*a = append(*a, &appended_value)
}
