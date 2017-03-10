package main

import "fmt"
import "net/http"
import "strings"
import "io/ioutil"
import "github.com/nlopes/slack"
import "os"

type SlackCommand struct {
	team    string
	domain  string
	channel string
	cname   string
	user    string
	uname   string
	command string
	text    string
	suser   slack.User
}

func process(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)
	data := string(b)
	sc := SlackCommand{}
	for _, v := range strings.Split(data, "&") {
		tokens := strings.Split(v, "=")
		field := tokens[0]
		if field == "team_id" {
			sc.team = tokens[1]
		} else if field == "team_domain" {
			sc.domain = tokens[1]
		} else if field == "channel_id" {
			sc.channel = tokens[1]
		} else if field == "channel_name" {
			sc.cname = tokens[1]
		} else if field == "user_id" {
			sc.user = tokens[1]
		} else if field == "user_name" {
			sc.uname = tokens[1]
		} else if field == "command" {
			sc.command = tokens[1]
		} else if field == "text" {
			sc.text = tokens[1]
		}
	}
	fmt.Println(sc)
	api := slack.New(os.Getenv("SLACK_PROPOSAL_BOT"))
	list, _ := api.GetUsers()
	for _, u := range list {
		if u.ID == sc.user {
			sc.suser = u
			break
		}
	}
	fmt.Println(sc.suser.Profile.Email)

	/*

		team_id T035N23CL
		team_domain mooon
		channel_id D1KD59XH9
		channel_name directmessage
		user_id U035LF6C1
		user_name aa
		command %2Fpmbot
		text wefwef

			token=FRZTqnWNkEWVWS0YDWEA8LYV&team_id=T035N23CL&team_domain=mooon&channel_id=D1KD59XH9&channel_name=directmessage&user_id=U035LF6C1&user_name=aa&command=%2Fpmbot&text=wefwef&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT035N23CL%2F153490741014%2FX4k1pHZ79iQnX7E0apdW4Aes

			token=FRZTqnWNkEWVWS0YDWEA8LYV&team_id=T035N23CL&team_domain=mooon&channel_id=G3KGJUX7A&channel_name=privategroup&user_id=U035LF6C1&user_name=aa&command=%2Fpmbot&text=here&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT035N23CL%2F152849098884%2F6Wxo8D86YYSXRHdUx9yg2rQb

	*/
	fmt.Fprintf(w, "Hello astaxie!")
}

func main() {
	http.HandleFunc("/pmbot", process)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}
