package json

import (
	"fmt"
	"strconv"
	"strings"
	"container/list"
	"net/url"
	"time"
	"os/exec"
	common "github.com/matehaxor03/holistic_common/common"
)

func cloneString(value *string) *string {
	if value == nil {
		return nil
	}

	temp := strings.Clone(*value)
	return &temp
}

func getTimeNow() *time.Time {
	now := time.Now()
	return &now
}

func Parse(s string) (*Map, []error) {
	var errors []error
	if s == "" {
		errors = append(errors, fmt.Errorf("error: value empty string"))
	}

	if !strings.HasPrefix(s, "{") {
		errors = append(errors, fmt.Errorf("error: json does not start with {"))
	}

	if !strings.HasSuffix(s, "}") {
		errors = append(errors, fmt.Errorf("error: json does not end with }"))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if s == "{}" {
		empty_map := Map{}
		return &empty_map, nil
	}

	runes := []rune(s)
	metrics := Map{}
	metrics.SetIntValue("{", 0)
	metrics.SetIntValue("}", 0)
	metrics.SetIntValue("[", 0)
	metrics.SetIntValue("]", 0)
	metrics.SetIntValue("opening_quote", 0)
	metrics.SetIntValue("closing_quote", 0)
	//metrics := Map{"{":Value{"value":0}, "}":Value{"value":0}, "[":Value{"value":0}, "]":Value{"value":0}, "opening_quote":Value{"value":0},"closing_quote":Value{"value":0}}
	mode := "looking_for_keys"
	parent_map_value := Map{}
	parent_map := Value{"value":&parent_map_value}
	list := list.New()
	list.PushFront(&parent_map)
	index := uint64(0)
	// parent map array and current map array etc
	result_error := parseJSONMap(&runes, &index, &mode, list, &metrics)

	opening_bracket_count, opening_bracket_count_errors := metrics.GetInt("{")
	closing_bracket_count, closing_bracket_count_errors := metrics.GetInt("}")
	opening_square_count, opening_square_count_errors := metrics.GetInt("[")
	closing_square_count, closing_square_count_errors := metrics.GetInt("]")
	opening_quote_count, opening_quote_count_errors := metrics.GetInt("opening_quote")
	closing_quote_count, closing_quote_count_errors := metrics.GetInt("closing_quote")


	if opening_bracket_count_errors != nil {
		errors = append(errors, opening_bracket_count_errors...)
	}

	if closing_bracket_count_errors != nil {
		errors = append(errors, closing_bracket_count_errors...)
	}

	if opening_square_count_errors != nil {
		errors = append(errors, opening_square_count_errors...)
	}

	if closing_square_count_errors != nil {
		errors = append(errors, closing_square_count_errors...)
	}

	if opening_quote_count_errors != nil {
		errors = append(errors, opening_quote_count_errors...)
	}

	if closing_quote_count_errors != nil {
		errors = append(errors, closing_quote_count_errors...)
	}

	if result_error != nil {
		errors = append(errors, result_error...)
	}

	if len(errors) > 0 {
		fmt.Println(s)
		return nil, errors
	}

	if *opening_bracket_count != *closing_bracket_count {
		errors = append(errors, fmt.Errorf("error: opening and closing brackets {} do not match, opening: %d closing: %d", *opening_bracket_count, *closing_bracket_count))
	}

	if *opening_square_count != *closing_square_count {
		errors = append(errors, fmt.Errorf("error: opening and closing squares [] do not match, opening: %d closing: %d", *opening_square_count, *closing_square_count))
	}

	if *opening_quote_count != *closing_quote_count {
		errors = append(errors, fmt.Errorf("error: opening and closing quotes do not match, opening: %d closing: %d", *opening_quote_count, *closing_quote_count))
	}

	temp_parent_map_value, temp_parent_map_value_errors := parent_map.GetMap()
	if temp_parent_map_value_errors != nil {
		errors = append(errors, temp_parent_map_value_errors...) 
	} else if common.IsNil(temp_parent_map_value) {
		errors = append(errors, fmt.Errorf("json.ParseJSON parent map is nil"))
	}

	if len(errors) > 0 {
		fmt.Println(s + " " + fmt.Sprintf("%s", errors))
		return nil, errors
	}

	
	return temp_parent_map_value, nil
}


func parseJSONMap(runes *[]rune, index *uint64, mode *string, list *list.List, metrics *Map) ([]error) {
	var errors []error
	if list == nil {
		errors = append(errors, fmt.Errorf("error: list is nil"))
	} else if (*list).Len() == 0 {
		errors = append(errors, fmt.Errorf("error: list is empty"))
	}

	if len(errors) > 0 {
		return errors
	}
	
	mode_looking_for_keys := "looking_for_keys"
	mode_looking_for_key_name := "looking_for_key_name"
	mode_looking_for_key_name_column := "looking_for_key_name_column"
	mode_looking_for_value := "looking_for_value"
	
	temp_key := ""
	temp_value := ""
	parsing_string := false
	found_value := false
	temp_mode := cloneString(mode)
	current_mode := *temp_mode

	for *index < uint64(len(*runes)) {
		i := *index
		value := (*runes)[*index]
		
		//fmt.Println(current_mode + " " + string(value))
		//fmt.Println(temp_key + " " + temp_value + " parsing:" + fmt.Sprintf("%s", parsing_string) + "found value: " +  fmt.Sprintf("%s", found_value))

		if !parsing_string {
			if string(value) == "{" {
				opening_count, _ := metrics.GetInt("{")
				*opening_count++
				metrics.SetInt("{", opening_count)
			}

			if string(value) == "}" {
				closing_count, _ := metrics.GetInt("}")
				*closing_count++
				metrics.SetInt("}", closing_count)
			}

			if string(value) == "[" {
				opening_count, _ := metrics.GetInt("[")
				*opening_count++
				metrics.SetInt("[", opening_count)
			}

			if string(value) == "]" {
				closing_count, _ := metrics.GetInt("]")
				*closing_count++
				metrics.SetInt("]", closing_count)
			}
		}
		

		if current_mode == mode_looking_for_keys {
			if string(value) == "\"" && string((*runes)[i-1]) != "\\"{
				current_mode = mode_looking_for_key_name
				parsing_string = true

				opening_count, _ := metrics.GetInt("opening_quote")
				*opening_count++
				metrics.SetInt("opening_quote", opening_count)
			} else if string(value) == "}" {
				if list.Len() > 1 {
					list.Remove(list.Front())
					front_value := (list.Front().Value).(*Value)
					if front_value.IsMap() {
						current_mode = mode_looking_for_keys
					} else {
						current_mode = mode_looking_for_value
					} /*else {
						errors = append(errors, fmt.Errorf("json.Parse front is neither a Map or Array", front_value.GetType()))
						return errors
					}*/
				}	
			} else if string(value) == "]" {
				// this should not occur for valid json todo throw error
				/*if list.Len() > 1 {
					list.Remove(list.Front())
					if fmt.Sprintf("%T", list.Front().Value) == "*json.Map" {
						current_mode = mode_looking_for_keys
					} else {
						current_mode = mode_looking_for_value
					}
				}*/
			}
		} else if current_mode == mode_looking_for_key_name {
			if string(value) == "\"" && string((*runes)[i-1]) != "\\" {
				current_mode = mode_looking_for_key_name_column
				parsing_string = false

				opening_count, _ := metrics.GetInt("closing_quote")
				*opening_count++
				metrics.SetInt("closing_quote", opening_count)
			} else {
				temp_key += string(value)
			}
		} else if current_mode == mode_looking_for_key_name_column {
			if string(value) == ":" {
				current_mode = mode_looking_for_value
			}
		} else if current_mode == mode_looking_for_value {
			
			if !found_value && (string(value) == " " || string(value) == "\r" || string(value) == "\n" || string(value) == "\t") ||  string(value) == "\\" {
				*index++
				continue
			} else {
				found_value = true
			}

			if !parsing_string {
				if i > 0 && !parsing_string && string(value) == "\"" && string((*runes)[i-1]) != "\\" {
					temp_value += string(value)
					parsing_string = true
					
					opening_count, _ := metrics.GetInt("opening_quote")
					*opening_count++
					metrics.SetInt("opening_quote", opening_count)

				} else if i == 0 && !parsing_string && string(value) == "\"" {
					temp_value += string(value)
					parsing_string = true
					
					opening_count, _ := metrics.GetInt("opening_quote")
					*opening_count++
					metrics.SetInt("opening_quote", opening_count)

				} else if string(value) == "{" {
					new_mode := mode_looking_for_keys
					new_map := Map{}
					new_map_value := Value{"value":new_map}

					front_value := (list.Front().Value).(*Value)
					if front_value.IsMap() {
						front_value.SetValue(temp_key, &new_map_value)	
					} else {
						temp_array, temp_array_errors := front_value.GetArray()
						if temp_array_errors != nil {
							return temp_array_errors
						} else if common.IsNil(temp_array) {
							errors = append(errors, fmt.Errorf("json.ParseJSON array is nil"))
							return errors
						}
						*temp_array = append(*temp_array, new_map_value)
					} 
					list.PushFront(&new_map_value)


					found_value = false
					*index++
					parseJSONMap(runes, index, &new_mode, list, metrics)
				} else if string(value) == "[" {
					new_mode := mode_looking_for_value
					new_array := Array{}
					new_array_value := Value{"value":new_array}

					front_value := (list.Front().Value).(*Value)
					if front_value.IsMap() {
						front_value.SetValue(temp_key, &new_array_value)	
					} else {
						temp_array, temp_array_errors := front_value.GetArray()
						if temp_array_errors != nil {
							return temp_array_errors
						} else if common.IsNil(temp_array) {
							errors = append(errors, fmt.Errorf("json.ParseJSON array is nil"))
							return errors
						}
						*temp_array = append(*temp_array, new_array_value)
					} 

					list.PushFront(&new_array_value)

					
					found_value = false
					*index++
					parseJSONMap(runes, index, &new_mode, list, metrics)
				} else if string(value) == "}" {
					parse_errors := parseJSONValue(temp_key, temp_value, list)
					if parse_errors != nil {
						errors = append(errors, parse_errors...)
					}

					temp_key = ""
					temp_value = ""

					found_value = false
					list.Remove(list.Front())

					if list.Len() > 0 {
						front_value := (list.Front().Value).(*Value)
						if front_value.IsMap() {
							current_mode = mode_looking_for_keys
						} else {
							current_mode = mode_looking_for_value
						}
					} 

				} else if string(value) == "]" {
					parse_errors := parseJSONValue(temp_key, temp_value, list)
					if parse_errors != nil {
						errors = append(errors, parse_errors...)
					}

					temp_key = ""
					temp_value = ""

					found_value = false
					list.Remove(list.Front())

					front_value := (list.Front().Value).(*Value)
					if front_value.IsMap() {
						current_mode = mode_looking_for_keys
					} else {
						current_mode = mode_looking_for_value
					} 

				} else if string(value) == "," {
					parse_errors := parseJSONValue(temp_key, temp_value, list)
					if parse_errors != nil {
						errors = append(errors, parse_errors...)
					}
					
					temp_key = ""
					temp_value = ""

					front_value := (list.Front().Value).(*Value)
					if front_value.IsMap() {
						current_mode = mode_looking_for_keys
					} else {
						current_mode = mode_looking_for_value
					} 
		
					found_value = false
				} else {
					temp_value += string(value)
				}
			} else if !parsing_string && string(value) == "\"" && string((*runes)[i-1]) != "\\" {
				temp_value += string(value)
				parsing_string = true

				opening_quote, _ := metrics.GetInt("opening_quote")
				*opening_quote++
				metrics.SetInt("opening_quote", opening_quote)

			} else if parsing_string && string(value) == "\"" && string((*runes)[i-1]) != "\\" {
				temp_value += string(value)
				parsing_string = false
				
				parse_errors := parseJSONValue(temp_key, temp_value, list)
				if parse_errors != nil {
					errors = append(errors, parse_errors...)
				}
				
				temp_key = ""
				temp_value = ""

				front_value := (list.Front().Value).(*Value)
				if front_value.IsMap() {
					current_mode = mode_looking_for_keys
				} else {
					current_mode = mode_looking_for_value
				} 
	
				found_value = false

				closing_quote, _ := metrics.GetInt("closing_quote")
				*closing_quote++
				metrics.SetInt("closing_quote", closing_quote)
			} else {
				temp_value += string(value)
			}

			
		}

		if len(errors) > 0 {
			return errors
		}

		*index++
	}

	return nil
}

func parseJSONValue(temp_key string, temp_value string, list *list.List) []error {
	var errors []error
	
	if list == nil {
		errors = append(errors, fmt.Errorf("error: list is nil"))
	} else if (*list).Len() == 0 {
		errors = append(errors, fmt.Errorf("error: list is empty"))
	}

	if len(errors) > 0 {
		return errors
	}

	temp_key = strings.ReplaceAll(temp_key, "\\\"", "\"")
	temp_value = strings.ReplaceAll(temp_value, "\\\"", "\"")
	
	data_type := ""
	string_value := cloneString(&temp_value)
	
	var boolean_value *bool
	var float64_value *float64
	var float32_value *float32
	var int64_value *int64
	var int32_value *int32
	var int16_value *int16
	var int8_value *int8
	var uint64_value *uint64
	var uint32_value *uint32
	var uint16_value *uint16
	var uint8_value *uint8

	//if strings.HasPrefix(*string_value, "\"") || strings.HasSuffix(*string_value, "\"") {
	string_temp := strings.TrimSpace(*string_value)
	string_value = &string_temp
	//}


	if strings.HasPrefix(*string_value, "\"") && strings.HasSuffix(*string_value, "\"") {
		data_type = "string"
		dequoted_value := (*string_value)[1:(len(*string_value)-1)]
		string_value = &dequoted_value	
	} else if strings.HasPrefix(*string_value, "\"") && !strings.HasSuffix(*string_value, "\"") {
		*string_value = strings.TrimSpace(*string_value)
		if !strings.HasPrefix(*string_value, "\"") && !strings.HasSuffix(*string_value, "\"") {
			errors = append(errors, fmt.Errorf("error: value has \" as prefix but not \" as suffix"))
			return errors
		} else {
			data_type = "string"
			dequoted_value := (*string_value)[1:(len(*string_value)-1)]
			string_value = &dequoted_value	
		}
	} else if !strings.HasPrefix(*string_value, "\"") && strings.HasSuffix(*string_value, "\"") {
		*string_value = strings.TrimSpace(*string_value)
		if !strings.HasPrefix(*string_value, "\"") && !strings.HasSuffix(*string_value, "\"") {
			errors = append(errors, fmt.Errorf("error: value has \" as suffix but not \" as prefix"))
			return errors
		} else {
			data_type = "string"
			dequoted_value := (*string_value)[1:(len(*string_value)-1)]
			string_value = &dequoted_value	
		}
	} else {
		// when parsing emtpy array []
		if *string_value == "" {
			return nil
		}

		if *string_value == "true" {
			data_type = "bool"
			boolean_value_true := true 
			boolean_value = &boolean_value_true
		} else if *string_value == "false" {
			data_type = "bool"
			boolean_value_false := false 
			boolean_value = &boolean_value_false
		} else if *string_value == "null" {
			data_type = "null"
		} else {
			var negative_number bool
			negative_number_count := strings.Count(*string_value, "-")
			if negative_number_count == 1 {
				negative_number = true
				if !strings.HasPrefix(*string_value, "-") {
					errors = append(errors, fmt.Errorf("error: negative symbol is not at the start of the number"))
				}
			} else if negative_number_count == 0 {
				negative_number = false
			} else {
				errors = append(errors, fmt.Errorf("error: value contained %d negative symbols expected 1", negative_number_count))
			}

			var decimal_number bool
			decimal_count := strings.Count(*string_value, ".")
			if decimal_count == 1 {
				decimal_number = true
			} else if decimal_count == 0 {
				decimal_number = false
			} else {
				errors = append(errors, fmt.Errorf("error: value contained %d decimal points expected 1", decimal_count))
			}

			if len(errors) > 0 {
				return errors
			}

			if decimal_number {
				data_type = "float64"
				float64_temp, float64_temp_error := strconv.ParseFloat(*string_value, 64)
				if float64_temp_error != nil {
					errors = append(errors, fmt.Errorf("error: strconv.ParseFloat(*string_value, 64) error"))
				} else {
					float64_value = &float64_temp
					
					float32_temp, float32_temp_error := strconv.ParseFloat(*string_value, 32)
					if float32_temp_error != nil {
					} else {
						data_type = "float32"
						float32_conv := float32(float32_temp)
						float32_value = &float32_conv
					}
				}

				if len(errors) > 0 {
					return errors
				}
			} else {
				if negative_number {
					data_type = "int64"
					int64_temp, int64_temp_error := strconv.ParseInt(*string_value, 10, 64)
					if int64_temp_error != nil {
						errors = append(errors, fmt.Errorf("error: strconv.ParseInt(*string_value, 10, 64) error"))
					} else {
						int64_value = &int64_temp
						if *int64_value >= -128 && *int64_value <= 127 {
							data_type = "int8"
							int8_temp, int8_temp_error := strconv.ParseInt(*string_value, 10, 8)
							if int8_temp_error != nil {
								errors = append(errors, fmt.Errorf("error: strconv.ParseInt(*string_value, 10, 8) error"))
							} else {
								int8_conv := int8(int8_temp)
								int8_value = &int8_conv
							}
						} else if *int64_value >= -32768 && *int64_value <= 32767 {
							data_type = "int16"
							int16_temp, int16_temp_error := strconv.ParseInt(*string_value, 10, 16)
							if int16_temp_error != nil {
								errors = append(errors, fmt.Errorf("error: strconv.ParseInt(*string_value, 10, 16) error"))
							} else {
								int16_conv := int16(int16_temp)
								int16_value = &int16_conv
							}
						} else if *int64_value >= -2147483648 && *int64_value <= 2147483647 {
							data_type = "int32"
							int32_temp, int32_temp_error := strconv.ParseInt(*string_value, 10, 32)
							if int32_temp_error != nil {
								errors = append(errors, fmt.Errorf("error: strconv.ParseInt(*string_value, 10, 32) error"))
							} else {
								int32_conv := int32(int32_temp)
								int32_value = &int32_conv
							}
						}
					}

					if len(errors) > 0 {
						return errors
					}

				} else {
					data_type = "uint64"
					uint64_temp, uint64_temp_error := strconv.ParseUint(*string_value, 10, 64)
					if uint64_temp_error != nil {
						errors = append(errors, fmt.Errorf("error: strconv.ParseUint(*string_value, 10, 64) error"))
					} else {
						uint64_value = &uint64_temp
						if *uint64_value >= 0 && *uint64_value <= 255 {
							data_type = "uint8"
							int8_temp, int8_temp_error := strconv.ParseUint(*string_value, 10, 8)
							if int8_temp_error != nil {
								errors = append(errors, fmt.Errorf("error:strconv.ParseUInt(*string_value, 10, 8) error"))
							} else {
								int8_conv := uint8(int8_temp)
								uint8_value = &int8_conv
							}
						} else if *uint64_value >= 256 && *uint64_value <= 65535 {
							data_type = "uint16"
							int16_temp, int16_temp_error := strconv.ParseUint(*string_value, 10, 16)
							if int16_temp_error != nil {
								errors = append(errors, fmt.Errorf("error: strconv.ParseUInt(*string_value, 10, 16) error"))
							} else {
								int16_conv := uint16(int16_temp)
								uint16_value = &int16_conv
							}
						} else if *uint64_value >= 65536  && *uint64_value <= 4294967295 {
							data_type = "uint32"
							int32_temp, int32_temp_error := strconv.ParseUint(*string_value, 10, 32)
							if int32_temp_error != nil {
								errors = append(errors, fmt.Errorf("error: strconv.ParseUInt(*string_value, 10, 32) error"))
							} else {
								int32_conv := uint32(int32_temp)
								uint32_value = &int32_conv
							}
						}
					}

					if len(errors) > 0 {
						return errors
					}

				}
			}

			if len(errors) > 0 {
				return errors
			}
		}
	}

	if data_type == "" {
		errors = append(errors, fmt.Errorf("error: data_type is unknown please implement"))
	}

	if len(errors) > 0 {
		return errors
	}

	front_value := (list.Front().Value).(*Value)

	if front_value.IsArray() {
		value_as_array, value_as_array_errors := front_value.GetArray()
		if value_as_array_errors != nil {
			errors = append(errors, value_as_array_errors...)
		} else if common.IsNil(value_as_array) {
			errors = append(errors, fmt.Errorf("json.parseJSONValue array is nil"))
		}

		if len(errors) > 0 {
			return errors
		}

		if data_type == "string" {
			value_as_array.AppendString(string_value)
		} else if data_type == "bool" {
			value_as_array.AppendBool(boolean_value)
		} else if data_type == "null" {
			value_as_array.AppendNil()
		} else if data_type == "float32" {
			value_as_array.AppendFloat32(float32_value)
		} else if data_type == "float64" {
			value_as_array.AppendFloat64(float64_value)
		} else if data_type == "int8" {
			value_as_array.AppendInt8(int8_value)
		} else if data_type == "int16" {
			value_as_array.AppendInt16(int16_value)
		} else if data_type == "int32" {
			value_as_array.AppendInt32(int32_value)
		}  else if data_type == "int64" {
			value_as_array.AppendInt64(int64_value)
		} else if data_type == "uint8" {
			value_as_array.AppendUInt8(uint8_value)
		} else if data_type == "uint16" {
			value_as_array.AppendUInt16(uint16_value)
		} else if data_type == "uint32" {
			value_as_array.AppendUInt32(uint32_value)
		} else if data_type == "uint64" {
			value_as_array.AppendUInt64(uint64_value)
		} else {
			errors = append(errors, fmt.Errorf("json.PaseJSONValue type is not supported for Array %s", data_type))
		}
	} else if front_value.IsMap() {
		value_as_map, value_as_map_errors := front_value.GetMap()
		if value_as_map_errors != nil {
			errors = append(errors, value_as_map_errors...)
		} else if common.IsNil(value_as_map) {
			errors = append(errors, fmt.Errorf("json.parseJSONValue map is nil"))
		}

		if len(errors) > 0 {
			return errors
		}

		if data_type == "string" {
			value_as_map.SetString(temp_key, string_value)
		} else if data_type == "bool" {
			value_as_map.SetBool(temp_key, boolean_value)
		} else if data_type == "null" {
			value_as_map.SetNil(temp_key)
		} else if data_type == "float64" {
			value_as_map.SetFloat64(temp_key, float64_value)
		} else if data_type == "float32" {
			value_as_map.SetFloat32(temp_key, float32_value)
		} else if data_type == "int8" {
			value_as_map.SetInt8(temp_key, int8_value)
		} else if data_type == "int16" {
			value_as_map.SetInt16(temp_key, int16_value)
		} else if data_type == "int32" {
			value_as_map.SetInt32(temp_key, int32_value)
		} else if data_type == "int64" {
			value_as_map.SetInt64(temp_key, int64_value)
		} else if data_type == "uint8" {
			value_as_map.SetUInt8(temp_key, uint8_value)
		} else if data_type == "uint16" {
			value_as_map.SetUInt16(temp_key, uint16_value)
		} else if data_type == "uint32" {
			value_as_map.SetUInt32(temp_key, uint32_value)
		} else if data_type == "uint64" {
			value_as_map.SetUInt64(temp_key, uint64_value)
		} else {
			errors = append(errors, fmt.Errorf("json.PaseJSONValue type is not supported for Map %s", data_type))
		}
	} else {
		errors = append(errors, fmt.Errorf("json.parseJSONValue is neither a map or an array it's %s", front_value.GetType()))
	}

	if len(errors) > 0 {
		return errors
	}


	return nil
}

func ConvertInterfaceValueToJSONStringValue(json *strings.Builder, value interface{}) ([]error) {
	var errors []error
	if json == nil {
		errors = append(errors, fmt.Errorf("error: *strings.Builder is nil"))
		return errors
	}

	if value == nil {
		json.WriteString("null")
		return nil
	}

	string_value := fmt.Sprintf("%s", value)

	if string_value == "<nil>" {
		json.WriteString("null")
		return nil
	}

	rep := fmt.Sprintf("%T", value)

	if string_value == "%!s("+rep+"=<nil>)" {
		json.WriteString("null")
		return nil
	}
	
	switch rep {
	case "json.Value":
		ConvertInterfaceValueToJSONStringValue(json, (value.(Value))["value"])
	case "*json.Value":
		ConvertInterfaceValueToJSONStringValue(json, (*(value.(*Value)))["value"])
	case "string":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll(value.(string), "\"", "\\\""))
		json.WriteString( "\"")
	case "*string":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll(*(value.(*string)), "\"", "\\\""))
		json.WriteString("\"")
	case "error":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll(value.(error).Error(), "\"", "\\\""))
		json.WriteString("\"")
	case "*error":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll((*(value.(*error))).Error(), "\"", "\\\""))
		json.WriteString("\"")
	case "*url.Error":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll((*(value.(*url.Error))).Error(), "\"", "\\\""))
		json.WriteString("\"")
	case "exec.ExitError":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll(fmt.Sprintf("%s", value), "\"", "\\\""))
		json.WriteString("\"")
	case "*exec.ExitError":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll(fmt.Sprintf("%s", value), "\"", "\\\""))
		json.WriteString("\"")
	case "errors.errorString":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll(fmt.Sprintf("%s", value), "\"", "\\\""))
		json.WriteString("\"")
	case "*errors.errorString":
		json.WriteString("\"")
		json.WriteString(strings.ReplaceAll(fmt.Sprintf("%s", value), "\"", "\\\""))
		json.WriteString("\"")
	case "bool":
		if value.(bool) {
			json.WriteString("true")
		} else {
			json.WriteString("false")
		}
	case "*bool":
		if *(value.(*bool)) {
			json.WriteString("true")
		} else {
			json.WriteString("false")
		}
	case "json.Map":
		x_error := value.(Map).ToJSONString(json)
		if x_error != nil {
			errors = append(errors, x_error...)
		}
	case "*json.Map":
		x_error := (*value.(*Map)).ToJSONString(json)
		if x_error != nil {
			errors = append(errors, x_error...)
		}
	case "json.Array":
		x_error := value.(Array).ToJSONString(json)
		if x_error != nil {
			errors = append(errors, x_error...)
		} 
	case "*json.Array":
		x_error := (*(value.(*Array))).ToJSONString(json)
		if x_error != nil {
			errors = append(errors, x_error...)
		}
	case "[]string":
		json.WriteString("[")
		array_length := len(value.([]string))
		for array_index, array_value := range value.([]string) {
			string_conversion_error := ConvertInterfaceValueToJSONStringValue(json, array_value)
			if string_conversion_error != nil {
				errors = append(errors, string_conversion_error...)
			} 

			if array_index < (array_length - 1) {
				json.WriteString(",")
			}
		}
		json.WriteString("]")
	case "*[]string":
		json.WriteString("[")
		array_length := len(*(value.(*[]string)))
		for array_index, array_value := range *(value.(*[]string)) {
			string_conversion_error := ConvertInterfaceValueToJSONStringValue(json, array_value)
			if string_conversion_error != nil {
				errors = append(errors, string_conversion_error...)
			}

			if array_index < (array_length - 1) {
				json.WriteString(",")
			}
		}
		json.WriteString("]")
	case "[]error":
		json.WriteString("[")
		array_length := len(value.([]error))
		for array_index, array_value := range value.([]error) {
			string_conversion_error := ConvertInterfaceValueToJSONStringValue(json, array_value)
			if string_conversion_error != nil {
				errors = append(errors, string_conversion_error...)
			} 

			if array_index < (array_length - 1) {
				json.WriteString(",")
			}
		}
		json.WriteString("]")
	case "*[]error":
		json.WriteString("[")
		array_length := len(*(value.(*[]error)))
		for array_index, array_value := range *(value.(*[]error)) {
			string_conversion_error := ConvertInterfaceValueToJSONStringValue(json, array_value)
			if string_conversion_error != nil {
				errors = append(errors, string_conversion_error...)
			} 

			if array_index < (array_length - 1) {
				json.WriteString(",")
			}
		}
		json.WriteString("]")
	case "func(string, *string, string, string) []error", "func(json.Map) []error", "*func(json.Map) []error":
		json.WriteString(strings.ReplaceAll(fmt.Sprintf("%s", rep), "\"", "\\\""))
	case "*time.Time":
		json.WriteString("\"")
		json.WriteString((*(value.(*time.Time))).Format("2006-01-02 15:04:05.000000"))
		json.WriteString("\"")
	case "map[string]map[string][][]string":
		json.WriteString("\"map[string]map[string][][]string\"")
	case "*uint64":
		json.WriteString(strconv.FormatUint(*(value.(*uint64)), 10))
	case "uint64":
		json.WriteString(strconv.FormatUint(value.(uint64), 10))
	case "*uint32":
		json.WriteString(strconv.FormatUint(uint64(*(value.(*uint32))), 10))
	case "uint32":
		json.WriteString(strconv.FormatUint(uint64(value.(uint32)), 10))
	case "*uint16":
		json.WriteString(strconv.FormatUint(uint64(*(value.(*uint16))), 10))
	case "uint16":
		json.WriteString(strconv.FormatUint(uint64(value.(uint16)), 10))
	case "*uint8":
		json.WriteString(strconv.FormatUint(uint64(*(value.(*uint8))), 10))
	case "uint8":
		json.WriteString(strconv.FormatUint(uint64(value.(uint8)), 10))
	case "*uint":
		json.WriteString(strconv.FormatUint(uint64(*(value.(*uint))), 10))
	case "uint":
		json.WriteString(strconv.FormatUint(uint64(value.(uint)), 10))
	case "*int64":
		json.WriteString(strconv.FormatInt(*(value.(*int64)), 10))
	case "int64":
		json.WriteString(strconv.FormatInt(value.(int64), 10))
	case "*int32":
		json.WriteString(strconv.FormatInt(int64(*(value.(*int32))), 10))
	case "int32":
		json.WriteString(strconv.FormatInt(int64((value.(int32))), 10))
	case "*int16":
		json.WriteString(strconv.FormatInt(int64(*(value.(*int16))), 10))
	case "int16":
		json.WriteString(strconv.FormatInt(int64((value.(int16))), 10))
	case "*int8":
		json.WriteString(strconv.FormatInt(int64(*(value.(*int8))), 10))
	case "int8":
		json.WriteString(strconv.FormatInt(int64((value.(int8))), 10))
	case "*int":
		json.WriteString(strconv.FormatInt(int64(*(value.(*int))), 10))
	case "int":
		json.WriteString(strconv.FormatInt(int64(value.(int)), 10))
	case "*float64":
		json.WriteString(fmt.Sprintf("%f", *(value.(*float64))))
	case "float64":
		json.WriteString(fmt.Sprintf("%f", (value.(float64))))
	case "*float32":
		json.WriteString(fmt.Sprintf("%f", *(value.(*float32))))
	case "float32":
		json.WriteString(fmt.Sprintf("%f", (value.(float32))))
	default:
		errors = append(errors, fmt.Errorf("error: JSON.ConvertInterfaceValueToJSONStringValue: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return errors
	}


	return nil
}

func ConvertInterfaceValueToStringValue(value interface{}) (*string, []error) {
	var errors []error
	var result string
	
	if value == nil {
		result = "null"
		return &result, nil
	}

	string_value := fmt.Sprintf("%s", value)

	if string_value == "<nil>" {
		result = "null"
		return &result, nil
	}

	rep := fmt.Sprintf("%T", value)

	if string_value == "%!s("+rep+"=<nil>)" {
		result = "null"
		return &result, nil
	}
	
	switch rep {
	case "string":
		result = value.(string)
	case "*string":
		result = *(value.(*string))
	case "error":
		result = fmt.Sprintf("%s",  value.(error).Error())
	case "*error":
		result = fmt.Sprintf("%s", (*(value.(*error))).Error())
	case "*url.Error":
		result = fmt.Sprintf("%s", (*(value.(*url.Error))).Error()) 
	case "exec.ExitError":
		result = fmt.Sprintf("%s", value.(exec.ExitError))
	case "*exec.ExitError":
		result = fmt.Sprintf("%s", *(value.(*exec.ExitError)))
	case "errors.errorString":
		result = fmt.Sprintf("%s", value)
	case "*errors.errorString":
		result = fmt.Sprintf("%s", value)
	case "bool":
		if value.(bool) {
			result = "true"
		} else {
			result = "false"
		}
	case "*bool":
		if *(value.(*bool)) {
			result = "true"
		} else {
			result = "false"
		}
	case "*time.Time":
		result = (*(value.(*time.Time))).Format("2006-01-02 15:04:05.000000")
	case "map[string]map[string][][]string":
		result = "map[string]map[string][][]string"
	case "*uint64":
		result = strconv.FormatUint(*(value.(*uint64)), 10)
	case "uint64":
		result = strconv.FormatUint(value.(uint64), 10)
	case "*uint32":
		result = strconv.FormatUint(uint64(*(value.(*uint32))), 10)
	case "uint32":
		result = strconv.FormatUint(uint64(value.(uint32)), 10)
	case "*uint16":
		result = strconv.FormatUint(uint64(*(value.(*uint16))), 10)
	case "uint16":
		result = strconv.FormatUint(uint64(value.(uint16)), 10)
	case "*uint8":
		result = strconv.FormatUint(uint64(*(value.(*uint8))), 10)
	case "uint8":
		result = strconv.FormatUint(uint64(value.(uint8)), 10)
	case "*uint":
		result = strconv.FormatUint(uint64(*(value.(*uint))), 10)
	case "uint":
		result = strconv.FormatUint(uint64(value.(uint)), 10)
	case "*int64":
		result = strconv.FormatInt(*(value.(*int64)), 10)
	case "int64":
		result = strconv.FormatInt(value.(int64), 10)
	case "*int32":
		result = strconv.FormatInt(int64(*(value.(*int32))), 10)
	case "int32":
		result = strconv.FormatInt(int64((value.(int32))), 10)
	case "*int16":
		result = strconv.FormatInt(int64(*(value.(*int16))), 10)
	case "int16":
		result = strconv.FormatInt(int64((value.(int16))), 10)
	case "*int8":
		result = strconv.FormatInt(int64(*(value.(*int8))), 10)
	case "int8":
		result = strconv.FormatInt(int64((value.(int8))), 10)
	case "*int":
		result = strconv.FormatInt(int64(*(value.(*int))), 10)
	case "int":
		result = strconv.FormatInt(int64(value.(int)), 10)
	case "*float64":
		result = fmt.Sprintf("%f", *(value.(*float64)))
	case "float64":
		result = fmt.Sprintf("%f", (value.(float64)))
	case "*float32":
		result = fmt.Sprintf("%f", *(value.(*float32)))
	case "float32":
		result = fmt.Sprintf("%f", (value.(float32)))
	case "json.Value":
		return ConvertInterfaceValueToStringValue((value.(Value))["value"])
	case "*json.Value":
		return ConvertInterfaceValueToStringValue((*(value.(*Value)))["value"])
	default:
		errors = append(errors, fmt.Errorf("error: JSON.ConvertInterfaceValueToJSONStringValue: type %s is not supported please implement", rep))
	}

	if len(errors) > 0 {
		return nil, errors
	}

	return &result, nil
}
