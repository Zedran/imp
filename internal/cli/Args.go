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

package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/Zedran/imp/internal/utils"
)

// Args struct binds, reads and validates all CLI arguments.
type Args struct {
	// Indicates whether the application should exit right after
	// calling cli.Parse. Set to true if -h or -G options are used.
	ExitEarly bool `json:"exit_early"`

	// Parameters required for the main functionality.
	Params utils.Params `json:"params"`
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
		help      = false
		genPreset = fs.Bool("G", false, "generate an empty preset file in user's home directory and exit")
		preset    = fs.String("P", "", "name of preset to be used instead of -0, -c, -e, -H, -l and -p")
		version   = fs.Bool("v", false, "display version information and exit")
	)

	if err := fs.Parse(args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			return a, fmt.Errorf("err: %w", err)
		}
		help = true
	}

	if help {
		a.ExitEarly = true
		a.usage(fs)
		return a, nil
	}

	if *version {
		a.ExitEarly = true
		a.version(fs)
		return a, nil
	}

	if *genPreset {
		a.ExitEarly = true
		return a, utils.GeneratePresetsFile()
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
	fs.StringVar(&a.Params.Input, "i", utils.USE_STD_STREAM, "input CSV file")
	fs.StringVar(&a.Params.Output, "o", utils.USE_STD_STREAM, "output CSV file")
	fs.StringVar(&a.Params.Encoding, "e", "utf-8", "input file encoding")
	fs.StringVar(&a.Params.Pattern, "p", "", "pattern that determines how to rewrite the input file")
	fs.StringVar(&a.Params.InputComma, "c", "", "comma character in the input file, if differs from output")
	fs.BoolVar(&a.Params.SkipHeader, "0", false, "omit the first row (header) from the input when rewriting")
	fs.StringVar(&a.Params.NewHeader, "H", "", "add this string as the first row")
	fs.BoolVar(&a.Params.Overwrite, "f", false, "overwrite output file if it exists")
	fs.BoolVar(&a.Params.UseCRLF, "l", false, "use CRLF instead of LF for line endings in the output file")
}

// loadPreset reads preset of the specified name and overwrites corresponding values
// in Args.Params.
func (a *Args) loadPreset(name string) error {
	preset, err := utils.LoadPreset(name)
	if err != nil {
		return err
	}

	a.Params.ApplyPreset(preset)

	return nil
}

// usage prints help message to os.Stderr.
func (a *Args) usage(fs *flag.FlagSet) {
	fs.SetOutput(os.Stdout)

	fmt.Printf("Usage: %s (-p PATTERN | -P PRESET) [OPTIONS]...\n\n", fs.Name())
	fmt.Printf("%s rewrites CSV files according to the specified set of parameters.\n\n", fs.Name())
	fmt.Printf("Options:\n")
	fs.PrintDefaults()
	fmt.Println(`
Pattern starts with the character serving as the CSV comma. The second character
will be interpreted as a tag prefix. Both the comma and tag prefix can be freely
chosen, but they need to be unique - they cannot be used anywhere else
in the pattern.

CSV content is shaped with a series of tags:
    - '/d<number>' causes imp to insert the column <number> from the input file.
      Comma character is allowed at the end.
    - '/s<text> causes imp to insert an arbitrary <text>.

There is also a special pattern ',*', indicating that imp should not perform any
column modifications. Other options ('-0ceHl') are still effective and encoding
is normalized.

Example:
    - input file structure:  'First name,Last name,Amount'
    - output file structure: 'Full name,Amount'
    - pattern:               ',/d0/s /d1,/d2'`,
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

// version prints version information.
func (a *Args) version(fs *flag.FlagSet) {
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("%s %s\n%s %s-%s\n\n", fs.Name(), info.Main.Version, info.GoVersion, runtime.GOOS, runtime.GOARCH)
		for _, d := range info.Deps {
			fmt.Printf("%s %s\n", d.Path, d.Version)
		}
	} else {
		fmt.Printf("%s (unknown version)\n", fs.Name())
	}

	fmt.Println(`
Copyright (C) 2026 Wojciech Głąb (github.com/Zedran)
License GNU GPL-3.0-only <https://gnu.org/licenses/gpl-3.0.html>
Third-party licenses: https://github.com/Zedran/imp/NOTICE.md`,
	)
}
