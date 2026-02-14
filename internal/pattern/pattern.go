package pattern

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Converts user input into a sequence of Tokens.
func ParsePattern(pattern string) ([]Token, error) {
	if len(pattern) < 1 {
		return []Token{}, errors.New("err: empty pattern")
	}

	sep := rune(pattern[0])

	if strings.HasSuffix(pattern, string(sep)) {
		return []Token{}, errors.New("err: incomplete group separator")
	}

	groups := strings.Split(pattern[1:], string(sep))

	tokens := make([]Token, 0, len(groups))

	for _, g := range groups {
		if len(g) < 2 {
			return []Token{}, fmt.Errorf("err: empty group: '%s'", g)
		}

		switch g[0] {
		case 'd':
			num, err := strconv.Atoi(g[1:])
			if err != nil {
				return []Token{}, fmt.Errorf("err: invalid character in column group: '%s'", g)
			}
			if num < 0 {
				return []Token{}, fmt.Errorf("err: negative number in column group: '%d'", num)
			}
			tokens = append(tokens, NewColumnToken(num))
		case 's':
			tokens = append(tokens, NewTextToken(g[1:]))
		}
	}

	return tokens, nil
}
