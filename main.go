package main

import (
	"fmt"

	"github.com/turnage/graw/reddit"
)

func main() {
	/*model, _ := sentiment.Restore()

	analysis := model.SentimentAnalysis("I like the beach", sentiment.English) // 0

	fmt.Println(analysis)*/

	bot, err := reddit.NewBotFromAgentFile(".secret", 0)

	harvest, err := bot.Listing("/r/golang", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/golang: ", err)
		return
	}

	for _, post := range harvest.Posts[:5] {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}

}
