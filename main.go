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

// func main() {

// 	accessToken := "vJ5VjJJ6IM15AxZqFPMHelDJ9AwGrYvxJTie98xSWKJ"
// 	message := "Found Axie#15326 on sale\nprice 0.12"
// 	imageURL := "https://storage.googleapis.com/assets.axieinfinity.com/axies/15326/axie/axie-full-transparent.png"
// 	if err := notify.SendImage(accessToken, message, imageURL); err != nil {
// 		panic(err)
// 	}
// 	// accessToken := "vJ5VjJJ6IM15AxZqFPMHelDJ9AwGrYvxJTie98xSWKJ"
// 	// message := "hello, world!"
// 	// imageURL := "image url. ex) https://..."

// 	// if err := notify.SendImage(accessToken, message); err != nil {
// 	// 	panic(err)
// 	// }

// 	app := fiber.New()

// 	app.Get("/posts", func(c *fiber.Ctx) error {
// 		var result []Post
// 		for _, post := range posts {
// 			result = append(result, post)
// 		}
// 		return c.JSON(result)
// 	})

// 	app.Listen(":3000")

// }

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
	go gocron.Every(10).Seconds().Do(Fetch, bot)
	go startService()
	<-gocron.Start()
	// startService()

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
