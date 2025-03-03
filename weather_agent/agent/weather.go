// agent/weather.go
package agent

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type WeatherClient struct{}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{}
}

func (w *WeatherClient) GetWeather(city string) string {
	client := resty.New()
	var resp struct {
		City string `json:"city"`
		Wea  string `json:"wea"`
		Tem  string `json:"tem"`
	}
	_, err := client.R().
		SetQueryParams(map[string]string{
			"appid":     "API ID",
			"appsecret": "API KEY",
			"city":      city,
		}).
		SetResult(&resp).
		Get("https://www.tianqiapi.com/free/day")
	if err != nil {
		return "查询天气失败"
	}
	return fmt.Sprintf("%s今天%s，温度%s。", resp.City, resp.Wea, resp.Tem)
}
