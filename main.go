package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type JsonObject struct {
	Id        string                 `json:"id"`
	Latitude  float64                `json:"latitude"`
	Longitude float64                `json:"longitude"`
	Tags      map[string]interface{} `json:"tags"`
}

func main() {
	config := flag.String("config", "", "Path to json file")
	outputPath := flag.String("output", "", "Path to output csv")
	flag.Parse()

	var err error
	var data []byte
	var objects []JsonObject

	if *config != "" {
		data, err = os.ReadFile(*config)
		if err != nil {
			fmt.Println("ReadFile error:", err)
			return
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		data, err = reader.ReadBytes(']')
		if err != nil {
			fmt.Println("Unmarshal error", err)
			return
		}
	}

	err = json.Unmarshal(data, &objects)
	if err != nil {
		fmt.Println("Unmarshal error", err)
		return
	}

	output := os.Stdout
	if *outputPath != "" {
		var file *os.File
		file, err = os.Create(*outputPath)
		if err != nil {
			fmt.Println("Create file error", err)
			return
		}
		defer file.Close()
		output = file
	}

	writer := csv.NewWriter(output)
	defer writer.Flush()

	for _, object := range objects {
		row := []string{
			object.Id,
			fmt.Sprintf("%f", object.Latitude),
			fmt.Sprintf("%f", object.Longitude),
		}

		for _, tag := range object.Tags {
			row = append(row, fmt.Sprintf("%v", tag))
		}

		err = writer.Write(row)
		if err != nil {
			fmt.Println("Create file error", err)
			return
		}
	}
}
