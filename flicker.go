package flicker

import (
	"database/sql"
	"errors"
)

type Generator struct {
	db   *sql.DB
	stmt *sql.Stmt
}

// NewGenerator returns a new Generator
func NewGenerator(db *sql.DB) (*Generator, error) {
	if db == nil {
		return nil, errors.New("nil *sql.DB")
	}
	stmt, err := db.Prepare("REPLACE INTO flicker(stub) VALUES('a')")
	if err != nil {
		return nil, err
	}
	return &Generator{
		db:   db,
		stmt: stmt,
	}, nil
}

// NextID returns a new ID.
func (p *Generator) NextID() (int64, error) {
	rslt, err := p.stmt.Exec()
	if err != nil {
		return 0, err
	}
	return rslt.LastInsertId()
}

// Close releases resources that opened by NewGenerator.
func (p *Generator) Close() error {
	if p.stmt != nil {
		return p.stmt.Close()
	}
	return nil
}
