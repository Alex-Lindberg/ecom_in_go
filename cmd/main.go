package main

import (
	"context"
	"ecom_in_go/api"
	"ecom_in_go/storage"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadDbEnv() string {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		fmt.Println("DB_HOST not set in the .env file")
	}
	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		fmt.Println("DB_USER not set in the .env file")
	}
	dbname, exists := os.LookupEnv("DB_NAME")
	if !exists {
		fmt.Println("DB_NAME not set in the .env file")
	}
	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		fmt.Println("DB_PASSWORD not set in the .env file")
	}
	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s", host, user, dbname, password)
}

func startHttpServer(db *gorm.DB) *http.Server {
	log.Println("Starting HTTP server initialization")
	router := mux.NewRouter()

	store := &storage.PGStore{DB: db}
	productHandler := &api.ProductHandler{ProductStore: store}
	customerHandler := &api.CustomerHandler{CustomerStore: store}
	orderHandler := &api.OrderHandler{
		OrderStore:    store,
		CustomerStore: store,
		ProductStore:  store,
		VariantStore:  store,
	}

	router.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", productHandler.GetProduct).Methods("GET")
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{id:[0-9]+}", customerHandler.GetCustomer).Methods("GET")
	router.HandleFunc("/orders", orderHandler.GetOrders).Methods("GET")
	router.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id:[0-9]+}", orderHandler.GetOrder).Methods("GET")

	http.Handle("/", router)
	log.Println("Creating HTTP server instance")
	server := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("Starting server on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	return server
}

func main() {
	dsn := loadDbEnv()
	if dsn == "" {
		log.Println("Error loading .env file")
		return
	}

	fmt.Println("Loaded .env file successfully")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error starting server: %v\n", err)
		return
	}

	server := startHttpServer(db)

	// Channel to listen for OS signals
	// This channel will block the main goroutine until it receives a signal
	// When it receives a signal it gracefully shuts down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	} else {
		log.Println("Server stopped gracefully.")
	}
}
