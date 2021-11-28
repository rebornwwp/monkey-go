package lexer

import (
	"testing"

	"example.com/m/token"
)

func TestReadChar(t *testing.T) {
	input := "hello"
	l := New(input)
	for i, x := range []byte(input) {
		if x != l.ch {
			t.Fatalf("test [%d], read character wrong: expected: %q, got: %q", i, x, l.ch)
		}
		l.readChar()
	}

	x := byte(0)
	if l.ch != x {
		t.Fatalf("test no char to consume, read character wrong: expected: %q, got: %q", x, l.ch)
	}
}

func TestNextToken(t *testing.T) {
	input := "(){}+=,;"
	tests := []struct {
		expectedToken   token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)
	for i, expected := range tests {
		tk := lexer.NextToken()
		t.Logf("%+v\n", tk)
		if expected.expectedToken != tk.Type {
			t.Fatalf("tests[%d], token type wrong, expected %q, got %q", i, expected.expectedToken, tk.Type)
		}

		if expected.expectedLiteral != tk.Literal {
			t.Fatalf("tests[%d], token literal wrong, expected %q, got %q", i, expected.expectedLiteral, tk.Literal)
		}
	}
	tk := lexer.NextToken()
	t.Logf("%+v\n", tk)
	x := lexer.NextToken()
	t.Logf("%+v\n", x)
}

func TestNextToken1(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
}

let result = add(five, ten);

!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
`
	tests := []struct {
		expectedToken   token.TokenType
		expectedLiteral string
	}{
		// let five = 5;
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// let ten = 10;
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		// let add = fn(x, y) {x+y;}
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// let result = add(five, ten);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		// !-/*5;
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// 5 < 10 > 5;
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// if (5 < 10) {
		// 	return true;
		// } else {
		// 	return false;
		// }`
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// 10 == 10;
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		// 10 != 9;
		{token.INT, "10"},
		{token.NOTEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		// string
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},

		// array
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},

		// hash
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}
	lexer := New(input)
	for i, expected := range tests {
		tk := lexer.NextToken()
		if expected.expectedToken != tk.Type {
			t.Fatalf("tests[%d], token type wrong, expected %q, got %q", i, expected.expectedToken, tk.Type)
		}

		if expected.expectedLiteral != tk.Literal {
			t.Fatalf("tests[%d], token literal wrong, expected %q, got %q", i, expected.expectedLiteral, tk.Literal)
		}
	}
}
