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

	// Overwrite output file if it exists.
	Overwrite bool `json:"overwrite"`

	// Collection of parameters describing the structure of input
	// and the desired output CSV data.
	Format Preset `json:"format"`
}

// ApplyPreset sets preset as the new value for Params.Format field.
func (p *Params) ApplyPreset(preset Preset) {
	p.Format = preset
}
