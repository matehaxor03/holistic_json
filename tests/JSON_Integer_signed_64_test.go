package tests
 
import (
    "testing"
	json "github.com/matehaxor03/holistic_json/json"
)

func TestCanParseNegativeInt64LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":-2147483649}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "int64" {
		t.Errorf("key is not a *int64: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetInt64("key") 

		if value_errors != nil {
			t.Errorf("map GetInt64 has errors %s", value_errors)
		} else if value == nil {
			t.Errorf("GetInt64 is nil")
		} else if *value != -2147483649 {
			t.Errorf("expected: -2147483649  actual: %d", *value)
		}
	}	
}

func TestCanParseNegativeInt64HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":-9223372036854775808}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "int64" {
		t.Errorf("key is not a *int64: %s", json_obj.GetType("key"))
	} else {
		value, value_errors := json_obj.GetInt64("key") 

		if value_errors != nil {
			t.Errorf("map GetInt64 has errors %s", value_errors)
		} else if value == nil {
			t.Errorf("GetInt64 is nil")
		} else if *value != -9223372036854775808 {
			t.Errorf("expected: -9223372036854775808  actual: %d", *value)
		}
	}	
}

func TestCannotParseNegativeInt64Overflow(t *testing.T) {
	json, json_errors := json.Parse("{\"key\":-9223372036854775809}")

	if json_errors == nil {
		t.Errorf("there were no errors")
	}

	if json != nil {
		t.Errorf("json value was returned")
	}
}