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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/demo", func(c *gin.Context) {
		res := []string{"foo", "bar"}
		c.JSON(200, res)
	})

	router.GET("/.well-known/apple-app-site-association", func(c *gin.Context) {
		result := map[string]interface{}{
    		"appclips": interface{} {
    			"apps": []interface{} {
    				"9HC298K985.com.orlov.cvapp.Clip"
    			}
    		}
		}
		c.JSON(200, result)
	})	

	router.Run(":" + port)
}
