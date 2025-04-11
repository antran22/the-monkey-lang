package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/token"
	"strconv"
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

func (p *Parser) parseExpression(precedence precedenceLevel) ast.Expression {
	prefix, ok := p.prefixParseFns[p.curToken.Type]
	if !ok {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

// expression types parser

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	value, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	return &ast.IntegerLiteral{Token: p.curToken, Value: value}
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
}

// parse expression

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}
