package flicker

import (
	"database/sql"
	"errors"
	"sync"
)

var (
	initOnce    sync.Once
	flickerStmt *sql.Stmt // 常驻内存, 不释放资源, 和进程共存亡
)

func Init(DB *sql.DB) {
	if DB == nil {
		panic("nil *sql.DB")
	}
	initOnce.Do(func() {
		var err error
		if flickerStmt, err = DB.Prepare("REPLACE INTO flicker(stub) VALUES('a')"); err != nil {
			panic(err.Error())
		}
	})
}

var errInitNotCalled = errors.New("Init Not Called")

func NextID() (int64, error) {
	if flickerStmt == nil {
		return 0, errInitNotCalled
	}
	rslt, err := flickerStmt.Exec()
	if err != nil {
		return 0, err
	}
	return rslt.LastInsertId()
}
