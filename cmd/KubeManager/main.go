package main

import (
	"KubeManager/cmd/KubeManager/app"
	"fmt"
	"net/http"
	"time"
)

func main() {
	application := app.NewApp()

	mux := application.MakeHandlers()

	Server := &http.Server{
		Addr:           "0.0.0.0:8085",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("[Server] Up and running on %s\n", Server.Addr)
	http.ListenAndServe(Server.Addr, Server.Handler)
}
