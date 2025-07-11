package server

import (
    "context"
    "io"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)


func h1(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello from server\n")
}

func h2(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, req.PathValue("id"))
}

func RunServer() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    srv := &http.Server{
	Addr: ":8080",
    }
    
    http.HandleFunc("/", h1)
    http.HandleFunc("/{id}/", h2)
    
    go func() {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	    log.Fatalf("Server failed to start: %v", err)
	}
    }()
    
    <-sigChan

    log.Println("Received shutdown signal, initiating graceful shutdown...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(ctx); err != nil {
	log.Fatalf("Server shutdown failed: %v", err)
    }
    
    log.Println("Server gracefully shut down")
}
