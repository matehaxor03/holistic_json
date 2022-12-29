package json

import (
	"fmt"
	"strings"
)

type Array [](Value)

func (a Array) ToJSONString(json *strings.Builder) ([]error) {
	var errors []error
	
	if json == nil {
		errors = append(errors, fmt.Errorf("error: *strings.Builder is nil"))
		return errors
	}
	
	length := len(a)

	if length == 0 {
		json.WriteString("[]")
		return nil
	}

	json.WriteString("[")
	for i, value := range a {
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
