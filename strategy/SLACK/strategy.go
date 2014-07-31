package slack

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../../personality"
	".."
	"bytes"
)


type SLACK_MESSAGE struct {
	Username string `json:"username"`
	IconUrl string `json:"icon_url"`
	Text string `json:"text"`
	Channel string `json:"channel"`
}

type Slack struct {
	Token string
}

func NewSlack(token string) Slack {
	return Slack{
		Token: token,
	}
}

const hookUrl string = "https://internlandia.slack.com/services/hooks/incoming-webhook?token="

func (s *Slack) SEND_IT_SEND_IT_SEND_IT(data *strategy.SWOLE_DATA) {

	text := "IT'S " + personality.TELL_ME_THE_HOUR(data.Time.Hour()) + "! "
	text += personality.INSPIRE_ME_PLEASE() + " "
	if data.Workout.Amount > 0 {
		text += string(data.Workout.Amount) + " "
	}
	text += data.Workout.Kind + ". "
	text += "GO!"

	msg := &SLACK_MESSAGE{
		Username: "SWOLEBOT_TEST_TEST_TEST (Golang™ Flavored)",
		IconUrl: "http://vaks.in/wp-content/uploads/2012/07/e14c.png",
		Text: text,
		// Channel: "swole",
		Channel: "#hey_guys",
	}

	body, _ := json.Marshal(msg)
	bodyReader := bytes.NewReader(body)
	resp, _ := http.Post(hookUrl + s.Token, "application/json", bodyReader)

	fmt.Printf("message sent: %s %s\n\n", resp.Status, msg);
}
