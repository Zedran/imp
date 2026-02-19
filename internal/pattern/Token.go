package pattern

// Indicates a type of value stored in a Token struct.
type TokenType byte

const (
	// Indicates that Token stores Column.
	TT_COLUMN TokenType = 'd'

	// Indicates that Token stores Text.
	TT_TEXT TokenType = 's'
)

// Token is an union-like type. It stores either CSV column number
// or an arbitrary text to be inserted.
type Token struct {
	// Indicates the value type stored in a Token struct.
	Type TokenType `json:"type"`

	// Column number. Content of the indicated column will be substituted
	// in place of this token during compilation process.
	Column int `json:"column"`

	// Arbitrary text that will be inserted as is during compilation process.
	Text string `json:"text"`
}

// Creates a new Token containing specified column number.
func NewColumnToken(columnNumber int) Token {
	return Token{TT_COLUMN, columnNumber, ""}
}

// Creates a new Token containing specified text.
func NewTextToken(text string) Token {
	return Token{TT_TEXT, -1, text}
}
