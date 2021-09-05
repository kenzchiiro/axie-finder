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
			// var contents []linebot.FlexComponent
			// text := linebot.TextComponent{
			// 	Type:   linebot.FlexComponentTypeText,
			// 	Text:   "Brown Cafe",
			// 	Weight: "bold",
			// 	Size:   linebot.FlexTextSizeTypeSm,
			// }
			// contents = append(contents, &text)
			// contents = append(contents, &text)
			// contents = append(contents, &text)

			// // Make Hero
			// hero := linebot.ImageComponent{
			// 	Type:        linebot.FlexComponentTypeImage,
			// 	URL:         "https://storage.googleapis.com/assets.axieinfinity.com/axies/15326/axie/axie-full-transparent.png",
			// 	Size:        "full",
			// 	AspectRatio: linebot.FlexImageAspectRatioType20to13,
			// 	AspectMode:  linebot.FlexImageAspectModeTypeCover,
			// 	Action:      linebot.NewMessageAction("left", "left clicked"),
			// }
			// // Make Body
			// body := linebot.BoxComponent{
			// 	Type:     linebot.FlexComponentTypeBox,
			// 	Layout:   linebot.FlexBoxLayoutTypeVertical,
			// 	Contents: contents,
			// }

			// // Make Body

			// footer := linebot.BoxComponent{
			// 	Type:     linebot.FlexComponentTypeButton,
			// 	Layout:   linebot.FlexBoxLayoutTypeVertical,
			// 	Contents: contents,
			// }

			// // Build Container
			// bubble := linebot.BubbleContainer{
			// 	Type:   linebot.FlexContainerTypeBubble,
			// 	Hero:   &hero,
			// 	Body:   &body,
			// 	Footer: &footer,
			// }

			// Unmarshal JSON
			flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(`{
				"type": "bubble",
				"direction": "ltr",
				"hero": {
				  "type": "image",
				  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_2_restaurant.png",
				  "size": "full",
				  "aspectRatio": "20:13",
				  "aspectMode": "cover",
				  "action": {
					"type": "uri",
					"label": "Action",
					"uri": "https://linecorp.com"
				  }
				},
				"body": {
				  "type": "box",
				  "layout": "vertical",
				  "spacing": "md",
				  "action": {
					"type": "uri",
					"label": "Action",
					"uri": "https://linecorp.com"
				  },
				  "contents": [
					{
					  "type": "text",
					  "text": "Brown's Burger",
					  "weight": "bold",
					  "size": "xl",
					  "contents": []
					},
					{
					  "type": "box",
					  "layout": "vertical",
					  "spacing": "sm",
					  "contents": [
						{
						  "type": "box",
						  "layout": "baseline",
						  "contents": [
							{
							  "type": "icon",
							  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_regular_32.png"
							},
							{
							  "type": "text",
							  "text": "$10.5",
							  "weight": "bold",
							  "margin": "sm",
							  "contents": []
							},
							{
							  "type": "text",
							  "text": "400kcl",
							  "size": "sm",
							  "color": "#AAAAAA",
							  "align": "end",
							  "contents": [
								{
								  "type": "span",
								  "text": "hello, world"
								}
							  ]
							}
						  ]
						},
						{
						  "type": "box",
						  "layout": "baseline",
						  "contents": [
							{
							  "type": "icon",
							  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_large_32.png"
							},
							{
							  "type": "text",
							  "text": "$15.5",
							  "weight": "bold",
							  "flex": 0,
							  "margin": "sm",
							  "contents": []
							},
							{
							  "type": "text",
							  "text": "550kcl",
							  "size": "sm",
							  "color": "#AAAAAA",
							  "align": "end",
							  "contents": []
							}
						  ]
						}
					  ]
					},
					{
					  "type": "text",
					  "text": "Sauce, Onions, Pickles, Lettuce & Cheese",
					  "size": "xxs",
					  "color": "#AAAAAA",
					  "wrap": true,
					  "contents": []
					},
					{
					  "type": "box",
					  "layout": "vertical",
					  "spacing": "sm",
					  "contents": [
						{
						  "type": "box",
						  "layout": "baseline",
						  "contents": [
							{
							  "type": "icon",
							  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_regular_32.png"
							},
							{
							  "type": "text",
							  "text": "$10.5",
							  "weight": "bold",
							  "margin": "sm",
							  "contents": []
							},
							{
							  "type": "text",
							  "text": "400kcl",
							  "size": "sm",
							  "color": "#AAAAAA",
							  "align": "end",
							  "contents": [
								{
								  "type": "span",
								  "text": "hello, world"
								}
							  ]
							}
						  ]
						},
						{
						  "type": "box",
						  "layout": "baseline",
						  "contents": [
							{
							  "type": "icon",
							  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_large_32.png"
							},
							{
							  "type": "text",
							  "text": "$15.5",
							  "weight": "bold",
							  "flex": 0,
							  "margin": "sm",
							  "contents": []
							},
							{
							  "type": "text",
							  "text": "550kcl",
							  "size": "sm",
							  "color": "#AAAAAA",
							  "align": "end",
							  "contents": []
							}
						  ]
						}
					  ]
					},
					{
					  "type": "box",
					  "layout": "vertical",
					  "spacing": "sm",
					  "contents": [
						{
						  "type": "box",
						  "layout": "baseline",
						  "contents": [
							{
							  "type": "icon",
							  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_regular_32.png"
							},
							{
							  "type": "text",
							  "text": "$10.5",
							  "weight": "bold",
							  "margin": "sm",
							  "contents": []
							},
							{
							  "type": "text",
							  "text": "400kcl",
							  "size": "sm",
							  "color": "#AAAAAA",
							  "align": "end",
							  "contents": [
								{
								  "type": "span",
								  "text": "hello, world"
								}
							  ]
							}
						  ]
						},
						{
						  "type": "box",
						  "layout": "baseline",
						  "contents": [
							{
							  "type": "icon",
							  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_large_32.png"
							},
							{
							  "type": "text",
							  "text": "$15.5",
							  "weight": "bold",
							  "flex": 0,
							  "margin": "sm",
							  "contents": []
							},
							{
							  "type": "text",
							  "text": "550kcl",
							  "size": "sm",
							  "color": "#AAAAAA",
							  "align": "end",
							  "contents": []
							}
						  ]
						}
					  ]
					}
				  ]
				},
				"footer": {
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "spacer"
					},
					{
					  "type": "button",
					  "action": {
						"type": "uri",
						"label": "Add to Cart",
						"uri": "https://linecorp.com"
					  },
					  "color": "#905C44",
					  "style": "primary"
					}
				  ]
				}
			  }`))
			if err != nil {
				log.Println(err)
			}
			// New Flex Message
			flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)

			// New Flex Message
			// flexMessage := linebot.NewFlexMessage("FlexWithCode", &bubble)
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
