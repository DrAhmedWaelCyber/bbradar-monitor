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
		log.Printf("Fetch Error (Temporary): %v", err)
		return // Exit cleanly so GitHub doesn't penalize the cron schedule
	}
	log.Printf("Fetched %d programs.", len(programs))

	state, err := LoadState()
	if err != nil {
		log.Printf("State Load Error: %v", err)
		return
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
		tgMsg.WriteString(fmt.Sprintf("🚨 <b>NEW BUG BOUNTY PROGRAMS DETECTED!</b> 🚨\n━━━━━━━━━━━━━━━━━━━━\n🔥 <b>Programs Found:</b> %d\n\n", len(newPrograms)))
		for _, p := range newPrograms {
			parts := strings.Split(p.Identifier, ":")
			platform := "Unknown"
			if len(parts) > 0 {
				platform = parts[0]
			}
			tgMsg.WriteString(fmt.Sprintf("🎯 <b>%s</b> <i>(%s)</i>\n🔗 <a href=\"%s\">View Scope & Details</a>\n\n", p.Name, platform, p.URL))
		}
		tgMsg.WriteString(fmt.Sprintf("━━━━━━━━━━━━━━━━━━━━\n⏱ <b>Time:</b> <i>%s</i>\n👨‍💻 <b>Developed by: Ahmed Wael</b>", time.Now().Format("2006-01-02 15:04:05")))
		
		BroadcastTelegramMessage(cfg, tgMsg.String())
		log.Println("Telegram notifications broadcasted.")

		// 3. Save State
		err = SaveState(state)
		if err != nil {
			log.Printf("State Save Error (Non-fatal): %v", err)
		} else {
			log.Println("State updated successfully.")
		}
		
	} else {
		log.Println("No new programs found. Sending heartbeat to Telegram...")
		
		// Send "No new programs" heartbeat every 5 minutes as requested
		heartbeatMsg := fmt.Sprintf("🟢 <b>BBRadar Monitor Status</b> 🟢\n━━━━━━━━━━━━━━━━━━━━\n📡 <b>Status:</b> Active & Hunting\n🔍 <b>Last Checked:</b> %s\n⚠️ <b>Result:</b> No new programs in this cycle.\n\n<i>Stay ready! 🎯</i>\n━━━━━━━━━━━━━━━━━━━━\n👨‍💻 <b>Developed by: Ahmed Wael</b>", time.Now().Format("2006-01-02 15:04:05"))
		BroadcastTelegramMessage(cfg, heartbeatMsg)
	}
}
