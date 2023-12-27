package iodata

import (
	"reflect"
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
	]

}`

func TestGetNameFromIOData(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.getValue("users[0].name")
	if err != nil {
		t.Error(err)
	}
	name, ok := value.(string)
	if !ok {
		t.Errorf("Type cast error. Variable type: %s", reflect.TypeOf(value))
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
	value, err := iodata.getValue("users[0].phoneNumbers[1]")
	if err != nil {
		t.Error(err)
	}
	phone, ok := value.(string)
	if !ok {
		t.Errorf("Type cast error. Variable type: %s", reflect.TypeOf(value))
	}
	if phone != "22222222" {
		t.Errorf("The phone is wrong")
	}
}

func TestGetAgeFromIOData(t *testing.T) {
	var iodata, err = NewIODataFromJSON([]byte(testJson))
	if err != nil {
		t.Error(err)
	}
	value, err := iodata.getValue("users[1].age")
	if err != nil {
		t.Error(err)
	}
	age, ok := value.(float64)
	if !ok {
		t.Errorf("Type cast error. Variable type: %s", reflect.TypeOf(value))
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
	value, err := iodata.getValue("value")
	if err != nil {
		t.Error(err)
	}
	v, ok := value.(float64)
	if !ok {
		t.Errorf("Type cast error. Variable type: %s", reflect.TypeOf(value))
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
	value, err := iodata.getValue("array[2]")
	if err != nil {
		t.Error(err)
	}
	v, ok := value.(string)
	if !ok {
		t.Errorf("Type cast error. Variable type: %s", reflect.TypeOf(value))
	}
	if v != "c" {
		t.Errorf("Value is wrong")
	}
}
