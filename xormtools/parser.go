package xormtools

import (
	"github.com/xwb1989/sqlparser"
)

type Converter interface {
	StmtToGo(stmt *sqlparser.DDL, pkgName, dbEngine string) (string, error)
}

func NewParser(c Converter) *Parser {
	p := new(Parser)
	p.c = c
	return p
}

type Parser struct {
	c Converter
}

func (c *Parser) Parse(sql string, model string, dbEngine string) (string, error) {
	statement, err := sqlparser.ParseStrictDDL(sql)
	if err != nil {
		return "", err
	}
	stmt, ok := statement.(*sqlparser.DDL)
	if !ok {
		return "", err
	}
	// convert to Go struct
	res, err := c.c.StmtToGo(stmt, model, dbEngine)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return res, nil
}
