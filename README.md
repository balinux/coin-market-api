# Coin Market API

A simple and efficient RESTful API wrapper for CoinGecko, built with Go and the Gin web framework. This project provides easy access to cryptocurrency market data and coin details.

## 🚀 Features

- **Get All Markets**: Retrieve a list of cryptocurrencies with their names, images, and current prices in USD.
- **Get Coin Detail**: Get in-depth information about a specific coin, including description and categories.
- **Standardized Response**: All API responses follow a consistent JSON structure.

## 🛠️ Tech Stack

- **Language**: [Go](https://golang.org/) (Golang)
- **Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Data Source**: [CoinGecko API](https://www.coingecko.com/en/api)

## 📋 Prerequisites

- Go (version 1.20 or later recommended)
- Internet connection (to fetch data from CoinGecko)

## 🔧 Installation & Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/balinux/coin-market-api.git
   cd coin-market-api
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8000`.

## 🛣️ API Endpoints

### 1. Market Overview
Returns a list of top cryptocurrencies.

- **URL**: `/api/v1/markets`
- **Method**: `GET`
- **Response Example**:
  ```json
  {
    "success": true,
    "message": "Successfully get all markets",
    "data": [
      {
        "name": "Bitcoin",
        "image": "https://...",
        "current_price": 65000.50
      }
    ]
  }
  ```

### 2. Coin Detail
Returns detailed information for a specific coin by its ID.

- **URL**: `/api/v1/markets/:id`
- **Method**: `GET`
- **Example**: `/api/v1/markets/bitcoin`
- **Response Example**:
  ```json
  {
    "success": true,
    "message": "Successfully get coin detail",
    "data": {
      "name": "Bitcoin",
      "image": "https://...",
      "current_price": 65000.50,
      "description": "Bitcoin is a decentralized digital currency...",
      "categories": ["Cryptocurrency", "Layer 1"]
    }
  }
  ```

## 🏗️ Project Structure

```text
.
├── common/
│   ├── client/      # HTTP client utility
│   └── response/    # Unified response structure
├── markets/
│   ├── dto.go       # Data Transfer Objects (Request/Response structs)
│   ├── router.go    # Route definitions and handlers
│   └── service.go   # Business logic and external API calls
├── main.go          # Application entry point
└── go.mod           # Go modules file
```

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
