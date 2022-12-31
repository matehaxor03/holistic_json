package tests
 
import (
    "testing"
	"fmt"
)

func TestCanParseEmptyMap(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":{}}")

	if !(json_obj.HasKey("key")) {
		t.Errorf("key not found available keys are %s", json_obj.Keys()[0])
	} else if json_obj.GetType("key") != "*json.Map" {
		t.Errorf("key is not a *json.Map: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetMap("key") 

		if value_errors != nil {
			t.Errorf("map GetMap has errors")
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len((value.Keys())) != 0 {
			t.Errorf("expected key length: length=0 actual: length=%d", len((value.Keys())))
		}
	}
}

func TestCanParseNestedMapWithStringValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":{\"key1\":\"value1\"}}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Map" {
		t.Errorf("key is not a *json.Map: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetMap("key") 
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len((value.Keys())) != 1 {
			t.Errorf("expected key length: length=1 actual: length=%d", len((value.Keys())))
		} else {
			inner_value, inner_value_errors := value.GetString("key1")
			if inner_value_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value_errors))
			} else if inner_value == nil {
				t.Errorf("key1 has nil value")
			} else if *inner_value != "value1" {
				t.Errorf("expected key1:\"value1\" actual:\"%s\"", *inner_value)
			}
		}
	}
}

func TestCanParseNestedMapWithMultipleStringValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":{\"key1\":\"value1\",\"key2\":\"value2\"}}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Map" {
		t.Errorf("key is not a *json.Map: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetMap("key") 
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len((value.Keys())) != 2 {
			t.Errorf("expected key length: length=1 actual: length=%d", len((value.Keys())))
		} else {
			inner_value, inner_value_errors := value.GetString("key1")
			if inner_value_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value_errors))
			} else if inner_value == nil {
				t.Errorf("key1 has nil value")
			} else if *inner_value != "value1" {
				t.Errorf("expected key1:\"value1\" actual:\"%s\"", *inner_value)
			}

			inner_value2, inner_value2_errors := value.GetString("key2")
			if inner_value2_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value2_errors))
			} else if inner_value2 == nil {
				t.Errorf("key2 has nil value")
			} else if *inner_value2 != "value2" {
				t.Errorf("expected key1:\"value2\" actual:\"%s\"", *inner_value2)
			}
		}
	}
}

func TestCanParseDoubleNestedMapWithStringValue(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":{\"key2\":{\"key3\":\"value3\"}}}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Map" {
		t.Errorf("key is not a *json.Map: %s", json_obj.GetType("key"))
	} else {			
		value, value_errors := json_obj.GetMap("key") 
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len((value.Keys())) != 1 {
			t.Errorf("expected key length: length=1 actual: length=%d", len((value.Keys())))
		} else {
			inner_value, inner_value_errors := value.GetMap("key2")
			if inner_value_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value_errors))
			} else if inner_value == nil {
				t.Errorf("key2 has nil value")
			} else if value.GetType("key2") != "*json.Map" {
				t.Errorf("key2 is not a *json.Map: %s", inner_value.GetType("key2"))
			} else {
				inner_value2, inner_value2_errors := inner_value.GetString("key3")
				if inner_value2_errors != nil {
					t.Errorf(fmt.Sprintf("%s", inner_value2_errors))
				} else if inner_value2 == nil {
					t.Errorf("key3 has nil value")
				} else if *inner_value2 != "value3" {
					t.Errorf("expected key3:\"value3\" actual:\"%s\"", *inner_value2)
				}
			}
		}
	}
}

func TestCanParseDoubleNestedMapWithStringValueAndStringValueAtRootLevelAfter(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{\"key\":{\"key2\":{\"key3\":\"value3\"}},\"key4\":\"value4\"}")

	if !json_obj.HasKey("key") {
		t.Errorf("key not found")
	} else if json_obj.GetType("key") != "*json.Map" {
		t.Errorf("key is not a *json.Map: %s", json_obj.GetType("key"))
	} else {		
		value, value_errors := json_obj.GetMap("key") 
		if value_errors != nil {
			t.Errorf(fmt.Sprintf("%s", value_errors))
		} else if value == nil {
			t.Errorf("GetMap is nil")
		} else if len((value.Keys())) != 1 {
			t.Errorf("expected keys length=1 actual: length=%d", len((value.Keys())))
		} else {
			inner_value, inner_value_errors := value.GetMap("key2")
			if inner_value_errors != nil {
				t.Errorf(fmt.Sprintf("%s", inner_value_errors))
			} else if inner_value == nil {
				t.Errorf("key2 has nil value")
			} else if value.GetType("key2") != "*json.Map" {
				t.Errorf("key2 is not a *json.Map: %s", inner_value.GetType("key2"))
			} else {
				inner_value2, inner_value2_errors := inner_value.GetString("key3")
				if inner_value2_errors != nil {
					t.Errorf(fmt.Sprintf("%s", inner_value2_errors))
				} else if inner_value2 == nil {
					t.Errorf("key3 has nil value")
				} else if *inner_value2 != "value3" {
					t.Errorf("expected key3:\"value3\" actual:\"%s\"", *inner_value2)
				}
			}
		}

		value4, inner_value4_errors := json_obj.GetString("key4")
		if inner_value4_errors != nil {
			t.Errorf(fmt.Sprintf("%s", inner_value4_errors))
		} else if value4 == nil {
			t.Errorf("key4 has nil value")
		} else if *value4 != "value4" {
			t.Errorf("expected key4:\"value4\" actual:\"%s\"", *value4)
		}
	}
}
