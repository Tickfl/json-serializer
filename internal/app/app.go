package app

import (
	"flag"
	"fmt"
	"io"
	"json-serializer/internal/models"
	"json-serializer/internal/parsers"
	"json-serializer/internal/writers"
	"os"
)

func Run() {
	config := flag.String("config", "", "Path to json file")
	outputPath := flag.String("output", "", "Path to output csv")
	format := flag.String("format", "csv", "Output format file")
	flag.Parse()

	var err error
	var data []byte
	var objects []models.JsonData

	if *config != "" {
		data, err = os.ReadFile(*config)
		if err != nil {
			fmt.Println("ReadFile error:", err)
			return
		}
	} else {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("ReadAll error:", err)
			return
		}
	}

	objects, err = parsers.ParseJson(data)
	if err != nil {
		fmt.Println("ParseJson error:", err)
		return
	}

	output := os.Stdout
	if *outputPath != "" {
		output, err = os.Create(*outputPath)
		if err != nil {
			fmt.Println("Error create file:", err)
			return
		}
		defer output.Close()
	}

	var writer writers.Writer

	switch *format {
	case "csv":
		writer = &writers.CsvWriter{Output: output, Data: objects}
	case "json":
		writer = &writers.JsonWriter{Output: output, Data: objects}
	case "yaml":
		writer = &writers.YamlWriter{Output: output, Data: objects}
	case "protobuf":
		writer = &writers.ProtobufWriter{Output: output, Data: objects}
	default:
		fmt.Println("Error format:", *format)
		return
	}

	err = writer.Write()
	if err != nil {
		fmt.Println("Error writer:", err)
		return
	}
}
