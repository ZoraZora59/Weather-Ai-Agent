# 天气查询 Ai Agent

只是学习AiAgent过程中的一个小项目，仅供参考。

## 运行步骤
```bash
go mod init go-ai-agent-demo
go get github.com/gin-gonic/gin
go get github.com/go-resty/resty/v2
go run main.go
```

## 测试 API
```bash
curl -X POST http://localhost:8080/ask \
-H "Content-Type: application/json" \
-d '{"message": "帮我查下北京天气"}'
```

## 结构
```
weather_agent/
├── main.go
├── agent/
│   ├── agent.go       # 核心逻辑
│   ├── llm.go         # GPT封装
│   └── weather.go     # 天气接口封装
├── go.mod
└── go.sum
```
