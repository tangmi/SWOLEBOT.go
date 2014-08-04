package main

import (
	"./strategy"
	"./strategy/slack"
	"./workouts"
	"fmt"
	"time"
	"os"
)

const timeFormat string = "Mon Jan 2 15:04:05 -0700 MST 2006"

func main() {

	fmt.Printf("swolebot server, lol\n")

	tickerChan := time.NewTicker(time.Second).C

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		fmt.Printf("env variable \"SLACK_TOKEN\" not set, set that please\n")
		os.Exit(1)
	}
	strategySlack := slack.NewSlack(token)

	// fencepost problem
	nextTime := GET_THE_NEXT_HOUR_PRETTY_PLEASE(time.Now())

	fmt.Printf("next event: %s\n", nextTime.Format(timeFormat))

	for {
		<-tickerChan // block like a baus

		current := time.Now()

		if current.After(nextTime) {
			fmt.Printf("sending message %s\n", current.Format(timeFormat))

			data := &strategy.SWOLE_DATA{
				Time:    current,
				Workout: workouts.GET_ME_THE_WORKOUT_FOR_RIGHT_NOW(nextTime),
			}
			strategySlack.SEND_IT_SEND_IT_SEND_IT(data)

			nextTime = GET_THE_NEXT_HOUR_PRETTY_PLEASE(time.Now())
			fmt.Printf("next event: %s\n", nextTime.Format(timeFormat))
		}
	}

}

func GET_THE_NEXT_HOUR_PRETTY_PLEASE(t time.Time) time.Time {
	n := t
	
	if n.Weekday() == time.Saturday || n.Weekday() == time.Sunday {
		// if we're on a weekend, push us to the next weekday at 9am
		n = UGH_WHENS_THE_NEXT_WEEKDAY_I_WANT_TO_WORKOUT(n)
		n = time.Date(n.Year(), n.Month(), n.Day(), 9, 0, 0, 0, n.Location());
	} else { // else weekday
		// if we're not during working hours... set the next event to 9am
		if t.Hour() >= 17 || t.Hour() < 9 {
			n = time.Date(n.Year(), n.Month(), n.Day(), 9, 0, 0, 0, n.Location());
			if t.Hour() >= 17 {
				// add a day if we're after 5pm
				n = n.Add(24 * time.Hour)

				// if we're	on a weekend, make it the next weekday
				n = UGH_WHENS_THE_NEXT_WEEKDAY_I_WANT_TO_WORKOUT(n)
			}
		} else {
			// otherwise set it to the next available hour
			n = t.Truncate(time.Hour).Add(time.Hour)
		}
	}

	return n
}

func UGH_WHENS_THE_NEXT_WEEKDAY_I_WANT_TO_WORKOUT(t time.Time) time.Time {
	n := t
	for n.Weekday() == time.Saturday || n.Weekday() == time.Sunday {
		n = n.Add(24 * time.Hour)
	}
	return n
}