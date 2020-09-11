package main

import (
	"log"
	"os"
	"os/signal"
	app "ps2eventlistener/internal"
)

func main() {
	censusServiceKey := os.Getenv("CensusKey")
	censusNamespace := os.Getenv("CensusNamespace")
	ps2CensusEndpoint := os.Getenv("PS2CensusEndpoint")
	natsEndpoint := os.Getenv("NatsEndpoint")

	eventManager, err := app.NewEventManager(censusServiceKey, censusNamespace, ps2CensusEndpoint, natsEndpoint)
	if err != nil {
		log.Fatalf("Failed to start application: %s", err)
	}

	eventNames := []string{"Death"}
	characterIDs := []string{"all"}
	worldIDs := []string{"all"}

	eventManager.Connect(eventNames, characterIDs, worldIDs)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

out:
	for {
		select {
		case <-c:
			log.Println("Shutting down...")
			eventManager.Dispose()
			break out
		}
	}
}
