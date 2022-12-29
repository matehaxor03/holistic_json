package tests
 
import (
    "testing"
)

func TestCanParseBoolTrue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":true}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolFalse(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":false}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != false {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseMultipleBools(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":true,\"key2\":false}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json_obj.GetType("key2") != "*bool" {
		t.Errorf("key2 is not a *bool: %s", json_obj.GetType("key2"))
	} else {
		value, value_errors := json_obj.GetBool("key2") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != false {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolSpaceBeforeKey(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{ \"key\":true}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolNewlineBeforeKey(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\n\"key\":true}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolDosNewlineBeforeKey(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\r\n\"key\":true}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolSpaceAfterKey(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\" :true}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolNewlineAfterKey(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\"\n:true}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolDosNewlineAfterKey(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\"\r\n:true}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolSpaceBeforeValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\": true}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolSpaceNewlineBeforeValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":\ntrue}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolSpaceDosNewlineBeforeValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":\r\ntrue}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolSpaceAfterValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":true }")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolNewlineAfterValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":true\n}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}

func TestCanParseBoolDosNewlineAfterValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":true\r\n}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "bool" {
		t.Errorf("key is not a *bool: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetBool("key") 

		if value_errors != nil {
			t.Errorf("map GetBool has errors")
		} else if value == nil {
			t.Errorf("GetBool is nil")
		} else if *value != true {
			t.Errorf("expected: value actual: %t", *value)
		}
	}
}
