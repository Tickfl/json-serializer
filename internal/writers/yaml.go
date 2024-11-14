package writers

import (
	"gopkg.in/yaml.v3"
	"io"
	"json-serializer/internal/models"
)

type YamlWriter struct{}

func (c *YamlWriter) Write(Output io.Writer, Data []models.JsonData) error {
	encoder := yaml.NewEncoder(Output)
	defer encoder.Close()

	return encoder.Encode(Data)
}
