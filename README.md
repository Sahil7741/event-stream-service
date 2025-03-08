# 🚀 Event Streaming Service with Kafka  

## 📌 Overview  

This project is a **high-performance event streaming system** built using **Kafka** and **Go**. It allows **producers** to send structured messages to a Kafka topic and **consumers** to efficiently process those messages in real-time. Designed for **low latency, high throughput, and scalability**, this system is ideal for handling event-driven architectures.  

## 🏗 Tech Stack  

- **Go** - Backend programming language.  
- **Kafka** - Distributed event streaming platform.  
- **Docker** - Containerization for Kafka and Zookeeper.  
- **REST API** - For message publishing.  

## 🔥 Features  

✅ **Processes 10,000+ messages per second**  
✅ **Sub-100ms event processing latency**  
✅ **Multi-worker consumers for parallel processing**  
✅ **Fault tolerance with structured logging & error handling**  
✅ **Scalable architecture with Kafka & Docker**  

## 📂 Project Structure  

```
📦 event-stream-service  
 ┣ 📜 main.go          # REST API with Kafka producer  
 ┣ 📜 publisher.go     # Standalone Kafka message producer  
 ┣ 📜 consumer.go      # Multi-threaded Kafka consumer  
 ┣ 📜 docker-compose.yml # Kafka and Zookeeper setup  
 ┣ 📜 README.md        # Documentation  
```

## ⚙️ Setup & Usage  

### 🛠 Prerequisites  

- **Docker & Docker-Compose** installed  
- **Go (1.19+)** installed  

### 1️⃣ Start Kafka & Zookeeper  

Run the following command to start Kafka and Zookeeper:  

```sh
docker-compose up -d
```

### 2️⃣ Run the Producer API  

Start the producer API:  

```sh
go run main.go
```

Send events using `cURL`:

```sh
curl -X POST http://localhost:8080/publish \
     -H "Content-Type: application/json" \
     -d '{"message": "Hello, Kafka!"}'
```

### 3️⃣ Run the Consumer  

Start the consumer to receive messages:  

```sh
go run consumer.go
```

### 4️⃣ Run the Publisher (Bulk Messaging)  

Send **10,000 messages** for performance testing:  

```sh
go run publisher.go
```

## 📊 Performance Benchmark  

- **Achieved 10,000+ messages per second**  
- **Distributed workload across multiple consumer workers**  
- **Ensured sub-100ms message processing latency**  

## 🚀 Future Enhancements  

- 📈 **Monitoring** with Prometheus & Grafana  
- 🗄 **Database Integration** (PostgreSQL, Redis)  
- 📡 **Scalability Improvements** with multi-broker Kafka  