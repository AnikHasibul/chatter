package util

import "testing"

func TestExtractSession(t *testing.T) {
	type model struct {
		in, trim, out string
	}
	var data = []model{
		{"/c/hello", "/c/", "hello"},
		{"/c/helloW", "/c/", "helloW"},
		{"/chat/hello", "/chat/", "hello"},
		{"/hello", "/", "hello"},
	}
	for _, v := range data {
		if out := ExtractSession(v.in, v.trim); out != v.out {
			t.Errorf("Expected %s, Got %s", v.out, out)
			t.Fail()
		}
	}
}
