package main
// See the sunset and current time for Berlin
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

	// Parse the timestamp and sunset strings into time.Time objects
	const layout = "15:04" // Adjust the layout to match your timestamp format
	currentTime, err := time.Parse(layout, data.Timestamp)
	if err != nil {
		fmt.Println("Error parsing current time:", err)
		return
	}
	sunsetTime, err := time.Parse(layout, data.Wheather.Sunset)
	if err != nil {
		fmt.Println("Error parsing sunset time:", err)
		return
	}

	// Calculate the duration of sunshine left
	duration := sunsetTime.Sub(currentTime)

	// Extract hours and minutes from the duration
	hours := duration / time.Hour
	minutes := (duration % time.Hour) / time.Minute

	fmt.Printf("Jetzt ist %v \nDie Sonne geht um %v unter\n", data.Timestamp, data.Wheather.Sunset)
	fmt.Printf("Verbleibende Sonnenstunden: %d Stunden und %d Minuten\n", hours, minutes)
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
