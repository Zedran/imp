package main

import (
	"log"

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
		return
	}

	err = icsv.RewriteCSV(args.Params)
	if err != nil {
		log.Fatal(err)
	}
}
