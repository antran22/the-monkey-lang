package parser

import (
	"monkey/ast"
	"monkey/token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expr := &ast.IfExpression{
		Token: p.curToken,
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	expr.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expr.ThenBranch = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expr.ElseBranch = p.parseBlockStatement()
	}

	return expr
}
