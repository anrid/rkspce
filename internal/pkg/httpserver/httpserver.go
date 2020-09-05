package httpserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server ...
type Server struct {
	e *echo.Echo
}

// New ...
func New() *Server {
	// Echo instance.
	e := echo.New()

	// Middleware.
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes.
	e.GET("/", root)
	e.GET("/health", health)

	return &Server{e}
}

// Start server.
func (s *Server) Start(addr string) {
	s.e.Logger.Fatal(s.e.Start(addr))
}

// AddRoutes expects a function that will add new routes
// to our server.
func (s *Server) AddRoutes(f func(e *echo.Echo)) {
	f(s.e)
}

// Default root handler.
func root(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Server time: %s", time.Now()))
}

// Default health check handler.
func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
