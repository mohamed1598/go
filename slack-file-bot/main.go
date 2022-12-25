package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4536953643543-4575494554480-uBjpUsEGb2yVkyxs5ZCyHAbp")
	os.Setenv("CHANNEL_ID", "C04G9TZ8GKE")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"oop.pdf"}
	for i:=0;i<len(fileArr);i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file ,err := api.UploadFile(params)
		if err !=nil{
			fmt.Printf("%s\n",err)
			return
		}
		fmt.Printf("Name:%s ,URL:%s\n",file.Name,file.URL)
	}
}
