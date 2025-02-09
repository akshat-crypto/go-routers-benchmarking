package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"runtime"
)

// MeasureLatency captures the request processing time.
func MeasureLatency(handler http.Handler, req *http.Request) time.Duration {
	w := httptest.NewRecorder()
	start := time.Now()
	handler.ServeHTTP(w, req)
	return time.Since(start)
}

// Benchmark framework for each router
func benchmarkFramework(b *testing.B, name string, handler http.Handler) {
	req, _ := http.NewRequest("GET", "/messages", nil)
	var latencies []time.Duration

	// Memory stats
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		latency := MeasureLatency(handler, req)
		latencies = append(latencies, latency)
	}
	b.StopTimer()

	runtime.ReadMemStats(&memStatsAfter)

	// Memory calculations
	totalAllocs := memStatsAfter.Mallocs - memStatsBefore.Mallocs
	totalBytes := memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc
	allocsPerOp := float64(totalAllocs) / float64(b.N)
	bytesPerOp := float64(totalBytes) / float64(b.N)

	// Latency calculations
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	totalLatency := time.Duration(0)
	for _, l := range latencies {
		totalLatency += l
	}

	avgLatency := totalLatency / time.Duration(len(latencies))
	p95 := latencies[int(float64(len(latencies))*0.95)]
	p99 := latencies[int(float64(len(latencies))*0.99)]

	// Final Output
	fmt.Printf("\n🚀 %s Benchmark Results 🚀\n", name)
	fmt.Printf("Average Latency:           %v\n", avgLatency)
	fmt.Printf("95th Percentile Latency:   %v\n", p95)
	fmt.Printf("99th Percentile Latency:   %v\n", p99)
	fmt.Printf("Memory Consumption (B/op): %.2f bytes\n", bytesPerOp)
	fmt.Printf("Allocations per Operation: %.2f allocs/op\n", allocsPerOp)
	fmt.Println("------------------------------")
}

// Gorilla Mux Benchmark
func BenchmarkMux(b *testing.B) {
	r := mux.NewRouter()
	r.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Fetching messages from Mux"))
	}).Methods("GET")

	benchmarkFramework(b, "Gorilla Mux", r)
}

// Gin Benchmark
func BenchmarkGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/messages", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"messages": "Fetching messages from Gin"})
	})

	benchmarkFramework(b, "Gin", r)
}

// Echo Benchmark
func BenchmarkEcho(b *testing.B) {
	e := echo.New()
	e.GET("/messages", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Fetching messages from Echo"})
	})

	benchmarkFramework(b, "Echo", e)
}

// Chi Benchmark
func BenchmarkChi(b *testing.B) {
	r := chi.NewRouter()
	r.Get("/messages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Fetching messages from Chi"))
	})

	benchmarkFramework(b, "Chi", r)
}
