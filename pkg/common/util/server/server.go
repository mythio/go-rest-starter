package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	e := gin.Default()
	e.GET("/", healthCheck)
	return e
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

// Config represents server specific config
type Config struct {
	Port                string
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	Debug               bool
}

// Start starts a gin server
func Start(e *gin.Engine) {
	s := &http.Server{
		Addr:         ":8080",
		Handler:      e,
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
	}
	// Start server
	go func() {
		// logger.Logger.Debug("Listening on port 8080")
		err := s.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	fmt.Println("Got signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
}
