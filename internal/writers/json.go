package writers

import (
	"encoding/json"
	"io"
	"json-serializer/internal/models"
)

type JsonWriter struct{}

func (c *JsonWriter) Write(Output io.Writer, Data []models.JsonData) error {
	encoder := json.NewEncoder(Output)
	encoder.SetIndent("", "  ")

	return encoder.Encode(Data)
}
