package writers

import (
	"encoding/csv"
	"fmt"
	"io"
	"json-serializer/internal/models"
)

type CsvWriter struct{}

func (c *CsvWriter) Write(Output io.Writer, Data []models.JsonData) error {
	writer := csv.NewWriter(Output)
	defer writer.Flush()

	for _, object := range Data {
		row := []string{
			object.Id,
			fmt.Sprintf("%f", object.Latitude),
			fmt.Sprintf("%f", object.Longitude),
		}

		for _, tag := range object.Tags {
			row = append(row, fmt.Sprintf("%v", tag))
		}

		err := writer.Write(row)
		if err != nil {
			return fmt.Errorf("error write to csv: %w", err)
		}
	}

	return nil
}
