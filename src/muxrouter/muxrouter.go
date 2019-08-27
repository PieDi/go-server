package muxrouter

import (
	"encoding/json"
	"io/ioutil"
	"user"

	//"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"io/ioutil"
	"net/http"
	"time"
	//"user"
)

type MuxRouter struct {
	Router *mux.Router
}


func MuxRouterInit() MuxRouter {
	r := mux.NewRouter()
	MuxRouter := MuxRouter{r}
	return MuxRouter
}

func (r MuxRouter) GetRequest(path string, port int, params map[string]interface{}) {

	var getParams = ""
	for k, v := range params{
		getParams = fmt.Sprintf("%s=%s", k, v)
	}
	if len(getParams) > 0 {
		path = path + "?" + getParams
	}
	fmt.Println(path)
	r.Router.HandleFunc(path, getHandel).Methods("GET")
	http.ListenAndServe(fmt.Sprintf(":%d", port), r.Router)
}

func (r MuxRouter) PostRequest(path string, port int) *http.Server{
	r.Router.HandleFunc(path, postHandel).Methods("POST")
	//http.ListenAndServe(fmt.Sprintf(":%d", port), r.Router)
	server := &http.Server{Addr:fmt.Sprintf(":%d", port), WriteTimeout: time.Second * 3, Handler: r.Router}
	server.ListenAndServe()
	return server
}


func getHandel(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, r.URL.Path)
	vars := mux.Vars(r)   // 路径中 {} 参数
	id := vars["id"]
	fmt.Println(id)
	queryVals := r.URL.Query()
	param, ok := queryVals["surname"]  // 路径中 ? 参数
	if ok {
		fmt.Println(queryVals, param)
	}
}

func postHandel(w http.ResponseWriter, r *http.Request)  {

	rHeader := r.Header
	var contentType string
	if len(rHeader["Content-Type"]) > 0 {
		contentType = rHeader["Content-Type"][0]
	}
	var reqParams map[string]interface{}
	var body []byte
	var res map[string]interface{}

	//var err error
	if contentType == "application/json" {
		// application/json 请求
		body, _ = ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &reqParams)
		//mysqlmanager := mysqlmanager.ShareMysqlManager("goMysqlTest", "goMysqlTable")
		switch r.URL.Path {
		case "/login":
			user.Login(reqParams, func(respJson map[string]interface{}) {
				fmt.Println(respJson)
			})
		case "/regist":
			user.Regist(reqParams, func(respJson map[string]interface{}) {
				fmt.Println(respJson)
			})
		case "/logout":
			user.Logout(reqParams, func(respJson map[string]interface{}) {
				fmt.Println(respJson)
			})
		}
	} else if contentType == "application/x-www-form-urlencoded" {
		fmt.Println("application/x-www-form-urlencoded 请求")
		/*  1、
		r.ParseForm()
		formDara := r.Form
		phoneNum := formDara["phoneNum"][0]
		nickName := formDara["nickName"][0]
		password := formDara["password"][0]
		*/

		/*  2、
		phoneNum := r.PostFormValue("phoneNum")
		nickName := r.PostFormValue("nickName")
		password := r.PostFormValue("password")
		*/

		/* 3、
		reBody := make([]byte, r.ContentLength)
		r.Body.Read(reBody)
		fmt.Println(string(reBody))
		params := strings.Split(string(reBody), "&")
		for _, str := range params {
			key := strings.Split(str, "=")[0]
			value := strings.Split(str, "=")[1]
			switch key {
			case "phoneNum":
				phoneNum = value
				break
			case "nickName":
				nickName = value
				break
			case "password":
				password = value
				break
			}
		}
		*/
	} else {
		// form-data 请求
		fmt.Println("form-data 请求")
		UploadHandler(w, r)
		//UrlHandler(w, r)
		/*
		r.ParseMultipartForm(32<<20)
		user := r.Form.Get("user")
		password := r.Form.Get("password")
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err, "--------1------------")//上传错误
		}
		defer file.Close()

		fmt.Println(user, password, handler.Filename) //test 123456 json.zip
		*/

	}

	// composite response body
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", contentType)
	w.Write(response)
}