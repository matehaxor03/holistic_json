package tests
 
import (
    "testing"
)

func TestCanParseEmptyArray(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArray("key") 

		if value_errors != nil {
			t.Errorf("map GetArray has errors")
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 0 {
			t.Errorf("expected: length=0 actual: length=%d", len(*value))
		}
	}
}

func TestCanParseArrayContainingSingleString(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[\"value\"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArray("key") 

		if value_errors != nil {
			t.Errorf("map GetArray has errors")
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0].(*string)) != "value" {
			t.Errorf("expected \"value\" actual: \"%s\"", (*(*value)[0].(*string)))
		}
	}	
}

func TestCanParseMultipleArraysContainingSingleString(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[\"value\"],\"key2\":[\"value2\"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArray("key") 

		if value_errors != nil {
			t.Errorf("map GetArray has errors")
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0].(*string)) != "value" {
			t.Errorf("expected \"value\" actual: \"%s\"", (*(*value)[0].(*string)))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json_obj.GetType("key2") != "*json.Array" {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArray("key2") 

		if value_errors != nil {
			t.Errorf("map GetArray has errors")
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0].(*string)) != "value2" {
			t.Errorf("expected \"value2\" actual: \"%s\"", (*(*value)[0].(*string)))
		}
	}
}

func TestCanParseArrayContainingMultipleStrings(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[\"value\",\"value2\"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArray("key") 

		if value_errors != nil {
			t.Errorf("map GetArray has errors")
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0].(*string)) != "value" {
			t.Errorf("expected \"value\" actual: \"%s\"", (*(*value)[0].(*string)))
		} else if (*(*value)[1].(*string)) != "value2" {
			t.Errorf("expected \"value2\" actual: \"%s\"", (*(*value)[0].(*string)))
		}
	}
}