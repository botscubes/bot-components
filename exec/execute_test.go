package exec

import (
	"encoding/json"
	"testing"

	"github.com/botscubes/bot-components/components"
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

var botComponents map[int64]string = map[int64]string{
	0: `{
		"type": "start",
		"path": "",
		"outputs": {
			"nextComponentId": 1
		},
		"data": {}
	}`,
	1: `{
		"type": "format",
		"path": "default",
		"outputs": {
			"nextComponentId": 2
		},
		"data": {
			"formatString": "@ ${default} ^"
		}
	}`,
	2: `{
		"type": "condition",
		"path": "condition",
		"data": {
			"expression": "condition"
		},
		"outputs": {
			"nextComponentId": 3,
			"idIfFalse": 4
		}
	}`,
	3: `{
		"type": "format",
		"path": "default",
		"outputs": {
			"nextComponentId": 4
		},
		"data": {
			"formatString": "( ${default} )"
		}
	}`,
	4: `{
		"type": "format",
		"path": "default",
		"outputs": {
			"nextComponentId": 5
		},
		"data": {
			"formatString": "{ ${default} }"
		}
	}`,
	5: `{
		"type": "code",
		"path": "result",
		"outputs": {
			"nextComponentId": 6
		},
		"data": {
			"code": "1 + 1"
		}
	}`,
	6: `{
		"type": "http",
		"path": "response",
		"outputs": {
			"nextComponentId": null
		},
		"data": {
			"method": "GET",
			"url": "https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=2",
			"header": ""
		}
	}`,
}
var contextData = `
{
	"default": "text",
	"condition": true
}
`

type testIO struct{}

func (*testIO) ReadText() *string {
	s := "test"
	return &s
}
func (*testIO) PrintText(text string) {

}
func (*testIO) PrintButtons(text string, buttons []*io.ButtonData) {

}
func TestExecute(t *testing.T) {

	var id int64 = 0
	var currentComponentId *int64 = &id
	ctx, err := context.NewContextFromJSON([]byte(contextData))
	if err != nil {
		t.Fatal(err)
	}
	var e = NewExecutor(ctx, &testIO{})
	for currentComponentId != nil {
		jsonData := []byte(botComponents[*currentComponentId])
		var d components.ComponentTypeData
		err := json.Unmarshal(jsonData, &d)
		if err != nil {
			t.Fatal(*currentComponentId, err)
		}
		tmp, err := components.NewComponentFromJSON(d.Type, jsonData)
		if err != nil {
			t.Fatal(*currentComponentId, err)
		}
		nextId, err := e.Execute(tmp)
		if err != nil {
			t.Fatalf("component id: %d, error: %s", *currentComponentId, err)
		}
		currentComponentId = nextId
	}
	checkString(ctx, t, "default", "{ ( @ text ^ ) }")
	checkString(ctx, t, "result", "2")
	checkString(ctx, t, "response.statusCode", "200")
}

func checkString(ctx *context.Context, t *testing.T, path string, s string) {
	v, err := ctx.GetValue(path)
	if err != nil {
		t.Fatal(err)
	}
	str, err := v.ToString()
	if err != nil {
		t.Fatal(err)
	}
	if str != s {
		t.Fatalf("Stings don't match: string: %s, expection: %s", str, s)
	}
}
