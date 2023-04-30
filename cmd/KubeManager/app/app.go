package app

type App struct {
	Ports []int
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
	return &App{
		Ports: []int{},
	}
}
