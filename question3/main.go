package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router := Router()
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("start http server listening 3000...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server Error: %s", err.Error())
	}
}

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "ok",
			"data":      "pong",
			"timestamp": time.Now().Unix(),
		})
	})

	// Step1: add a signup API
	// 1. check email is exist
	// 2. create account
	// 3. response account info which contains token

	// Step2: add an API to get account info
	// 1. check token is valid
	// 2. get account info
	// 3. response account info

	return router
}
