package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	Icon = "/usr/share/icons/gnome/48x48/status/appointment-soon.png"
	Me = "me"
)

func parseMessage(sentence string, timeText string) (string, string) {
	body := strings.TrimSpace(strings.Replace(sentence, timeText, "", 1))
	if strings.HasPrefix(body, "to ") {
		return strings.Replace(body, "to ", "", 1), "to"
	}
	if strings.HasPrefix(body, "that ") {
		return strings.Replace(body, "that ", "", 1), "that"
	}
	return body, ""
}

func parseTime(text string) *when.Result {
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

func notifyShell(args []string) {
	sentence := strings.Join(args, " ")
	parsedTime := parseTime(sentence)
	todo, word := parseMessage(sentence, parsedTime.Text)
	// https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go
	echo := exec.Command("echo", "notify-send", "-i", Icon, "'"+strings.Title(todo)+"'", "'… your friendly GoReminder'")
	at := exec.Command("at", parsedTime.Time.Format("15:04 02.01.06"))
	read, write := io.Pipe()
	echo.Stdout = write
	at.Stdin = read
	var b2 bytes.Buffer
	at.Stdout = &b2
	echo.Start()
	at.Start()
	echo.Wait()
	write.Close()
	echo.Wait()
	io.Copy(os.Stdout, &b2)
	fmt.Println("Okay, I will remind you " + word + " " + todo + " at " + parsedTime.Time.Format("2006-01-02 15:04"))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Me + " subcommand is required")
		os.Exit(1)
	}
	meCommand := flag.NewFlagSet(Me, flag.ExitOnError)
	switch os.Args[1] {
	case Me:
		meCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
	if meCommand.Parsed() {
		notifyShell(meCommand.Args())
	}
}
