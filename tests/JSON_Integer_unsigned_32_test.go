package tests
 
import (
    "testing"
)


func TestCanParseUInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t,"{\"key\":65536}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint32" {
		t.Errorf("key is not a *uint32: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt32("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt32 is nil")
		} else if *value != 65536 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCanParseUInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":4294967295}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint32" {
		t.Errorf("key is not a *uint32: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt32("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt32 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt32 is nil")
		} else if *value != 4294967295 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}
