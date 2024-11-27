package main

import (
	"go-code-challenge/cmd"
)

// @title Go Code Challenge API
// @version 1.0
// @description This is the Go server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/jubarodrigo/go-code-challenge
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /

func main() {
	cmd.StartApp()
}
