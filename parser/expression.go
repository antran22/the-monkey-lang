package parser

import (
	"monkey/ast"
	"monkey/token"
)

type precedenceLevel int

const (
	_ precedenceLevel = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var tokenPrecendence = map[token.TokenType]precedenceLevel{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

// check token precedence

func (p *Parser) peekPrecedence() precedenceLevel {
	if p, ok := tokenPrecendence[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() precedenceLevel {
	if p, ok := tokenPrecendence[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

// parse function

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
