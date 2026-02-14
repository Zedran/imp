package pattern

import (
	"slices"
	"strings"
	"testing"

	"github.com/Zedran/imp/internal/tests"
)

func TestParsePattern(t *testing.T) {
	type testData struct {
		Input    string `json:"input"`
		Expected Spec   `json:"expected"`
		Err      string `json:"err"`
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
			if len(c.Err) > 0 {
				t.Fatalf("no error returned for %s, expected: '%v'", c.Input, c.Err)
			}
			if !slices.Equal(out.Tokens, c.Expected.Tokens) || out.Comma != c.Expected.Comma {
				t.Fatalf("incorrect output for '%s': %v != %v", c.Input, out, c.Expected)
			}
		}

	}
}
