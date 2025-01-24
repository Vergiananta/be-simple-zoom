package main

import (
	"fmt"

	"github.com/Vergiananta/be-simple-zoom/api/router"
	"github.com/Vergiananta/be-simple-zoom/config"
	"github.com/Vergiananta/be-simple-zoom/db/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello auth")
	r := gin.Default()
	router.GetRoute(r)

	r.Run()
}
