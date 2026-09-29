package main

import (
	"fmt"
	"os"
)

// Developed by: Ahmed Wael
type Config struct {
	SMTPHost         string
	SMTPPort         string
	SMTPUser         string
	SMTPPass         string
	NotifyEmail      string
	TelegramToken    string
	TelegramChatID   string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		SMTPHost:         os.Getenv("SMTP_HOST"),
		SMTPPort:         os.Getenv("SMTP_PORT"),
		SMTPUser:         os.Getenv("SMTP_USER"),
		SMTPPass:         os.Getenv("SMTP_PASS"),
		NotifyEmail:      os.Getenv("NOTIFY_EMAIL"),
		TelegramToken:    os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChatID:   os.Getenv("TELEGRAM_CHAT_ID"),
	}

	if cfg.SMTPHost == "" || cfg.SMTPPort == "" || cfg.SMTPUser == "" || cfg.SMTPPass == "" || cfg.NotifyEmail == "" {
		return nil, fmt.Errorf("missing one or more required email environment variables")
	}
	
	if cfg.TelegramToken == "" || cfg.TelegramChatID == "" {
		return nil, fmt.Errorf("missing telegram environment variables (TELEGRAM_BOT_TOKEN, TELEGRAM_CHAT_ID)")
	}

	return cfg, nil
}
