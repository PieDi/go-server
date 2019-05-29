package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"muxrouter"
	"net/http"
	"strings"
)

const htmlStr = `<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    <input type="submit" value="登陆">
</form>
<select name="fruit">
<option value="apple">apple</option>
<option value="pear">pear</option>
<option value="banane">banane</option>
</select>
<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球
</body>
</html>`

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Host)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, htmlStr) //这个写入到w的是输出到客户端的
}

func handel(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, r.URL.Path)
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	queryVals := r.URL.Query()
	param, ok := queryVals["surname"]
	if ok {
		fmt.Println(queryVals, param)
	}
}

func main() {

	r := muxrouter.MuxRouterInit()
	//r.GetRequest("/articles/{id}", nil)
	r.PostRequest("/login")

	//r := mux.NewRouter()
	//r.HandleFunc("/", handel)
	//r.HandleFunc("/products", handel).Methods("POST")
	//r.HandleFunc("/articles", handel)
	//r.HandleFunc("/articles/{id}", handel).Methods("GET")
	//r.HandleFunc("/authors", handel).Queries("surname", "{surname}")
	//err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	//	pathTemplate, err := route.GetPathTemplate()
	//	if err == nil {
	//		fmt.Println("ROUTE:", pathTemplate)
	//	}
	//	pathRegexp, err := route.GetPathRegexp()
	//	if err == nil {
	//		fmt.Println("Path regexp:", pathRegexp)
	//	}
	//	queriesTemplates, err := route.GetQueriesTemplates()
	//	if err == nil {
	//		fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
	//	}
	//	queriesRegexps, err := route.GetQueriesRegexp()
	//	if err == nil {
	//		fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
	//	}
	//	methods, err := route.GetMethods()
	//	if err == nil {
	//		fmt.Println("Methods:", strings.Join(methods, ","))
	//	}
	//	fmt.Println()
	//	return nil
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//http.ListenAndServe(":3000", r)

	//mysqlmanager := mysqlmanager.ShareMysqlManager("goMysqlTest", "goMysqlTable")
	//fmt.Println(mysqlmanager)
	//fmt.Println("\n")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
