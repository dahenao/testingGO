package server

import (
	"functional/prey"
	"functional/shark"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	shark shark.Shark
	prey  prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/shark

func (h *Handler) ConfigureShark() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request SharkRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			resp := ResponseSuccess{Success: false}
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		h.shark.Configure([2]float64{request.XPosition, request.YPosition}, request.Speed)
		resp := ResponseSuccess{Success: true}
		c.JSON(http.StatusOK, resp)
		return
	}
}

// PUT: /v1/prey

func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request PreyRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			resp := ResponseSuccess{Success: false}
			c.JSON(http.StatusBadRequest, resp)
			return
		}

		h.prey.SetSpeed(request.Speed)
		resp := ResponseSuccess{Success: true}
		c.JSON(http.StatusOK, resp)
		return

	}
}

// POST: /v1/simulate

func (h *Handler) SimulateHunt() gin.HandlerFunc {
	return func(c *gin.Context) {

		err, time := h.shark.Hunt(h.prey)
		if err != nil {
			resp := ResponseSimulate{
				Success: false,
				Message: err.Error(),
				Time:    time}
			c.JSON(http.StatusOK, resp)
			return
		}
		resp := ResponseSimulate{Success: true,
			Message: "",
			Time:    time}
		c.JSON(http.StatusOK, resp)
		return

	}
}
