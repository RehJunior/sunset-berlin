package main
// See the sunset and current time for Berlin
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type responseData struct {
	Timestamp string `json:"timestamp"`
	Wheather  struct {
		ApparentTemperatureMax      float64 `json:"apparent_temperature_max"`
		ApparentTemperatureMin      float64 `json:"apparent_temperature_min"`
		PrecipitationProbabilityMax int     `json:"precipitation_probability_max"`
		Sunset                      string  `json:"sunset"`
		Temperature                 float64 `json:"temperature"`
	} `json:"wheather"`
}

func main() {
	url := "https://bvg.fly.dev"
	data := endpoint(url)

	fmt.Printf("Jetzt ist %v \nDie Sonne geht um %v unter\n", data.Timestamp, data.Wheather.Sunset)
}

func endpoint(url string) responseData {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var all responseData
	json.Unmarshal(body, &all)
	return all
}
