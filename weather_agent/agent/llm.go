
// agent/llm.go
package agent

import (
    "fmt"
    "github.com/go-resty/resty/v2"
)

type LLM struct {
    apiKey string
}

func NewLLM(apiKey string) *LLM {
    return &LLM{apiKey: apiKey}
}

func (l *LLM) callGPT(prompt string) string {
    client := resty.New()
    var result struct {
        Choices []struct {
            Message struct{ Content string `json:"content"` }
        } `json:"choices"`
    }
    client.R().
        SetHeader("Authorization", "Bearer "+l.apiKey).
        SetHeader("Content-Type", "application/json").
        SetBody(map[string]interface{}{
            "model": "gpt-3.5-turbo",
            "messages": []map[string]string{
                {"role": "user", "content": prompt},
            },
        }).
        SetResult(&result).
        Post("https://api.openai.com/v1/chat/completions")
    if len(result.Choices) > 0 {
        return result.Choices[0].Message.Content
    }
    return ""
}

func (l *LLM) ParseIntent(text string) string {
    prompt := fmt.Sprintf(`判断用户意图，如果是查询天气返回"查询天气"，否则返回"其他"：%s`, text)
    return l.callGPT(prompt)
}

func (l *LLM) ExtractCity(text string) string {
    prompt := fmt.Sprintf(`从文本中提取城市名称，例如“帮我查北京天气”，返回"北京"，只返回城市名：%s`, text)
    return l.callGPT(prompt)
}
