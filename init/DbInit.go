package init

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DbCli *gorm.DB

func init() {
	fmt.Printf("start DB Init ...\n")
	open, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/Hest?parseTime=true&loc=Asia%2FShanghai"))
	if err != nil {
		panic(err)
	}
	db, _ := open.DB()
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxIdleTime(time.Minute)
	DbCli = open
}
