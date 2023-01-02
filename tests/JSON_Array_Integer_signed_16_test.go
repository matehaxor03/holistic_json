package tests
 
import (
    "testing"
)

func TestCanParseArrayContainingSingleInt16LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-129]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -129 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}
}

func TestCanParseMultipleArraysContainingSingleLowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-129],\"key2\":[-130]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if(*(*value)[0]) != -129 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -130 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt16LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-129,-130],\"key2\":[-131,-132]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -129 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -130 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -131 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -132 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}

func TestCanParseArrayContainingMultipleInt16LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-129,-130]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -129 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -130 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}

func TestCanParseArrayContainingSingleInt16HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-32768]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32768 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt16HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-32768],\"key2\":[-32767]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32768 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32767 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt16HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-32768,-32767],\"key2\":[-32766,-32765]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32768 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -32767 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32766 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -32765 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}

func TestCanParseArrayContainingMultipleInt16HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-32768,-32767]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt16("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32768 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -32767 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}
