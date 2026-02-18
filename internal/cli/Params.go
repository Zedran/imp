package cli

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
