# AI Log Monitor

A distributed, fault-tolerant log monitoring system built with Go, Python, Redis, and Docker. It simulates log ingestion, detects anomalies, and supports CI/CD and containerized deployment.

---

## 🧠 Overview

This project demonstrates how to build and connect microservices across different languages to monitor system logs in real-time. Logs are pushed to a Redis queue, analyzed by a Python-based anomaly detection service, and categorized as normal or anomalous.

---

## 🧱 Architecture

```
┌────────────┐ ┌────────────┐ ┌──────────────┐
│ Collector │──────▶│ Redis │──────▶│ Worker │
└────────────┘ └────────────┘ └──────┬───────┘
                                     │
                            ┌────────▼────────┐
                            │ Anomaly Detector│
                            │ (FastAPI, Py) │
                            └─────────────────┘
```


---

## 🛠 Tech Stack

- **Go** – Collector and Worker services
- **Python (FastAPI)** – Anomaly detection API
- **Redis** – Queue for log message buffering
- **Docker + Docker Compose** – Containerization
- **GitHub Actions** – CI pipeline for Go services

---

## 🚀 How to Run It (with Docker)

### 1. Clone the repo

```bash
git clone https://github.com/ckang21/ai-log-monitor.git
cd ai-log-monitor
```

### 2. Build and start all services

```
docker-compose up --build
```

### 3. Watch output in terminal

- The collector pushes logs to Redis
- The worker dequeues logs and sends them to the anomaly detector
- If the log contains "level": "error", it is flagged as an anomaly

### Sample Output
```
🔁 Starting log worker...
✅ Normal log: Server started
🚨 Anomaly Detected: Database timeout
✅ Normal log: Health check passed
```
---
### Project Highlights
- Fault-tolerant log queue with retry logic
- REST API communication across services
- Distributed service orchestration via Docker Compose
- GitHub Actions CI pipeline for Go services
- Easily extendable with alerts, dashboards, or real ML models

### Author

**ckang21**
[github.com/ckang21](https://github.com/ckang21)
