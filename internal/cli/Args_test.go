// imp -- tool for rewriting CSV files and normalizing encoding.
// Copyright (C) 2026  Wojciech Głąb (github.com/Zedran)
//
// This file is part of imp.
//
// imp is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 only.
//
// imp is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with imp.  If not, see <https://www.gnu.org/licenses/>.

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
				t.Fatalf("no error returned for case %d, expected: '%v'", i, c.Err)
			}
			if a != c.Expected {
				t.Fatalf("incorrect output for case %d:\n'%#v' !=\n'%#v'", i, a, c.Expected)
			}
		}
	}
}
