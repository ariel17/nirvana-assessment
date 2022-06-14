package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/ariel17/nirvana-assessment/api"
	"github.com/ariel17/nirvana-assessment/pkg/configs"
)

// StartServer creates a new instance of HTTP server with indicated handlers
// configured and begins serving content.
func StartServer() {
	r := gin.Default()
	r.GET("/", CoalesceHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(fmt.Sprintf(":%d", configs.GetPort())); err != nil {
		panic(err)
	}
}