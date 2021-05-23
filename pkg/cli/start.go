package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/rest"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/simulator"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

const (
	startCommandHelp = `
	Command starts up the fleet simulation. Format is "start [arguments]"
	
	Arguments:
		--environment, -e	- Dynatrace SaaS or Managed tenant and domain. You don't need https:// or the ending slash
		--token, -t		- Dynatrace API Token with Metrics (V2) permission
		--fleets, -f		- Number of fleets to simulate (max. 10) (default: 2)
		--taxisPerFleet, -tpf	- Number of taxis per fleet to simulate. Ranges supported using the 'min-max' format for more variety (default: 5)
	
	Example:
		start -e abc123.live.dynatrace.com -t abcdefg1234567 -f 3 -tpf 2-5`
)

func startCommand(args []string) {
	var (
		environment string
		token       string
		taxis       string
		fleets      int
	)
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "ERROR: No arguments provided. Type 'help' to see usage.")
		return
	}
	if args[0] == "help" {
		printStartCommandHelp()
		return
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--environment", "-e":
			environment = parseFlagEnvironment(args[i+1])
		case "--token", "-t":
			token = args[i+1]
		case "--fleet", "-f":
			nfleets, err := strconv.Atoi(args[i+1])
			if err != nil {
				util.PrintError(err)
			}
			fleets = parseFlagNumFleets(nfleets)
		case "--taxisPerFleet", "-tpf":
			taxis = args[i+1]
		}
	}
	if fleets == 0 {
		fleets = 2
	}
	if taxis == "" {
		taxis = "5"
	}
	client := rest.NewDTClient(environment, token)
	simulator.StartSimulation(client, int(fleets), taxis)
}

func printStartCommandHelp() {
	fmt.Fprintln(os.Stdout, startCommandHelp)
}
