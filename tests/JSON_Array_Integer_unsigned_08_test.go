package tests
 
import (
    "testing"
)

func TestCanParseArrayContainingSingleUInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[0]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 0 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[0],\"key2\":[1]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 0 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 1 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[0,1],\"key2\":[2,3]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 0 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 1 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 2 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 3 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[0,1]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 0 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 1 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}
}

func TestCanParseArrayContainingSingleUInt8HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[255]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 255 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleUInt8HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[254],\"key2\":[255]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 254 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 255 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleUInt8HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[252,253],\"key2\":[254,255]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 252 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 253 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 254 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 255 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultipleUInt8HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[254,255]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfUInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 254 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[0]))
		} else if (*(*value)[1]) != 255 {
			t.Errorf("expected \"value\" actual: %d", (*(*value)[1]))
		}
	}
}
