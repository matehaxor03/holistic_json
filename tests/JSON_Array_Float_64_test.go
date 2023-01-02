package tests
 
import (
    "testing"
	"fmt"
	"math"
)

// low boundary
func TestCanParseArrayContainingSingleFloat64Boundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[" + fmt.Sprintf("%f", (math.MaxFloat32 * 10)) +"]}")

	{		
		value, value_errors := json_obj.GetArrayOfFloat64("key") 

		if value_errors != nil {
			t.Errorf("error: %s", value_errors)
		} else if value == nil {
			t.Errorf("error: GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("error: expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64(math.MaxFloat32 * 10) {
			t.Errorf("error: expected: %f actual: %f", float64(math.MaxFloat32 * 10), *((*value)[0]))
		}
	}

	{		
		value, value_errors := json_obj.GetArrayOfFloat64Value("key") 

		if value_errors != nil {
			t.Errorf("error: %s", value_errors)
		} else if value == nil {
			t.Errorf("error: GetArray is nil")
		} else if len(value) != 1 {
			t.Errorf("error: expected: length=1 actual: length=%d", len(value))
		} else if ((value)[0]) != float64(math.MaxFloat32 * 10) {
			t.Errorf("error: expected: %f actual: %f", float64(math.MaxFloat32 * 10), ((value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleFloat64LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[" + fmt.Sprintf("%f", (math.MaxFloat32 * 10)) +"],\"key2\":[" + fmt.Sprintf("%f", (-1 * math.MaxFloat32 * 10)) +"]}")

	{		
		value, value_errors := json_obj.GetArrayOfFloat64("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64(math.MaxFloat32 * 10) {
			t.Errorf("expected %f actual: %f", float64(math.MaxFloat32 * 10), *((*value)[0]))
		}
	}

	{		
		value, value_errors := json_obj.GetArrayOfFloat64Value("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(value))
		} else if ((value)[0]) != float64(math.MaxFloat32 * 10) {
			t.Errorf("expected %f actual: %f", float64(math.MaxFloat32 * 10), ((value)[0]))
		}
	}

	{		
		value, value_errors := json_obj.GetArrayOfFloat64("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64(-1 * math.MaxFloat32 * 10) {
			t.Errorf("expected %f actual: %f", float64(-1 * math.MaxFloat32 * 10), *((*value)[0]))
		}
	}

	{		
		value, value_errors := json_obj.GetArrayOfFloat64Value("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(value))
		} else if ((value)[0]) != float64(-1 * math.MaxFloat32 * 10) {
			t.Errorf("expected %f actual: %f", float64(-1 * math.MaxFloat32 * 10), ((value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleFloat64LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[" + fmt.Sprintf("%f", (math.MaxFloat32 * 10)) +"," + fmt.Sprintf("%f", (math.MaxFloat32 * 100)) +"],\"key2\":[" + fmt.Sprintf("%f", (-1 * math.MaxFloat32 * 10)) +"," + fmt.Sprintf("%f", (-1 * math.MaxFloat32 * 100)) +"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64(math.MaxFloat32 * 10) {
			t.Errorf("expected %f actual: %f", (math.MaxFloat32 * 10), *((*value)[0]))
		} else if *((*value)[1]) != float64(math.MaxFloat32 * 100) {
			t.Errorf("expected %f actual: %f", float64(math.MaxFloat32 * 100), *((*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64(-1 * math.MaxFloat32 * 10) {
			t.Errorf("expected %f actual: %f", float64(-1 * math.MaxFloat32 * 10), *((*value)[0]))
		} else if *((*value)[1]) != float64(-1 * math.MaxFloat32 * 100) {
			t.Errorf("expected %f actual: %f", float64(-1 * math.MaxFloat32 * 100), *((*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultipleFloat64LowBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[" + fmt.Sprintf("%f", (math.MaxFloat32 * 10)) +"," + fmt.Sprintf("%f", (-1 * math.MaxFloat32 * 10)) +"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64(math.MaxFloat32 * 10) {
			t.Errorf("expected %f actual: %f", (math.MaxFloat32 * 10), *((*value)[0]))
		} else if *((*value)[1]) != float64(-1 * math.MaxFloat32 * 10) {
			t.Errorf("expected %f actual: %f", float64(-1 * math.MaxFloat32 * 10), *((*value)[1]))
		}
	}
}

// high boundary 
func TestCanParseArrayContainingSingleFloat64HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[" + fmt.Sprintf("%f", (math.MaxFloat64 / 10.0)) +"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64(math.MaxFloat64 / 10.0) {
			t.Errorf("expected %f actual: %f", float64(math.MaxFloat64 / 10.0), *((*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingSingleFloat64HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[" + fmt.Sprintf("%f", (math.MaxFloat64 / 10.0)) +"],\"key2\":[" + fmt.Sprintf("%f", (-1 * math.MaxFloat64 / 100.0)) +"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64(math.MaxFloat64 / 10.0) {
			t.Errorf("expected %f actual: %f", float64(math.MaxFloat64 / 10.0), *((*value)[0]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 1 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) !=  float64(-1 * math.MaxFloat64 / 100.0) {
			t.Errorf("expected %f actual: %f",  float64(-1 * math.MaxFloat64 / 100.0), *((*value)[0]))
		}
	}
}

func TestCanParseMultipleArraysContainingMultipleFloat64HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[" + fmt.Sprintf("%f", (math.MaxFloat64 / 10.0)) +"," + fmt.Sprintf("%f", (-1 * math.MaxFloat64 / 100.0)) +"],\"key2\":[" + fmt.Sprintf("%f", (math.MaxFloat64 / 1000.0)) +"," + fmt.Sprintf("%f", (-1 * math.MaxFloat64 / 10000.0)) +"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64((math.MaxFloat64 / 10.0)) {
			t.Errorf("expected %f actual: %f", float64((math.MaxFloat64 / 10.0)), *((*value)[0]))
		} else if *((*value)[1]) != float64((-1 * math.MaxFloat64 / 100.0)) {
			t.Errorf("expected %f actual: %f",  float64((-1 * math.MaxFloat64 / 100.0)), *((*value)[1]))
		}
	}

	if !json_obj.HasKey("key2") {
		t.Errorf("key2 not found")
	} else if !json_obj.IsArray("key2") {
		t.Errorf("key2 is not a *json.Array: %s", json_obj.GetType("key2"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key2") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=1 actual: length=%d", len(*value))
		} else if *((*value)[0]) != float64((math.MaxFloat64 / 1000.0)) {
			t.Errorf("expected %f actual: %f", float64((math.MaxFloat64 / 1000.0)), *((*value)[0]))
		} else if *((*value)[1]) != float64((-1 * math.MaxFloat64 / 10000.0)) {
			t.Errorf("expected %f actual: %f", float64((-1 * math.MaxFloat64 / 10000.0)), *((*value)[1]))
		}
	}
}

func TestCanParseArrayContainingMultiplFloat64HighBoundary(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[" + fmt.Sprintf("%f", (math.MaxFloat64 / 10.0)) +"," + fmt.Sprintf("%f", (-1 * math.MaxFloat64 / 100.0)) +"]}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if !json_obj.IsArray("key") {
		t.Errorf("key is not a *json.Array: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetArrayOfFloat64("key") 

		if value_errors != nil {
			t.Errorf("%s", value_errors)
		} else if value == nil {
			t.Errorf("GetArray is nil")
		} else if len(*value) != 2 {
			t.Errorf("expected: length=2 actual: length=%d", len(*value))
		} else if (*(*value)[0]) != float64(math.MaxFloat64 / 10.0) {
			t.Errorf("expected %f actual: %f", float64(math.MaxFloat64 / 10.0), *(*value)[0])
		} else if (*(*value)[1]) != float64((-1 * math.MaxFloat64 / 100.0)) {
			t.Errorf("expected %f actual: %f", float64(-1 * math.MaxFloat64 / 100.0), *(*value)[1])
		}
	}
}

func TestCannotParseArrayContainingSingleFloat64BoundaryWhenNilValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":[null]}")
	
	{
		value, value_errors := json_obj.GetArrayOfFloat64Value("key") 
		if value_errors == nil {
			t.Errorf("expected to have errors actual: no errors")
		} 

		if value != nil {
			t.Errorf("expect array to be nil")
		}
	}
}
