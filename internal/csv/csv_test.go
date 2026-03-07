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

package csv

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Zedran/imp/internal/pattern"
	"github.com/Zedran/imp/internal/tests"
	"github.com/Zedran/imp/internal/utils"
)

func TestRewriteRows(t *testing.T) {
	type testData struct {
		Input    string       `json:"input"`
		Expected string       `json:"expected"`
		Preset   utils.Preset `json:"preset"`
		Err      string       `json:"err"`
	}

	cases := make([]testData, 0)

	err := tests.ReadData("TestRewriteRows.json", &cases)
	if err != nil {
		t.Fatalf("failed to load test data: '%v'", err)
	}

	for _, c := range cases {
		var params utils.Params
		params.ApplyPreset(c.Preset)

		spec, err := pattern.ParsePattern(params.Pattern, c.Preset.CurrSep)
		if err != nil {
			t.Fatalf("failed to parse pattern: '%v'", err)
		}

		var (
			ir  = strings.NewReader(c.Input)
			out bytes.Buffer
		)

		err = rewriteRows(ir, &out, spec, params)

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
