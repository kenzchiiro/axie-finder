package services

import (
	"axie-notify/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func GetAxie(data *models.Payload) (res []byte) {
	// data := &models.Payload{
	// 	OperationName: `GetAxieBriefList`,
	// 	Query: `query GetAxieBriefList(
	// 		$auctionType: AuctionType,
	// 		$criteria: AxieSearchCriteria,
	// 		$from: Int,
	// 		$sort: SortBy,
	// 		$size: Int,
	// 		$owner: String)
	// 		{axies(
	// 			auctionType: $auctionType,
	// 			criteria: $criteria,
	// 			from: $from,
	// 			sort: $sort,
	// 			size: $size,
	// 			owner: $owner)
	// 			{total
	// 			results {
	// 				...AxieBrief
	// 				__typename  }
	// 			  __typename
	// 			  }      }
	// 			    fragment AxieBrief on Axie {
	// 					id
	// 					name
	// 					stage
	// 					class
	// 					breedCount
	// 					image
	// 					title
	// 					genes
	// 					battleInfo {  banned  __typename}
	// 					auction {
	// 						currentPrice
	// 						currentPriceUSD
	// 						__typename
	// 					}stats {
	// 						...AxieStats  __typename
	// 					}parts {
	// 						id
	// 						name
	// 						class
	// 						type
	// 						specialGenes
	// 						__typename
	// 						}
	// 						__typename
	// 					}
	// 					fragment AxieStats on AxieStats
	// 					{
	// 						hp
	// 						speed
	// 						skill
	// 						morale
	// 						__typename
	// 					}`,
	// 	Variables: models.Variables{
	// 		AuctionType: "Sale",
	// 		Criteria: models.Criteria{
	// 			Classes: []string{"Plant"},
	// 			Parts: []string{
	// 				"mouth-serious",
	// 				"mouth-humorless",
	// 				"horn-little-branch",
	// 				"horn-winter-branch",
	// 				"back-pumpkin",
	// 				"tail-carrot",
	// 				"tail-namek-carrot",
	// 			},
	// 			Hp:         nil,
	// 			Speed:      nil,
	// 			Skill:      nil,
	// 			Morale:     nil,
	// 			BreedCount: nil,
	// 			Pureness:   []int{5, 6},
	// 			NumMystic:  nil,
	// 			Title:      nil,
	// 			Region:     nil,
	// 			Stages:     []int{3, 4},
	// 		},
	// 		From:  0,
	// 		Size:  1,
	// 		Sort:  "PriceAsc",
	// 		Owner: nil,
	// 	},
	// }
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://axieinfinity.com/graphql-server-v2/graphql", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	res, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
	return res
}

func SetVariablesAxie(variables *models.Variables) (res []byte) {
	data := &models.Payload{
		OperationName: `GetAxieBriefList`,
		Query: `query GetAxieBriefList(
			$auctionType: AuctionType,
			$criteria: AxieSearchCriteria,
			$from: Int,
			$sort: SortBy,
			$size: Int,
			$owner: String)
			{axies(
				auctionType: $auctionType,
				criteria: $criteria,
				from: $from,
				sort: $sort,
				size: $size,
				owner: $owner)
				{total
				results {
					...AxieBrief
					__typename  }
				  __typename
				  }      }
				    fragment AxieBrief on Axie {
						id
						name
						stage
						class
						breedCount
						image
						title
						genes
						battleInfo {  banned  __typename}
						auction {
							currentPrice
							currentPriceUSD
							__typename
						}stats {
							...AxieStats  __typename
						}parts {
							id
							name
							class
							type
							specialGenes
							__typename
							}
							__typename
						}
						fragment AxieStats on AxieStats
						{
							hp
							speed
							skill
							morale
							__typename
						}`,
		Variables: *variables,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://axieinfinity.com/graphql-server-v2/graphql", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	res, _ = ioutil.ReadAll(resp.Body)

	axies := models.Data{}
	json.Unmarshal(res, &axies)

	for _, v := range axies.Axies.Results {
		fmt.Println(v.Auction)
		fmt.Println(v.Name)
		fmt.Println(v.Stats)
		fmt.Println(v.ID)
		fmt.Println(v.Genes)
		fmt.Println(v.Class)
		fmt.Println(v.Image)
		fmt.Println(v.Parts)
	}
	return res
}

func SetParameterAxieFromMessage(params string) (data []byte) {
	param := strings.Split(params, ";")

	_type := strings.Split(param[0], ",")
	_part := strings.Split(param[1], ",")

	variables := models.Variables{
		AuctionType: "Sale",
		Criteria: models.Criteria{
			Classes:    _type,
			Parts:      _part,
			Hp:         nil,
			Speed:      nil,
			Skill:      nil,
			Morale:     nil,
			BreedCount: nil,
			Pureness:   []int{1, 2, 3, 4, 5, 6},
			NumMystic:  nil,
			Title:      nil,
			Region:     nil,
			Stages:     []int{3, 4},
		},
		From:  0,
		Size:  1,
		Sort:  "PriceAsc",
		Owner: nil,
	}
	data = SetVariablesAxie(&variables)
	return
}

func SetAxieToFlexMessage() (flexMessage *linebot.FlexMessage) {
	// Open our jsonFile
	jsonFile, err := os.Open("data/message.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// Unmarshal JSON
	byteValue, _ := ioutil.ReadAll(jsonFile)
	flexContainer, err := linebot.UnmarshalFlexMessageJSON(byteValue)
	// New Flex Message
	flexMessage = linebot.NewFlexMessage("FlexWithJSON", flexContainer)

	return
}

func AddQueue(userID, msg string) (err error) {
	_queueList := make([]models.Queue, 0)
	_queue := models.Queue{
		Name:    userID,
		Command: msg,
	}

	_queueList = append(_queueList, _queue)
	file, err := json.MarshalIndent(_queueList, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile("data/queue.json", file, 0644)
	return
}

func NewAxieFlexMessageTemplate() (flexMessage *linebot.FlexMessage) {
	// Make Hero
	hero := linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		URL:         "https://storage.googleapis.com/assets.axieinfinity.com/axies/15326/axie/axie-full-transparent.png",
		Size:        "lg",
		AspectRatio: linebot.FlexImageAspectRatioType20to13,
		AspectMode:  linebot.FlexImageAspectModeTypeCover,
		Action:      linebot.NewMessageAction("left", "left clicked"),
	}

	// Make Title
	var titleContents []linebot.FlexComponent
	var span []*linebot.SpanComponent
	titleText := linebot.TextComponent{
		Type:     linebot.FlexComponentTypeText,
		Text:     "AXIE NAME",
		Weight:   "bold",
		Size:     linebot.FlexTextSizeTypeLg,
		Contents: span,
	}
	titleContents = append(titleContents, &titleText)
	titleIcon := linebot.IconComponent{
		Type: linebot.FlexComponentTypeIcon,
		URL:  "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_beast.svg",
	}
	titleContents = append(titleContents, &titleIcon)
	titleBox := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: titleContents,
	}
	var contentsBoxGroup []linebot.FlexComponent
	contentsBoxGroup = append(contentsBoxGroup, &titleBox)

	titleBoxGroup := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: contentsBoxGroup,
	}

	var bodyContents []linebot.FlexComponent
	bodyContents = append(bodyContents, &titleBoxGroup)
	bodyContents = append(bodyContents, &linebot.SeparatorComponent{})

	//dd
	var statContentsBaseline []linebot.FlexComponent

	statIcon := linebot.IconComponent{
		Type: linebot.FlexComponentTypeIcon,
		URL:  "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/stat_health.png?alt=media&token=c928f31f-54c5-4828-a414-de680b6a0e25",
	}
	span = append(span, &linebot.SpanComponent{Type: linebot.FlexComponentTypeSpan, Text: "27"})
	statText := linebot.TextComponent{
		Type:     linebot.FlexComponentTypeText,
		Text:     "HP",
		Weight:   "bold",
		Align:    "start",
		Size:     linebot.FlexTextSizeTypeLg,
		Margin:   linebot.FlexComponentMarginTypeSm,
		Contents: span,
	}

	// span = append(span, &linebot.SpanComponent{Type: linebot.FlexComponentTypeSpan, Text: "27"})
	// statValue := linebot.TextComponent{
	// 	Type:     linebot.FlexComponentTypeText,
	// 	Text:     "27",
	// 	Align:    "start",
	// 	Size:     linebot.FlexTextSizeTypeLg,
	// 	Margin:   linebot.FlexComponentMarginTypeSm,
	// 	Contents: span,
	// }
	statContentsBaseline = append(statContentsBaseline, &statIcon)
	statContentsBaseline = append(statContentsBaseline, &statText)
	// statContentsBaseline = append(statContentsBaseline, &statValue)

	bodyBaseline := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: statContentsBaseline,
	}

	var statContents []linebot.FlexComponent
	statContents = append(statContents, &bodyBaseline)

	bodyBox := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Spacing:  linebot.FlexComponentSpacingTypeSm,
		Contents: statContents,
	}
	bodyContents = append(bodyContents, &bodyBox)

	// Make Body
	body := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: bodyContents,
	}

	// Make Contents Footer
	var contentsFooter []linebot.FlexComponent
	// Make Footer
	button := linebot.ButtonComponent{
		Type: linebot.FlexComponentTypeButton,
		Action: &linebot.URIAction{
			Label: "VIEW",
			URI:   "https://linecorp.com",
		},
		Color: "#40C9ABFF",
		Style: "primary",
	}
	contentsFooter = append(contentsFooter, &button)

	// Make Footer
	footer := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: contentsFooter,
	}
	// Build Container
	bubble := linebot.BubbleContainer{
		Type:   linebot.FlexContainerTypeBubble,
		Hero:   &hero,
		Body:   &body,
		Footer: &footer,
	}

	bubbleList := []*linebot.BubbleContainer{}
	bubbleList = append(bubbleList, &bubble)
	bubbleList = append(bubbleList, &bubble)
	bubbleList = append(bubbleList, &bubble)

	carousal := linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeBubble,
		Contents: bubbleList}
	// New Flex Message
	flexMessage = linebot.NewFlexMessage("FlexWithCode", &carousal)
	return
}
