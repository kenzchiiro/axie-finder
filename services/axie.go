package services

import (
	"axie-notify/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func GetAxie() (res []byte) {
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
		Variables: models.Variables{
			AuctionType: "Sale",
			Criteria: models.Criteria{
				Classes: []string{"Plant"},
				Parts: []string{
					"mouth-serious",
					"mouth-humorless",
					"horn-little-branch",
					"horn-winter-branch",
					"back-pumpkin",
					"tail-carrot",
					"tail-namek-carrot",
				},
				Hp:         nil,
				Speed:      nil,
				Skill:      nil,
				Morale:     nil,
				BreedCount: nil,
				Pureness:   []int{5, 6},
				NumMystic:  nil,
				Title:      nil,
				Region:     nil,
				Stages:     []int{3, 4},
			},
			From:  0,
			Size:  12,
			Sort:  "PriceAsc",
			Owner: nil,
		},
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
	fmt.Println(string(res))
	return res
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
