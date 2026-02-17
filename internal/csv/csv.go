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

	"github.com/Zedran/imp/internal/cli"
	"github.com/Zedran/imp/internal/encoding"
	"github.com/Zedran/imp/internal/pattern"
)

// RewriteCSV is the top function of the application's internals.
// Returns an error if any stage of the rewriting process fails.
func RewriteCSV(params cli.Params) error {
	spec, err := pattern.ParsePattern(params.Pattern)
	if err != nil {
		return err
	}

	input, err := encoding.OpenUTF8(params.Input, params.Encoding)
	if err != nil {
		return err
	}
	defer input.Close()

	if !params.Overwrite {
		if _, err := os.Stat(params.Output); err == nil {
			return fmt.Errorf("err: file '%s' already exists", params.Output)
		} else if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("err: unexpected error returned from os.Stat: %w", err)
		}
	}

	output, err := os.Create(params.Output)
	if err != nil {
		return err
	}
	defer output.Close()

	return rewriteRows(input, output, spec, params)
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
func rewriteRows(input io.Reader, output io.Writer, spec pattern.Spec, params cli.Params) error {
	comma, _ := utf8.DecodeRuneInString(spec.Comma)

	r := csv.NewReader(input)
	if len(params.InputComma) > 0 {
		r.Comma, _ = utf8.DecodeRuneInString(params.InputComma)
	} else {
		r.Comma = comma
	}

	w := csv.NewWriter(output)
	w.Comma = comma
	w.UseCRLF = params.UseCRLF
	defer w.Flush()

	if params.SkipHeader {
		_, err := r.Read()
		if err != nil {
			return fmt.Errorf("err: unexpected error on header skip: %w", err)
		}
	}

	if len(params.NewHeader) > 0 {
		w.Write(strings.Split(params.NewHeader, spec.Comma))
	}

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("err: unexpected error on read: %w", err)
		}

		newRecord, err := buildRow(record, spec)
		if err != nil {
			return err
		}
		w.Write(newRecord)
	}

	return nil
}
