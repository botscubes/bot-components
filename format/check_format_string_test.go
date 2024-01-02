package format

import "testing"

func TestCheckFormatString(t *testing.T) {
	for _, pair := range strsToCheck {

		err := CheckFormatString(pair[0])
		if err != nil {
			t.Fatal(err)
		}
	}
}
