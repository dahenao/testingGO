package server

type ResponseSuccess struct {
	Success bool `json:"success"`
}

type ResponseSimulate struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Time    float64 `json:"time"`
}
