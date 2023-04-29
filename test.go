package main

import (
	"fmt"

	"github.com/Craftserve/mcstatus"
)

func main() {
	address := "erwan.fun"
	port := 25577
	status, err := mcstatus.QueryServer(address, port)
	if err != nil {
		fmt.Println(err)
	}
	numPlayers := status.Players.Online
	fmt.Printf("Le serveur a actuellement %d joueurs en ligne", numPlayers)
}
