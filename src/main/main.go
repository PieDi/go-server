package main

import "muxrouter"

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

	//aesCFB := en_decrypt.InstanceAesCFB("6368616e67652074")
	//encStr := aesCFB.CFBAesEncrypt("么么哒")
	//fmt.Println(encStr)
	//fmt.Println(aesCFB.CFBAesDecrypter(encStr))

	//6368616e67652074
	//aesCBC := en_decrypt.InstanceAesCBC("6368616e67652074")
	//aesCBCStr := aesCBC.CBCAesEncrypt("么么哒-么么哒")
	//fmt.Println(aesCBCStr)
	//fmt.Println(aesCBC.CBCAesDecrypt(aesCBCStr))

	//aesECB := en_decrypt.InstanceAesECB("6368616e676520746368616e67652074")
	//aesEncStr := aesECB.ECBAesEncrypt("么么哒")
	//fmt.Println(aesEncStr)
	//fmt.Println(aesECB.ECBAesDecrypter(aesEncStr))


	//des := en_decrypt.InstanceDes("sasasasa")
	//desStr := des.DesEncrypt("么么哒-么么哒")
	//fmt.Println(desStr)
	//fmt.Println(des.DesDecrypt(desStr))
	//
	//triDes := en_decrypt.InstanceTripDes("123456789012345678901234")
	//triDesStr := triDes.TripleDesEncrypt("么么哒")
	//fmt.Println(triDesStr)
	//fmt.Println(triDes.TripleDesDecrypt(triDesStr))

	//rsa := en_decrypt.InstanceRsa(privateKey, publicKey)
	//rsaEncStr := rsa.RsaEnceypt("么么哒")
	//fmt.Println(rsaEncStr)
	//rsaDecStr := rsa.RsaDecrypt(rsaEncStr)
	//fmt.Println(rsaDecStr)


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


