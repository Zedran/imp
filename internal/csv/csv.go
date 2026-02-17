package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/Zedran/imp/internal/encoding"
	"github.com/Zedran/imp/internal/pattern"
)

// RewriteCSV is the top function of the application's internals.
// It accepts 4 arguments:
//   - inputPath     - original CSV file path
//   - outputPath    - new CSV file path
//   - encoding      - the encoding of the original CSV file
//   - patternString - instructions on how to rewrite the CSV
//   - skipHeader    - skip the first row of input CSV file when rewriting
//   - newHeader     - the string to be inserted as the first row
//   - overwrite     - if true, output file is overwritten if it exists
//
// Returns an error if any stage of the rewriting process fails.
func RewriteCSV(inputPath, outputPath, enc, patternString string, skipHeader, overwrite bool, newHeader string) error {
	spec, err := pattern.ParsePattern(patternString)
	if err != nil {
		return err
	}

	input, err := encoding.OpenUTF8(inputPath, enc)
	if err != nil {
		return err
	}
	defer input.Close()

	if !overwrite {
		if _, err := os.Stat(outputPath); err == nil {
			return fmt.Errorf("err: file '%s' already exists", outputPath)
		} else if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("err: unexpected error returned from os.Stat: %w", err)
		}
	}

	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	return rewriteRows(input, output, spec, skipHeader, newHeader)
}

// buildRow returns a single, rewritten row in a string form.
func buildRow(old []string, spec pattern.Spec) ([]string, error) {
	var (
		b      strings.Builder
		newRow = make([]string, 0, len(spec.Tokens))
	)

	for _, token := range spec.Tokens {
		switch token.Type {
		case pattern.TT_COLUMN:
			if len(old) <= token.Column {
				return nil, fmt.Errorf("err: column number out of range: '%d'", token.Column)
			}
			b.WriteString(old[token.Column])
		case pattern.TT_TEXT:
			if token.Text == string(spec.Comma) {
				newRow = append(newRow, b.String())
				b.Reset()
			} else {
				b.WriteString(token.Text)
			}
		}
	}
	if b.Len() > 0 {
		newRow = append(newRow, b.String())
	}

	return newRow, nil
}

// rewriteRows coordinates the rewriting process. It accepts input Reader
// and output writer, as well as Spec struct compiled by pattern.ParsePattern.
func rewriteRows(input io.Reader, output io.Writer, spec pattern.Spec, skipHeader bool, newHeader string) error {
	comma, _ := utf8.DecodeRuneInString(spec.Comma)

	r := csv.NewReader(input)
	r.Comma = comma

	w := csv.NewWriter(output)
	w.Comma = comma
	defer w.Flush()

	if skipHeader {
		_, err := r.Read()
		if err != nil {
			return fmt.Errorf("err: unexpected error on header skip: %w", err)
		}
	}

	if len(newHeader) > 0 {
		w.Write(strings.Split(newHeader, spec.Comma))
	}

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
		w.Write(newRecord)
	}

	return nil
}
