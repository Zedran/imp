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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// presetsFile is the default name of the presets file.
const presetsFile string = ".imp-presets.json"

// Preset stores a combination of pattern and encoding. Presets can be written
// to the presets file to facilitate reuse.
type Preset struct {
	// Input file encoding.
	Encoding string `json:"encoding"`

	// Pattern for output CSV file.
	Pattern string `json:"pattern"`

	// Character serving as comma in the input file.
	InputComma string `json:"input_comma"`

	// Indicates whether the header should be skipped.
	SkipHeader bool `json:"skip_header"`

	// This string will be inserted as the first row.
	NewHeader string `json:"new_header"`

	// Use CRLF instead of LF in the output file
	UseCRLF bool `json:"crlf"`

	// Decimal separator for currency.
	CurrSep string `json:"curr_sep"`
}

// GeneratePresetsFile writes an empty presets file to current user's home
// directory. If the file is already present, it returns an error.
func GeneratePresetsFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("err: failed to locate user's home directory: %w", err)
	}

	path := filepath.Join(home, presetsFile)

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("err: file '%s' already exists", path)
	} else if !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("err: unexpected error returned from os.Stat: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	empty := map[string]Preset{"default": {}}

	stream, err := json.MarshalIndent(empty, "", "    ")
	if err != nil {
		return err
	}
	stream = append(stream, '\n')

	if _, err := f.Write(stream); err != nil {
		return err
	}

	fmt.Printf("presets file created at %s\n", path)
	return nil
}

// LoadPreset returns a Preset of the specified name from the presets file.
// If the Preset is not found, an error message is returned.
func LoadPreset(name string) (Preset, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Preset{}, fmt.Errorf("err: failed to locate user's home directory: %w", err)
	}

	f, err := os.Open(filepath.Join(home, presetsFile))
	if err != nil {
		return Preset{}, fmt.Errorf("err: failed to open the presets file: %w", err)
	}
	defer f.Close()

	presets, err := readPresets(f)
	if err != nil {
		return Preset{}, fmt.Errorf("err: failed to parse the presets file: %w", err)
	}

	p, found := presets[name]
	if !found {
		return Preset{}, fmt.Errorf("err: preset '%s' not found", name)
	}

	return p, nil
}

// readPresets unmarshalls Presets into a map.
func readPresets(f *os.File) (map[string]Preset, error) {
	stream, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var p map[string]Preset

	err = json.Unmarshal(stream, &p)
	return p, err
}
