package main

import (
	"fmt"
	"os"
)

type Config struct {
	SMTPHost    string
	SMTPPort    string
	SMTPUser    string
	SMTPPass    string
	NotifyEmail string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),
		SMTPUser:    os.Getenv("SMTP_USER"),
		SMTPPass:    os.Getenv("SMTP_PASS"),
		NotifyEmail: os.Getenv("NOTIFY_EMAIL"),
	}

	if cfg.SMTPHost == "" || cfg.SMTPPort == "" || cfg.SMTPUser == "" || cfg.SMTPPass == "" || cfg.NotifyEmail == "" {
		return nil, fmt.Errorf("missing one or more required environment variables (SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASS, NOTIFY_EMAIL)")
	}

	return cfg, nil
}
