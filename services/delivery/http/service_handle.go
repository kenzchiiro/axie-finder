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
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg := strings.Split(message.Text, " ")
				err := services.AddQueue(event.Source.UserID, message.Text)
				if msg[0] == "#find" {
					axiesData := services.SetParameterAxieFromMessage(msg[1])
					flexMessage := services.SetAxieToFlexMessage(axiesData)

					if _, err = handler.Bot.ReplyMessage(event.ReplyToken, flexMessage).Do(); err != nil {
						log.Print(err)
					}
				} else {
					if _, err = handler.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("command invalid")).Do(); err != nil {
						log.Print(err)
					}
				}

			}
		}
	}
	return c.JSON(200, "")
}
