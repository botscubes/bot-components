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

var strsToCheck [][2]string = [][2]string{
	{
		"\\\\тест } test\\\\ } {{ \\$ \\$ \\$\\$",
		"\\тест } test\\ } {{ $ $ $$",
	},
	{
		"{ post_id = ${ posts[currentPostIndex].id }, post_title = ${   posts[currentPostIndex].title  }, post_description = ${posts[currentPostIndex].description}, posted = ${posts[currentPostIndex].posted} }",
		"{ post_id = 1, post_title = Post1, post_description = Post bla bla bla, posted = false }",
	},
}

func TestFormat(t *testing.T) {
	var ctx, err = context.NewContextFromJSON([]byte(testJson))
	if err != nil {
		t.Fatal(err)
	}
	for _, pair := range strsToCheck {

		s, err := Format(pair[0], ctx)
		if err != nil {
			t.Fatal(err)
		}
		if s != pair[1] {
			t.Fatalf("Not equal: current string: %s", s)
		}
	}
}
