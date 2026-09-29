package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Developed by: Ahmed Wael
func main() {
	log.Println("Starting BBRadar Monitor - Developed by Ahmed Wael")
	
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
		log.Printf("Found %d new programs! Sending notifications...", len(newPrograms))
		
		// 1. Send Email Notification
		err = SendNotification(cfg, newPrograms)
		if err != nil {
			log.Printf("Email Notification Error: %v", err)
		} else {
			log.Println("Email notification sent successfully.")
		}

		// 2. Formulate & Send Telegram Notification
		var tgMsg strings.Builder
		tgMsg.WriteString(fmt.Sprintf("🚨 <b>New Bug Bounty Programs Detected! (%d)</b>\n\n", len(newPrograms)))
		for _, p := range newPrograms {
			tgMsg.WriteString(fmt.Sprintf("🔹 <b>%s</b>\n🔗 <a href='%s'>Program Link</a>\n\n", p.Name, p.URL))
		}
		tgMsg.WriteString(fmt.Sprintf("🕒 <i>Checked at: %s</i>\n👨‍💻 <i>Dev: Ahmed Wael</i>", time.Now().Format("15:04:05")))
		
		err = SendTelegramMessage(cfg, tgMsg.String())
		if err != nil {
			log.Printf("Telegram Notification Error: %v", err)
		} else {
			log.Println("Telegram notification sent successfully.")
		}

		// 3. Save State
		err = SaveState(state)
		if err != nil {
			log.Fatalf("State Save Error: %v", err)
		}
		log.Println("State updated successfully.")
		
	} else {
		log.Println("No new programs found. Sending heartbeat to Telegram...")
		
		// Send "No new programs" heartbeat every 5 minutes as requested
		heartbeatMsg := fmt.Sprintf("✅ <b>BBRadar Monitor (Alive)</b>\n\n🔍 Checked at: <i>%s</i>\n⚠️ No new programs found in this cycle.\n👨‍💻 <i>Dev: Ahmed Wael</i>", time.Now().Format("15:04:05"))
		err = SendTelegramMessage(cfg, heartbeatMsg)
		if err != nil {
			log.Printf("Telegram Heartbeat Error: %v", err)
		}
	}
}
