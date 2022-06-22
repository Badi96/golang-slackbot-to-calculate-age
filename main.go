package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChanel <-chan *slacker.CommandEvent) {
	for event := range analyticsChanel {
		fmt.Println("command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	year_today, _, _ := time.Now().Date()

	os.Setenv("SLACK_BOT_TOKEN", "Replace_with_Bot_token_here")
	os.Setenv("SLACK_APP_TOKEN", "Replace_with_channel_id_token") //socket token
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	//Comamand with example for the bot
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year) // needed to convert from string to number
			if err != nil {
				println("error")
			}
			age := year_today - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //defer makes sure that canel is called towards the end

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
