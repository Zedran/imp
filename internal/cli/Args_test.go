package cli

import (
	"os"
	"strings"
	"testing"

	"github.com/Zedran/imp/internal/tests"
)

func TestParse(t *testing.T) {
	type testData struct {
		Args     []string `json:"args"`
		Expected Args     `json:"expected"`
		Err      string   `json:"err"`
	}

	cases := make([]testData, 0)

	err := tests.ReadData("TestParse.json", &cases)
	if err != nil {
		t.Fatalf("failed to load test data: '%v'", err)
	}

	for i, c := range cases {
		a, err := Parse(append([]string{os.Args[0]}, c.Args...))

		if err != nil {
			if len(c.Err) == 0 {
				t.Fatalf("unexpected error message for case %d: '%v'", i, err)
			}
			if !strings.HasPrefix(err.Error(), c.Err) {
				t.Fatalf("incorrect error message for case %d: '%v' != '%v*'", i, err, c.Err)
			}
		} else {
			if len(c.Err) > 0 {
				t.Log(a.Params.Encoding)
				t.Fatalf("no error returned for case %d, expected: '%v'", i, c.Err)
			}
			if a != c.Expected {
				t.Fatalf("incorrect output for case %d:\n'%#v' !=\n'%#v'", i, a, c.Expected)
			}
		}
	}
}
