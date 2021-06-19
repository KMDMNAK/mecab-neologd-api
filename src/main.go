package main

import "github.com/gin-gonic/gin"

import "project.com/mecabapi/mecab"
import (
	"fmt"
)

func handleOwakati(c *gin.Context) {
	word := c.Param("word")
	fmt.Printf("%d\n", word)
	println(word)
	key := mecab.GetNode(word)
	// key := word
	c.JSON(200, gin.H{"key": key})
}

func main() {
	r := gin.Default()
	r.GET("/owakati/:word", handleOwakati)
	// mecab.getNode()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
