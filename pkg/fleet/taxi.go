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

type Taxi interface {
	GetId() int
	GetClass() string
	GetFleetID() int
	GetSpeed() float64
	GetEngineTemp() float64
	GetDaysToRevision() int
}
type taxiImpl struct {
	id             int
	class          string
	fleetID        int
	speed          float64
	engineTemp     float64
	daysToRevision int
}

func (t taxiImpl) GetId() int {
	return t.id
}
func (t taxiImpl) GetClass() string {
	return t.class
}
func (t taxiImpl) GetFleetID() int {
	return t.fleetID
}
func (t taxiImpl) GetSpeed() float64 {
	d := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(60)
	return d + float64(i)
}
func (t taxiImpl) GetEngineTemp() float64 {
	d := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(20) + 90
	return d + float64(i)
}
func (t taxiImpl) GetDaysToRevision() int {
	return t.daysToRevision
}

func NewTaxi(id int, class string, fleet int) Taxi {
	return &taxiImpl{
		id:             id,
		class:          class,
		fleetID:        fleet,
		speed:          0,
		engineTemp:     90,
		daysToRevision: 365,
	}
}
