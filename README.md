# Aggregator

This is a simple RSS Feed Aggregator built in Go. It collects, stores, and allows users to follow RSS feeds. The system handles feeds, feed follows, posts, and users.

## Features

- Feeds: Collect and manage RSS feeds.
- Feed Follows: Users can follow specific feeds.
- Posts: Store and fetch posts from followed feeds.
- Users: Basic user management (creation, API key).

## Installation

1. Clone the repository.
2. Connect to a PostgreSQL database.(You have set two environment variables: PORT and DB_URL)
3. Run the following commands:
```bash
go mod vendor
go build -o aggregator ./cmd/aggregator && ./aggregator
```

## API

Please see routes.go for the available API endpoints.

## License

This project is licensed under the MIT License.