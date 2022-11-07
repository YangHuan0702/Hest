package init

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DbCli *sql.DB

func init() {
	fmt.Printf("start DB Init ...\n")
	open, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/Hest?parseTime=true&loc=Asia%2FShanghai"))
	if err != nil {
		panic(err)
	}
	db, _ := open.DB()
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxIdleTime(time.Minute)
	DbCli = db
}
