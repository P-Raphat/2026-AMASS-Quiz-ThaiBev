package routes

import (
	"github.com/P-Raphat/2026-AMASS-Quiz-ThaiBev/handlers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

}
