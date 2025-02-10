package utils

type JSONResponse struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Result      interface{} `json:"result,omitempty"`
}

func GenerateJSONResponse(status string, description string) JSONResponse {
	return JSONResponse{
		Status:      status,
		Description: description,
	}
}

func GenerateJSONResultResponse(status string, description string, result interface{}) JSONResponse {
	return JSONResponse{
		Status:      status,
		Description: description,
		Result:      result,
	}
}
