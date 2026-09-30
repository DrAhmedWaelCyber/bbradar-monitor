<div align="center">
  <h1>🚀 BBRadar Monitor</h1>
  <p><b>An ultra-fast, automated Bug Bounty Reconnaissance Tool to monitor new programs.</b></p>
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
  [![GitHub Actions](https://img.shields.io/badge/Automated-GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions)](https://github.com/features/actions)
  [![Telegram](https://img.shields.io/badge/Alerts-Telegram-2CA5E0?style=for-the-badge&logo=telegram)](https://telegram.org/)
  
  <h3>👨‍💻 Developed by: <b>Ahmed Wael</b></h3>
</div>

<hr>

## 🎯 Overview
**BBRadar Monitor** is a highly efficient Go-based automation tool designed for Bug Bounty Hunters. It actively tracks [bbradar.io](https://bbradar.io) to fetch newly launched bug bounty programs (HackerOne, Bugcrowd, YesWeHack, etc.) and delivers **instant notifications** directly to your Email and Telegram.

By utilizing raw `JSON-LD` extraction instead of fragile HTML scraping, it guarantees robust and blazing-fast data parsing.

## ✨ Key Features
- **⚡ Zero-Delay Parsing:** Extracts structured SEO JSON-LD natively.
- **🤖 Fully Automated:** Pre-configured to run continuously in the background.
- **🧠 Stateful Memory:** Keeps track of previously alerted programs via `state.json` to prevent duplicate alerts.
- **📡 Multi-Channel Alerting:** 
  - 📧 Beautifully formatted HTML Email alerts.
  - 💬 Instant **Telegram Bot** notifications (Supports multiple users).
  - ⏱️ Heartbeat status checks to ensure the monitor is alive.

## ⚙️ Configuration & Setup

To deploy this monitor for your own hunting automation, fork this repository and configure the following in **Settings ➡️ Secrets and variables ➡️ Actions**:

### 🔐 Environment Secrets
| Secret Name | Description | Example |
|---|---|---|
| `SMTP_HOST` | Your email SMTP server | `smtp.gmail.com` |
| `SMTP_PORT` | Your SMTP port | `587` |
| `SMTP_USER` | Your email address | `hunter@gmail.com` |
| `SMTP_PASS` | Your Email App Password | `appblpkzpegobkjwjgg` *(No spaces)* |
| `NOTIFY_EMAIL` | The email receiving the alerts | `your.email@gmail.com` |
| `TELEGRAM_BOT_TOKEN` | Telegram Bot Token | `8693343243:AAF7J...` |
| `TELEGRAM_CHAT_ID` | Primary Telegram Chat ID | `123456789` |
| `TELEGRAM_CHAT_ID_2`| *(Optional)* Second Chat ID | `987654321` |

### 🛡️ Granting Action Permissions
Because the tool updates its internal state memory (`state.json`), you must allow GitHub Actions to commit changes:
1. Go to **Settings** ➡️ **Actions** ➡️ **General**.
2. Scroll to **Workflow permissions**.
3. Select **Read and write permissions**.
4. Click **Save**.

## 🚀 Triggering the Monitor
Once configured, the tool runs on its configured schedule. You can also instantly trigger it by:
1. Clicking the **Actions** tab.
2. Selecting **Monitor BBRadar**.
3. Clicking **Run workflow**.

<hr>

<div align="center">
  <i>Happy Hunting! May the bounties be ever in your favor. 🎯💰</i><br>
  <b>© 2026 Ahmed Wael</b>
</div>
