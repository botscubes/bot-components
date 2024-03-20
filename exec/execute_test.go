package exec

import (
	"encoding/json"
	"testing"

	"github.com/botscubes/bot-components/components"
	"github.com/botscubes/bot-components/context"
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
			"nextComponentId": null
		},
		"data": {
			"formatString": "{ ${default} }"
		}
	}`,
}
var contextData = `
{
	"default": "text",
	"condition": true
}
`

type textIO struct{}

func (*textIO) InputText() *string {
	s := "test"
	return &s
}
func (*textIO) OutputText(text string) {

}

func TestExecute(t *testing.T) {

	var id int64 = 0
	var currentComponentId *int64 = &id
	ctx, err := context.NewContextFromJSON([]byte(contextData))
	if err != nil {
		t.Fatal(err)
	}
	var e = NewExecutor(ctx, &textIO{})
	for currentComponentId != nil {
		t.Logf("%d", *currentComponentId)
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
			t.Fatal(*currentComponentId, err)
		}
		currentComponentId = nextId
	}
	v, err := ctx.GetValue("default")
	if err != nil {
		t.Fatal(err)
	}
	s, err := v.ToString()
	if err != nil {
		t.Fatal(err)
	}
	es := "{ ( @ text ^ ) }"
	if s != es {
		t.Fatalf("Stings don't match: string: %s, expection: %s", s, es)
	}

}
