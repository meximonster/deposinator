package utils

import (
	"strconv"
	"strings"
)

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

func ParseArray(array string) ([]int, error) {
	array = strings.Trim(array, "{}")
	if array == "" {
		return []int{}, nil
	}

	parts := strings.Split(array, ",")
	result := make([]int, len(parts))
	for i, part := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}
