package tests
 
import (
    "testing"
)

func TestCanParseArrayContainingSingleInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-32769]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32769 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-32769],\"key2\":[-32770]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32769 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32770 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-32769,-32770],\"key2\":[-32771,-32772]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32769 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -32770 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32771 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -32772 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}

func TestCanParseArrayContainingMultipleInt32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-32769,-32770]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -32769 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -32770 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}

func TestCanParseArrayContainingSingleInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-2147483648]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -2147483648 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-2147483648],\"key2\":[-2147483647]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -2147483648 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -2147483647 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-2147483648,-2147483647],\"key2\":[-2147483646,-2147483645]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -2147483648 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -2147483647 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -2147483646 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -2147483645 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}

func TestCanParseArrayContainingMultipleInt32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-2147483648,-2147483647]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -2147483648 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -2147483647 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}
