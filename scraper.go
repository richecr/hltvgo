package main

import (
	"fmt"
	"time"

	"github.com/richecr/hltv-go/lib/operations"
)

func main() {
	start := time.Now()
	matches, _ := operations.GetMatches()
	// operations.GetMatches()
	duration := time.Since(start)
	fmt.Println("Time duration:", duration.Milliseconds())
	fmt.Println(len(matches))
	for _, match := range matches {
		fmt.Println(match.Id)
		fmt.Println(match.Team1.Id)
		fmt.Println(match.Team1.Name)
		fmt.Println(match.Team2.Id)
		fmt.Println(match.Team2.Name)
		fmt.Println(match.Event)
		fmt.Println(match.DateHour)
		fmt.Println(match.Live)
		fmt.Println("----------")
	}
}
