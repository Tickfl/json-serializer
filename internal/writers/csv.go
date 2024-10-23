package writers

import (
	"encoding/csv"
	"fmt"
	"io"
	"json-serializer/internal/models"
)

type CsvWriter struct {
	Output io.Writer
	Data   []models.JsonData
}

func (c *CsvWriter) Write() error {
	writer := csv.NewWriter(c.Output)
	defer writer.Flush()

	for _, object := range c.Data {
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
