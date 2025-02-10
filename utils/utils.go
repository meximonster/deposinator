package utils

type JSONResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

func GenerateJSONResponse(status string, description string) JSONResponse {
	return JSONResponse{
		Status:      status,
		Description: description,
	}
}
