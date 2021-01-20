package conn

import (
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// SampleConn is the sample connection
	SampleConn = &MysqlConn{
		Dsn:             "user:password@/demo",
		ConnMaxLiftTime: 3,
		MaxOpenConns:    10,
		MaxIdleConns:    2,
	}
)

type MysqlConn struct {
	Dsn             string
	ConnMaxLiftTime int
	MaxOpenConns    int
	MaxIdleConns    int
}

// Connectdb allow user to generate a connection
func Connectdb(conn *MysqlConn) *sqlx.DB {
	db, err := sqlx.Open("mysql", conn.Dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(conn.ConnMaxLiftTime))
	db.SetMaxOpenConns(conn.MaxOpenConns)
	db.SetMaxIdleConns(conn.MaxIdleConns)

	return db
}
