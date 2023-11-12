package service

import (
	"math/rand"
	"time"

	"github.com/Rod-Way/game/pkg/life"
)

// для хранения состояния
type LifeService struct {
	currentWorld *life.World
	nextWorld    *life.World
}

func New(height, width int) (*LifeService, error) {
	rand.NewSource(time.Now().UTC().UnixNano())

	currentWorld:= life.NewWorld(height, width)
	
	// для упрощения примера хаотично заполним
	currentWorld.RandInit(40)

	newWorld := life.NewWorld(height, width)
	

	ls := LifeService{
		currentWorld: currentWorld,
		nextWorld:    newWorld,
	}

	return &ls, nil
}

// получение очередного состояния игры
func (ls *LifeService) NewState() *life.World {
	life.NextState(ls.currentWorld, ls.nextWorld)

	ls.currentWorld = ls.nextWorld

	return ls.currentWorld
}
