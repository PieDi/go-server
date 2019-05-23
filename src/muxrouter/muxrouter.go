package muxrouter

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
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

func (r MuxRouter) PostRequest(path string, params map[string]interface{}) {
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
	vars := mux.Vars(r)
	servicename := vars["servicename"] // 路径中 {} 参数

	// parse JSON body
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	servicetype := req["servicetype"].(string)

	// composite response body
	var res = map[string]string{"result":"succ", "name":servicename, "type":servicetype}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}