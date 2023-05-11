package server

type PreyRequest struct {
	Speed float64 `json:"speed"`
}

type SharkRequest struct {
	XPosition float64 `json:"x_position"`
	YPosition float64 `json:"y_position"`
	Speed     float64 `json:"speed"`
}
