package main

import (
	"KubeManager/cmd/KubeManager/app"
	"fmt"
	"net/http"
)

func main() {
	app := app.NewApp()

	fmt.Printf("[Server] Up and running on %s\n", app.Server.Addr)
	http.ListenAndServe(app.Server.Addr, app.Server.Handler)
}
