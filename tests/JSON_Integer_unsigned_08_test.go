package tests
 
import (
    "testing"
)

func TestCanParseUInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":0}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint8" {
		t.Errorf("key is not a *uint8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt8 is nil")
		} else if *value != 0 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCanParseUInt8LowBoundaryWithSpaceBeforeValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\": 0}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint8" {
		t.Errorf("key is not a *uint8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt8 is nil")
		} else if *value != 0 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCanParseUInt8LowBoundaryWithSpaceAfterValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":0 }")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint8" {
		t.Errorf("key is not a *uint8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt8 is nil")
		} else if *value != 0 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}	
}

func TestCanParseUInt8LowBoundaryWithNewlineBeforeValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":\n0}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint8" {
		t.Errorf("key is not a *uint8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt8 is nil")
		} else if *value != 0 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCanParseUInt8LowBoundaryWithNewlineAfterValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":0\n}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint8" {
		t.Errorf("key is not a *uint8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt8 is nil")
		} else if *value != 0 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}

func TestCanParseUInt8LowBoundaryWithDosNewlineBeforeValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":\r\n0}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint8" {
		t.Errorf("key is not a *uint8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt8 is nil")
		} else if *value != 0 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}	
}

func TestCanParseUInt8LowBoundaryWithDosNewlineAfterValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":0\r\n}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint8" {
		t.Errorf("key is not a *uint8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt8 is nil")
		} else if *value != 0 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}


func TestCanParseUInt8HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":255}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "uint8" {
		t.Errorf("key is not a *uint8: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetUInt8("key") 

		if value_errors != nil {
			t.Errorf("map GetUInt8 has errors: " + value_errors[0].Error())
		} else if value == nil {
			t.Errorf("GetUInt8 is nil")
		} else if *value != 255 {
			t.Errorf("expected: value actual: %d", *value)
		}
	}
}