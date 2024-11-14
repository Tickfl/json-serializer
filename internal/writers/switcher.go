package writers

import (
	"fmt"
	"io"
	"json-serializer/internal/models"
)

type Writer interface {
	Write(Output io.Writer, Data []models.JsonData) error
}

func GetWriter(format string) (Writer, error) {
	var writer Writer

	switch format {
	case "csv":
		writer = &CsvWriter{}
	case "json":
		writer = &JsonWriter{}
	case "yaml":
		writer = &YamlWriter{}
	case "protobuf":
		writer = &ProtobufWriter{}
	default:
		return nil, fmt.Errorf("error format " + format)
	}

	return writer, nil
}
