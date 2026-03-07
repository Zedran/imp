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

package utils

import "strings"

// FormatCurrency removes all "cosmetic" separators from the amount and changes
// decimal separators to the decimalSeparator. The function assumes that
// the decimal separator in the input is either ',' or '.'.
func FormatCurrency(amount, decimalSeparator string) string {
	if len(amount) == 0 {
		return "0" + decimalSeparator + "00"
	}

	var b strings.Builder

	if strings.HasPrefix(amount, "(") && strings.HasSuffix(amount, ")") {
		if len(amount) == 2 {
			return "0" + decimalSeparator + "00"
		}
		amount = amount[1 : len(amount)-1]
		b.WriteByte('-')
	} else if strings.HasPrefix(amount, "-") {
		b.WriteByte('-')
		amount = amount[1:]
	} else if strings.HasPrefix(amount, "+") {
		amount = amount[1:]
	}

	var (
		liComma = strings.LastIndex(amount, ",")
		liDot   = strings.LastIndex(amount, ".")

		curDecSep rune
	)

	if liComma > liDot {
		curDecSep = ','
	} else {
		curDecSep = '.'
	}

	for _, c := range amount {
		if c >= '0' && c <= '9' {
			b.WriteRune(c)
		} else if c == curDecSep {
			b.WriteString(decimalSeparator)
		}
	}

	s := b.String()

	if !strings.Contains(s, decimalSeparator) {
		if len(s) == 0 {
			return "0" + decimalSeparator + "00"
		}
		return s + decimalSeparator + "00"
	}

	if strings.HasPrefix(s, decimalSeparator) {
		s = "0" + s
	}

	if strings.HasSuffix(s, decimalSeparator) {
		s += "00"
	}

	return s
}
