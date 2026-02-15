package pattern

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParsePattern converts user input into a Specification struct.
func ParsePattern(pattern string) (Spec, error) {
	if len(pattern) < 2 {
		return Spec{}, errors.New("err: empty pattern")
	}

	comma := rune(pattern[0])
	pref := rune(pattern[1])

	if comma == pref {
		return Spec{}, errors.New("err: comma and prefix are not unique characters")
	}

	if strings.HasSuffix(pattern, string(pref)) {
		return Spec{}, errors.New("err: incomplete group separator")
	}

	groups := strings.Split(pattern[2:], string(pref))

	tokens := make([]Token, 0, len(groups))

	for _, g := range groups {
		if len(g) < 2 {
			return Spec{}, fmt.Errorf("err: empty group: '%s'", g)
		}

		switch g[0] {
		case 'd':
			sub := g[1:]
			appendComma := false
			if strings.HasSuffix(sub, string(comma)) {
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
				tokens = append(tokens, NewTextToken(string(comma)))
			}
		case 's':
			tokens = append(tokens, NewTextToken(g[1:]))
		default:
			return Spec{}, fmt.Errorf("err: unknown group type specifier: '%s'", g)
		}
	}

	return Spec{tokens, comma}, nil
}
