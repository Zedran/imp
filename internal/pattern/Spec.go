package pattern

// Spec is derived from a user-provided string pattern. It contains
// information about the structure of the output CSV file.
type Spec struct {
	// A sequence of Token structs that indicates the new row layout.
	Tokens []Token `json:"tokens"`

	// CSV separator character.
	Comma string `json:"comma"`
}
