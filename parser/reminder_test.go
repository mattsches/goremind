package parser

import (
	"github.com/olebedev/when"
	"testing"
	"time"
)

type testPair struct {
	WhenResult *when.Result
	Reminder   *Reminder
}

var mockTime = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

var tests = []testPair{
	{
		&when.Result{Index: 0, Source: "to go home in 1 minute", Text: "in 1 minute", Time: mockTime},
		&Reminder{Preposition: "to", Body: "go home", WhenResult: when.Result{Index: 0, Source: "to go home in 1 minute", Text: "in 1 minute", Time: mockTime}},
	},
	{
		&when.Result{Index: 0, Source: "tomorrow at 4pm that it is time to go home", Text: "tomorrow at 4pm", Time: mockTime},
		&Reminder{Preposition: "that", Body: "it is time to go home", WhenResult: when.Result{Index: 0, Source: "tomorrow at 4pm that it is time to go home", Text: "tomorrow at 4pm", Time: mockTime}},
	},
}

func TestMessage(t *testing.T) {
	for _, pair := range tests {
		v := Message(pair.WhenResult)
		if v.Preposition != pair.Reminder.Preposition {
			t.Error("For '", pair.WhenResult.Source, "' expected '", pair.Reminder.Preposition, "' got '", v.Preposition, "'")
		}
		if v.Body != pair.Reminder.Body {
			t.Error("For '", pair.WhenResult.Source, "' expected '", pair.Reminder.Body, "' got '", v.Body, "'")
		}
	}
}
