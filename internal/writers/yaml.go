package writers

import (
	"gopkg.in/yaml.v3"
	"io"
	"json-serializer/internal/models"
)

type YamlWriter struct {
	Output io.Writer
	Data   []models.JsonData
}

func (c *YamlWriter) Write() error {
	encoder := yaml.NewEncoder(c.Output)
	defer encoder.Close()

	return encoder.Encode(c.Data)
}
