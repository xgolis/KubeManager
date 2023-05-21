package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
	Port    string `json:"port,omitempty"`
}

func (a *App) MakeHandlers() *http.ServeMux {
	mux := *http.NewServeMux()
	mux.HandleFunc("/", a.initialDeployment)
	return &mux
	// a.Server.HandleFunc("/", getGit(w http.ResponseWriter, req *http.Request)})
}

func sendError(w *http.ResponseWriter, err error) {
	(*w).Header().Set("Content-Type", "application/json")

	status := Response{
		Message: err.Error(),
	}
	fmt.Print(status.Message)
	statusJson, err := json.Marshal(status)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

	(*w).Write(statusJson)
}

func (a *App) initialDeployment(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	usersRequest, err := getUsersRequest(req)
	if err != nil {
		sendError(&w, fmt.Errorf("error reading users request: %v\n", err))
		return
	}

	usersRequest.Port = a.findAvailablePort()

	helmPath, err := getHelmPath(usersRequest)
	if err != nil {
		sendError(&w, fmt.Errorf("error while getting helm: %v\n", err))
		return
	}

	err = applyHelm(usersRequest, helmPath)
	if err != nil {
		sendError(&w, fmt.Errorf("error while applying: %v\n", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	status := Response{
		Message: "The application deployed successfully\nThe application is accessible at 35.240.30.14:" + usersRequest.Port,
		Port:    usersRequest.Port,
	}
	statusJson, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(statusJson)
}

func (a *App) findAvailablePort() string {
	i := 31940

	for _, port := range a.Ports {
		if port == i {
			i = i + 1
		}
	}

	a.Ports = append(a.Ports, i)
	fmt.Print(a.Ports)

	return strconv.Itoa(i)
}
