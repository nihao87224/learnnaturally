package main

import (
	"github.com/gin-gonic/gin"
	"learnNaturally/configuration"
	"learnNaturally/remini"
	"log"
)

type Prompt struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/chat", func(c *gin.Context) {
		cfg, err := configuration.LoadConfig("./data/")
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		var prompt Prompt
		if err := c.BindJSON(&prompt); err != nil {
			log.Fatalf("Error binding JSON: %v", err)
		}
		content := remini.ChatAndProofread(prompt.Message, cfg.Common.ProxyUrl)
		c.JSON(200, gin.H{
			"message": content,
		})
	})

	err := router.Run("0.0.0.0:8091")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

//func main() {
//	ctx := context.Background()
//	client, err := commands.MyGeminiClient(ctx, "http://127.0.0.1:10809")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer client.Close()
//	model := client.GenerativeModel("gemini-1.5-flash")
//	session := model.StartChat()
//	session.History = []*genai.Content{}
//	resp, err := session.SendMessage(ctx, genai.Text("hello\n"))
//	if err != nil {
//		log.Fatalf("Error sending message: %v", err)
//	}
//
//	for _, part := range resp.Candidates[0].Content.Parts {
//		fmt.Printf("%v\n", part)
//	}
//}
