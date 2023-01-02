package tests
 
import (
    "testing"
	"fmt"
)

func TestCanParseBoolArrayTrue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[true]}")

	{
		values, values_errors := json_obj.GetArrayOfBool("key") 

		if values_errors != nil {
			t.Errorf(fmt.Sprintf("%s", values_errors))
		} else if values == nil {
			t.Errorf("values is nil")
		} else if len(*values) != 1 {
			t.Errorf("expected elements: %d actual: %d", 1, len(*values))
		} else if (*(*values)[0]) != true {
			t.Errorf("expected: %t actual: %t", true, (*(*values)[0]))
		}
	}

	{
		values, values_errors := json_obj.GetArrayOfBoolValue("key") 

		if values_errors != nil {
			t.Errorf(fmt.Sprintf("%s", values_errors))
		} else if values == nil {
			t.Errorf("values is nil")
		} else if len(values) != 1 {
			t.Errorf("expected elements: %d actual: %d", 1, len(values))
		} else if ((values)[0]) != true {
			t.Errorf("expected: %t actual: %t", true, ((values)[0]))
		}
	}
}

func TestCanParseBoolArrayFalse(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[false]}")

	{
		values, values_errors := json_obj.GetArrayOfBool("key") 

		if values_errors != nil {
			t.Errorf(fmt.Sprintf("%s", values_errors))
		} else if values == nil {
			t.Errorf("values is nil")
		} else if len(*values) != 1 {
			t.Errorf("expected elements: %d actual: %d", 1, len(*values))
		} else if (*(*values)[0]) != false {
			t.Errorf("expected: %t actual: %t", false, (*(*values)[0]))
		}
	}

	{
		values, values_errors := json_obj.GetArrayOfBoolValue("key") 

		if values_errors != nil {
			t.Errorf(fmt.Sprintf("%s", values_errors))
		} else if values == nil {
			t.Errorf("values is nil")
		} else if len(values) != 1 {
			t.Errorf("expected elements: %d actual: %d", 1, len(values))
		} else if ((values)[0]) != false {
			t.Errorf("expected: %t actual: %t", false, ((values)[0]))
		}
	}
}

func TestCanParseBoolArrayMulitple(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[true, false]}")

	{
		values, values_errors := json_obj.GetArrayOfBool("key") 

		if values_errors != nil {
			t.Errorf(fmt.Sprintf("%s", values_errors))
		} else if values == nil {
			t.Errorf("values is nil")
		} else if len(*values) != 2 {
			t.Errorf("expected elements: %d actual: %d", 1, len(*values))
		} else if (*(*values)[0]) != true {
			t.Errorf("expected: %t actual: %t", true, (*(*values)[0]))
		} else if (*(*values)[1]) != false {
			t.Errorf("expected: %t actual: %t", false, (*(*values)[1]))
		}
	}

	{
		values, values_errors := json_obj.GetArrayOfBoolValue("key") 

		if values_errors != nil {
			t.Errorf(fmt.Sprintf("%s", values_errors))
		} else if values == nil {
			t.Errorf("values is nil")
		} else if len(values) != 2 {
			t.Errorf("expected elements: %d actual: %d", 1, len(values))
		} else if ((values)[0]) != true {
			t.Errorf("expected: %t actual: %t", true, ((values)[0]))
		}  else if ((values)[1]) != false {
			t.Errorf("expected: %t actual: %t", false, ((values)[1]))
		}
	}
}

func TestCanParseBoolArrayMulitpleWithNil(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[true, false, null]}")

	{
		values, values_errors := json_obj.GetArrayOfBool("key") 

		if values_errors != nil {
			t.Errorf(fmt.Sprintf("%s", values_errors))
		} else if values == nil {
			t.Errorf("values is nil")
		} else if len(*values) != 3 {
			t.Errorf("expected elements: %d actual: %d", 1, len(*values))
		} else if (*(*values)[0]) != true {
			t.Errorf("expected: %t actual: %t", true, (*(*values)[0]))
		} else if (*(*values)[1]) != false {
			t.Errorf("expected: %t actual: %t", false, (*(*values)[1]))
		} else if (*values)[2] != nil {
			t.Errorf("expected: %T actual: %T", nil, (*values)[2])
		}
	}
}

func TestCannotParseBoolArrayMulitpleWithNil(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[true, false, null]}")

	{
		values, values_errors := json_obj.GetArrayOfBoolValue("key") 

		if values_errors == nil {
			t.Errorf("expect there to be errors")
		} 

		if values != nil {
			t.Errorf("values to be nil")
		}
	}
}


