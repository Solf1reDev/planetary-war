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
	flag.StringVar(&TickDurationMinutes, "t", "60", "Tick Duration in Minutes")
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
	tick := time.Tick(time.Duration(tickMinutes) * time.Second)

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
	// get queue of tick actions
	jobqueue := [4]string{
		"moveShip",
		"moveShip",
		"combat",
	}
	//
	for i, jobName := range jobqueue {
		fmt.Printf("2**%d = %d\n", i, jobName)
		executeAction(jobName)
	}

}

func executeAction(action string) {
	switch action {
	case "shipCombat":
		fmt.Println(action)
	case "moveShip":
		fmt.Println(action)
	case "stuff":
		fmt.Println(action)
	default:
		fmt.Println("default fall through")
	}
}

type baseAction interface {
}

type shipUpdateAction struct {
	baseAction baseAction
}

func (*shipUpdateAction) updateShipPositionAction() {

}
