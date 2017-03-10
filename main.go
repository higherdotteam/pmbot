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
	fmt.Fprintf(w, "Processing...")

	if strings.Contains(sc.text, "fscore") {
	}
}

func main() {
	http.HandleFunc("/pmbot", process)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}
