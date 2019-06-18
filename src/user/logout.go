package user

import (
	"mysqlmanager"
)

func Logout(params map[string]interface{}, callBack func(respJson map[string]interface{}))  {
	mysqlmanager := mysqlmanager.ShareMysqlManager("goMysqlTest", "goMysqlTable")
	Query(mysqlmanager, params, func(json map[string]interface{}) {
		loginInfo := (json["data"].([]map[string]interface{}))[0]

		if loginInfo["resultCode"] == 100 {
			response["resultCode"] = "100"
			response["msg"] = "用户不存在"
			callBack(response)
		} else {
			if len(loginInfo["token"].(string)) > 0 {
				Update(mysqlmanager, "token", "", "phoneNum", params["phoneNum"].(string), func(i map[string]interface{}) {
					if i["resultCode"] == "000" {
						response["resultCode"] = "000"
						response["msg"] = "退出登录成功"
						callBack(response)
					}
				})

			} else {
				response["resultCode"] = "100"
				response["msg"] = "未登录"
				callBack(response)
			}
		}
	})
}