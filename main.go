package main

import (
	"github.com/cristiano-pacheco/go-starter-kit/cmd"
	"github.com/cristiano-pacheco/go-starter-kit/internal/shared/modules/config"
)

// @title           Go Starter API
// @version         1.0
// @description     Go Starter API

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer <token>**

// @BasePath  /
func main() {
	config.Init()
	cmd.Execute()
}
