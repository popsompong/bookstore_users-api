package app

import (
	"github.com/gin-gonic/gin"
	"github.com/popsompong/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8080")
}
