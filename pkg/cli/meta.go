package cli

const (
	version   = "v0.0.1"
	title     = "Interactive Fleet Simulator"
	shortHelp = "This is an interactive app. Type 'help' for usage information and examples."
	longHelp  = `
	Interactive Fleet Simulator
	---------------------------
	This tool is used to simulate generic entities, their relationships, and metrics. The scenario simulates
	a fleet of smart taxies which generate metrics that are then pushed out to a Dynatrace tenant via API.

	Commands:
		start - starts the simulation, with given parameters
		stop  - stops any running simulation
		help  - prints this help message
		exit  - stops any running simulation and exits the app
	
	Additional 'help' is available for each command.
	`
	inputStr = "> "
)
