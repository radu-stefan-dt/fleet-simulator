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

package fleet

import (
	"math/rand"
	"time"
)

type Fleet interface {
	GetId() int
	GetLocation() string
	GetAvailableCars() int
	GetBusyCars() int
	GetTotalCars() int
	GetCustomerQueue() int
	GetTaxis() []Taxi
	UpdateCars(int, int, int)
	MakeCarBusy()
	UpdateQueue(int)
	RegisterTaxi(Taxi)
	InitialiseFleet()
}

type fleetImpl struct {
	id            int
	location      string
	carsAvailable int
	carsBusy      int
	carsTotal     int
	customerQueue int
	taxis         []Taxi
}

func (f fleetImpl) GetId() int {
	return f.id
}
func (f fleetImpl) GetLocation() string {
	return f.location
}
func (f fleetImpl) GetTotalCars() int {
	return f.carsTotal
}
func (f fleetImpl) GetAvailableCars() int {
	return f.carsAvailable
}
func (f fleetImpl) GetBusyCars() int {
	return f.carsBusy
}
func (f fleetImpl) GetCustomerQueue() int {
	return f.customerQueue
}
func (f fleetImpl) GetTaxis() []Taxi {
	return f.taxis
}

func (f *fleetImpl) MakeCarBusy() {
	if f.carsAvailable-1 >= 0 {
		f.carsAvailable--
		f.carsBusy++
	}
}
func (f *fleetImpl) UpdateCars(available, busy, total int) {
	f.carsAvailable = available
	f.carsBusy = busy
	f.carsTotal = total
}
func (f *fleetImpl) UpdateQueue(q int) {
	f.customerQueue = q
}
func (f *fleetImpl) RegisterTaxi(t Taxi) {
	f.taxis = append(f.taxis, t)
}
func (f *fleetImpl) InitialiseFleet() {
	for i := 0; i < f.GetTotalCars(); i++ {
		var class string
		switch {
		case i%3 == 0:
			class = "limo"
		case i%3 == 1:
			class = "exec"
		default:
			class = "casual"
		}
		tID := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(89_999_999) + 10_000_000
		t := NewTaxi(tID, class, f.GetId())
		f.RegisterTaxi(t)
		time.Sleep(time.Nanosecond) // ensures next random seed is different
	}
}

func NewFleet(id int, location string, carsTotal int) Fleet {
	return &fleetImpl{
		id:            id,
		location:      location,
		carsAvailable: carsTotal,
		carsBusy:      0,
		carsTotal:     carsTotal,
		customerQueue: 0,
		taxis:         []Taxi{},
	}
}