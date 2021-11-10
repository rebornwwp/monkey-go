package parser

import (
	"fmt"

	"example.com/m/ast"
	"example.com/m/lexer"
	"example.com/m/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: make([]string, 0),
	}

	// set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: make([]ast.Statement, 0),
	}

	for !p.curTokenTypeIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// just parser let, identifier and '=', skip follows
	for !p.curTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenTypeIs(tk token.TokenType) bool {
	return p.curToken.Type == tk
}

func (p *Parser) peekTokenTypeIs(tk token.TokenType) bool {
	return p.peekToken.Type == tk
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenTypeIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

// todo: peekError in parser Statment
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
