package api

import (
	"github.com/gin-gonic/gin"
)

func createEngine() *gin.Engine {
	r := gin.Default()
	ph := CreateProperHandler()
	r.GET("/proper/extract/:text", ph.extractProper)
	r.GET("/proper/count/:text", ph.countProper)
	return r
}

func Bootstrap() {
	r := createEngine()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
