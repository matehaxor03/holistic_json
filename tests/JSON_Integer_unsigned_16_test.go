package tests
 
import (
    "testing"
)

func TestCanParseUInt16LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t,"{\"key\":256}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*uint16" {
		t.Errorf("key is not a *uint16: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt16("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt16 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt16 is nil")
		} else if *value != 256 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}	
}

func TestCanParseUInt16HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t,"{\"key\":65535}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*uint16" {
		t.Errorf("key is not a *uint16: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt16("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt16 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt16 is nil")
		} else if *value != 65535 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}
