package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

// Variables used for command line parameters
var (
	TickDurationMinutes string
	RoundNo             string
	TickCount           int
	// ShipUpdateQueue     []shipUpdateAction
)

func init() {

	flag.StringVar(&RoundNo, "r", "", "Round Number")
	flag.StringVar(&TickDurationMinutes, "t", "", "Tick Duration in Minutes")
	flag.Parse()
	TickCount = 0
}

func main() {
	i, _ := strconv.Atoi(TickDurationMinutes)
	gameLoop(i)
}

func gameLoop(tickMinutes int) {

	fmt.Println(time.Now())
	// Set tick timer
	tick := time.Tick(time.Duration(tickMinutes) * time.Minute)

	for {
		select {
		case <-tick:
			TickCount++
			fmt.Print("Tick ")
			fmt.Print(TickCount, " - ")
			fmt.Println(time.Now())
		}
	}
}

func actionProcessor() {

}

type baseAction interface {
}

type shipUpdateAction struct {
	baseAction baseAction
}

func (*shipUpdateAction) updateShipPositionAction() {

}
