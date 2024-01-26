package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
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
