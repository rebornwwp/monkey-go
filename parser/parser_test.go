package parser

import (
	"testing"

	"example.com/m/ast"
	"example.com/m/lexer"
)

func TestLetStatement(t *testing.T) {

	input := `
let x = 5;
let y = 10;
let foobar = 123456;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	if len(program.Statements) != len(tests) {
		t.Fatalf("program.Statements does not contain %d statements. got=%d", len(tests), len(program.Statements))
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

// todo: add more tests
func testLetStatement(t *testing.T, stmt ast.Statement, expectedName string) bool {
	t.Helper()
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral expected 'let'. got %q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt type is not ast.LetStatement. got %T", stmt)
	}

	if letStmt.Name.Value != expectedName {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", expectedName, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != expectedName {
		t.Errorf("s.Name not '%s'. got=%s", expectedName, letStmt.Name)
		return false
	}
	return true
}

func checkParserError(t *testing.T, p *Parser) {
	t.Helper()
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("Parser get %d errors", len(errors))
	for _, e := range errors {
		t.Errorf("Parser error: %s", e)
	}
	t.FailNow()
}
