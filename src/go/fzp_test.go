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

var TestTable_Read = []struct {
	src        string
	readError  bool
	checkError bool
}{
	{"../fixture/empty.fzp", true, true},
	{"../fixture/empty_xml.fzp", true, true},
	{"../fixture/empty_fzp_0.fzp", false, true},
	// {"../fixture/empty_fzp_1.fzp", false, true},
	// {"../fixture/empty_fzp_2.fzp", false, true},

	// passed
	// {"../fixture/pass/pass_0.fzp", false, false},

	// failed
	// {"../fixture/fail/fail_0_all.fzp", false, true},
	// {"../fixture/fail/fail_1_title.fzp", false, true},
}

func Test_ReadFzp_Ok(t *testing.T) {
	totalTests := len(TestTable_Read)
	for k, v := range TestTable_Read {
		t.Logf("Run Test %v of %v - %q\n", k, totalTests, v.src)
		f, _, errFile := ReadFzp(v.src)

		if v.readError {
			if errFile == nil {
				t.Error("Fzp.ReadFzp missing error", errFile)
			}
		}

		errCheck := f.Check()
		if v.checkError {
			if errCheck == nil {
				t.Error("Fzp.Check missing error:", errCheck)
			}
		}
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
