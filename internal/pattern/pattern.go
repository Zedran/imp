package pattern

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Spec is derived from a user-provided string pattern. It contains
// information about the structure of the output CSV file.
type Spec struct {
	// A sequence of Token structs that indicates the new row layout.
	Tokens []Token `json:"tokens"`

	// CSV separator character.
	Comma rune `json:"comma"`
}

// ParsePattern converts user input into a Specification struct.
func ParsePattern(pattern string) (Spec, error) {
	if len(pattern) < 2 {
		return Spec{}, errors.New("err: empty pattern")
	}

	comma := rune(pattern[0])
	sep := rune(pattern[1])

	if strings.HasSuffix(pattern, string(sep)) {
		return Spec{}, errors.New("err: incomplete group separator")
	}

	groups := strings.Split(pattern[2:], string(sep))

	tokens := make([]Token, 0, len(groups))

	for _, g := range groups {
		if len(g) < 2 {
			return Spec{}, fmt.Errorf("err: empty group: '%s'", g)
		}

		switch g[0] {
		case 'd':
			num, err := strconv.Atoi(g[1:])
			if err != nil {
				return Spec{}, fmt.Errorf("err: invalid character in column group: '%s'", g)
			}
			if num < 0 {
				return Spec{}, fmt.Errorf("err: negative number in column group: '%d'", num)
			}
			tokens = append(tokens, NewColumnToken(num))
		case 's':
			tokens = append(tokens, NewTextToken(g[1:]))
		default:
			return Spec{}, fmt.Errorf("err: unknown group type specifier: '%s'", g)
		}
	}

	return Spec{tokens, comma}, nil
}
