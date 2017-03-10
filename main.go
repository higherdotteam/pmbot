package main

import "fmt"
import "net/http"
import "strings"
import "io/ioutil"
import "github.com/nlopes/slack"
import "net/url"
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

func get_fscore(sc SlackCommand) string {
	endpoint := "https://api.spotify.com/v1/search?q=guns&type=artist&limit=30user=" +
		url.QueryEscape(sc.suser.RealName) + "&channel=" + url.QueryEscape(sc.cname)
	fmt.Println(endpoint)
	response, err := http.Get(endpoint)
	if err != nil {
		return err.Error()
	}

	bytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err.Error()
	}

	if response.StatusCode == 200 {
		return fmt.Sprintf("API got: %d bytes", len(bytes))
	} else {
		return fmt.Sprintf("%d %s", response.StatusCode, string(bytes))
	}
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

	if strings.Contains(sc.text, "fscore") {
		fscore := get_fscore(sc)
		fmt.Fprintf(w, fscore)
	} else {
		fmt.Fprintf(w, "Unknown command")
	}
}

func main() {
	http.HandleFunc("/pmbot", process)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}
