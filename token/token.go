package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}


/*
*/
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	LT = "<"
	GT = "<"

	FUNCTION = "FUNCTION"
	LET = "LET"
)


