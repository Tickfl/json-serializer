package models

type JsonData struct {
	Id        string                 `json:"id"`
	Latitude  float64                `json:"latitude"`
	Longitude float64                `json:"longitude"`
	Tags      map[string]interface{} `json:"tags"`
}
