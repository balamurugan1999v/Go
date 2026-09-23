package main

import (
	controllers "crudOperation/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.SetTrustedProxies(nil)
	controllers.Routers(server)
	server.Run("localhost:8081")

}
