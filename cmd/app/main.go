package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"json-serializer/internal/models"
	"json-serializer/internal/writers"
	"os"
)

func main() {
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

	err = json.Unmarshal(data, &objects)
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
	writer, err = writers.GetWriter(*format)
	if err != nil {
		fmt.Println("Error writer:", err)
		return
	}

	err = writer.Write(output, objects)
	if err != nil {
		fmt.Println("Error writer:", err)
		return
	}
}
