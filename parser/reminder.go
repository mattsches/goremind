package parser

import (
	"fmt"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"os"
	"strings"
	"time"
)

var prepositions = []string{
	"to",
	"that",
}

// Reminder contains information about the timing and the contents of the reminder
type Reminder struct {
	WhenResult  when.Result
	Body        string
	Preposition string
}

// Message returns the Reminder with all properties set
func Message(whenResult *when.Result) *Reminder {
	r := &Reminder{
		WhenResult:  *whenResult,
		Body:        strings.TrimSpace(strings.Replace(whenResult.Source, whenResult.Text, "", 1)),
		Preposition: "",
	}
	for _, prep := range prepositions {
		if strings.HasPrefix(r.Body, prep+" ") {
			r.Preposition = prep
			r.Body = strings.Replace(r.Body, prep+" ", "", 1)
			break
		}
	}
	return r
}

// Time tries to parse and extract the date and time for the reminder from the input string
func Time(text string) *when.Result {
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)
	r, err := w.Parse(text, time.Now())
	if err != nil {
		fmt.Println("An error has occurred :-(")
		os.Exit(1)
	}
	if r == nil {
		fmt.Println("Could not parse time and date, fall back to default (now + 1 hour)")
		r := new(when.Result)
		r.Time = time.Now().Add(time.Hour)
		return r
	}
	return r
}
