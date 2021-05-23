/**
 * Copyright (c) 2021 Radu Stefan
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 **/

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
