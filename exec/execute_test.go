package exec

import (
	"encoding/json"
	"testing"

	"github.com/botscubes/bot-components/components"
	"github.com/botscubes/bot-components/context"
)

var botComponents map[int]string = map[int]string{
	1: `{
		"type": "format",
		"path": "default",
		"nextComponentId": 2,
		"formatString": "@ ${default} ^"
	}`,
	2: `{
		"type": "condition",
		"path": "condition",
		"nextComponentId": 3,
		"idIfFalse": 4
	}`,
	3: `{
		"type": "format",
		"path": "default",
		"nextComponentId": 4,
		"formatString": "( ${default} )"
	}`,
	4: `{
		"type": "format",
		"path": "default",
		"nextComponentId": null,
		"formatString": "{ ${default} }"
	}`,
}
var contextData = `
{
	"default": "text",
	"condition": true
}
`

func TestExecute(t *testing.T) {
	var id = 1
	var currentComponentId *int = &id
	ctx, err := context.NewContextFromJSON([]byte(contextData))
	if err != nil {
		t.Fatal(err)
	}
	for currentComponentId != nil {
		t.Logf("%d", *currentComponentId)
		jsonData := []byte(botComponents[*currentComponentId])
		var d components.ComponentTypeData
		err := json.Unmarshal(jsonData, &d)
		if err != nil {
			t.Fatal(*currentComponentId, err)
		}
		tmp, err := components.NewActionOrControlComponentFromJSON(d.Type, jsonData)
		if err != nil {
			t.Fatal(*currentComponentId, err)
		}
		nextId, err := Execute(ctx, tmp)
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
	e := "{ ( @ text ^ ) }"
	if s != e {
		t.Fatalf("Stings don't match: string: %s, expection: %s", s, e)
	}

}
