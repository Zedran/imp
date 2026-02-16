package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	icsv "github.com/Zedran/imp/internal/csv"
	"github.com/Zedran/imp/internal/presets"
)

func main() {
	log.SetFlags(0)

	var (
		input     = flag.String("i", "", "input CSV file")
		output    = flag.String("o", "", "output CSV file")
		encoding  = flag.String("c", "utf-8", "input file encoding")
		pattern   = flag.String("p", "", "pattern that determines how to rewrite the input file")
		skipHeader  = flag.Bool("1", false, "omit the first line (header) from the input when rewriting")
		overwrite = flag.Bool("f", false, "overwrite output file if it exists")
		genPreset = flag.Bool("G", false, "generate an empty preset file in user's home directory and exit")
		preset    = flag.String("P", "", "preset name to be used instead of -e and -p")
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
			log.Fatal(err)
		}
		os.Exit(0)
	}

	if len(*input) == 0 {
		log.Fatal("err: input file not specified")
	}

	if len(*output) == 0 {
		log.Fatal("err: output file not specified")
	}

	if len(*preset) > 0 {
		preset, err := presets.LoadPreset(*preset)
		if err != nil {
			log.Fatal(err)
		}
		*encoding = preset.Encoding
		*pattern = preset.Pattern
		*skipHeader = preset.SkipHeader
	}

	if len(*encoding) == 0 {
		log.Fatal("err: input encoding not specified")
	}

	err := icsv.RewriteCSV(*input, *output, *encoding, *pattern, *skipHeader, *overwrite)
	if err != nil {
		log.Fatal(err)
	}
}
