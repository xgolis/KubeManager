package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"k8s.io/client-go/kubernetes"
)

type Manager struct {
	KubeClient *kubernetes.Clientset
	HelmChart  string
}

type Response struct {
	PathToHelm string `json:"pathToHelm"`
}

func getUsersRequest(req *http.Request) (*User, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading request: %v", err)
	}

	var usersRequest User
	err = json.Unmarshal(body, &usersRequest)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return &usersRequest, nil
}

func getHelmPath(usersRequest *User) (string, error) {
	posturl := "http://10.102.243.209:8081"

	// ziskaj available port
	body, err := json.Marshal(usersRequest)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error while creating request: %v", err)
	}

	r.Header.Add("Content-Type", "application/x-tar")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", fmt.Errorf("error while creating request: %v", err)
	}

	defer res.Body.Close()

	file, err := os.Create(usersRequest.UserName + ".tar")
	if err != nil {
		return "", fmt.Errorf("error while creating request: %v", err)
	}
	defer file.Close()

	// Copy the content from the response body to the file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", fmt.Errorf("error while copying request: %v", err)
	}

	tarFilename := usersRequest.UserName + ".tar"

	// Create a new tar command to untar the file
	cmd := exec.Command("tar", "-xf", tarFilename)

	// Set the working directory for the command (optional)
	cmd.Dir = "./"

	// Set the standard output and error for the command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and wait for it to finish
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	os.Remove(tarFilename)
	return "", nil
	// return response.PathToHelm, nil
}

func getKubeClient(pathToHelm string) *Manager {
	return &Manager{}
}
