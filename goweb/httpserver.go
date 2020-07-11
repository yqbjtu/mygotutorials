package main

import (
	"fmt"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func pathVariableAndParameter(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/pathvarandparam/")
	// http://127.0.0.1/pathvarandparam/5/abc?name=a1&name=a4,  id 会是 5/abc
	fmt.Fprintf(w, "id: %v\n", id)

	vars := req.URL.Query()
	name := vars["name"]

	name, ok := vars["name"]
	if !ok {
		fmt.Fprintf(w, "there is no param 'name'\n")
	} else {
		fmt.Fprintf(w, "name: %v\n", name)
	}
}

/*
   http://127.0.0.1/hello
   http://127.0.0.1/headersdemo
   http://127.0.0.1/pathvarandparam/5
 http://127.0.0.1/pathvarandparam/5?name=a1	param a value is [[a1]]
 http://127.0.0.1/pathvarandparam/5?name=a1&name=a2	param a value is [[a1 a2]]


*/
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headersdemo", headers)
	//http.HandleFunc("/pathvarandparam/:id/props", pathVariableAndParameter)
	http.HandleFunc("/pathvarandparam/", pathVariableAndParameter)

	http.ListenAndServe(":80", nil)
}
