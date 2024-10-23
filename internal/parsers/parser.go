package parsers

import (
	"encoding/json"
	"json-serializer/internal/models"
)

func ParseJson(data []byte) ([]models.JsonData, error) {
	var jsonData []models.JsonData

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
