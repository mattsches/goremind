package command

import (
	"bytes"
	"fmt"
	"github.com/mattsches/goremind/parser"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	Icon = "/usr/share/icons/gnome/48x48/status/appointment-soon.png" // Icon to displayed alongside the reminder
)

// Me takes the Reminder and passes it on the the system
// https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go
func Me(r *parser.Reminder) {
	if r.Body == "" {
		fmt.Println("Empty reminder body, aborting!")
		os.Exit(1)
	}
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
