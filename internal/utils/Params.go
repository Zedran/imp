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

// Params is a collection of parameters required for CSV file rewriting.
type Params struct {
	// Path to the input CSV file.
	Input string `json:"input"`

	// Path to the output CSV file.
	Output string `json:"output"`

	// Encoding of the input file.
	Encoding string `json:"encoding"`

	// Pattern that determines how to rewrite the input file.
	Pattern string `json:"pattern"`

	// Character that serves as CSV comma in the input file.
	// If empty, it is assumed that it matches comma in the output.
	InputComma string `json:"input_comma"`

	// String inserted as the first row.
	NewHeader string `json:"new_header"`

	// Omit the first row from the input.
	SkipHeader bool `json:"skip_header"`

	// Overwrite output file if it exists.
	Overwrite bool `json:"overwrite"`

	// Use CRLF instead of LF for line endings.
	UseCRLF bool `json:"crlf"`
}

// ApplyPreset copies values from Preset struct to the corresponding members.
func (p *Params) ApplyPreset(preset Preset) {
	p.Encoding = preset.Encoding
	p.Pattern = preset.Pattern
	p.InputComma = preset.InputComma
	p.SkipHeader = preset.SkipHeader
	p.NewHeader = preset.NewHeader
	p.UseCRLF = preset.UseCRLF
}
