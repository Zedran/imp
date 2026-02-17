package csv

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Zedran/imp/internal/pattern"
	"github.com/Zedran/imp/internal/tests"
)

func TestRewriteRows(t *testing.T) {
	type testData struct {
		Input      string `json:"input"`
		Expected   string `json:"expected"`
		Pattern    string `json:"pattern"`
		SkipHeader bool   `json:"skip_header"`
		Err        string `json:"err"`
	}

	cases := make([]testData, 0)

	err := tests.ReadData("TestRewriteRows.json", &cases)
	if err != nil {
		t.Fatalf("failed to load test data: '%v'", err)
	}

	for _, c := range cases {
		spec, err := pattern.ParsePattern(c.Pattern)
		if err != nil {
			t.Fatalf("failed to parse pattern: '%v'", err)
		}

		var (
			ir  = strings.NewReader(c.Input)
			out bytes.Buffer
		)

		err = rewriteRows(ir, &out, spec, c.SkipHeader)

		if err != nil {
			if len(c.Err) == 0 {
				t.Fatalf("unexpected error message for '%s': '%v'", c.Input, err)
			}
			if !strings.HasPrefix(err.Error(), c.Err) {
				t.Fatalf("incorrect error message for '%s': '%v' != '%v*'", c.Input, err, c.Err)
			}
		} else {
			if len(c.Err) > 0 {
				t.Fatalf("no error returned for %s, expected: '%v'", c.Input, c.Err)
			}
			if os := out.String(); os != c.Expected {
				t.Fatalf("incorrect output: '%s' != '%s'", os, c.Expected)
			}
		}
	}
}
