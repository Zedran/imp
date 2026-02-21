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

package tests

import (
	"crypto/sha256"
	"encoding/json"
	"os"
	"path/filepath"
)

const TEST_DATA_DIR string = "testdata"

// ReadData unmarshals JSON file containing test data into an arbitrary
// data structure.
func ReadData(fname string, v any) error {
	stream, err := os.ReadFile(filepath.Join(TEST_DATA_DIR, fname))
	if err != nil {
		return err
	}
	return json.Unmarshal(stream, v)
}

// SHA256 returns SHA256 sum of data.
func SHA256(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}
