package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lending/handler/loan"
	"github.com/lending/internal/core/services"
	"github.com/lending/queue"
	"github.com/lending/repositories"
)

func main() {
	messageQueueing := queue.New("127.0.0.1:4150")
	loanRepository := repositories.NewMemKVS()
	loanService := services.New(loanRepository, messageQueueing)
	loanHandler := loan.NewHTTPHandler(loanService)
	router := gin.New()
	router.GET("/loan/read/:id", loanHandler.Read)
	router.POST("/loan/create", loanHandler.Create)
	router.PUT("/loan/approve", loanHandler.Approve)
	router.PUT("/loan/update", loanHandler.Update)

	router.Run(":8080")
}
