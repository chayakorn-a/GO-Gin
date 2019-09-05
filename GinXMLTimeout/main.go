package main

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type Login struct {
	XMLName  xml.Name `xml:"login"`
	User     string   `form:"user" xml:"user"  binding:"required"`
	Password string   `form:"password" xml:"password" binding:"required"`
	Data     string   `form:"data" xml:"data" binding:"required"`
	Info     Info     `form:"info" xml:"info" binding:"required"`
}

type Info struct {
	XMLName xml.Name `xml:"info"`
	Name    string   `xml:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in", "DataResp": xml.Data, "infoName": xml.Info.Name})
	})
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8090")
}
