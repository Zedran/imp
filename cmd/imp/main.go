package main

import (
	"flag"
	"log"

	icsv "github.com/Zedran/imp/internal/csv"
)

func main() {
	log.SetFlags(0)

	var (
		input    = flag.String("i", "", "input CSV file")
		output   = flag.String("o", "", "output CSV file")
		encoding = flag.String("c", "utf-8", "input file encoding")
		pattern  = flag.String("p", "", "pattern that determines how to rewrite the input file")
	)

	flag.Parse()

	if len(*input) == 0 {
		log.Fatal("err: input file not specified")
	}

	if len(*output) == 0 {
		log.Fatal("err: output file not specified")
	}

	if len(*encoding) == 0 {
		log.Fatal("err: input encoding not specified")
	}

	err := icsv.RewriteCSV(*input, *output, *encoding, *pattern)
	if err != nil {
		log.Fatal(err)
	}
}
