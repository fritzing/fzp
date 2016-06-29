package ext

import (
	"testing"
)

var testData = []struct {
	filename string
	valid    bool
}{
	{"test.fzp", true},
	{"test.FZP", false},
	{"test.json", true},
	{"test.JSON", false},
	{"test.yml", true},
	{"test.YML", false},
	{"test.yaml", false},
	{"test.YAML", false},
	{"test.xml", false},
	{"test.XML", false},
	{"test.unknown", false},

	{"path/to/the/fritzing/part.fzp", true},
	{"path/to/the/fritzing/part/but/not/fzp.file", false},
}

func Test_IsValid(t *testing.T) {
	for _, tt := range testData {
		v, _ := IsValid(tt.filename)
		if v != tt.valid {
			t.Errorf("%s is not equal\n", tt.filename)
		}
	}
}
