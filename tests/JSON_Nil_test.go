package tests
 
import (
    "testing"
)

func TestCanParseNil(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":null}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "nil" {
		t.Errorf("key is not a nil: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetString("key") 

		if value_errors != nil {
			t.Errorf("map GetString has errors")
		} else if value != nil {
			t.Errorf("GetString is not nil")
		}
	}
}

func TestCanParseMultipleNil(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":null,\"key2\":null}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "nil" {
		t.Errorf("key is not nil: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetString("key") 
		if value_errors != nil {
			t.Errorf("map GetString has errors")
		} else if value != nil {
			t.Errorf("GetString is not nil")
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json_obj.GetType("key2") != "nil" {
		t.Errorf("key2 is not nil: %s", json_obj.GetType("key2"))
	} else {
		value, value_errors := json_obj.GetString("key2") 
		if value_errors != nil {
			t.Errorf("map GetString has errors")
		} else if value != nil {
			t.Errorf("GetString is not nil")
		}
	}
	
}
