package main

import "fmt"
import "net/http"

import "io/ioutil"

func process(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)
	/*

		token=FRZTqnWNkEWVWS0YDWEA8LYV&team_id=T035N23CL&team_domain=mooon&channel_id=D1KD59XH9&channel_name=directmessage&user_id=U035LF6C1&user_name=aa&command=%2Fpmbot&text=wefwef&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT035N23CL%2F153490741014%2FX4k1pHZ79iQnX7E0apdW4Aes
	*/
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
