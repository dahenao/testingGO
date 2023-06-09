package tests

import (
	"bytes"
	"functional/cmd/server"
	"functional/prey"
	"functional/shark"
	"functional/simulator"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	r := gin.Default()
	sim := simulator.NewCatchSimulator(35.4)

	whiteShark := shark.CreateWhiteShark(sim)
	tuna := prey.CreateTuna()

	handler := server.NewHandler(whiteShark, tuna)

	g := r.Group("/v1")

	g.PUT("/shark", handler.ConfigureShark())
	g.PUT("/prey", handler.ConfigurePrey())
	g.POST("/simulate", handler.SimulateHunt())

	return r
}

func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
