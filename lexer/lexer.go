package lexer

import "go-hacks/token"

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func New(str string) *Lexer {
	l := &Lexer{input: str}

	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()

	switch l.ch {
	case '=' :
		tok = newToken(token.ASSIGN, l.ch)
	case ';' :
		tok = newToken(token.SEMICOLON, l.ch)
	case '(' :
		tok = newToken(token.LPAREN, l.ch)
	case ')' :
		tok = newToken(token.RPAREN, l.ch)
	case '{' :
		tok = newToken(token.LBRACE, l.ch)
	case '}' :
		tok = newToken(token.RBRACE, l.ch)
	case ',' :
		tok = newToken(token.COMMA, l.ch)
	case '+' :
		tok = newToken(token.PLUS, l.ch)
	case '-' :
		tok = newToken(token.MINUS, l.ch)
	case '<' :
		tok = newToken(token.LT, l.ch)
	case '>' :
		tok = newToken(token.GT, l.ch)
	case '*' :
		tok = newToken(token.ASTERISK, l.ch)
	case '/' :
		tok = newToken(token.SLASH, l.ch)
	case '!' :
		tok = newToken(token.BANG, l.ch)
	case 0 :
		tok.Literal = ""
		tok.Type = token.EOF
	default :
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch){
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_'
}

func isDigit(s byte) bool {
	return s >= '0' && s <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

var keywords = map[string]token.TokenType {
	"fn" : token.FUNCTION,
	"let" : token.LET,
}

func LookupIdent(ident string) token.TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return token.IDENT
}


func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.readChar()
	}
}





