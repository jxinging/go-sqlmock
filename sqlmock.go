package sqlmock

import (
	"fmt"
	"io"

	"database/sql"

	. "database/sql/driver"
)

type mockDriver struct {
}

func (*mockDriver) Open(name string) (Conn, error) {
	return &mockConn{}, nil
}

type mockConn struct {
}

func (*mockConn) Prepare(query string) (Stmt, error) {
	return &mockStmt{queryStr: query}, nil
}

func (*mockConn) Close() error {
	return nil
}

func (*mockConn) Begin() (Tx, error) {
	return &mockTx{}, nil
}

type mockStmt struct {
	queryStr string
}

func (*mockStmt) Close() error {
	return nil
}

func (*mockStmt) NumInput() int {
	return -1
}

func (s *mockStmt) Exec(args []Value) (Result, error) {
	fmt.Printf("%s, %v\n", s.queryStr, args)
	return &mockResult{}, nil
}

func (*mockStmt) Query(args []Value) (Rows, error) {
	return &mockRows{}, nil
}

type mockResult struct{}

func (*mockResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (*mockResult) RowsAffected() (int64, error) {
	return 0, nil
}

type mockRows struct{}

func (*mockRows) Columns() []string {
	return []string{}
}

func (*mockRows) Close() error {
	return nil
}

func (*mockRows) Next(dest []Value) error {
	return io.EOF
}

type mockTx struct {
}

func (*mockTx) Commit() error {
	return nil
}

func (*mockTx) Rollback() error {
	return nil
}

func init() {
	sql.Register("iac-sqlmock", &mockDriver{})
}

