package pattern

import (
	"slices"
	"strings"
	"testing"

	"github.com/Zedran/imp/internal/tests"
)

func TestParsePattern(t *testing.T) {
	type testData struct {
		Input    string  `json:"input"`
		Expected []Token `json:"expected"`
		Err      string  `json:"err"`
	}

	cases := make([]testData, 0)

	err := tests.ReadData("TestParsePattern.json", &cases)
	if err != nil {
		t.Fatalf("failed to load test data: '%v'", err)
	}

	for _, c := range cases {
		out, err := ParsePattern(c.Input)

		if err != nil {
			if !strings.HasPrefix(err.Error(), c.Err) {
				t.Fatalf("incorrect error message for '%s': '%v' != '%v*'", c.Input, err, c.Err)
			}
		} else {
			if !slices.Equal(out, c.Expected) {
				t.Fatalf("incorrect output for '%s': %v != %v", c.Input, out, c.Expected)
			}
		}

	}
}
