package app

import (
	"fmt"
	"net/http"
	"strconv"
)

func (a *App) MakeHandlers() *http.ServeMux {
	mux := *http.NewServeMux()
	mux.HandleFunc("/", a.initialDeployment)
	return &mux
	// a.Server.HandleFunc("/", getGit(w http.ResponseWriter, req *http.Request)})
}

func (a *App) initialDeployment(w http.ResponseWriter, req *http.Request) {
	usersRequest, err := getUsersRequest(req)
	if err != nil {
		fmt.Fprintf(w, "error reading users request: %v", err)
		return
	}

	usersRequest.Port = a.findAvailablePort()

	helmPath, err := getHelmPath(usersRequest)
	if err != nil {
		fmt.Fprintf(w, "error while getting helm: %v", err)
		return
	}

	err = applyHelm(usersRequest, helmPath)
	if err != nil {
		fmt.Fprintf(w, "error while applying: %v", err)
		return
	}
}

func (a *App) findAvailablePort() string {
	i := 31940

	for port := range a.Ports {
		if port == i {
			i++
		}
	}

	a.Ports = append(a.Ports, i)

	return strconv.Itoa(i)
}
