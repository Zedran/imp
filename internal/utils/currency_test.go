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

package utils

import (
	"testing"

	"github.com/Zedran/imp/internal/tests"
)

func TestFormatCurrency(t *testing.T) {
	type testData struct {
		Input    string `json:"input"`
		Decimal  string `json:"decimal"`
		Expected string `json:"expected"`
	}

	cases := make([]testData, 0)

	err := tests.ReadData("TestFormatCurrency.json", &cases)
	if err != nil {
		t.Fatalf("failed to load test data: '%v'", err)
	}

	for i, c := range cases {
		out := FormatCurrency(c.Input, c.Decimal)

		if out != c.Expected {
			t.Fatalf("failed for case %d ['%s' '%s']: '%s' != '%s'", i, c.Input, c.Decimal, out, c.Expected)
		}
	}
}
