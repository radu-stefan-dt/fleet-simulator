package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

func NewCli() {
	fmt.Println("Welcome to " + title + " " + version)
	fmt.Println(shortHelp)
	for {
		fmt.Print(inputStr)
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			util.PrintError(err)
		}
		cmdFields := strings.Fields(strings.TrimSuffix(cmdString, "\n"))
		cmd := cmdFields[0]
		args := cmdFields[1:]

		switch cmd {
		case "help":
			help()
		case "start":
			startCommand(args)
		case "stop":
			stopCommand()
		case "exit":
			exit()
		}
	}
}

func help() {
	fmt.Fprintln(os.Stdout, longHelp)
}

func stopCommand() {
	fmt.Fprintln(os.Stdout, "You ran the STOP command")
}

func exit() {
	os.Exit(0)
}
