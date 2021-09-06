package services

import (
	"axie-notify/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/shopspring/decimal"
)

const (
	CLASS_PLANT_ICON   = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_plant.png?alt=media&token=6de5e7f3-7af5-493d-a753-ae4cfa26ffdf"
	CLASS_AQUATIC_ICON = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_aquatic.png?alt=media&token=a1feecca-9f1c-44bf-8171-b2080be6c599"
	CLASS_BEAST_ICON   = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_beast.png?alt=media&token=16bcd963-0ce9-4410-8373-fc820491cdb1"
	CLASS_BIRD_ICON    = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_bird.png?alt=media&token=51135e98-ee39-48bf-84be-ecbe74df4904"
	CLASS_BUG_ICON     = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_bug.png?alt=media&token=e7587bb2-add4-4640-acf1-655392b206d5"
	CLASS_REPTILE_ICON = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_reptile.png?alt=media&token=3083c354-6cba-4a90-b0cd-526f55618c31"
	CLASS_DAWN_ICON    = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_dawn.png?alt=media&token=7ad6f189-7519-49eb-8a8f-ae05b35d7b7c"
	CLASS_DUSK_ICON    = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_dusk.png?alt=media&token=2987be86-23cd-4bd6-aa8f-d51e4e126058"
	CLASS_MECH_ICON    = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/class_mech.png?alt=media&token=60b725f5-cbc4-47a7-a1c0-ec1ed71826ed"
	ETH_ICON           = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/eth.png?alt=media&token=587218db-28b2-41e7-b200-d583038112f7"
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

func SetVariablesAxie(variables *models.Variables) (result *models.DataRespone) {
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
	res, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(res, &result)
	return result
}

func SetParameterAxieFromMessage(params string) (result *models.DataRespone) {
	param := strings.Split(params, ";")

	_type := strings.Split(param[0], ",")
	_part := strings.Split(param[1], ",")
	_limit, _ := strconv.Atoi(param[2])

	fmt.Println("type", _type)
	fmt.Println("part", _part)

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
		Size:  _limit,
		Sort:  "PriceAsc",
		Owner: nil,
	}
	result = SetVariablesAxie(&variables)
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

func SetAxieFlexMessage(axieData *models.Results) (bubble *linebot.BubbleContainer) {
	// Make Hero
	hero := linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		URL:         axieData.Image,
		Size:        "lg",
		AspectRatio: linebot.FlexImageAspectRatioType20to13,
		AspectMode:  linebot.FlexImageAspectModeTypeCover,
		// Action:      linebot.NewMessageAction("left", "left clicked"),
	}

	// Make Title
	var titleContents []linebot.FlexComponent
	var span []*linebot.SpanComponent
	titleText := linebot.TextComponent{
		Type:     linebot.FlexComponentTypeText,
		Text:     axieData.Name,
		Weight:   "bold",
		Size:     linebot.FlexTextSizeTypeXs,
		Contents: span,
	}
	url_class := ""
	if axieData.Class == "Plant" {
		url_class = CLASS_PLANT_ICON
	} else if axieData.Class == "Beast" {
		url_class = CLASS_BEAST_ICON

	} else if axieData.Class == "Bird" {
		url_class = CLASS_BIRD_ICON

	} else if axieData.Class == "Aquatic" {
		url_class = CLASS_AQUATIC_ICON

	} else if axieData.Class == "Reptile" {
		url_class = CLASS_PLANT_ICON

	} else if axieData.Class == "Dawn" {
		url_class = CLASS_DAWN_ICON

	} else if axieData.Class == "Dusk" {
		url_class = CLASS_DUSK_ICON
	} else {
		url_class = CLASS_MECH_ICON
	}
	titleIcon := linebot.IconComponent{
		Type: linebot.FlexComponentTypeIcon,
		Size: linebot.FlexIconSizeTypeSm,
		URL:  url_class,
	}
	titleContents = append(titleContents, &titleText)
	titleContents = append(titleContents, &titleIcon)
	titleBox := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: titleContents,
	}
	var contentsBoxGroup []linebot.FlexComponent
	contentsBoxGroup = append(contentsBoxGroup, &titleBox)

	// Make Title ID
	var titleContentsID []linebot.FlexComponent
	titleTextID := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "#" + axieData.ID,
		Weight: "bold",
		Color:  "#AAAAAA",
		Size:   linebot.FlexTextSizeTypeXxs,
	}

	titleContentsID = append(titleContentsID, &titleTextID)
	titleBoxID := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: titleContentsID,
	}

	contentsBoxGroup = append(contentsBoxGroup, &titleBoxID)
	titleBoxGroup := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: contentsBoxGroup,
	}

	var bodyContents []linebot.FlexComponent
	bodyContents = append(bodyContents, &linebot.SeparatorComponent{})

	//HP
	var statContents []linebot.FlexComponent
	var statContentsBaselineHP []linebot.FlexComponent

	statIconHP := linebot.IconComponent{
		Type: linebot.FlexComponentTypeIcon,
		URL:  "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/stat_health.png?alt=media&token=c928f31f-54c5-4828-a414-de680b6a0e25",
	}
	statTextHP := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "HP",
		Weight: "bold",
		Align:  "start",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}

	statValueHP := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   strconv.Itoa(axieData.Stats.Hp),
		Align:  "end",
		Color:  "#AAAAAA",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}
	statContentsBaselineHP = append(statContentsBaselineHP, &statIconHP)
	statContentsBaselineHP = append(statContentsBaselineHP, &statTextHP)
	statContentsBaselineHP = append(statContentsBaselineHP, &statValueHP)

	bodyBaselineHP := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: statContentsBaselineHP,
	}

	statContents = append(statContents, &bodyBaselineHP)
	//----------------------- end hp ----------------------

	//SPD
	var statContentsBaselineSPD []linebot.FlexComponent

	statIconSPD := linebot.IconComponent{
		Type: linebot.FlexComponentTypeIcon,
		URL:  "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/stat_speed.png?alt=media&token=5c285c92-bab9-4cfa-a43d-bfdb7e1ae0a8",
	}
	statTextSPD := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "SPEED",
		Weight: "bold",
		Align:  "start",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}

	statValueSPD := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   strconv.Itoa(axieData.Stats.Speed),
		Align:  "end",
		Color:  "#AAAAAA",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}
	statContentsBaselineSPD = append(statContentsBaselineSPD, &statIconSPD)
	statContentsBaselineSPD = append(statContentsBaselineSPD, &statTextSPD)
	statContentsBaselineSPD = append(statContentsBaselineSPD, &statValueSPD)

	bodyBaselineSPD := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: statContentsBaselineSPD,
	}
	statContents = append(statContents, &bodyBaselineSPD)
	//----------------------- end spd ----------------------
	//SKL
	var statContentsBaselineSKL []linebot.FlexComponent

	statIconSKL := linebot.IconComponent{
		Type: linebot.FlexComponentTypeIcon,
		URL:  "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/stat_skill.png?alt=media&token=51bec4cd-00b5-4a8b-bca9-5bff9c17d5e4",
	}
	statTextSKL := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "SKILL",
		Weight: "bold",
		Align:  "start",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}

	statValueSKL := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   strconv.Itoa(axieData.Stats.Skill),
		Align:  "end",
		Color:  "#AAAAAA",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}
	statContentsBaselineSKL = append(statContentsBaselineSKL, &statIconSKL)
	statContentsBaselineSKL = append(statContentsBaselineSKL, &statTextSKL)
	statContentsBaselineSKL = append(statContentsBaselineSKL, &statValueSKL)

	bodyBaselineSKL := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: statContentsBaselineSKL,
	}
	statContents = append(statContents, &bodyBaselineSKL)
	//----------------------- end skill ----------------------
	//MR
	var statContentsBaselineMR []linebot.FlexComponent

	statIconMR := linebot.IconComponent{
		Type: linebot.FlexComponentTypeIcon,
		URL:  "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/stat_morale.png?alt=media&token=56733460-caa6-406f-b380-49f63485958e",
	}
	statTextMR := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "MORALE",
		Weight: "bold",
		Align:  "start",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}

	statValueMR := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   strconv.Itoa(axieData.Stats.Morale),
		Align:  "end",
		Color:  "#AAAAAA",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}
	statContentsBaselineMR = append(statContentsBaselineMR, &statIconMR)
	statContentsBaselineMR = append(statContentsBaselineMR, &statTextMR)
	statContentsBaselineMR = append(statContentsBaselineMR, &statValueMR)

	bodyBaselineMR := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: statContentsBaselineMR,
	}
	statContents = append(statContents, &bodyBaselineMR)
	statContents = append(statContents, &linebot.SpacerComponent{})

	//----------------------- end mr ----------------------

	bodyBox := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Spacing:  linebot.FlexComponentSpacingTypeSm,
		Contents: statContents,
	}
	bodyContents = append(bodyContents, &bodyBox)

	//----------------------- end stat ----------------------
	bodyContents = append(bodyContents, &linebot.SeparatorComponent{})

	// Make Price

	var priceContents []linebot.FlexComponent
	var priceContentsBaseline []linebot.FlexComponent

	statIconETH := linebot.IconComponent{
		Size: linebot.FlexIconSizeTypeXxs,
		Type: linebot.FlexComponentTypeIcon,
		URL:  ETH_ICON,
	}

	price, err := decimal.NewFromString(axieData.Auction.CurrentPrice)
	if err != nil {
		fmt.Println(err)
	}
	div, err := decimal.NewFromString("1000000000000000000")
	if err != nil {
		fmt.Println(err)
	}
	price = price.Div(div)

	priceValue := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   price.String(),
		Align:  "start",
		Weight: "bold",
		Size:   linebot.FlexTextSizeTypeXs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}

	priceText := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "WETH",
		Align:  "end",
		Color:  "#AAAAAA",
		Size:   linebot.FlexTextSizeTypeXs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}
	priceContentsBaseline = append(priceContentsBaseline, &statIconETH)
	priceContentsBaseline = append(priceContentsBaseline, &priceValue)
	priceContentsBaseline = append(priceContentsBaseline, &priceText)

	priceBaseline := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: priceContentsBaseline,
	}

	priceContents = append(priceContents, &linebot.SpacerComponent{})
	priceContents = append(priceContents, &priceBaseline)

	//----------------------- end price ----------------------
	bodyBoxPrice := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Spacing:  linebot.FlexComponentSpacingTypeSm,
		Contents: priceContents,
	}

	bodyContents = append(bodyContents, &bodyBoxPrice)
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
		Type:   linebot.FlexComponentTypeButton,
		Height: linebot.FlexButtonHeightTypeSm,
		Action: &linebot.URIAction{
			Label: "VIEW",
			URI:   "https://marketplace.axieinfinity.com/axie/" + axieData.ID,
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
	_bubble := linebot.BubbleContainer{
		Header:    &titleBoxGroup,
		Type:      linebot.FlexContainerTypeBubble,
		Direction: linebot.FlexBubbleDirectionTypeLTR,
		Size:      linebot.FlexBubbleSizeTypeMicro,
		Hero:      &hero,
		Body:      &body,
		Footer:    &footer,
	}

	bubble = &_bubble
	return
}

func SetAxieToFlexMessage(axiesData *models.DataRespone) (flexMessage *linebot.FlexMessage) {
	bubbleList := []*linebot.BubbleContainer{}
	for _, axie := range axiesData.Data.Axies.Results {
		bubble := SetAxieFlexMessage(&axie)
		bubbleList = append(bubbleList, bubble)
	}
	carousal := linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeBubble,
		Contents: bubbleList}
	// New Flex Message
	flexMessage = linebot.NewFlexMessage("FlexWithCode", &carousal)
	return
}

// func SetAxieToFlexMessage() (flexMessage *linebot.FlexMessage) {
// 	// Open our jsonFile
// 	jsonFile, err := os.Open("data/message.json")
// 	// if we os.Open returns an error then handle it
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// defer the closing of our jsonFile so that we can parse it later on
// 	defer jsonFile.Close()
// 	// Unmarshal JSON
// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	flexContainer, err := linebot.UnmarshalFlexMessageJSON(byteValue)
// 	// New Flex Message
// 	flexMessage = linebot.NewFlexMessage("FlexWithJSON", flexContainer)

// 	return
// }
