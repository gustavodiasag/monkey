package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"else":   ELSE,
	"false":  FALSE,
	"fn":     FUNCTION,
	"if":     IF,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	// Operators.
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	STAR   = "*"
	SLASH  = "/"
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NEQ    = "!="
	// Identifiers and literals.
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	// Keywords.
	ELSE     = "ELSE"
	FALSE    = "FALSE"
	FUNCTION = "FUNCTION"
	IF       = "IF"
	LET      = "LET"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	// Illegal.
	ILLEGAL = "ILLEGAL"
	// Eof.
	EOF = "EOF"
)
