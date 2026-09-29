# 🚀 BBRadar Monitor

**BBRadar Monitor** is a highly-efficient, automated Go tool that actively monitors [bbradar.io](https://bbradar.io/) for newly published bug bounty programs. When new programs are detected, it instantly sends an HTML Email and a Telegram notification directly to your device!

> 👨‍💻 **Developed by:** Ahmed Wael

## ✨ Features
- **Scrapes Intelligently:** Parses native SEO JSON-LD data instead of brittle HTML structure to ensure stability.
- **GitHub Actions Automation:** Runs automatically every 5 minutes completely for free.
- **Stateful Memory:** Remembers programs it has already alerted you about using a `state.json` file.
- **Multi-Channel Alerts:** 
  - 📧 Clean, formatted HTML Emails.
  - 💬 Telegram Bot Notifications.
  - ⏱️ Heartbeat notifications every 5 minutes via Telegram indicating that the monitor is active and hunting.

## ⚙️ Configuration & Setup

1. Fork or Clone this repository as a Private repository.
2. Go to your repository **Settings** ➡️ **Secrets and variables** ➡️ **Actions**.
3. Add the following **New repository secrets**:

| Secret Name | Description |
|---|---|
| `SMTP_HOST` | E.g., `smtp.gmail.com` |
| `SMTP_PORT` | E.g., `587` |
| `SMTP_USER` | Your email address |
| `SMTP_PASS` | Your Email App Password (no spaces) |
| `NOTIFY_EMAIL` | The email address to receive alerts |
| `TELEGRAM_BOT_TOKEN` | Your Telegram Bot Token from BotFather |
| `TELEGRAM_CHAT_ID` | Your personal Telegram Chat ID |
| `TELEGRAM_CHAT_ID_2` | (Optional) A second Telegram Chat ID to receive the same alerts |

### 🔑 Important: Grant Write Permissions
To prevent the script from sending you the same programs repeatedly, it saves its state to `state.json` and pushes it back to GitHub.
1. Go to repository **Settings** ➡️ **Actions** ➡️ **General**.
2. Scroll to **Workflow permissions**.
3. Select **Read and write permissions**.
4. Click **Save**.

## 🚀 Running the Monitor
The monitor will start running automatically every 5 minutes. You can also trigger it manually:
1. Go to the **Actions** tab.
2. Select **Monitor BBRadar** on the left.
3. Click **Run workflow**.

---
*Happy Hunting! 🎯*
