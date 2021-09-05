package http

import (
	"axie-notify/services"
	"context"
	"strings"

	"log"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

type HTTPCallBackHanlder struct {
	Bot *linebot.Client
}

// NewServiceHTTPHandler provide the inititail set up service path to handle request
func NewServiceHTTPHandler(e *echo.Echo, linebot *linebot.Client) {

	hanlders := &HTTPCallBackHanlder{Bot: linebot}
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Service is online")
	})

	e.POST("/callback", hanlders.Callback)

}

// Callback provides the function to handle request from line
func (handler *HTTPCallBackHanlder) Callback(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	events, err := handler.Bot.ParseRequest(c.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.String(400, linebot.ErrInvalidSignature.Error())
		} else {
			c.String(500, "internal")
		}
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// data := services.GetAxie()

			// New Flex Message
			// Reply Message
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg := strings.Split(message.Text, " ")
				services.AddQueue(event.Source.UserID, msg[1])
				if msg[0] == "track" {
					// jsonFile, err := os.Open("data/message.json")
					// // if we os.Open returns an error then handle it
					// if err != nil {
					// 	fmt.Println(err)
					// }
					// // defer the closing of our jsonFile so that we can parse it later on
					// defer jsonFile.Close()
					// // Unmarshal JSON
					// byteValue, _ := ioutil.ReadAll(jsonFile)
					// flexContainer, err := linebot.UnmarshalFlexMessageJSON(byteValue)
					// // New Flex Message
					// flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)

					test := services.AxieFlexMessage()
					if _, err = handler.Bot.ReplyMessage(event.ReplyToken, test).Do(); err != nil {
						log.Print(err)
					}
				} else {
					if _, err = handler.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("track")).Do(); err != nil {
						log.Print(err)
					}
				}

			}
		}
	}
	return c.JSON(200, "")
}
