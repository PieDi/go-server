package user

import (
	"fmt"
	"mysqlmanager"
	"sync"
)

type User struct {
	NickName string `json:"nickName"`
	PhoneNum string `json:"phoneNum"`

}
var user *User

var once sync.Once
var response map[string]interface{}

func GetInstance() *User  {

	once.Do(func() {
		m := &User {}
		user = m
	})
	return user
}

func (p User)User()  {
	fmt.Println(p, p.NickName)
}

// 用户信息的查询
func Query(mysql *mysqlmanager.MysqlManager, params map[string]interface{}, callBack func(json map[string]interface{})) {

	defer mysql.DB.Close()
	tx, _ := mysql.DB.Begin()
	queryStr := fmt.Sprintf("SELECT * FROM %s where phoneNum=%s", mysql.Tname , params["phoneNum"])
	rows, _ := tx.Query(queryStr)
	defer rows.Close()
	var count = 0
	var dataList [] map[string]interface{}
	for rows.Next(){
		count++
		var id, password, phoneNum, nickName, token []byte
		item := make(map[string]interface{})
		if err := rows.Scan(&id, &password, &phoneNum, &nickName, &token); err != nil {
			fmt.Println("数据库查询")
			panic(err)
		}
		item["id"] = string(id)
		item["password"] = string(password)
		item["phoneNum"] = string(phoneNum)
		item["nickName"] = string(nickName)
		item["token"] = string(token)
		dataList = append(dataList, item)
	}
	if count > 0 {
		response = map[string]interface{}{"resultCode": "000", "msg": "查询成功", "data": dataList}
	} else {
		response = map[string]interface{}{"resultCode": "100", "msg": "用户不存在"}
	}
	callBack(response)
	tx.Commit()
}

//用户数据更新
func Update(mysql *mysqlmanager.MysqlManager, updateKey, updateValue, whereKey, whereValue string, callBack func(map[string]interface{})()) {

	defer mysql.DB.Close()
	_, useErr := mysql.DB.Exec("USE " + mysql.DBname)
	if useErr != nil {
		fmt.Println("使用数据库失败")
		panic(useErr)
	}
	tx, _ := mysql.DB.Begin()
	updateStr := fmt.Sprintf("UPDATE %s set %s=? where %s=?", mysql.Tname, updateKey, whereKey)
	fmt.Println(updateStr)
	_, err := tx.Exec(updateStr, updateValue, whereValue)
	if err != nil {
		panic(err)
	}
	callBack(map[string]interface{}{"resultCode": "000", "msg": "更新数据成功"})
	tx.Commit()
}

// 用户数据写入
func Insert (mysql *mysqlmanager.MysqlManager, params map[string]interface{}, callBack func(map[string]interface{})) {

	defer mysql.DB.Close()
	// 在插入数据时可能 数据库已经关闭
	_, useErr := mysql.DB.Exec("USE " + mysql.DBname)
	if useErr != nil {
		fmt.Println("使用数据库失败")
		panic(useErr)
	}

	tx, _ := mysql.DB.Begin()
	queryStr := fmt.Sprintf("INSERT INTO %s (password, phoneNum, nickName) values(?, ?, ?)", mysql.Tname)
	_, err := tx.Exec(queryStr, params["password"], params["phoneNum"], params["nickName"])
	phoneNum := params["phoneNum"].(string)
	nickName := params["nickName"].(string)
	if err != nil {
		callBack(map[string]interface{}{"resultCode": "100", "msg": "注册失败", "phoneNum": phoneNum, "nickName": nickName})
		fmt.Println("数据插入失败", err)
	}
	callBack(map[string]interface{}{"resultCode": "000", "msg": "注册成功", "phoneNum": phoneNum, "nickName": nickName})
	tx.Commit()
}