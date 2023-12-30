package iodata

import (
	"testing"
)

var testJson = `
{
	"value": 10,
	"users": [
		{
			"name": "Alex",
			"age": 20,
			"phoneNumbers": [
				"11111111",
				"22222222"
			]
		},
		{
			"name": "Andrey",
			"age": 24
		}
	],
	"array": [
		"a",
		"b",
		"c"
	],
	"arrayIndex": 1,
	"indices": {
		"users": [1, 0],
		"phoneNumber": 0
	}
}`

func TestGetNameFromIOData(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.GetValue("users[0].name")
	if err != nil {
		t.Error(err)
	}
	name, err := value.ToString()
	if err != nil {
		t.Error(err)
	}
	if name != "Alex" {
		t.Errorf("The name is wrong")
	}
}

func TestGetPhoneNumberFromIOData(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.GetValue("users[0].phoneNumbers[1]")
	if err != nil {
		t.Error(err)
	}
	phone, err := value.ToString()
	if err != nil {
		t.Error(err)
	}
	if phone != "22222222" {
		t.Errorf("The phone is wrong. Value: %s", phone)
	}
}

func TestGetAgeFromIOData(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.GetValue("users[1].age")
	if err != nil {
		t.Error(err)
	}
	age, err := value.ToInt()
	if err != nil {
		t.Error(err)
	}
	if int(age) != 24 {
		t.Errorf("The phone is wrong")
	}
}

func TestValueFromIOData(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.GetValue("value")
	if err != nil {
		t.Error(err)
	}
	v, err := value.ToInt()
	if err != nil {
		t.Error(err)
	}
	if int(v) != 10 {
		t.Errorf("Value is wrong")
	}
}

func TestArrayValueFromIOData(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.GetValue("array[2]")
	if err != nil {
		t.Error(err)
	}
	v, err := value.ToString()
	if err != nil {
		t.Error(err)
	}
	if v != "c" {
		t.Errorf("Value is wrong")
	}
}

func TestImplicitAccessToArrayElement(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.GetValue("array[arrayIndex]")
	if err != nil {
		t.Error(err)
	}
	v, err := value.ToString()
	if err != nil {
		t.Error(err)
	}
	if v != "b" {
		t.Errorf("Value is wrong")
	}
}
func TestImplicitlyGetPhoneNumberFromIOData(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.GetValue("users[indices.users[1]].phoneNumbers[indices.phoneNumber]")
	if err != nil {
		t.Error(err)
	}
	phone, err := value.ToString()
	if err != nil {
		t.Error(err)
	}
	if phone != "11111111" {
		t.Errorf("The phone is wrong. Value: %s", phone)
	}
}
