package context

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
		},
		{

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
	},
	"property": "users",
	"property2": "phoneNumber"
}`

func TestGetNameFromContext(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("users[0].name")
	if err != nil {
		t.Fatal(err)
	}
	name, err := value.ToString()
	if err != nil {
		t.Fatal(err)
	}
	if name != "Alex" {
		t.Fatalf("The name is wrong")
	}
}

func TestGetPhoneNumberFromContext(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("users[0].phoneNumbers[1]")
	if err != nil {
		t.Fatal(err)
	}
	phone, err := value.ToString()
	if err != nil {
		t.Fatal(err)
	}
	if phone != "22222222" {
		t.Fatalf("The phone is wrong. Value: %s", phone)
	}
}

func TestGetAgeFromContext(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("users[1].age")
	if err != nil {
		t.Fatal(err)
	}
	age, err := value.ToInt()
	if err != nil {
		t.Fatal(err)
	}
	if int(age) != 24 {
		t.Fatalf("The phone is wrong")
	}
}

func TestValueFromContext(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("value")
	if err != nil {
		t.Fatal(err)
	}
	v, err := value.ToInt()
	if err != nil {
		t.Fatal(err)
	}
	if int(v) != 10 {
		t.Fatalf("Value is wrong")
	}
}

func TestArrayValueFromContext(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("array[2]")
	if err != nil {
		t.Fatal(err)
	}
	v, err := value.ToString()
	if err != nil {
		t.Fatal(err)
	}
	if v != "c" {
		t.Fatalf("Value is wrong")
	}
}

func TestImplicitAccessToArrayElement(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("array[arrayIndex]")
	if err != nil {
		t.Fatal(err)
	}
	v, err := value.ToString()
	if err != nil {
		t.Fatal(err)
	}
	if v != "b" {
		t.Fatalf("Value is wrong")
	}
}
func TestImplicitlyGetPhoneNumberFromContext(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("users[indices.users[1]].phoneNumbers[indices.phoneNumber]")
	if err != nil {
		t.Fatal(err)
	}
	phone, err := value.ToString()
	if err != nil {
		t.Fatal(err)
	}
	if phone != "11111111" {
		t.Fatalf("The phone is wrong. Value: %s", phone)
	}
}
func TestGetUserAgeFromContextUsingImplicitPropery(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("[property][0].age")
	if err != nil {
		t.Fatal(err)
	}
	age, err := value.ToInt()
	if err != nil {
		t.Fatal(err)
	}
	if age != 20 {
		t.Fatalf("The phone is wrong. Value: %d", age)
	}
}
func TestGetPhoneNumberFromContextUsingImplicitPropery(t *testing.T) {
	var context, err = NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	value, err := context.GetValue("[property][indices.[property][1]].phoneNumbers[indices.[property2]]")
	if err != nil {
		t.Fatal(err)
	}
	phone, err := value.ToString()
	if err != nil {
		t.Fatal(err)
	}
	if phone != "11111111" {
		t.Fatalf("The phone is wrong. Value: %s", phone)
	}
}
