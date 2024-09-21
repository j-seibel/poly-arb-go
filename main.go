package main

import (
	"os"
	"poly/arb/client"
)

func main() {

	//command line arguemnt to check for market refresh
	if len(os.Args) > 1 {
		if os.Args[1] == "r" {
			println("Refreshing markets")
			client.FindNegRiskMarkets()
			return
		}
	}

	client.StartSubscription()
}
