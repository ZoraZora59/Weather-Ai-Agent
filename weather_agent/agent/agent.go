// agent/agent.go
package agent

import (
)

type Agent struct {
    llm        *LLM
    weatherAPI *WeatherClient
}

func NewAgent() *Agent {
    return &Agent{
        llm:        NewLLM("你的OpenAI Key"),
        weatherAPI: NewWeatherClient(),
    }
}

func (a *Agent) HandleRequest(input string) string {
    intent := a.llm.ParseIntent(input)
    if intent != "查询天气" {
        return "暂时只能帮你查天气哦。"
    }
    city := a.llm.ExtractCity(input)
    if city == "" {
        return "请告诉我要查询哪个城市的天气。"
    }
    weather := a.weatherAPI.GetWeather(city)
    return weather
}