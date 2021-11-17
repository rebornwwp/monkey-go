package ast

import (
	"bytes"

	"example.com/m/token"
)

type Node interface {
	TokenLiteral() string
	// String every node to source code
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// root of program
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var buf bytes.Buffer
	for _, s := range p.Statements {
		buf.WriteString(s.String())
	}
	return buf.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

// why identifier is expression; eg: let x = valueProducingIdentifier
// in this situation. so
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var buf bytes.Buffer
	buf.WriteString(ls.TokenLiteral() + " ")
	buf.WriteString(ls.Name.String())
	buf.WriteString(" = ")

	if ls.Value != nil {
		buf.WriteString(ls.Value.String())
	}
	buf.WriteString(";")
	return buf.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var buf bytes.Buffer
	buf.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		buf.WriteString(rs.ReturnValue.String())
	}
	buf.WriteString(";")
	return buf.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

// !true, +1, -1
func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(pe.Operator)
	buf.WriteString(pe.Right.String())
	buf.WriteString(")")
	return buf.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// 1+1, 1*1
func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(ie.Left.String())
	buf.WriteString(" " + ie.Operator + " ")
	buf.WriteString(ie.Right.String())
	buf.WriteString(")")
	return buf.String()
}

var _ Statement = &LetStatement{}
var _ Statement = &ReturnStatement{}
var _ Statement = &ExpressionStatement{}
var _ Expression = &Identifier{}
var _ Expression = &IntegerLiteral{}
