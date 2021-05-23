package util

import (
	"fmt"
	"os"
)

var locations = [10]string{
	"Bristol",
	"London",
	"Manchester",
	"Birmingham",
	"Liverpool",
	"Leeds",
	"Newcastle",
	"Nottingham",
	"Basingstoke",
	"Reading",
}

func PrintError(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func Locations() [10]string {
	return locations
}
