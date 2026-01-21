# JRI-BOT-2

A Telegram bot that helps you decide what to eat using weighted random selection from curated food presets.

## Features

- **Weighted Random Selection** - Food items have different weights based on recommended frequency (daily, weekly, monthly)
- **Multiple Food Presets** - Choose from three curated food packs:
  - **Based Pack** - Diverse default food mix
  - **Thai Pack** - Thai cuisine focused
  - **Georgian Pack** - Georgian cuisine focused
- **Per-User Preferences** - The bot remembers each user's selected food pack
- **Simple Interface** - One button to get a food suggestion

## Tech Stack

- **Go 1.23** - Core language
- **telebot v4** - Telegram Bot API wrapper
- **weightedrand** - Weighted random selection
- **Docker** - Containerized deployment
- **GitHub Actions** - CI/CD pipeline

## Project Structure

```
jri-bot-2/
├── src/
│   ├── main.go              # Bot entry point and handlers
│   ├── jri/
│   │   ├── food.go          # Food selection logic
│   │   ├── presets.go       # Food preset definitions
│   │   └── storage.go       # User preference storage
│   └── frequency/
│       └── frequency.go     # Frequency weight constants
├── Dockerfile
├── docker-compose.yml
└── .github/workflows/
    └── docker-image.yml     # CI/CD pipeline
```

## Configuration

| Variable | Description | Required |
|----------|-------------|----------|
| `TOKEN`  | Telegram Bot API token | Yes |

Get your bot token from [@BotFather](https://t.me/BotFather) on Telegram.

## Running Locally

```bash
cd src
go build -o bot
TOKEN=your_telegram_token ./bot
```

## Docker

### Using Docker Compose

```bash
# Set token in environment or .env file
export TOKEN=your_telegram_token
docker-compose up -d
```

### Using Docker directly

```bash
docker build -t jri-bot .
docker run -e TOKEN=your_telegram_token jri-bot
```

## Bot Commands

| Command | Description |
|---------|-------------|
| `/start` | Initialize the bot and show the food button |
| `/packs` | Select a food preset |
| `Че сожрать` (button) | Get a random food suggestion |

## How It Works

The bot uses a weighted random algorithm where each food item has a weight based on how often it should be suggested:

- Items meant to be eaten daily have higher weights
- Items meant for weekly consumption have medium weights
- Monthly treats have lower weights

This creates a natural distribution that suggests variety while respecting meal frequency preferences.

## CI/CD

The GitHub Actions workflow automatically builds and pushes Docker images to Docker Hub on every push to `main`. Required secrets:

- `DOCKER_USERNAME`
- `DOCKER_PASSWORD`

## License

This project is provided as-is for personal use.
