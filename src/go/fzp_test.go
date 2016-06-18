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

//////////////

// func Test_HasExtFzp(t *testing.T) {
// 	result := HasExtFzp("data.fzp")
// 	if !result {
// 		t.Error("HasExtFzp failed")
// 	}
// 	result = HasExtFzp("data.notafzp")
// 	if result {
// 		t.Error("IsFileFzp failed")
// 	}
// }
//
// func Test_HasExtJSON(t *testing.T) {
// 	result := HasExtJSON("data.json")
// 	if !result {
// 		t.Error("HasExtJSON failed")
// 	}
// 	result = HasExtJSON("data.notajson")
// 	if result {
// 		t.Error("IsFileJSON failed")
// 	}
// }
//
// func Test_HasExtYAML(t *testing.T) {
// 	result := HasExtYAML("data.yml")
// 	if !result {
// 		t.Error("HasExtYAML failed")
// 	}
// 	result = HasExtYAML("data.notayml")
// 	if result {
// 		t.Error("IsFileYAML failed")
// 	}
// }

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
