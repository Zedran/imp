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
