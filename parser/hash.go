package parser

import (
	"monkey/ast"
	"monkey/token"
)

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{
		Token: p.curToken,
		Pairs: make([]ast.HashKeyValuePair, 0),
	}

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)

		hash.Pairs = append(hash.Pairs, ast.HashKeyValuePair{
			Key:   key,
			Value: value,
		})

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return hash
}
