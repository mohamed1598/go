package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)
func printCommantEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4536953643543-4544965394374-5i5YCg1J9WFGeQ72m9g5CTFj")
	os.Setenv("SLAK_APP_TOKEN","xapp-1-A04G0T29286-4575221549296-94f14c0d14836160c738f4c89140877bb20191f457c9054dcd17f1de57bd7a3e")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLAK_APP_TOKEN"))
	go printCommantEvents(bot.CommandEvents())
	// 
	bot.Command("my yob is <year>",&slacker.CommandDefinition{
		Description: "yob calculator",
		Examples: []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext,request slacker.Request,response slacker.ResponseWriter){
			year := request.Param("year")
			yob,err :=strconv.Atoi(year)
			if err!= nil{
				println("error")
			}
			age := 2021-yob
			r:= fmt.Sprintf("age is %d",age)
			response.Reply(r)
		},
	})

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}
