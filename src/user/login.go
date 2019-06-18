package user

import (
	"crypto/md5"
	"encoding/hex"
	"mysqlmanager"
	"strconv"
	"time"
)


func Login(params map[string]interface{}, callBack func(respJson map[string]interface{}))  {

	//phoneNum := params["phoneNum"].(string)
	//nickName := params["nickName"].(string)
	//password := params["password"].(string)
	mysqlmanager := mysqlmanager.ShareMysqlManager("goMysqlTest", "goMysqlTable")
	Query(mysqlmanager, params, func(json map[string]interface{}) {
		loginInfo := (json["data"].([]map[string]interface{}))[0]

		if loginInfo["resultCode"] == 100 {
			response["resultCode"] = "100"
			response["msg"] = "用户不存在"
			callBack(response)
		} else {
			if loginInfo["phoneNum"] == params["phoneNum"] && loginInfo["password"] == params["password"] {
				timeUnix := time.Now().Unix()
				timStr := strconv.FormatInt(timeUnix, 10)
				timStr = timStr[len(timStr)-4 : len(timStr)]
				idStr := loginInfo["id"].(string)
				w := md5.New()
				w.Write([]byte(idStr + "gotest" + timStr))
				tokenStr := hex.EncodeToString(w.Sum([]byte("gotest")))
				Update(mysqlmanager, "token", tokenStr, "phoneNum", params["phoneNum"].(string), func(i map[string]interface{}) {
					if i["resultCode"] == "000" {
						response["resultCode"] = "000"
						response["msg"] = "登录成功"
						response["token"] = tokenStr
						callBack(response)
					}
				})

			} else {
				response["resultCode"] = "100"
				response["msg"] = "密码错误"
				callBack(response)
			}
		}
	})
}
