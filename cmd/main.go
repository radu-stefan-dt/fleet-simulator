package main

import (
	"github.com/radu-stefan-dt/fleet-simulator/pkg/cli"
)

func main() {
	cli.NewCli()
}

// func buildCli() *cli.App {
// 	app := cli.NewApp()

// 	app.Usage = "Simulates fleets of taxies operating and pushing metrics to Dynatrace via the API."
// 	app.Version = version

// 	cli.VersionPrinter = func(c *cli.Context) {
// 		fmt.Println(c.App.Version)
// 	}
// 	cli.VersionFlag = &cli.BoolFlag{
// 		Name:  "version",
// 		Usage: "print the version and exit",
// 	}

// 	app.Description = `
// 	Tool used to simulate generic entities, their relationships, and metrics. The scenario simulates a
// 	fleet of smart taxies which generate metrics that are then pushed out to a Dynatrace tenant via API.
// 	`
// 	startCommand := getStartCommand()
// 	app.Commands = []*cli.Command{&startCommand}

// 	return app
// }

// func getStartCommand() cli.Command {
// 	command := cli.Command{
// 		Name:      "start",
// 		Usage:     "starts the simulation",
// 		UsageText: "start [command options]",
// 		Before: func(c *cli.Context) error {
// 			fmt.Println("Fleet Simulator v" + version)
// 			return nil
// 		},
// 		Flags: []cli.Flag{
// 			&cli.StringFlag{
// 				Name:     "environment",
// 				Aliases:  []string{"e"},
// 				Usage:    "Dynatrace SaaS or Managed tenant and domain. You don't need https:// or the ending slash.",
// 				Required: true,
// 			},
// 			&cli.StringFlag{
// 				Name:     "token",
// 				Aliases:  []string{"t"},
// 				Usage:    "Dynatrace API Token with Metrics (V2) permission",
// 				Required: true,
// 			},
// 			&cli.IntFlag{
// 				Name:    "fleets",
// 				Aliases: []string{"f"},
// 				Usage:   "Number of fleets to simulate (max. 10). Will impact total number of taxis and therefore metrics ingested.",
// 				Value:   2,
// 			},
// 			&cli.StringFlag{
// 				Name:    "taxisPerFleet",
// 				Aliases: []string{"tpf"},
// 				Usage:   "Number of taxis per fleet to simulate. Can give a range using the 'min-max' format for more variety.",
// 				Value:   "5",
// 			},
// 		},
// 		Action: func(ctx *cli.Context) error {
// 			if ctx.Args().Present() {
// 				fmt.Printf("Error:\n\tFound argument: %s\n\tThis command uses only flags, no arguments are needed.\n", ctx.Args().First())
// 				cli.ShowAppHelpAndExit(ctx, 1)
// 			}

// 			env := parseFlagEnvironment(ctx.String("environment"))
// 			fleets := parseFlagNumFleets(ctx.Int("fleets"))
// 			client := newDTClient(env, ctx.String("token"))

// 			return startSimulation(
// 				client,
// 				fleets,
// 				ctx.String("taxisPerFleet"),
// 			)
// 		},
// 	}
// 	return command
// }
