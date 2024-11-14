package writers

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"json-serializer/internal/models"
	"json-serializer/proto/writers"
)

type ProtobufWriter struct{}

func (c *ProtobufWriter) Write(Output io.Writer, Data []models.JsonData) error {
	for _, item := range Data {
		protoItem := &writers.JsonData{
			Id:        item.Id,
			Latitude:  item.Latitude,
			Longitude: item.Longitude,
			Tags:      make(map[string]string),
		}

		for key, tag := range item.Tags {
			protoItem.Tags[key] = fmt.Sprintf("%v", tag)
		}

		data, err := proto.Marshal(protoItem)
		if err != nil {
			return fmt.Errorf("error marshal protobuf: %w", err)
		}

		if _, err := Output.Write(data); err != nil {
			return fmt.Errorf("error write protobuf: %w", err)
		}
	}
	return nil
}
