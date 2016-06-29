package fzp

import (
	"testing"
)

func Test_GetFormat(t *testing.T) {
	format, _ := GetFormat("FZP")
	if format != FormatFzp {
		t.Error("not equal")
	}

	format, _ = GetFormat("data.fzp")
	if format != FormatFzp {
		t.Error("not equal")
	}

	format, _ = GetFormat("data.json")
	if format != FormatJSON {
		t.Error("not equal")
	}

	format, _ = GetFormat("data.yaml")
	if format != FormatYAML {
		t.Error("not equal")
	}
}

func Test_ReadFzp_Ok(t *testing.T) {
	f, _, err := ReadFzp("../../docs/template.fzp")
	if err != nil {
		t.Error("Fzp.ReadFzp broken")
	}

	errCheck := f.Check()
	if errCheck != nil {
		t.Error("Fzp.Check broken:", errCheck)
	}
}

func Test_ReadFzp_Failed(t *testing.T) {
	_, _, err := ReadFzp("../not.found")
	if err == nil {
		t.Error("Fzp.ReadFzp (that doesn't exists) broken")
	}
}

func Test_ReadFzp_CheckTags(t *testing.T) {
	// fake data
	fzpData := Fzp{}
	fzpData.Tags = append(fzpData.Tags, "")
	// was an error returned?
	err, _ := fzpData.CheckTags()
	if err == nil {
		t.Error("Fzp.CheckTags broken")
	}
}
