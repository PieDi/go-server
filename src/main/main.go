package main

import (
	"en_decrypt"
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

var privateKey  = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDfw1/P15GQzGGYvNwVmXIGGxea8Pb2wJcF7ZW7tmFdLSjOItn9
kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQolDOkEzNP0B8XKm+Lxy4giwwR5
LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TUGQD6QzKY1Y8xS+FoQQIDAQAB
AoGAbSNg7wHomORm0dWDzvEpwTqjl8nh2tZyksyf1I+PC6BEH8613k04UfPYFUg1
0F2rUaOfr7s6q+BwxaqPtz+NPUotMjeVrEmmYM4rrYkrnd0lRiAxmkQUBlLrCBiF
u+bluDkHXF7+TUfJm4AZAvbtR2wO5DUAOZ244FfJueYyZHECQQD+V5/WrgKkBlYy
XhioQBXff7TLCrmMlUziJcQ295kIn8n1GaKzunJkhreoMbiRe0hpIIgPYb9E57tT
/mP/MoYtAkEA4Ti6XiOXgxzV5gcB+fhJyb8PJCVkgP2wg0OQp2DKPp+5xsmRuUXv
720oExv92jv6X65x631VGjDmfJNb99wq5QJBAMSHUKrBqqizfMdOjh7z5fLc6wY5
M0a91rqoFAWlLErNrXAGbwIRf3LN5fvA76z6ZelViczY6sKDjOxKFVqL38ECQG0S
pxdOT2M9BM45GJjxyPJ+qBuOTGU391Mq1pRpCKlZe4QtPHioyTGAAMd4Z/FX2MKb
3in48c0UX5t3VjPsmY0CQQCc1jmEoB83JmTHYByvDpc8kzsD8+GmiPVrausrjj4p
y2DQpGmUic2zqCxl6qXMpBGtFEhrUbKhOiVOJbRNGvWW
-----END RSA PRIVATE KEY-----`

var publicKey  = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfw1/P15GQzGGYvNwVmXIGGxea
8Pb2wJcF7ZW7tmFdLSjOItn9kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQol
DOkEzNP0B8XKm+Lxy4giwwR5LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TU
GQD6QzKY1Y8xS+FoQQIDAQAB
-----END PUBLIC KEY-----`
func main() {
	//aesCFB := en_decrypt.InstanceAesCFB("6368616e676520746869732070617373")
	//encStr := aesCFB.CFBAesEncrypt("么么哒")
	//fmt.Println(encStr)
	//
	//decStr := aesCFB.CFBAesDecrypter(encStr)
	//fmt.Println(decStr)

	rsa := en_decrypt.InstanceRsa(privateKey, publicKey)
	rsaEncStr := rsa.RsaEnceypt("么么哒")
	fmt.Println(rsaEncStr)
	rsaDecStr := rsa.RsaDecrypt(rsaEncStr)
	fmt.Println(rsaDecStr)
	r := muxrouter.MuxRouterInit()
	go r.PostRequest("/regist", 3000)
	go r.PostRequest("/login", 3000)
	go r.PostRequest("/logout", 3000)
	select {

	}

	//fmt.Printf("%p", &cUser)
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

}

