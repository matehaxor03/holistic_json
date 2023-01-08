package tests
 
import (
    "testing"
	"strings"
	"fmt"
	json "github.com/matehaxor03/holistic_json/json"
)

func ParseJSONSuccessfully(t *testing.T, json_string string) (*json.Map) {
	json, json_errors := json.Parse(json_string)

	if json_errors != nil {
		t.Errorf("%s", json_errors)
		t.FailNow()
	} else if json == nil {
		t.Errorf("json is nil")
		t.FailNow()
	} else {
		PrintJSON(t, json)
	}
	return json
}

func PrintJSON(t *testing.T, json *json.Map) {
	if json == nil {
		t.Errorf("json is nil")
		t.FailNow()
	} else {
		var json_string strings.Builder
		json_string_errors := json.ToJSONString(&json_string)
		if json_string_errors != nil {
			t.Errorf("%s", json_string_errors)
			t.FailNow()
		} else {
			fmt.Println(json_string.String())
			fmt.Println(json.GetKeys())
		}
	}
}

func TestCanParseMinimal(t *testing.T) {
	json_obj := ParseJSONSuccessfully(t, "{}")

	if len((json_obj.GetKeys())) != 0  {
		t.Errorf("keys is not zero")
	}
}

func TestCannotParseEmptyString(t *testing.T) {
	value, value_errors := json.Parse("")

	if value_errors == nil {
		t.Errorf("no error on parse")
	} 

	if value != nil {
		t.Errorf("map is not nil")
	}
}

func TestCannotParseMalformedBrackets1(t *testing.T) {
	value, value_errors := json.Parse("{")

	if value_errors == nil {
		t.Errorf("no error on parse")
	} 

	if value != nil {
		t.Errorf("map is not nil")
	}
}

func TestCannotParseMalformedBrackets2(t *testing.T) {
	value, value_errors := json.Parse("}")

	if value_errors == nil {
		t.Errorf("no error on parse")
	} 

	if value != nil {
		t.Errorf("map is not nil")
	}
}



