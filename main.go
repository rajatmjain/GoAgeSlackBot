package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main(){
	SLACK_BOT_TOKEN := goDotEnvVariable("SLACK_BOT_TOKEN")
	SLACK_APP_TOKEN := goDotEnvVariable("SLACK_APP_TOKEN")

    bot := slacker.NewClient(SLACK_BOT_TOKEN,SLACK_APP_TOKEN)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>",&slacker.CommandDefinition{
		Description: "age calculator",
		Examples : []string{"my yob is 1997"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob,err := strconv.Atoi(year)
			if(err!=nil){
				fmt.Println(err)
			}
			age := 2022-yob
			r := fmt.Sprintf("Your age is %d years",age)
			response.Reply(r)
		},
	})

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if(err!=nil){
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("-------------------")
		fmt.Println("Command Events")
		fmt.Println("-------------------")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("-------------------")
	}
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }