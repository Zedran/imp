package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/Zedran/imp/internal/presets"
)

// Args struct binds, reads and validates all CLI arguments.
type Args struct {
	// Indicates whether the application should exit right after
	// calling cli.Parse. Set to true if the -G option is used.
	ExitEarly bool

	// Parameters required for the main functionality.
	Params Params
}

// Parse binds CLI flags to the corresponding variables, calls flag.Parse
// and validates user input.
func Parse() (Args, error) {
	if flag.Parsed() {
		return Args{}, errors.New("arguments already parsed")
	}

	var a Args

	a.BindParams()

	var (
		genPreset = flag.Bool("G", false, "generate an empty preset file in user's home directory and exit")
		preset    = flag.String("P", "", "preset name to be used instead of -e, -h, -l and -p")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
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

	flag.Parse()

	if *genPreset {
		if err := presets.GeneratePresetsFile(); err != nil {
			return a, err
		}
		a.ExitEarly = true
		return a, nil
	}

	if len(*preset) > 0 {
		err := a.LoadPreset(*preset)
		if err != nil {
			return a, err
		}
	}

	return a, a.Validate()
}

// BindParams binds CLI args to members of Args.Params.
func (a *Args) BindParams() {
	flag.StringVar(&a.Params.Input, "i", "", "input CSV file")
	flag.StringVar(&a.Params.Output, "o", "", "output CSV file")
	flag.StringVar(&a.Params.Encoding, "e", "utf-8", "input file encoding")
	flag.StringVar(&a.Params.Pattern, "p", "", "pattern that determines how to rewrite the input file")
	flag.BoolVar(&a.Params.SkipHeader, "0", false, "omit the first row (header) from the input when rewriting")
	flag.StringVar(&a.Params.NewHeader, "H", "", "add this string as the first row")
	flag.BoolVar(&a.Params.Overwrite, "f", false, "overwrite output file if it exists")
	flag.BoolVar(&a.Params.UseCRLF, "l", false, "use CRLF instead of LF for line endings in the output file")
}

// LoadPreset reads preset of the specified name and overwrites corresponding values
// in Args.Params.
func (a *Args) LoadPreset(name string) error {
	preset, err := presets.LoadPreset(name)
	if err != nil {
		return err
	}

	a.Params.Encoding = preset.Encoding
	a.Params.Pattern = preset.Pattern
	a.Params.SkipHeader = preset.SkipHeader
	a.Params.NewHeader = preset.NewHeader
	a.Params.UseCRLF = preset.UseCRLF

	return nil
}

// Validate enforces invariants of the Args struct.
func (a *Args) Validate() error {
	if len(a.Params.Input) == 0 {
		return errors.New("err: input file not specified")
	}

	if len(a.Params.Output) == 0 {
		return errors.New("err: output file not specified")
	}

	if len(a.Params.Encoding) == 0 {
		return errors.New("err: input encoding not specified")
	}

	return nil
}
