package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Zedran/imp/internal/encoding"
	"github.com/Zedran/imp/internal/pattern"
)

// RewriteCSV is the top function of the application's internals.
// It accepts 4 arguments:
//   - inputPath     - original CSV file path
//   - outputPath    - new CSV file path
//   - encoding      - the encoding of the original CSV file
//   - patternString - instructions on how to rewrite the CSV
//
// Returns an error if any stage of the rewriting process fails.
func RewriteCSV(inputPath, outputPath, enc, patternString string) error {
	spec, err := pattern.ParsePattern(patternString)
	if err != nil {
		return err
	}

	input, err := encoding.OpenUTF8(inputPath, enc)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	return rewriteRows(input, output, spec)
}

// buildRow returns a single, rewritten row in a string form.
func buildRow(old []string, spec pattern.Spec) (string, error) {
	var b strings.Builder

	for _, token := range spec.Tokens {
		switch token.Type {
		case pattern.TT_COLUMN:
			if len(old) <= token.Column {
				return "", fmt.Errorf("err: column number out of range: '%d'", token.Column)
			}
			b.WriteString(old[token.Column])
		case pattern.TT_TEXT:
			b.WriteString(token.Text)
		}
	}
	b.WriteByte('\n')

	return b.String(), nil
}

// rewriteRows coordinates the rewriting process. It accepts input Reader
// and output writer, as well as Spec struct compiled by pattern.ParsePattern.
func rewriteRows(input io.Reader, output io.StringWriter, spec pattern.Spec) error {
	r := csv.NewReader(input)
	r.Comma = spec.Comma

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		newRecord, err := buildRow(record, spec)
		if err != nil {
			return err
		}
		output.WriteString(newRecord)
	}

	return nil
}
