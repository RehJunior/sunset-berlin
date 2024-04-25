package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)
type ResponseData struct {
	Timestamp string `json:"timestamp"`
	Wheather  struct {
		ApparentTemperatureMax      float64 `json:"apparent_temperature_max"`
		ApparentTemperatureMin      float64 `json:"apparent_temperature_min"`
		PrecipitationProbabilityMax int     `json:"precipitation_probability_max"`
		Sunset                      string  `json:"sunset"`
		Temperature                 float64 `json:"temperature"`
	} `json:"wheather"`
}
func main()  {
url:= "https://bvg.fly.dev"
data := endpoint(url)


fmt.Printf("Jetzt ist %v \nDie Sonne geht um %v unter\n",data.Timestamp,data.Wheather.Sunset)
}

func endpoint(url string)ResponseData{
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body , err := io.ReadAll(resp.Body)
	
	var all ResponseData
	err = json.Unmarshal(body,&all)
	if err != nil {
		log.Fatal(err)
	}
	return all
}