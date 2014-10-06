package gophercloud

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	th "github.com/rackspace/gophercloud/testhelper"
)

func TestMaybeString(t *testing.T) {
	testString := ""
	var expected *string
	actual := MaybeString(testString)
	th.CheckDeepEquals(t, expected, actual)

	testString = "carol"
	expected = &testString
	actual = MaybeString(testString)
	th.CheckDeepEquals(t, expected, actual)
}

func TestMaybeInt(t *testing.T) {
	testInt := 0
	var expected *int
	actual := MaybeInt(testInt)
	th.CheckDeepEquals(t, expected, actual)

	testInt = 4
	expected = &testInt
	actual = MaybeInt(testInt)
	th.CheckDeepEquals(t, expected, actual)
}

func TestBuildQueryStringWithPointerToStruct(t *testing.T) {
	expected := &url.URL{
		RawQuery: "j=2&r=red",
	}

	type Opts struct {
		J int    `q:"j"`
		R string `q:"r"`
		C bool
	}

	opts := Opts{J: 2, R: "red"}

	actual, err := BuildQueryString(&opts)
	if err != nil {
		t.Errorf("Error building query string: %v", err)
	}

	th.CheckDeepEquals(t, expected, actual)
}

func TestBuildQueryStringWithoutRequiredFieldSet(t *testing.T) {
	type Opts struct {
		J int    `q:"j"`
		R string `q:"r,required"`
		C bool
	}

	opts := Opts{J: 2, C: true}

	_, err := BuildQueryString(&opts)
	if err == nil {
		t.Error("Unexpected result: There should be an error thrown when a required field isn't set.")
	}

	t.Log(err)
}

func TestBuildHeaders(t *testing.T) {
	testStruct := struct {
		Accept string `h:"Accept"`
		Num    int    `h:"Number"`
		Style  bool   `h:"Style"`
	}{
		Accept: "application/json",
		Num:    4,
		Style:  true,
	}
	expected := map[string]string{"Accept": "application/json", "Number": "4", "Style": "true"}
	actual, err := BuildHeaders(&testStruct)
	th.CheckNoErr(t, err)
	th.CheckDeepEquals(t, expected, actual)
}

func TestIsZero(t *testing.T) {
	var testMap map[string]interface{}
	testMapValue := reflect.ValueOf(testMap)
	expected := true
	actual := isZero(testMapValue)
	th.CheckEquals(t, expected, actual)
	testMap = map[string]interface{}{"empty": false}
	testMapValue = reflect.ValueOf(testMap)
	expected = false
	actual = isZero(testMapValue)
	th.CheckEquals(t, expected, actual)

	var testArray [2]string
	testArrayValue := reflect.ValueOf(testArray)
	expected = true
	actual = isZero(testArrayValue)
	th.CheckEquals(t, expected, actual)
	testArray = [2]string{"one", "two"}
	testArrayValue = reflect.ValueOf(testArray)
	expected = false
	actual = isZero(testArrayValue)
	th.CheckEquals(t, expected, actual)

	var testStruct struct {
		A string
		B time.Time
	}
	testStructValue := reflect.ValueOf(testStruct)
	expected = true
	actual = isZero(testStructValue)
	th.CheckEquals(t, expected, actual)
	testStruct = struct {
		A string
		B time.Time
	}{
		B: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}
	testStructValue = reflect.ValueOf(testStruct)
	expected = false
	actual = isZero(testStructValue)
	th.CheckEquals(t, expected, actual)

}
