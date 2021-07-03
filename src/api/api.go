package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func createEngine() *gin.Engine {
	r := gin.Default()
	ph := CreateProperHandler()
	r.GET("/proper/extract/:text", ph.extractProper)
	r.GET("/proper/count/:text", ph.countProper)
	r.POST("/proper/extract", ph.extractProper)
	r.POST("/proper/count", ph.countProper)
	return r
}

func Bootstrap(port int) {
	r := createEngine()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	s := &http.Server{
		Addr:           ":" + fmt.Sprint(port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
