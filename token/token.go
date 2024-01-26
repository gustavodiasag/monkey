package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let":LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

const (
    // Delimiters.
    COMMA       = ","
    SEMICOLON   = ";"
    LPAREN      = "("
    RPAREN      = ")"
    LBRACE      = "{"
    RBRACE      = "}"
    // Operators.
    ASSIGN      = "="
    PLUS        = "+"
    MINUS       = "-"
    BANG        = "!"
    STAR        = "*"
    SLASH       = "/"
    LT          = "<"
    GT          = ">"
    // Identifiers and literals.
    IDENT       = "IDENT"
    INT         = "INT"
    // Keywords.
    FUNCTION    = "FUNCTION"
    LET         = "LET"
    // Illegal.
    ILLEGAL = "ILLEGAL"
    // Eof.
    EOF     = "EOF"
)
