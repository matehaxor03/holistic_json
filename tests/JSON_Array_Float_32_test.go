package tests
 
import (
    "testing"
)

func TestCanParseArrayContainingSingleFloat32Boundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[0.0]}")
	{
		value, value_errors := json_obj.GetArrayOfFloat32("key") 

		if value_errors != nil {
			t.Errorf("error: %s", value_errors)
		} else if value == nil {
			t.Errorf("error: GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("error: expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != 0.0 {
			t.Errorf("error: expected: %f actual: %f", 0.0, *((*value)[0]))
		}
	}

	{
		value, value_errors := json_obj.GetArrayOfFloat32Value("key") 
		if value_errors != nil {
			t.Errorf("error: %s", value_errors)
		} else if value == nil {
			t.Errorf("error: GetArray is nil")
		} else if len(value) != 1 {
			t.Errorf("error: expected: length=1 actual: length=%d", len(value))
		} else if ((value)[0]) != 0.0 {
			t.Errorf("error: expected: %f actual: %f", 0.0, ((value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleFloat32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[0.123457],\"key2\":[-0.7654321]}")

	{
		value, value_errors := json_obj.GetArrayOfFloat32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != 0.123457 {
			t.Errorf("expected %f actual: %f", 0.123457, *((*value)[0]))
		}
	}

	{
		value, value_errors := json_obj.GetArrayOfFloat32Value("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(value))
		} else if ((value)[0]) != 0.123457 {
			t.Errorf("expected %f actual: %f", 0.123457, ((value)[0]))
		}
	}

	{
		value, value_errors := json_obj.GetArrayOfFloat32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -0.765432 {
			t.Errorf("expected %f actual: %f", -0.765432, *((*value)[0]))
		}
	}

	{
		value, value_errors := json_obj.GetArrayOfFloat32Value("key2") 
		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(value))
		} else if ((value)[0]) != -0.765432 {
			t.Errorf("expected %f actual: %f", -0.765432, ((value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleFloat32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-0.7654,-1.2345],\"key2\":[-3.4343,-4.3434]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -0.7654 {
			t.Errorf("expected %f actual: %f", -0.7654, *((*value)[0]))
		} else if *((*value)[1]) != -1.2345 {
			t.Errorf("expected %f actual: %f", -1.2345, *((*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -3.4343 {
			t.Errorf("expected %f actual: %f", -3.4343, *((*value)[0]))
		} else if *((*value)[1]) != -4.3434 {
			t.Errorf("expected %f actual: %f", -4.3434, *((*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultipleFloat32LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[-7.2323,-8.3477]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if *((*value)[0]) != -7.2323 {
			t.Errorf("expected %f actual: %f", -7.2323, *((*value)[0]))
		} else if *((*value)[1]) != -8.3477 {
			t.Errorf("expected %f actual: %f", -8.3477, *((*value)[1]))
		}
	}
}

func TestCanParseArrayContainingSingleFloat32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[340000000000000000000000000000000000000.00]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != 340000000000000000000000000000000000000.00 {
			t.Errorf("expected %f actual: %f", 340000000000000000000000000000000000000.00, *((*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleFloat32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[332000000000000000000000000000000000000.00],\"key2\":[31000000000000000000000000000000000000.00]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != 332000000000000000000000000000000000000.00 {
			t.Errorf("expected %f actual: %f", 332000000000000000000000000000000000000.00, *((*value)[0]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != 31000000000000000000000000000000000000.00 {
			t.Errorf("expected %f actual: %f", 31000000000000000000000000000000000000.00, *((*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleFloat32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[2345000000000.44, 234880000.55],\"key2\":[20000.66,56556465.88]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float32(2345000000000.44) {
			t.Errorf("expected %f actual: %f", float32(2345000000000.44), *((*value)[0]))
		} else if *((*value)[1]) != float32(234880000.55) {
			t.Errorf("expected %f actual: %f",  float32(234880000.55), *((*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float32(20000.66) {
			t.Errorf("expected %f actual: %f", float32(20000.66), *((*value)[0]))
		} else if *((*value)[1]) != float32(56556465.88) {
			t.Errorf("expected %f actual: %f", float32(56556465.88), *((*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultiplFloat32HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[564355345465.88, 324234234.344]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat32("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != 564355345465.88 {
			t.Errorf("expected %f actual: %f", 564355345465.88, *(*value)[0])
		} else if (*(*value)[1]) != 324234234.344 {
			t.Errorf("expected %f actual: %f", 324234234.344, *(*value)[1])
		}
	}
}

func TestCannotParseArrayContainingSingleFloat32BoundaryWhenNilValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[null]}")
	
	{
		value, value_errors := json_obj.GetArrayOfFloat32Value("key") 
		if value_errors == nil {
			t.Errorf("expected to have errors actual: no errors")
		} 

		if value != nil {
			t.Errorf("expect array to be nil")
		}
	}
}
