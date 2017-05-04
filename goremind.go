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
	Me = "me"
)

func parseMessage(args []string) (string, string) {
	for idx, element := range args {
		if strings.Compare(element, "to") == 0 {
			body := args[idx+1:]
			return strings.Join(body, " "), "to"
		}
		if strings.Compare(element, "that") == 0 {
			body := args[idx+1:]
			return strings.Join(body, " "), "that"
		}
	}
	return "", ""
}

func parseTime(text string) time.Time {
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
		return time.Now().Add(time.Hour)
	}
	return r.Time
}

func notifyShell(args []string) {
	parsedTime := parseTime(strings.Join(args, " "))
	todo, word := parseMessage(args)
	// https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go
	echo := exec.Command("echo", "notify-send", "-i", "/usr/share/icons/gnome/48x48/status/appointment-soon.png", "'"+strings.Title(todo)+"'", "'… your friendly GoReminder'")
	at := exec.Command("at", parsedTime.Format("15:04 02.01.06"))
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
	fmt.Println("Okay, I will remind you at " + parsedTime.Format("2006-01-02 15:04") + " " + word + " " + todo)
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
