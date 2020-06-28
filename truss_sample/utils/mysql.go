package utils

// docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	DRIVE_MYSQL          = "mysql"
	MYSQL_LINK_FORMATTER = "%v:%v@tcp(%v)/%v"
)

func OpenMysql(username, password, hostPort, dbName string) (db *sql.DB, err error) {
	db, err = sql.Open(DRIVE_MYSQL, fmt.Sprintf(MYSQL_LINK_FORMATTER, username, password, hostPort, dbName))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Second * 5)
	return db, nil
}
