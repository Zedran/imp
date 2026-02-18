package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// Args struct binds, reads and validates all CLI arguments.
type Args struct {
	// Indicates whether the application should exit right after
	// calling cli.Parse. Set to true if -h or -G options are used.
	ExitEarly bool `json:"exit_early"`

	// Parameters required for the main functionality.
	Params Params `json:"params"`
}

// Parse binds CLI options to the corresponding variables, parses user input
// and validates it. If args is nil, os.Args is used.
func Parse(args []string) (Args, error) {
	if args == nil {
		args = os.Args
	}

	var a Args

	fs := flag.NewFlagSet("imp", flag.ContinueOnError)

	// Suppress FlagSet.Parse messages.
	fs.SetOutput(io.Discard)

	a.bindParams(fs)

	var (
		help      = fs.Bool("h", false, "displays help message")
		genPreset = fs.Bool("G", false, "generate an empty preset file in user's home directory and exit")
		preset    = fs.String("P", "", "preset name to be used instead of -e, -h, -l and -p")
	)
	fs.BoolVar(help, "help", false, "displays help message")

	if err := fs.Parse(args[1:]); err != nil {
		return a, fmt.Errorf("err: %w", err)
	}

	if *help {
		a.ExitEarly = true
		a.usage(fs)
		return a, nil
	}

	if *genPreset {
		a.ExitEarly = true
		return a, generatePresetsFile()
	}

	if len(*preset) > 0 {
		err := a.loadPreset(*preset)
		if err != nil {
			return a, err
		}
	}

	return a, a.validate()
}

// bindParams binds CLI args to members of Args.Params.
func (a *Args) bindParams(fs *flag.FlagSet) {
	fs.StringVar(&a.Params.Input, "i", "", "input CSV file")
	fs.StringVar(&a.Params.Output, "o", "", "output CSV file")
	fs.StringVar(&a.Params.Encoding, "e", "utf-8", "input file encoding")
	fs.StringVar(&a.Params.Pattern, "p", "", "pattern that determines how to rewrite the input file")
	fs.StringVar(&a.Params.InputComma, "c", "", "comma character in the input file - specify if input and output use different characters")
	fs.BoolVar(&a.Params.SkipHeader, "0", false, "omit the first row (header) from the input when rewriting")
	fs.StringVar(&a.Params.NewHeader, "H", "", "add this string as the first row")
	fs.BoolVar(&a.Params.Overwrite, "f", false, "overwrite output file if it exists")
	fs.BoolVar(&a.Params.UseCRLF, "l", false, "use CRLF instead of LF for line endings in the output file")
}

// loadPreset reads preset of the specified name and overwrites corresponding values
// in Args.Params.
func (a *Args) loadPreset(name string) error {
	preset, err := loadPreset(name)
	if err != nil {
		return err
	}

	a.Params.ApplyPreset(preset)

	return nil
}

// usage prints help message to os.Stderr.
func (a *Args) usage(fs *flag.FlagSet) {
	fs.SetOutput(os.Stderr)

	fmt.Fprintf(os.Stderr, "Usage of %s:\n", fs.Name())
	fs.PrintDefaults()
	fmt.Fprintln(
		os.Stderr,
		"\n",
		"Pattern starts with the character serving as the CSV comma.\n",
		"The second character will be interpreted as a tag prefix. Both\n",
		"the comma and tag prefix can be freely chosen, but they need\n",
		"to be unique - they cannot be used anywhere else in the pattern.\n\n",
		"CSV content is specified with a series of tags:\n",
		"  - '/d<number>' causes imp to insert the column <number> from\n",
		"    the input file. Comma character is allowed at the end.\n",
		"  - '/s<text> causes imp to insert an arbitrary <text>.\n\n",
		"Example:\n",
		"  - input file header:  'First name,Last name,Amount'\n",
		"  - output file header: 'Full name,Amount'\n",
		"  - pattern:            ',/d0/s /d1,/d2'",
	)
}

// validate enforces invariants of the Args struct.
func (a *Args) validate() error {
	if len(a.Params.Input) == 0 {
		return errors.New("err: input file not specified")
	}

	if len(a.Params.Output) == 0 {
		return errors.New("err: output file not specified")
	}

	if len(a.Params.Encoding) == 0 {
		return errors.New("err: input encoding not specified")
	}

	if len(a.Params.InputComma) > 1 {
		return errors.New("err: value of -c is not a single character")
	}

	return nil
}
