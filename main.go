package main

import "fmt"
import "net/http"
import "strings"

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

func main() {
	http.HandleFunc("/pmbot", process)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
