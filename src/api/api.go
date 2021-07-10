package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"project.com/mecabapi/parser"
)

func createEngine() *gin.Engine {
	r := gin.Default()
	neoM, err := parser.CreateMecabTagger(true)
	ipadicM, err := parser.CreateMecabTagger(false)
	if err != nil {
		panic(err)
	}
	eph := CreateExtractProperHandler(neoM)
	r.GET("/proper/extract/:text", extractRequestText, eph.handle)
	r.POST("/proper/extract", extractRequestText, eph.handle)

	cph := CreateCountProperHandler(neoM)
	r.GET("/proper/count/:text", extractRequestText, cph.handle)
	r.POST("/proper/count", extractRequestText, cph.handle)

	ech := CreateExtractCountryHandler(ipadicM)
	r.POST("/country/extract", extractRequestText, ech.handle)

	egh := CreateExtractGeoHandler(ipadicM)
	r.POST("/geo/extract", extractRequestText, egh.handle)

	eh := CreateExtractHandler(neoM)
	r.POST("/extract", extractRequestText, eh.handle)
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
