package main

import (
	"gogogo/cmd/iojson"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PATH_HTML = "C:/programms/Go/Goprojects/src/gogogo/ui/html/"
const PATH_STATIC = "C:/programms/Go/Goprojects/src/gogogo/ui/static/"

func main() {
	var net_connection =true
	response,err:=http.Get("https://www.google.com")
	if err !=nil||response.StatusCode != http.StatusOK{
		net_connection=false
		log.Println("Нет подключения к интернету!")
	}
	route := gin.Default()
	if(net_connection){
		iojson.Start()
		log.Println("Метеоданные обновлены!")
	}

	route.LoadHTMLGlob(PATH_HTML + "*")
	route.Static("/static", PATH_STATIC)

	route.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Next()

	})

	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	route.GET("/d3map", func(c* gin.Context){
		c.HTML(http.StatusOK, "d3map.html", gin.H{})
	})
	log.Println("Git")
	route.Run(":8080")
}
