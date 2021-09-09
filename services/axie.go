package services

import (
	"axie-notify/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	for i, v := range _type {
		_type[i] = strings.Title(strings.ToLower(v))
	}
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

	cmd := strings.Split(msg, " ")
	_queue := models.Queue{
		Command:   cmd[0],
		Parameter: cmd[1],
	}

	queueList[userID] = _queue

	file, err := json.MarshalIndent(queueList, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile("data/queue.json", file, 0644)
	return
}

func SetAxieFlexMessage(axieData *models.Results) (bubble *linebot.BubbleContainer) {
	decimal.DivisionPrecision = 4
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

	//Part
	var partContentsBaseline []linebot.FlexComponent

	partEyes := models.IconComponent{
		Type:        linebot.FlexComponentTypeIcon,
		AspectRatio: linebot.FlexIconAspectRatioType2to1,
		URL:         FindPartAxieIcon(&axieData.Parts[0]),
		OffsetStart: "24px",
	}
	partEars := models.IconComponent{
		Type:        linebot.FlexComponentTypeIcon,
		Size:        linebot.FlexIconSizeTypeMd,
		AspectRatio: linebot.FlexIconAspectRatioType2to1,
		URL:         FindPartAxieIcon(&axieData.Parts[1]),
		OffsetStart: "12px",
	}
	partBack := models.IconComponent{
		Type:        linebot.FlexComponentTypeIcon,
		Size:        linebot.FlexIconSizeTypeMd,
		AspectRatio: linebot.FlexIconAspectRatioType2to1,
		URL:         FindPartAxieIcon(&axieData.Parts[2]),
		OffsetEnd:   "6px",
	}
	partMouth := models.IconComponent{
		Type:        linebot.FlexComponentTypeIcon,
		Size:        linebot.FlexIconSizeTypeMd,
		AspectRatio: linebot.FlexIconAspectRatioType2to1,
		URL:         FindPartAxieIcon(&axieData.Parts[3]),
		OffsetEnd:   "12px",
	}
	partHorn := models.IconComponent{
		Type:        linebot.FlexComponentTypeIcon,
		AspectRatio: linebot.FlexIconAspectRatioType2to1,
		URL:         FindPartAxieIcon(&axieData.Parts[4]),
		OffsetEnd:   "24px",
	}
	partTail := models.IconComponent{
		Type:        linebot.FlexComponentTypeIcon,
		AspectRatio: linebot.FlexIconAspectRatioType2to1,
		URL:         FindPartAxieIcon(&axieData.Parts[5]),
		OffsetEnd:   "48px",
	}

	partContentsBaseline = append(partContentsBaseline, &linebot.FillerComponent{})
	partContentsBaseline = append(partContentsBaseline, &partEyes)
	partContentsBaseline = append(partContentsBaseline, &partEars)
	partContentsBaseline = append(partContentsBaseline, &partBack)
	partContentsBaseline = append(partContentsBaseline, &partMouth)
	partContentsBaseline = append(partContentsBaseline, &partHorn)
	partContentsBaseline = append(partContentsBaseline, &partTail)
	partContentsBaseline = append(partContentsBaseline, &linebot.FillerComponent{})

	partGroup := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: partContentsBaseline,
	}

	// end part

	var bodyContents []linebot.FlexComponent

	bodyContents = append(bodyContents, &partGroup)

	//HP
	var statContents []linebot.FlexComponent
	statContents = append(statContents, &linebot.SpacerComponent{})
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

	// Make Price ETH

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

	bodyBoxPrice := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Spacing:  linebot.FlexComponentSpacingTypeSm,
		Contents: priceContents,
	}

	bodyContents = append(bodyContents, &bodyBoxPrice)

	//----------------------- end price ----------------------

	// Make Price USD

	var priceContentsUSD []linebot.FlexComponent
	var priceContentsBaselineUSD []linebot.FlexComponent

	priceValueUSD := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "$ " + axieData.Auction.CurrentPriceUSD,
		Align:  "start",
		Color:  "#AAAAAA",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}

	priceTextUSD := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "USD",
		Align:  "end",
		Color:  "#AAAAAA",
		Size:   linebot.FlexTextSizeTypeXxs,
		Margin: linebot.FlexComponentMarginTypeSm,
	}
	priceContentsBaselineUSD = append(priceContentsBaselineUSD, &priceValueUSD)
	priceContentsBaselineUSD = append(priceContentsBaselineUSD, &priceTextUSD)

	priceBaselineUSD := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeBaseline,
		Contents: priceContentsBaselineUSD,
	}
	priceContentsUSD = append(priceContentsUSD, &priceBaselineUSD)

	bodyBoxPriceUSD := linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Spacing:  linebot.FlexComponentSpacingTypeSm,
		Contents: priceContentsUSD,
	}

	bodyContents = append(bodyContents, &bodyBoxPriceUSD)

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
		Header: &titleBoxGroup,
		Type:   linebot.FlexContainerTypeBubble,
		Size:   linebot.FlexBubbleSizeTypeMicro,
		Hero:   &hero,
		Body:   &body,
		Footer: &footer,
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

func FindPartAxieIcon(part *models.Parts) (iconURL string) {
	if part.Class == "Plant" {
		if part.Type == "Mouth" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_mouth_plant.png?alt=media&token=4f387978-2d39-46cb-82d7-f7136868b9d2"
		} else if part.Type == "Back" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_back_plant.png?alt=media&token=56a4830a-0bc2-47e0-a7cb-b654e6743c98"
		} else if part.Type == "Eyes" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_eyes_plant.png?alt=media&token=bdfaf303-6155-479f-820b-4d1e6ef161b8"
		} else if part.Type == "Ears" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_ears_plant.png?alt=media&token=58e9813a-e84e-4fa6-b84f-e47156cae04b"
		} else if part.Type == "Horn" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_horn_plant.png?alt=media&token=56c3a8aa-06c9-4f93-bac8-98ff4d26dec7"
		} else {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_tail_plant.png?alt=media&token=20ffae72-b0f2-421f-b8f7-bee0d72097fc"
		}
	} else if part.Class == "Aquatic" {
		if part.Type == "Mouth" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_mouth_aquatic.png?alt=media&token=5229f995-95cc-45e8-b68f-b151db60839a"
		} else if part.Type == "Back" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_back_aquatic.png?alt=media&token=28123caa-7cfc-4cc8-a8af-cf375f0889b1"
		} else if part.Type == "Eyes" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_eyes_aquatic.png?alt=media&token=24dd2b63-5dc0-4f94-83e5-5c202ebf08d6"
		} else if part.Type == "Ears" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_ears_aquatic.png?alt=media&token=67e0cc9d-33da-46da-9712-864f3b22f1b7"
		} else if part.Type == "Horn" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_horn_aquatic.png?alt=media&token=51a4f3c3-5f89-453f-819b-6723e4932c32"
		} else {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_tail_aquatic.png?alt=media&token=07192a3b-4741-41c9-877d-2cc7f7dfe633"
		}
	} else if part.Class == "Beast" {
		if part.Type == "Mouth" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_mouth_beast.png?alt=media&token=f212d08b-1740-4c4d-a55e-48f4e4d21127"
		} else if part.Type == "Back" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_back_beast.png?alt=media&token=34f1171f-2617-4f1d-85f6-bfd873da688b"
		} else if part.Type == "Eyes" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_eyes_beast.png?alt=media&token=cbe1f930-82da-4d9d-ab56-6035a6812cd5"
		} else if part.Type == "Ears" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_ears_beast.png?alt=media&token=9b3f861b-db05-441e-a0f9-377eb2538f7a"
		} else if part.Type == "Horn" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_horn_beast.png?alt=media&token=8f5ff055-926f-49fe-827e-9ed518eb9679"
		} else {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_tail_beast.png?alt=media&token=2340c100-c937-4121-84a0-f6d9d5f88864"
		}
	} else if part.Class == "Bird" {
		if part.Type == "Mouth" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_mouth_bird.png?alt=media&token=6c341688-7d2a-491f-a8e2-0393b5da610c"
		} else if part.Type == "Back" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_back_bird.png?alt=media&token=a8572ee0-f98a-4fa4-9a3c-9f81f22bbe94"
		} else if part.Type == "Eyes" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_eyes_bird.png?alt=media&token=1c5adee4-8ff9-428f-aeab-72dfca6e0cee"
		} else if part.Type == "Ears" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_ears_bird.png?alt=media&token=38d8620c-cd2f-4725-8c33-a6fa1fcc4b2c"
		} else if part.Type == "Horn" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_horn_bird.png?alt=media&token=0f41cd7e-9d66-41ff-8802-e556042ec5d0"
		} else {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_tail_bird.png?alt=media&token=230038ed-e6b1-43b0-9501-20b62b9d1fe3"
		}
	} else if part.Class == "Bug" {
		if part.Type == "Mouth" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_mouth_bug.png?alt=media&token=44a749ca-ea36-434d-9505-b0d272471422"
		} else if part.Type == "Back" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_back_bug.png?alt=media&token=545bd980-5ec3-452e-819e-1a65d827202a"
		} else if part.Type == "Eyes" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_eyes_bug.png?alt=media&token=b7ddb5ae-c0dc-4c4e-8fa4-489ffac4b1c3"
		} else if part.Type == "Ears" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_ears_bug.png?alt=media&token=a9893f8c-c71c-4bff-9de2-2c56e9548da1"
		} else if part.Type == "Horn" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_horn_bug.png?alt=media&token=4d7262c5-8aa8-40ec-a802-c132609d15f6"
		} else {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_tail_bug.png?alt=media&token=5f032df5-70df-468a-b8ba-eed709143e5f"
		}
	} else {
		if part.Type == "Mouth" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_mouth_reptile.png?alt=media&token=4fc0c8bd-b6be-4604-8932-9eeaa5af8045"
		} else if part.Type == "Back" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_back_reptile.png?alt=media&token=2a9922d9-d29d-435f-b196-f1610026c2eb"
		} else if part.Type == "Eyes" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_eyes_reptile.png?alt=media&token=008eb036-390c-4ed3-8b46-e5882f2ef83c"
		} else if part.Type == "Ears" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_ears_reptile.png?alt=media&token=16d878bd-dd45-48e1-8853-fcad2427d8dc"
		} else if part.Type == "Horn" {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_horn_reptile.png?alt=media&token=ef636ede-af20-4fbe-a47d-53d7a7a74e5a"
		} else {
			iconURL = "https://firebasestorage.googleapis.com/v0/b/filestore-1d8e6.appspot.com/o/part_tail_reptile.png?alt=media&token=dc1f5d89-1e89-4201-a832-d2c4a3ce2067"
		}
	}
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
