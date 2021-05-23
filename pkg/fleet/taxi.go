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
