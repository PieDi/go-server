package mysqlmanager

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type MysqlManager struct {
	DB *sql.DB
}

var Manger *MysqlManager
var Once sync.Once
var DBname string //数据库名字
var Tname string  //表名字

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

		_, tErr := db.Exec(fmt.Sprintf("create table if not exists %s (id int primary key, number char(9), name char(10))", tname))
		if tErr != nil {
			fmt.Println("数据库中建表失败")
			panic(tErr)
		}
	}
	Manger.DB = db
	//})
	DBname = dbname
	Tname = tname
	return Manger
}


func (mysql *MysqlManager) Insert (params []map[string]interface{}) {
	defer Manger.DB.Close()
	tx, _ := Manger.DB.Begin()
	for k, v := range params{
		fmt.Println(k, v)
	}
	tx.Commit()
}

func (mysql *MysqlManager) Delete(dbname string) {

	defer Manger.DB.Close()
	tx, _ := Manger.DB.Begin()
	if len(dbname) > 0 {

	}
	tx.Commit()
}

func (mysql *MysqlManager) Update(dbname string) {

	defer Manger.DB.Close()
	tx, _ := Manger.DB.Begin()
	if len(dbname) > 0 {

	}
	tx.Commit()
}

func (mysql *MysqlManager) Query(dbname string) {

	defer Manger.DB.Close()
	tx, _ := Manger.DB.Begin()
	if len(dbname) > 0 {

	}
	tx.Commit()
}