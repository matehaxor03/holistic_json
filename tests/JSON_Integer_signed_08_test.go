package tests
 
import (
    "testing"
)

func TestCanParseNegativeInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":-1}")
	
	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsInt8("key") {
		t.Errorf("key is not a *int8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetInt8 has errors")
		} else if value == nil {
			t.Errorf("GetInt8 is nil")
		} else if *value != -1 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCanParseNegativeInt8HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":-128}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsInt8("key") {
		t.Errorf("key is not a *int8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetInt8 has errors")
		} else if value == nil {
			t.Errorf("GetInt8 is nil")
		} else if *value != -128 {
			t.Errorf("expected: -128 actual: %d", *value)
		}
	}
}
	