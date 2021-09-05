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

func AxieFlexMessage() (flexMessage *linebot.FlexMessage) {
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
