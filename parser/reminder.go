package parser

import (
	"github.com/olebedev/when"
	"strings"
	"github.com/olebedev/when/rules/en"
	"github.com/olebedev/when/rules/common"
	"time"
	"fmt"
	"os"
)

var prepositions = []string{
	"to",
	"that",
}

type Reminder struct {
	WhenResult  when.Result
	Body        string
	Preposition string
}

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
