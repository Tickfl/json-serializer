package writers

import (
	"encoding/json"
	"io"
	"json-serializer/internal/models"
)

type JsonWriter struct {
	Output io.Writer
	Data   []models.JsonData
}

func (c *JsonWriter) Write() error {
	encoder := json.NewEncoder(c.Output)
	encoder.SetIndent("", "  ")

	return encoder.Encode(c.Data)
}
