package app

import (
	"net/http"
	"time"
)

type App struct {
	Server *http.Server
}

type User struct {
	// app
	Name string `json:"name"`
	// namespace
	UserName string `json:"username"`
	Image    string `json:"image"`
	Port     string `json:"port,omitempty"`
	AppPort  string `json:"appport"`
}

func NewApp() *App {

	mux := makeHandlers()

	return &App{
		Server: &http.Server{
			Addr:           "localhost:8085",
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}
