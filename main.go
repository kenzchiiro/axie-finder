package main

import (
	"axie-notify/models"
	"axie-notify/services"
	"axie-notify/services/delivery/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jasonlvhit/gocron"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Fetch(bot *linebot.Client) {
	queueFile, err := os.Open("data/queue.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer queueFile.Close()
	// Unmarshal JSON
	byteValue, _ := ioutil.ReadAll(queueFile)
	queueList := map[string]models.Queue{}

	err = json.Unmarshal(byteValue, &queueList)

	for userID, v := range queueList {
		if v.Command == "#find" {
			axiesData := services.SetParameterAxieFromMessage(v.Parameter)
			flexMessage := services.SetAxieToFlexMessage(axiesData)
			if _, err := bot.PushMessage(userID, flexMessage).Do(); err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}

func main() {
	bot := connectLineBot()
	go gocron.Every(5).Seconds().Do(Fetch, bot)
	go startService()
	<-gocron.Start()

}

func connectLineBot() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func startService() {
	e := echo.New()
	http.NewServiceHTTPHandler(e, connectLineBot())
	e.Logger.Fatal(e.Start(getPort()))
}

func getPort() string {
	var port = os.Getenv("PORT") // ----> (A)
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port // ----> (B)
}
