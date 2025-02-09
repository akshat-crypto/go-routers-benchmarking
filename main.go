package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

func main() {
	// Start all servers concurrently
	go startMuxServer()
	go startGinServer()
	go startEchoServer()
	go startChiServer()

	log.Println("Servers are running on different ports...")
	select {} // Block forever
}

// Gorilla Mux Server
func startMuxServer() {
	r := mux.NewRouter()
	r.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Fetching messages from Mux"))
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}

// Gin Server
func startGinServer() {
	r := gin.Default()
	r.GET("/messages", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"messages": "Fetching messages from Gin"})
	})
	log.Fatal(r.Run(":8082"))
}

// Echo Server
func startEchoServer() {
	e := echo.New()
	e.GET("/messages", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Fetching messages from Echo"})
	})
	log.Fatal(e.Start(":8083"))
}

// Chi Server
func startChiServer() {
	r := chi.NewRouter()
	r.Get("/messages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Fetching messages from Chi"))
	})
	log.Fatal(http.ListenAndServe(":8084", r))
}

