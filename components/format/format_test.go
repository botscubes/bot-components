package format

import (
	"testing"

	"github.com/botscubes/bot-components/context"
)

var testJson = `
{
	"posts": [
		{
			"id": 1,
			"title": "Post1",
			"description": "Post bla bla bla",
			"posted": false
		},
		{
			"id": 2,
			"title": "Post2",
			"description": "Post2 bla bla bla",
			"posted": true
		}
	],
	"currentPostIndex": 0
}`

func TestCheckBackslash(t *testing.T) {
	var ctx, err = context.NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	s, err := Format("\\\\тест } test\\\\ } {{ \\$ \\$ \\$\\$", ctx)
	if err != nil {
		t.Fatal(err)
	}
	if s != "\\тест } test\\ } {{ $ $ $$" {
		t.Fatalf("Not equal: current string: %s", s)
	}
}

func TestCheckReplacements(t *testing.T) {
	var ctx, err = context.NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	s, err := Format("{ post_id = ${ posts[currentPostIndex].id }, post_title = ${   posts[currentPostIndex].title  }, post_description = ${posts[currentPostIndex].description}, posted = ${posts[currentPostIndex].posted} }", ctx)
	if err != nil {
		t.Fatal(err)
	}
	if s != "{ post_id = 1, post_title = Post1, post_description = Post bla bla bla, posted = false }" {
		t.Fatalf("Not equal: current string: %s", s)
	}

}
