package main
// See the sunset and current time for Berlin
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/fatih/color"
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

	color.Cyan("Jetzt ist %s\n", data.Timestamp)
    color.Red("Die Sonne geht um %s unter\n", data.Wheather.Sunset)
    color.Green("Verbleibende Sonnenstunden: %d Stunden und %d Minuten\n", hours, minutes)

}

func endpoint(url string) responseData {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Failed to fetch data:", err)
        return responseData{}
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Failed to read response body:", err)
        return responseData{}
    }

    var all responseData
    json.Unmarshal(body, &all)
    return all
}

