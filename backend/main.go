package main

import (
	"log"
	"net/http"

	"github.com/P-Raphat/2026-AMASS-Quiz-ThaiBev/database"
	"github.com/P-Raphat/2026-AMASS-Quiz-ThaiBev/routes"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	database.Connect()

	r := gin.Default()
	r.Use(CORSMiddleware())

	routes.Setup(r)

	log.Fatal(r.Run(":8080"))
}
