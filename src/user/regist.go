package user

import (
	"mysqlmanager"
)

func Regist(params map[string]interface{}, callBack func(respJson map[string]interface{}))  {
	mysqlmanager := mysqlmanager.ShareMysqlManager("goMysqlTest", "goMysqlTable")
	Query(mysqlmanager, params, func(json map[string]interface{}) {
		loginInfo := (json["data"].([]map[string]interface{}))[0]

		if loginInfo["resultCode"] == 100 {
			Insert(mysqlmanager, params, func(i map[string]interface{}) {
				callBack(i)
			})
		} else {
			response["resultCode"] = "100"
			response["msg"] = "用户已存在"
			callBack(response)
		}
	})
}