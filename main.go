package main

import "fmt"
import "github.com/nlopes/slack"
import "os"
import "time"
import "strings"

var Me string

func handleRtm(rtm *slack.RTM) {

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Println(ev.Msg.Type, ev.Msg.Channel, ev.Msg.User, ev.Msg.Text)

				if ev.Msg.User != Me && strings.HasPrefix(ev.Msg.Channel, "D") {
					from := ev.Msg.User
					fmt.Println(from)
				}
			}
		}
	}

}

func main() {
	fmt.Println("listening for proposals...")
	api := slack.New(os.Getenv("SLACK_PROPOSAL_BOT"))
	list, _ := api.GetUsers()
	for _, u := range list {
		if u.Name == "pmbot" {
			Me = u.ID
			break
		}
	}

	rtm := api.NewRTM()

	go rtm.ManageConnection()
	go handleRtm(rtm)

	for {
		time.Sleep(time.Second)
	}
}
