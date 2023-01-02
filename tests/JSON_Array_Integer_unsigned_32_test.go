package tests
 
import (
    "testing"
)

func TestCanParseArrayContainingSingleUInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[65536]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 65536 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[65536],\"key2\":[65537]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 65536 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 65537 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[65536,65537],\"key2\":[65538,65539]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 65536 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 65537 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 65538 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 65539 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[65536,65537]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 65536 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 65537 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}
}

func TestCanParseArrayContainingSingleUInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[4294967295]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 4294967295 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[4294967294],\"key2\":[4294967295]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 4294967294 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 4294967295 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[4294967292,4294967293],\"key2\":[4294967294,4294967295]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 4294967292 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 4294967293 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 4294967294 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 4294967295 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[4294967294,4294967295]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 4294967294 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 4294967295 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}
}
