package parser

import (
	"github.com/taichi-kobaya/go-monkey-interpreter/monkey/ast"
	"github.com/taichi-kobaya/go-monkey-interpreter/monkey/lexer"
	"github.com/taichi-kobaya/go-monkey-interpreter/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 2つトークンを読み込む。curTokenとpeekTokenの両方がセットされる。
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
