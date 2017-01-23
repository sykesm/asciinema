package util_test

import (
	"testing"

	"github.com/sykesm/asciinema/util"
)

func TestIsUTF8Locale(t *testing.T) {
	var tests = []struct {
		lcAll          string
		lcCtype        string
		lang           string
		expectedResult bool
	}{
		{"pl_PL.UTF-8", "pl_PL.ISO-8859-1", "pl_PL.ISO-8859-2", true},
		{"cz_CS.utf8", "pl_PL.ISO-8859-1", "pl_PL.ISO-8859-2", true},
		{"", "pl_PL.ISO-8859-1", "pl_PL.ISO-8859-2", false},
		{"", "", "pl_PL.ISO-8859-2", false},
		{"", "", "", false},
		{"UTF-8", "pl_PL.ISO-8859-1", "pl_PL.ISO-8859-2", false},
		{"", "ISO-8859-1", "pl_PL.ISO-8859-2", false},
		{"", "", "ISO-8859-2", false},
	}

	for _, test := range tests {
		env := map[string]string{
			"LC_ALL":   test.lcAll,
			"LC_CTYPE": test.lcCtype,
			"LANG":     test.lang,
		}
		if util.IsUTF8Locale(env) != test.expectedResult {
			t.Errorf("expected %v for %v", test.expectedResult, test)
		}
	}
}
