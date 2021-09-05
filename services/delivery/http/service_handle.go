package http

import (
	"axie-notify/models"
	"axie-notify/services"
	"context"
	"fmt"

	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

type HTTPCallBackHanlder struct {
	Bot          *linebot.Client
	ServicesInfo *models.ServicesInfo
}

// NewServiceHTTPHandler provide the inititail set up service path to handle request
func NewServiceHTTPHandler(e *echo.Echo, linebot *linebot.Client, servicesInfo *models.ServicesInfo) {

	hanlders := &HTTPCallBackHanlder{Bot: linebot, ServicesInfo: servicesInfo}
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
			// Make Contents
			var contents []linebot.FlexComponent
			text := linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   "Brown Cafe",
				Weight: "bold",
				Size:   linebot.FlexTextSizeTypeXl,
			}
			contents = append(contents, &text)
			// Make Hero
			hero := linebot.ImageComponent{
				Type:        linebot.FlexComponentTypeImage,
				URL:         "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
				Size:        "full",
				AspectRatio: linebot.FlexImageAspectRatioType20to13,
				AspectMode:  linebot.FlexImageAspectModeTypeCover,
				Action:      linebot.NewMessageAction("left", "left clicked"),
			}
			// Make Body
			body := linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeBox,
				Layout:   linebot.FlexBoxLayoutTypeVertical,
				Contents: contents,
			}
			// Build Container
			bubble := linebot.BubbleContainer{
				Type: linebot.FlexContainerTypeBubble,
				Hero: &hero,
				Body: &body,
			}
			// New Flex Message
			flexMessage := linebot.NewFlexMessage("FlexWithCode", &bubble)
			// Reply Message
			switch message := event.Message.(type) {

			case *linebot.TextMessage:
				messageFromPing := services.PingService(message.Text, handler.ServicesInfo, time.Second*1)
				fmt.Println(messageFromPing)
				if _, err = handler.Bot.ReplyMessage(event.ReplyToken, flexMessage).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
	return c.JSON(200, "")
}
