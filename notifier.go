package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

// Developed by: Ahmed Wael

func SendNotification(cfg *Config, newPrograms []Program) error {
	if len(newPrograms) == 0 {
		return nil
	}

	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)

	var body bytes.Buffer
	body.WriteString(fmt.Sprintf("To: %s\r\n", cfg.NotifyEmail))
	body.WriteString("Subject: 🚀 New Bug Bounty Programs Detected!\r\n")
	body.WriteString("MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n")

	body.WriteString(`
		<html>
		<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
			<h2 style="color: #2c3e50;">New Bug Bounty Programs Detected</h2>
			<p>The following new programs matching your filters (wildcard, domain, api) were detected on bbradar.io:</p>
			<ul style="list-style-type: none; padding: 0;">
	`)

	for _, p := range newPrograms {
		parts := strings.Split(p.Identifier, ":")
		platform := "Unknown"
		if len(parts) > 0 {
			platform = parts[0]
		}
		
		item := fmt.Sprintf(`
			<li style="background: #f8f9fa; border: 1px solid #dee2e6; margin-bottom: 15px; padding: 15px; border-radius: 5px;">
				<h3 style="margin-top: 0; color: #007bff;">%s <span style="font-size: 0.8em; color: #6c757d;">(%s)</span></h3>
				<strong>Scope / Link:</strong> <a href="%s" style="color: #0056b3; text-decoration: none;">%s</a><br/>
				<strong>Detected At:</strong> %s<br/>
				<strong>Tags matched:</strong> wildcard, domain, api
			</li>
		`, p.Name, platform, p.URL, p.URL, time.Now().Format(time.RFC1123))
		body.WriteString(item)
	}

	body.WriteString(`
			</ul>
			<p style="font-size: 0.9em; color: #7f8c8d; margin-top: 20px;">Automated via Go BBRadar Monitor | Developer: Ahmed Wael</p>
		</body>
		</html>
	`)

	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)
	err := smtp.SendMail(addr, auth, cfg.SMTPUser, []string{cfg.NotifyEmail}, body.Bytes())
	return err
}

func SendTelegramMessage(cfg *Config, message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.TelegramToken)

	payload := map[string]string{
		"chat_id":    cfg.TelegramChatID,
		"text":       message,
		"parse_mode": "HTML",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("telegram API returned status: %d", resp.StatusCode)
	}

	return nil
}
