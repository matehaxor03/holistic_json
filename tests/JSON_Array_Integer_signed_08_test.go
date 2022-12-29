package tests
 
import (
    "testing"
)

// int8 boundary low boundary
func TestCanParseArrayContainingSingleInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-1]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key") 

		if value_errors != nil {
			t.Errorf("error: %s", value_errors)
		} else if value == nil {
			t.Errorf("error: GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("error: expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -1 {
			t.Errorf("error: expected \"value\" actual: %d", *((*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-1],\"key2\":[-2]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json_obj.GetType("key2") != "*json.Array" {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -2 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-1,-2],\"key2\":[-3,-4]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		} else if *((*value)[1]) != -2 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json_obj.GetType("key2") != "*json.Array" {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -3 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		} else if *((*value)[1]) != -4 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultipleInt8LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-1,-2]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -1 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		} else if *((*value)[1]) != -2 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1]))
		}
	}
}

// int8 boundary high
func TestCanParseArrayContainingSingleInt8HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-128]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -128 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleHighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-128],\"key2\":[-127]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -128 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json_obj.GetType("key2") != "*json.Array" {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -127 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleHighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-128,-127],\"key2\":[-126,-125]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -128 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		} else if *((*value)[1]) != -127 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if json_obj.GetType("key2") != "*json.Array" {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -126 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		} else if *((*value)[0]) != -125 {
			t.Errorf("expected \"value\" actual: %d", *((*value)[0]))
		}
	}
}

func TestCanParseArrayContainingMultipleInt8HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-128,-127]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Array" {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfInt8("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != -128 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[0])
		} else if (*(*value)[1]) != -127 {
			t.Errorf("expected \"value\" actual: %d", *(*value)[1])
		}
	}
}
