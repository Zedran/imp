package main

import (
	"log"
	"os"

	"github.com/Zedran/imp/internal/cli"
	icsv "github.com/Zedran/imp/internal/csv"
)

func main() {
	log.SetFlags(0)

	args, err := cli.Parse(nil)
	if err != nil {
		log.Fatal(err)
	}

	if args.ExitEarly {
		os.Exit(0)
	}

	err = icsv.RewriteCSV(args.Params)
	if err != nil {
		log.Fatal(err)
	}
}
