package main

import (
	"log"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Config Error: %v", err)
	}

	log.Println("Fetching programs from bbradar.io...")
	programs, err := FetchPrograms()
	if err != nil {
		log.Fatalf("Fetch Error: %v", err)
	}
	log.Printf("Fetched %d programs.", len(programs))

	state, err := LoadState()
	if err != nil {
		log.Fatalf("State Load Error: %v", err)
	}

	var newPrograms []Program
	for _, p := range programs {
		if !state[p.Identifier] {
			newPrograms = append(newPrograms, p)
			state[p.Identifier] = true
		}
	}

	if len(newPrograms) > 0 {
		log.Printf("Found %d new programs! Sending notification...", len(newPrograms))
		err = SendNotification(cfg, newPrograms)
		if err != nil {
			log.Fatalf("Notification Error: %v", err)
		}
		log.Println("Notification sent successfully.")

		err = SaveState(state)
		if err != nil {
			log.Fatalf("State Save Error: %v", err)
		}
		log.Println("State updated successfully.")
	} else {
		log.Println("No new programs found. State remains unchanged.")
	}
}
