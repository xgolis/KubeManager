package app

import (
	"fmt"
	"net/http"
)

func makeHandlers() *http.ServeMux {
	mux := *http.NewServeMux()
	mux.HandleFunc("/", initialDeployment)
	return &mux
	// a.Server.HandleFunc("/", getGit(w http.ResponseWriter, req *http.Request)})
}

func initialDeployment(w http.ResponseWriter, req *http.Request) {
	usersRequest, err := getUsersRequest(req)
	if err != nil {
		fmt.Fprintf(w, "error reading users request: %v", err)
		return
	}
	usersRequest.Port = "31930"
	_, err = getHelmPath(usersRequest)
	if err != nil {
		fmt.Fprintf(w, "error while getting helm: %v", err)
		return
	}
}
