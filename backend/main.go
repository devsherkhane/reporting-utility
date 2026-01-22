package main

import (
	"log"
	"net/http"
	"reporting-utility/internal/db"
	"reporting-utility/internal/route"
	"reporting-utility/internal/utils/config"

	"github.com/gorilla/mux"
)

func main() {

	
    // 1. Load YAML Config FIRST
    if err := config.LoadConfig(); err != nil {
        log.Fatal("Failed to load config:", err)
    }

    db.ConnectDB() // Uses DSN from config 
    config.InitSession() // Uses key from config [cite: 4]

    router := mux.NewRouter()
    route.SetupRoutes(router)

    // Use the port from config.yaml
    port := config.AppConfig.Server.Port 
    log.Printf("Server starting on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, enableCORS(router)))
}

func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Use the origin from config.yaml [cite: 5]
        w.Header().Set("Access-Control-Allow-Origin", config.AppConfig.Server.AllowedOrigin)
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

        if r.Method == "OPTIONS" {
            return
        }
        next.ServeHTTP(w, r)
    })
}