package utils

import (
	"encoding/json"
	"fmt"
	"go-course/demo/app/shared/log"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
)

func IsNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func JsonToEntity(jsonIn string, entity interface{}) {
	err := json.Unmarshal([]byte(jsonIn), entity)

	if err != nil {
		entity = nil
	}
}

func EntityToJson(entity interface{}) string {
	str, err := json.MarshalIndent(entity,"","  ")
	if err != nil {
		return "{}"
	}
	fmt.Printf("%s\n", string(str))
	return string(str)
}

func ConvertEntity(in, out interface{}) interface{} {
	str, _ := json.Marshal(in)
	err2 := json.Unmarshal(str, out)

	if err2 != nil {
		return nil
	}
	return out
}

func ConvertHttpResponseBodyToString(response *http.Response) string {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	return string(bodyBytes)
}

func ConvertInterfaceToMap(i interface{}) map[string]interface{} {
	var result map[string]interface{}
	r, _ := json.Marshal(i)
	json.Unmarshal(r, &result)
	return result
}

func StringToInteger(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

func F64ToInteger(f float64) int {
	result := int(f)
	return result
}

func FormatDate(format string) string {
	switch {
	case format == "yyyy-MM-dd":
		return "2006-01-02"

	case format == "yyyy/MM/dd":
		return "2006/01/02"
	case format == "dd-mm-yyyy HH:mm:ss":
		return "02 Jan 06 15:04 -0700"
	default:
		return "2006-01-02"
	}
}

func DeepEqualString(array []string, val string) bool {
	var equal bool
	for _, value := range array {
		if reflect.DeepEqual(value, val) {
			equal = true
		}
	}
	return equal
}

func MatchAndReplace(regex, originalString, replace string) string {
	reg, err := regexp.Compile(regex)
	if err != nil {
		log.Fatal("Could not compile regex expression")
	}

	var newString string
	if reg != nil {
		newString = reg.ReplaceAllString(originalString, replace)
	}

	return newString
}
