package cli

// Params is a collection of parameters required for CSV file rewriting.
type Params struct {
	// Path to the input CSV file.
	Input string

	// Path to the output CSV file.
	Output string

	// Encoding of the input file.
	Encoding string

	// Pattern that determines how to rewrite the input file.
	Pattern string

	// Character that serves as CSV comma in the input file.
	// If empty, it is assumed that it matches comma in the output.
	InputComma string

	// String inserted as the first row.
	NewHeader string

	// Omit the first row from the input.
	SkipHeader bool

	// Overwrite output file if it exists.
	Overwrite bool

	// Use CRLF instead of LF for line endings.
	UseCRLF bool
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
