package pattern

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// ParsePattern converts user input into a Specification struct.
func ParsePattern(pattern string) (Spec, error) {
	if len(pattern) < 2 {
		return Spec{}, errors.New("err: empty pattern")
	}

	var (
		comma = pattern[:1]
		pref  = pattern[1:2]
	)

	if comma == pref {
		return Spec{}, errors.New("err: comma and prefix are not unique characters")
	}

	if strings.HasSuffix(pattern, pref) {
		return Spec{}, errors.New("err: incomplete group separator")
	}

	groups := strings.Split(pattern[2:], pref)

	tokens := make([]Token, 0, len(groups))

	commaRune, _ := utf8.DecodeRuneInString(comma)

	for _, g := range groups {
		if len(g) < 2 {
			return Spec{}, fmt.Errorf("err: empty group: '%s'", g)
		}

		switch TokenType(g[0]) {
		case TT_COLUMN:
			sub := g[1:]
			appendComma := false
			if strings.HasSuffix(sub, comma) {
				sub = sub[:len(sub)-1]
				appendComma = true
			}
			num, err := strconv.Atoi(sub)
			if err != nil {
				return Spec{}, fmt.Errorf("err: invalid character in column group: '%s'", g)
			}
			if num < 0 {
				return Spec{}, fmt.Errorf("err: negative number in column group: '%d'", num)
			}
			tokens = append(tokens, NewColumnToken(num))
			if appendComma {
				tokens = append(tokens, NewTextToken(comma))
			}
		case TT_TEXT:
			var b strings.Builder
			for _, c := range g[1:] {
				if c == commaRune {
					if b.Len() > 0 {
						tokens = append(tokens, NewTextToken(b.String()))
						b.Reset()
					}
					tokens = append(tokens, NewTextToken(comma))
				} else {
					b.WriteRune(c)
				}
			}
			if b.Len() > 0 {
				tokens = append(tokens, NewTextToken(b.String()))
			}
		default:
			return Spec{}, fmt.Errorf("err: unknown group type specifier: '%s'", g)
		}
	}

	return Spec{tokens, comma}, nil
}
