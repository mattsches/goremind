package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/mattsches/goremind/parser"
	"os/exec"
	"io"
	"bytes"
)

const (
	Icon = "/usr/share/icons/gnome/48x48/status/appointment-soon.png"
	Me   = "me"
)

// https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go
func notifyShell(r *parser.Reminder) {
	echo := exec.Command("echo", "notify-send", "-i", Icon, "'"+strings.Title(r.Body)+"'", "'… your friendly GoReminder'")
	at := exec.Command("at", r.WhenResult.Time.Format("15:04 02.01.06"))
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
	fmt.Println("Okay, I will remind you " + r.Preposition + " \"" + r.Body + "\" at " + r.WhenResult.Time.Format("2006-01-02 15:04"))
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
		notifyShell(parser.Message(parser.Time(strings.Join(meCommand.Args(), " "))))
	}
}
