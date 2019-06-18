package mysqlmanager

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type MysqlManager struct {
	DB *sql.DB
	DBname string //数据库名字
	Tname string
}

type User struct {
	PhoneNum string
	NickName string
	Token string
}

var Manger *MysqlManager
var Once sync.Once

func ShareMysqlManager(dbname, tname string) *MysqlManager  {
	//once.Do(func() {  //单例写法
	Manger = &MysqlManager{}
	var  mysqlInfo string = "root:Mysql123!@tcp(0.0.0.0:3306)/"
	db, openErr := sql.Open("mysql", mysqlInfo)
	if len(dbname) > 0 {
		if openErr != nil {
			fmt.Println("打开数据库失败", mysqlInfo)
			panic(openErr)
		}
		_, dbErr := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
		if dbErr != nil {
			fmt.Println("创建数据库失败")
			panic(dbErr)
		}
		_, useErr := db.Exec("USE " + dbname)
		if useErr != nil {
			fmt.Println("使用数据库失败")
			panic(useErr)
		}
		_, tErr := db.Exec(fmt.Sprintf("create table if not exists %s (id int primary key AUTO_INCREMENT, password TEXT, phoneNum TEXT, nickName TEXT, token TEXT NULL) AUTO_INCREMENT = 10000", tname))
		if tErr != nil {
			fmt.Println("数据库中建表失败")
			panic(tErr)
		}
	}
	Manger.DB = db
	//})
	Manger.DBname = dbname
	Manger.Tname = tname
	return Manger
}


func (mysql *MysqlManager) Delete(dbname string) {

	defer Manger.DB.Close()
	tx, _ := Manger.DB.Begin()
	if len(dbname) > 0 {

	}
	tx.Commit()
}




