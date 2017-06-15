package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/mattsches/goremind/parser"
	"github.com/mattsches/goremind/command"
	"os/exec"
)

const (
	Me   = "me"
	List = "list"
)

func checkExesExist(exes []string) {
	for _, exe := range exes {
		_, err := exec.LookPath(exe)
		if err != nil {
			fmt.Printf("didn't find required '%s' executable\n", exe)
			os.Exit(1)
		}
	}
}

func main() {
	checkExesExist([]string{"echo", "notify-send", "at"})
	if len(os.Args) < 2 {
		fmt.Println("A subcommand is required")
		os.Exit(1)
	}
	meCommand := flag.NewFlagSet(Me, flag.ExitOnError)
	listCommand := flag.NewFlagSet(List, flag.ExitOnError)
	switch os.Args[1] {
	case Me:
		meCommand.Parse(os.Args[2:])
		if meCommand.Parsed() {
			command.Me(parser.Message(parser.Time(strings.Join(meCommand.Args(), " "))))
		}
	case List:
		listCommand.Parse(os.Args[2:])
		if listCommand.Parsed() {
			command.List()
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
