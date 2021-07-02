package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.StaticFile("/.well-known/apple-app-site-association", "./static/apple-app-site-association")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/demo", func(c *gin.Context) {
		res := []string{"foo", "bar"}
		c.JSON(200, res)
	})

	// router.GET("/.well-known/apple-app-site-association", func(c *gin.Context) {
	// 	appclips := map[string]interface{}{
	// 		"apps": []interface{}{
 //        		"9HC298K985.com.orlov.cvapp.Clip",
 //    		},
	// 	}
	// 	result := map[string]interface{}{
 //    		"appclips": appclips,
	// 	}

	// 	c.Header("Content-Type", "application/json")
	// 	c.JSON(200, result)
	// })	

	router.Run(":" + port)
}
