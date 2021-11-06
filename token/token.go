package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL" // A token we don't know about/understand
	EOF = "EOF" // End of file

	// Identifiers and literals
	IDENT = "IDENT" // add, x, y, etc.
	INT = "INT" // 123456

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

type Token struct {
	Type TokenType
	Literal string
}
