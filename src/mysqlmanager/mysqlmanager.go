package mysqlmanager

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"sync"
	"time"
)

type MysqlManager struct {
	DB *sql.DB
}

type User struct {
	PhoneNum string
	NickName string
	Token string
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
		_, tErr := db.Exec(fmt.Sprintf("create table if not exists %s (id int primary key AUTO_INCREMENT, password TEXT, phoneNum TEXT, nickName TEXT, token TEXT NULL) AUTO_INCREMENT = 10000", tname))
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

// 注册行为
func (mysql *MysqlManager) Regist (params map[string]interface{}, callBack func(json string))  {

	mysql.query(params, func(i map[string]interface{}) {
		if i["resultCode"] == "000"{
			// 查询到了对应的结果表明已经注册的有账号
			jsonObj := map[string]interface{}{"resultCode": "100", "msg": "用户存在"}
			b, _ := json.Marshal(jsonObj)
			callBack(string(b))
		} else {
			mysql.insert(params, func(json string) {
				callBack(json)
			})
		}
	})
}

// 登录行为
func (mysql *MysqlManager) Login (params map[string]interface{}, callBack func(json string)) {
	mysql.query(params, func(i map[string]interface{}) {
		jsonObj := make(map[string]interface{})
		if i["resultCode"] == "000"{
			loginInfo := (i["data"].([]map[string]interface{}))[0]
			if loginInfo["phoneNum"] == params["phoneNum"] && loginInfo["password"] == params["password"] {
				timeUnix := time.Now().Unix()
				timStr := strconv.FormatInt(timeUnix, 10)
				timStr = timStr[len(timStr)-4 : len(timStr)]
				idStr := loginInfo["id"].(string)
				w := md5.New()
				w.Write([]byte(idStr + "gotest" + timStr))
				tokenStr := hex.EncodeToString(w.Sum([]byte("gotest")))
				mysql.update(tokenStr, loginInfo["phoneNum"].(string), func(i map[string]interface{}) {
					if i["resultCode"] == "000" {
						jsonObj["resultCode"] = "000"
						jsonObj["msg"] = "登录成功"
						jsonObj["token"] = tokenStr
					}
				})
			} else {
				jsonObj["resultCode"] = "100"
				jsonObj["msg"] = "密码错误"
			}
		} else {
			jsonObj["resultCode"] = "100"
			jsonObj["msg"] = "还没有注册"

		}
		b, _ := json.Marshal(jsonObj)
		callBack(string(b))
	})
}

func (mysql *MysqlManager) insert (params map[string]interface{}, callBack func(json string)) {

	defer Manger.DB.Close()
	// 在插入数据时可能 数据库已经关闭
	_, useErr := Manger.DB.Exec("USE " + DBname)
	if useErr != nil {
		fmt.Println("使用数据库失败")
		panic(useErr)
	}

	tx, _ := Manger.DB.Begin()
	queryStr := fmt.Sprintf("INSERT INTO %s (password, phoneNum, nickName) values(?, ?, ?)", Tname)
	_, err := tx.Exec(queryStr, params["password"], params["phoneNum"], params["nickName"])
	if err != nil {
		fmt.Println("数据插入失败", err)
	}
	phoneNum := params["phoneNum"].(string)
	nickName := params["nickName"].(string)
	//var user User
	//user.PhoneNum = phoneNum
	//user.NickName = nickName
	//b, _ := json.Marshal(user)
	b, _ := json.Marshal(map[string]interface{}{"phoneNum": phoneNum, "nickName": nickName})
	callBack(string(b))
	tx.Commit()
}

func (mysql *MysqlManager) delete(dbname string) {

	defer Manger.DB.Close()
	tx, _ := Manger.DB.Begin()
	if len(dbname) > 0 {

	}
	tx.Commit()
}

func (mysql *MysqlManager) update(token, phoneNum string, callBack func(map[string]interface{})()) {

	defer Manger.DB.Close()
	_, useErr := Manger.DB.Exec("USE " + DBname)
	if useErr != nil {
		fmt.Println("使用数据库失败")
		panic(useErr)
	}
	tx, _ := Manger.DB.Begin()
	updateStr := fmt.Sprintf("UPDATE %s set token=%s where phoneNum=%s", Tname, token, phoneNum)
	fmt.Println(updateStr)
	_, err := tx.Exec(updateStr)
	if err != nil {
		panic(err)
	}
	callBack(map[string]interface{}{"resultCode": "000", "msg": "success"})
	tx.Commit()
}

func (mysql *MysqlManager) query(params map[string]interface{}, callBack func(map[string]interface{})) {

	defer Manger.DB.Close()
	tx, _ := Manger.DB.Begin()
	queryStr := fmt.Sprintf("SELECT * FROM %s where phoneNum=%s", Tname, params["phoneNum"])
	rows, _ := tx.Query(queryStr)
	defer rows.Close()
	var count = 0
	var dataList [] map[string]interface{}
	for rows.Next(){
		count++
		var id, password, phoneNum, nickName, token []byte
		item := make(map[string]interface{})
		if err := rows.Scan(&id, &password, &phoneNum, &nickName, &token); err != nil {
			panic(err)
		}
		item["id"] = string(id)
		item["password"] = string(password)
		item["phoneNum"] = string(phoneNum)
		item["nickName"] = string(nickName)
		item["token"] = string(token)
		dataList = append(dataList, item)
	}

	if count == 0{
		callBack(map[string]interface{}{"resultCode": "100", "msg": "fail"})
	} else {
		callBack(map[string]interface{}{"resultCode": "000", "msg": "success", "data": dataList})
	}
	tx.Commit()
}