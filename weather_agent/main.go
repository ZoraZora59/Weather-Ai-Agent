// main.go
package main

import (
	"zorazora59/weather_ai_agent/agent"

	"github.com/gin-gonic/gin"
)

func main() {
	ag := agent.NewAgent()

	r := gin.Default()
	r.POST("/ask", func(c *gin.Context) {
		var req struct {
			Message string `json:"message"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		reply := ag.HandleRequest(req.Message)
		c.JSON(200, gin.H{"reply": reply})
	})

	r.Run(":8080")
}
