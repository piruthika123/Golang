# Sales Analytics API

## Project Overview

This project is a backend system designed to ingest, normalize, and analyze large-scale sales data from a CSV source. It provides RESTful API endpoints for data refresh and revenue-based insights.

---

## 🧰 Tech Stack

* **Language:** Go (Golang)
* **Database:** PostgreSQL / MySQL / SQLite (choose and update below)
* **Framework:** Gin for REST API
* **ORM:** GORM

---

## 🗃️ Database Schema Overview

![Database Schema](schema.jpg) *(include a diagram of your schema in JPG/PDF format here)*

Main Tables:

* `customers`
* `products`
* `orders`
* `order_items`

---

## 🚀 How to Run the Project

### Prerequisites

* Go 1.20+
* PostgreSQL/MySQL/SQLite

### Step-by-Step Setup

1. **Clone the repository**

```bash
git clone https://github.com/yourusername/sales-analytics.git
cd sales-analytics
```

2. **Configure Database**
   Update `config/config.go` with your DB connection settings.

3. **Run Migrations**

```bash
go run cmd/main.go --migrate
```

4. **Run the Server**

```bash
go run cmd/main.go
```

---

## 🔁 Daily Refresh Job

The job runs every 24 hours using `time.Ticker` in the background. It loads CSV data, overwrites duplicates, and logs status to `logs/refresh.log`.

You can also trigger it on demand using the API.

---

## 📊 API Documentation

| Route                                                      | Method | Body | Response                                     | Description                     |
| ---------------------------------------------------------- | ------ | ---- | -------------------------------------------- | ------------------------------- |
| `/refresh`                                                 | POST   | None | `{ message: "Data refreshed successfully" }` | Triggers CSV data reload        |
| `/analytics/total-revenue?start=YYYY-MM-DD&end=YYYY-MM-DD` | GET    | None | `{ total_revenue: 12345.67 }`                | Total revenue between two dates |

> Add more rows as you build more endpoints.

---

## ⚠️ Error Handling

* API returns meaningful messages and appropriate HTTP codes.
* Background job logs both success and failure.

---

## 🧪 Testing

* Unit and integration tests can be added under `internal/tests/`.
* Mock data CSV included under `data/` for testing purposes.

---

## 📄 License

MIT License.

---

## 👤 Author

Piruthika — [GitHub](https://github.com/yourusername)

---

## 📬 Contact

Feel free to contact me for any issues or enhancements.

---

## ✨ Notes

This project focuses on scalable backend architecture, database normalization, and real-time analytics with proper refresh handling. It's built with clean code principles and RESTful API design.

---
