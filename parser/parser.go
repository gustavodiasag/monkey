package parser

import (
    "fmt"

	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type (
    prefixParseFn func() ast.Expression
    infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l         *lexer.Lexer
	errors    []string

	curToken  token.Token
	peekToken token.Token

    prefixParseFns map[token.TokenType]prefixParseFn
    infixParseFns map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// Read two tokens, so curToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStmt()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
    p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
    p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseStmt() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStmt()
    case token.RETURN:
        return p.parseReturnStmt()
	default:
		return p.parseExpressionStmt()
	}
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
    stmt := &ast.ReturnStatement{Token: p.curToken}

    p.nextToken()

    for !p.curTokenIs(token.SEMICOLON) {
        p.nextToken()
    }
    return stmt
}

func (p *Parser) parseExpressionStmt() *ast.ExpressionStatement {
    stmt := &ast.ExpressionStatement{Token: p.curToken}
    stmt.Expression = p.parseExpression(LOWEST)

    if p.peekTokenIs(token.SEMICOLON) {
        p.nextToken()
    }
    return stmt

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if !p.peekTokenIs(t) {
        p.peekError(t)
		return false
	}
	p.nextToken()
	return true
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected token: %s, got '%s'.", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
