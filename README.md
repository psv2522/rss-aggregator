# RSS Aggregator

A simple RSS aggregator written in Go.

## Features

- Fetches RSS feeds from a list of URLs
- Stores the fetched data in a PostgreSQL database
- Provides a RESTful API for fetching the data
- Uses a Go template for rendering the data

## Setup

1. Clone the repository
```bash
git clone https://github.com/psv252/rss-aggregator.git
cd rss-aggregator
```

2. Install the dependencies
```bash
go install github.com/air-verse/air@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest

```
3. Create a `.env` file similar to `.env.example` and fill in the required environment variables

4. To start the development server run:
```bash
air -c .air.toml
```