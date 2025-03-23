package server

import (
    "context"
    "net/http"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
    "primix/logger"
    "primix/utils"
)

// Server represents the HTTP server configuration
type Server struct {
    mux         *http.ServeMux
    middlewares []Middleware
    server      *http.Server
    handlers    map[string]http.HandlerFunc
    wg          sync.WaitGroup
}

// Middleware type for handler wrapping
type Middleware func(http.HandlerFunc) http.HandlerFunc

// NewServer creates a new Primix server instance
func NewServer() *Server {
    return &Server{
        mux:      http.NewServeMux(),
        handlers: make(map[string]http.HandlerFunc),
        middlewares: []Middleware{
            loggingMiddleware, // Default middleware for request logging
        },
    }
}

// Use adds middleware to the server
func (s *Server) Use(m Middleware) {
    s.middlewares = append(s.middlewares, m)
}

// Handle registers a new route with handler
func (s *Server) Handle(method, pattern string, handler http.HandlerFunc) {
    wrappedHandler := handler
    for _, mw := range s.middlewares {
        wrappedHandler = mw(wrappedHandler)
    }
    s.handlers[method+" "+pattern] = wrappedHandler
    s.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
        wrappedHandler(w, r)
    })
}

// Get is a convenience method for GET requests
func (s *Server) Get(pattern string, handler http.HandlerFunc) {
    s.Handle(http.MethodGet, pattern, handler)
}

// Post is a convenience method for POST requests
func (s *Server) Post(pattern string, handler http.HandlerFunc) {
    s.Handle(http.MethodPost, pattern, handler)
}

// loggingMiddleware logs each request
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        utils.Info("Request: " + r.Method + " " + r.URL.Path)
        next(w, r)
        logger.Logger.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
    }
}

// Start runs the server with graceful shutdown
func (s *Server) Start(addr string) error {
    s.server = &http.Server{
        Addr:         addr,
        Handler:      s.mux,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    // Channel to listen for OS signals
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    // Run server in a goroutine
    s.wg.Add(1)
    go func() {
        defer s.wg.Done()
        utils.Info("Starting server on " + addr)
        if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Logger.Printf("Server error: %v", err)
        }
    }()

    // Wait for shutdown signal
    <-stop
    utils.Warn("Shutting down server...")

    // Graceful shutdown with 5-second timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := s.server.Shutdown(ctx); err != nil {
        logger.Logger.Printf("Shutdown error: %v", err)
        return err
    }

    s.wg.Wait()
    utils.Success("Server stopped gracefully")
    return nil
}

// Static serves static files from a directory
func (s *Server) Static(prefix, dir string) {
    s.mux.Handle(prefix+"/", http.StripPrefix(prefix, http.FileServer(http.Dir(dir))))
}