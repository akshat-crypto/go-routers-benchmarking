# Go Router Benchmark 🚀

## 📋 **Requirements**
- **Go (Golang)** version 1.18 or higher
- **Git** for version control (optional)
- Supported Routers:
  - Gorilla Mux
  - Gin
  - Echo
  - Chi

## ⚡ **Setup Instructions**

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/go-router-benchmark.git
   cd go-router-benchmark
   ```

2. **Start the Servers:**
   ```bash
   go run main.go
   ```
   This will start the servers on the following ports:
   - **Gorilla Mux:** http://localhost:8081/messages
   - **Gin:** http://localhost:8082/messages
   - **Echo:** http://localhost:8083/messages
   - **Chi:** http://localhost:8084/messages

3. **Run the Benchmarks:**
   ```bash
   go test -bench=. -benchmem -count=1
   ```
   This command will execute benchmarks for all routers with memory profiling.

## 📊 **Benchmark Results**

### 🚀 **Gorilla Mux Benchmark Results:**
- **Average Latency:** 2.117µs
- **95th Percentile Latency:** 3.982µs
- **99th Percentile Latency:** 8.068µs
- **Memory Consumption (B/op):** 719.70 bytes
- **Allocations per Operation:** 13 allocs/op

### 🚀 **Gin Benchmark Results:**
- **Average Latency:** 3.41µs
- **95th Percentile Latency:** 5.614µs
- **99th Percentile Latency:** 12.639µs
- **Memory Consumption (B/op):** 906.13 bytes
- **Allocations per Operation:** 15 allocs/op

### 🚀 **Echo Benchmark Results:**
- **Average Latency:** 3.5µs
- **95th Percentile Latency:** 5.786µs
- **99th Percentile Latency:** 12.847µs
- **Memory Consumption (B/op):** 895.13 bytes
- **Allocations per Operation:** 15 allocs/op

### 🚀 **Chi Benchmark Results:**
- **Average Latency:** 1.139µs
- **95th Percentile Latency:** 2.256µs
- **99th Percentile Latency:** 4.423µs
- **Memory Consumption (B/op):** 457.61 bytes
- **Allocations per Operation:** 8 allocs/op

## 🙌 **Contributing**
Feel free to fork the repository, create feature branches, and submit pull requests to improve benchmarks or add new routers.
