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

package simulator

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/fleet"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/rest"
	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

func StartSimulation(dtc rest.DTClient, numFleets int, numTaxis string) error {
	for i := 0; i < numFleets; i++ {
		f := fleet.NewFleet(
			rand.New(rand.NewSource(time.Now().UnixNano())).Intn(899_999)+100_000,
			util.Locations()[i],
			parseNumTaxis(numTaxis),
		)
		f.InitialiseFleet()
		go sendFleetMetrics(dtc, f)
		go sendTaxiMetrics(dtc, f)
	}

	select {}
}

func sendFleetMetrics(dtc rest.DTClient, f fleet.Fleet) {
	var (
		metricData string
		dimensions string
		id         string = fmt.Sprintf("%d", f.GetId())
		loc        string = f.GetLocation()
	)
	for {
		var (
			metrics = map[string]string{
				"fleet.cars.available": fmt.Sprintf("%d", f.GetAvailableCars()),
				"fleet.cars.busy":      fmt.Sprintf("%d", f.GetBusyCars()),
				"fleet.cars.total":     fmt.Sprintf("%d", f.GetTotalCars()),
				"fleet.queue":          fmt.Sprintf("%d", f.GetCustomerQueue()),
			}
		)

		dimensions = "fleetid=" + id + ",location=" + loc
		for mKey, mVal := range metrics {
			metricData += mKey + "," + dimensions + " " + mVal + "\n"
		}

		dtc.PostMetrics(metricData)
		fmt.Println(time.Now().Format("02.01.2006 - 15:04:05"), ": Sent fleet metrics for fleet", id)
		time.Sleep(2 * time.Minute)
	}
}
func sendTaxiMetrics(dtc rest.DTClient, f fleet.Fleet) {
	for {
		for _, t := range f.GetTaxis() {
			var (
				metricData string
				dimensions string
				id         string = fmt.Sprintf("%d", t.GetId())
				class      string = t.GetClass()
				fleetId    string = fmt.Sprintf("%d", t.GetFleetID())
			)

			var (
				metrics = map[string]string{
					"taxi.speed":                 fmt.Sprintf("%f", t.GetSpeed()),
					"taxi.engine.temperature":    fmt.Sprintf("%f", t.GetEngineTemp()),
					"taxi.engine.daystorevision": fmt.Sprintf("%d", t.GetDaysToRevision()),
				}
			)

			dimensions = "taxiid=" + id + ",class=" + class + ",fleetid=" + fleetId
			for mKey, mVal := range metrics {
				metricData += mKey + "," + dimensions + " " + mVal + "\n"
			}

			dtc.PostMetrics(metricData)
			fmt.Println(time.Now().Format("02.01.2006 - 15:04:05"), ": Sent taxi metrics for taxi", id)
		}
		time.Sleep(1 * time.Minute)
	}
}

func parseNumTaxis(nt string) int {
	if strings.Contains(nt, "-") {
		splits := strings.Split(nt, "-")
		min, err := strconv.ParseInt(splits[0], 0, 0)
		if err != nil {
			util.PrintError(err)
		}
		max, err := strconv.ParseInt(splits[1], 0, 0)
		if err != nil {
			util.PrintError(err)
		}
		return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(int(max-min)) + int(min)
	}
	num, err := strconv.ParseInt(nt, 0, 0)
	if err != nil {
		util.PrintError(err)
	}
	return int(num)
}
