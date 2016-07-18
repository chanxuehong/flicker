package flicker

import (
	"database/sql"
	"errors"
	"sync"
)

var (
	initOnce sync.Once
	stmt     *sql.Stmt // 常驻内存, 不释放资源, 和进程共存亡
)

func Init(DB *sql.DB) {
	if DB == nil {
		panic("nil *sql.DB")
	}
	var err error
	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}
	initOnce.Do(func() {
		if stmt, err = DB.Prepare("REPLACE INTO flicker(stub) VALUES('a')"); err != nil {
			panic(err.Error())
		}
	})
}

var ErrInitNotCalled = errors.New("Init Not Called")

func NextID() (int64, error) {
	if stmt == nil {
		return 0, ErrInitNotCalled
	}
	rslt, err := stmt.Exec()
	if err != nil {
		return 0, err
	}
	return rslt.LastInsertId()
}
