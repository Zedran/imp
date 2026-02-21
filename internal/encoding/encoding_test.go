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

package encoding

import (
	"io"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/Zedran/imp/internal/tests"
)

func TestOpenUTF8(t *testing.T) {
	reference, err := os.ReadFile(filepath.Join(tests.TEST_DATA_DIR, "utf-8.txt"))
	if err != nil {
		t.Fatalf("failed to load test data: '%v'", err)
	}

	refSum := tests.SHA256(reference)

	cases := []string{
		"windows-1250",
		"utf-8",
	}

	for _, c := range cases {
		reader, err := OpenUTF8(filepath.Join(tests.TEST_DATA_DIR, c+".txt"), c)
		if err != nil {
			t.Fatalf("failed to load test data: '%v'", err)
		}
		defer reader.Close()

		stream, err := io.ReadAll(reader)
		if err != nil {
			t.Fatalf("failed to read normalized data: '%v'", err)
		}

		outSum := tests.SHA256(stream)

		if !slices.Equal(refSum, outSum) {
			t.Fatalf("normalized '%s' file is not equal to the UTF-8 reference", c)
		}
	}

}
