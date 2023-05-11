package products

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestGetProductsOK(t *testing.T) {
	//arrange
	sv := NewServiceMock()
	sv.On("GetAllBySeller", "FEX112AC").
		Return([]Product{{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		}}, nil)
	hd := NewHandler(sv)

	r := gin.Default()
	g := r.Group("/v1")
	g.GET("/products", hd.GetProducts)
	expectedHTTPStatusCode := http.StatusOK
	expectedPrds := `[{"ID": "mock",
			"SellerID": "FEX112AC",
			"Description": "generic product",
			"Price": 123.55}]`
	//act
	request, response := CreateRequestTest("GET", "/v1/products?seller_id=FEX112AC", "")
	r.ServeHTTP(response, request)

	//assert
	assert.Equal(t, expectedHTTPStatusCode, response.Code)
	assert.JSONEq(t, expectedPrds, response.Body.String(), "configure of shark incorrect")
	sv.AssertExpectations(t)

}

func TestGetProductsServerError(t *testing.T) {
	//arrange
	sv := NewServiceMock()
	sv.On("GetAllBySeller", "F").
		Return([]Product{}, errors.New("error in repository unexpected sellerId: F"))
	hd := NewHandler(sv)

	r := gin.Default()
	g := r.Group("/v1")
	g.GET("/products", hd.GetProducts)
	expectedHTTPStatusCode := http.StatusInternalServerError
	expected := `{"error": "error in repository unexpected sellerId: F"}`
	//act
	request, response := CreateRequestTest("GET", "/v1/products?seller_id=F", "")
	r.ServeHTTP(response, request)

	//assert
	assert.Equal(t, expectedHTTPStatusCode, response.Code)
	assert.JSONEq(t, expected, response.Body.String(), "invalid response")
	sv.AssertExpectations(t)

}

func TestGetProductsInvalidID(t *testing.T) {
	//arrange
	sv := NewServiceMock()
	hd := NewHandler(sv)

	r := gin.Default()
	g := r.Group("/v1")
	g.GET("/products", hd.GetProducts)
	expectedHTTPStatusCode := http.StatusBadRequest
	expected := `{"error": "seller_id query param is required"}`
	//act
	request, response := CreateRequestTest("GET", "/v1/products?seller_id=", "")
	r.ServeHTTP(response, request)

	//assert
	assert.Equal(t, expectedHTTPStatusCode, response.Code)

	assert.JSONEq(t, expected, response.Body.String(), "seller id incorrect")

}
