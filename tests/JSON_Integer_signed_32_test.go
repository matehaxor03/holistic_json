package tests
 
import (
    "testing"
)

func TestCanParseNegativeInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":-32769}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsInt32("key") {
		t.Errorf("key is not a *int32: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetInt32("key") 

		if value_errors != nil {
			t.Errorf("map GetInt32 has errors")
		} else if value == nil {
			t.Errorf("GetInt32 is nil")
		} else if *value != -32769 {
			t.Errorf("expected: -32769  actual: %d", *value)
		}
	}
}

func TestCanParseNegativeInt32HighBoundarfy(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":-2147483648}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsInt32("key") {
		t.Errorf("key is not a *int32: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetInt32("key") 

		if value_errors != nil {
			t.Errorf("map GetInt32 has errors")
		} else if value == nil {
			t.Errorf("GetInt32 is nil")
		} else if *value != -2147483648 {
			t.Errorf("expected: -2147483648  actual: %d", *value)
		}
	}	
}
