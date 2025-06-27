# 🎰 go-jackpot
Jackpot Probability Calculator

A small HTTP server in Go that exposes a single endpoint to simulate a jackpot draw. Each request represents a bet that may win or not, based on a probabilistic logic.

---

## ⚙️ Requirements

- Go 1.20 or higher (latest stable recommended)

---

## 📦 Installation

Make sure Go is installed:

```bash
go version
```

---

## ▶️ Running the server

Start the server with:

```bash
go run .
```


The server will start on:
```
http://localhost:8080/jackpot-draw
```

---

## 🧪 Running tests

Run all tests in the project with:

```bash
go test ./... -v
```

This will execute unit tests for all modules: logic, service, server and storage.

---

## 📡 How to use the endpoint

### POST ```/jackpot-draw```

Simulates a jackpot bet.

* **Request body (JSON):**
```json
{
  "bet": 10
}
```
* **Response (JSON)**
```json
{
  "is_won": true
}
```

* **Example using ```curl```:**
```bash
curl -X POST http://localhost:8080/jackpot-draw \
     -H "Content-Type: application/json" \
     -d '{"bet": 10}'
```
Alternatively, you can use
```bash
python3 scripts/test_connection.py
```

---

## 💾 Persistence layer
All bets are logged to a local file named ```jackpot-log.jsonl``` in the project root. Each line is a JSON object representing a single bet result:
```json
{"timestamp":"2025-06-27T13:00:51+02:00","bet":1,"is_won":false}
```
No external database is required. The file is automatically created if it does not exist.

---

## 🗂️ Project structure

```
.
├── main.go                  # Entry point
├── go.mod                   # Go modules
├── internal/
│   ├── logic/               # Jackpot logic
│   ├── service/             # Application service layer
│   ├── storage/             # Persistence (JSONL)
│   └── server/              # HTTP interface
├── config/                  # Config file
├── scripts/                 # Test script (optional)
└── README.md
```