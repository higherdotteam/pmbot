package main

import "fmt"
import "net/http"

import "io/ioutil"

func process(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println("path", r.URL.Path, string(b))
	fmt.Fprintf(w, "Hello astaxie!")
}

func main() {
	http.HandleFunc("/pmbot", process)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}
