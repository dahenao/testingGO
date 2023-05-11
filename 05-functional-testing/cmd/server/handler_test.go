package server

import (
	"bytes"
	"functional/prey"
	"functional/shark"
	"functional/simulator"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CreateServer() *gin.Engine {
	r := gin.Default()
	sim := simulator.NewCatchSimulator(35.4)

	whiteShark := shark.CreateWhiteShark(sim)
	tuna := prey.CreateTuna()

	handler := NewHandler(whiteShark, tuna)

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

func TestHunt(t *testing.T) {
	t.Run("Escape prey, slow shark", func(t *testing.T) {
		//arrange

		expectConfigurePrey := `{"success": true}`
		expectConfigureShark := `{"success": true}`
		expectConfigureSimul := `{"success": true,"message": "could not catch it","time": 0}`
		servertest := CreateServer()
		preyBody := `{"speed": 299.0}`
		sharkbody := `{
			"x_position": 2.5,
			"y_position": 1.0,
			"speed": 200.0
			}`

		requestPrey, responsePrey := CreateRequestTest("PUT", "/v1/prey", preyBody)
		requestShark, responseShark := CreateRequestTest("PUT", "/v1/shark", sharkbody)
		requestSimu, responseSimu := CreateRequestTest("POST", "/v1/simulate", "")

		//act
		servertest.ServeHTTP(responseShark, requestShark)
		servertest.ServeHTTP(responsePrey, requestPrey)
		servertest.ServeHTTP(responseSimu, requestSimu)

		//assert
		assert.JSONEq(t, expectConfigurePrey, responsePrey.Body.String(), "configure of prey incorrect")

		assert.JSONEq(t, expectConfigureShark, responseShark.Body.String(), "configure of shark incorrect")
		t.Log(responseSimu.Body.String())
		assert.JSONEq(t, expectConfigureSimul, responseSimu.Body.String(), "error while hunt")

	})

	t.Run("Escape prey, shark is a big distance", func(t *testing.T) {
		//arrange

		expectConfigurePrey := `{"success": true}`
		expectConfigureShark := `{"success": true}`
		expectConfigureSimul := `{"success": true,"message": "could not catch it","time": 0}`
		servertest := CreateServer()
		preyBody := `{"speed": 90.0}`
		sharkbody := `{
			"x_position": 500.0,
			"y_position": 500.0,
			"speed": 100.0
			}`

		requestPrey, responsePrey := CreateRequestTest("PUT", "/v1/prey", preyBody)
		requestShark, responseShark := CreateRequestTest("PUT", "/v1/shark", sharkbody)
		requestSimu, responseSimu := CreateRequestTest("POST", "/v1/simulate", "")

		//act
		servertest.ServeHTTP(responseShark, requestShark)
		servertest.ServeHTTP(responsePrey, requestPrey)
		servertest.ServeHTTP(responseSimu, requestSimu)

		//assert
		assert.JSONEq(t, expectConfigurePrey, responsePrey.Body.String(), "configure of prey incorrect")

		assert.JSONEq(t, expectConfigureShark, responseShark.Body.String(), "configure of shark incorrect")
		t.Log(responseSimu.Body.String())
		assert.JSONEq(t, expectConfigureSimul, responseSimu.Body.String(), "error while hunt")

	})

	t.Run("shark hunt after 24 min", func(t *testing.T) {
		//arrange

		expectConfigurePrey := `{"success": true}`
		expectConfigureShark := `{"success": true}`
		expectConfigureSimul := `{"success": true,"message": "","time": 24}`
		servertest := CreateServer()
		preyBody := `{"speed": 80.0}`
		sharkbody := `{
			"x_position": 400.0,
			"y_position": 400.0,
			"speed": 103.0
			}`

		requestPrey, responsePrey := CreateRequestTest("PUT", "/v1/prey", preyBody)
		requestShark, responseShark := CreateRequestTest("PUT", "/v1/shark", sharkbody)
		requestSimu, responseSimu := CreateRequestTest("POST", "/v1/simulate", "")

		//act
		servertest.ServeHTTP(responseShark, requestShark)
		servertest.ServeHTTP(responsePrey, requestPrey)
		servertest.ServeHTTP(responseSimu, requestSimu)

		//assert
		assert.JSONEq(t, expectConfigurePrey, responsePrey.Body.String(), "configure of prey incorrect")

		assert.JSONEq(t, expectConfigureShark, responseShark.Body.String(), "configure of shark incorrect")
		t.Log(responseSimu.Body.String())
		assert.JSONEq(t, expectConfigureSimul, responseSimu.Body.String(), "error while hunt")

	})
}
