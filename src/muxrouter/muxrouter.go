package muxrouter

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

type MuxRouter struct {
	Router *mux.Router
}


func MuxRouterInit() MuxRouter {
	r := mux.NewRouter()
	MuxRouter := MuxRouter{r}
	return MuxRouter
}

func (r MuxRouter) GetRequest(path string, params map[string]interface{}) {

	var getParams = ""
	for k, v := range params{
		getParams = fmt.Sprintf("%s=%s", k, v)
	}
	if len(getParams) > 0 {
		path = path + "?" + getParams
	}
	fmt.Println(path)
	r.Router.HandleFunc(path, getHandel).Methods("GET")
	http.ListenAndServe(":3000", r.Router)
}

func (r MuxRouter) PostRequest(path string) {
	r.Router.HandleFunc(path, postHandel).Methods("POST")
	http.ListenAndServe(":3000", r.Router)
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
	fmt.Println(contentType)
	var reqParams map[string]interface{}
	var body []byte
	var phoneNum, nickName, password string
	//var err error
	if contentType == "multipart/form-data" {
		// form-data 请求
		//r.ParseMultipartForm(32<<20)
		r.ParseMultipartForm(32<<20)
		phoneNum = r.FormValue("phoneNum")
		nickName = r.PostFormValue("nickName")
		password = r.PostFormValue("password")
		aa, bb, _ := r.FormFile("phoneNum")
		fmt.Println(phoneNum, nickName, password,"\n", aa, bb)
	} else if contentType == "application/json" {
		// application/json 请求
		body, _ = ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &reqParams)
		fmt.Println(reqParams)
		phoneNum = reqParams["phoneNum"].(string)
		nickName = reqParams["nickName"].(string)
		password = reqParams["password"].(string)
	} else if contentType == "application/x-www-form-urlencoded" {
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
	}
	var res map[string]interface{}
	if len(password) > 0 {
		res = map[string]interface{}{"result": 0, "nickName": nickName, "type": phoneNum}
	}
	// composite response body
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", contentType)
	w.Write(response)
}