package models

import (
	"encoding/json"

	"github.com/line/line-bot-sdk-go/linebot"
)

type ServiceInfo struct {
	IPAddress   string
	Port        string
	ServiceName string
}

type ServicesInfo []ServiceInfo

type Queue struct {
	Command   string `json:"command"`
	Parameter string `json:"parameter"`
}

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
// curl \
// -X POST \
// -H "Content-Type: application/json" \
// --data '{"operationName":"GetAxieBriefList","query":"query GetAxieBriefList($auctionType: AuctionType, $criteria: AxieSearchCriteria, $from: Int, $sort: SortBy, $size: Int, $owner: String) {\naxies(auctionType: $auctionType, criteria: $criteria, from: $from, sort: $sort, size: $size, owner: $owner) {\n  total\n  results {\n    ...AxieBrief\n    __typename\n  }\n  __typename\n}\n      }\n\n      fragment AxieBrief on Axie {\nid\nname\nstage\nclass\nbreedCount\nimage\ntitle\ngenes\nbattleInfo {\n  banned\n  __typename\n}\nauction {\n  currentPrice\n  currentPriceUSD\n  __typename\n}\nstats {\n  ...AxieStats\n  __typename\n}\nparts {\n  id\n  name\n  class\n  type\n  specialGenes\n  __typename\n}\n__typename\n      }\n    \n      fragment AxieStats on AxieStats {\n       hp\n       speed\n       skill\n       morale\n__typename\n      }","variables":{"auctionType":"Sale","criteria":{"classes":["Plant"],"parts":["mouth-serious","mouth-humorless","horn-little-branch","horn-winter-branch","back-pumpkin","tail-carrot","tail-namek-carrot"],"hp":null,"speed":null,"skill":null,"morale":null,"breedCount":null,"pureness":[5,6],"numMystic":[],"title":null,"region":null,"stages":[3,4]},"from":0,"size":12,"sort":"PriceAsc","owner":null}}' \
// https://axieinfinity.com/graphql-server-v2/graphql

type Payload struct {
	OperationName string    `json:"operationName"`
	Query         string    `json:"query"`
	Variables     Variables `json:"variables"`
}
type Criteria struct {
	Classes    []string      `json:"classes"`
	Parts      []string      `json:"parts"`
	Hp         interface{}   `json:"hp"`
	Speed      interface{}   `json:"speed"`
	Skill      interface{}   `json:"skill"`
	Morale     interface{}   `json:"morale"`
	BreedCount interface{}   `json:"breedCount"`
	Pureness   []int         `json:"pureness"`
	NumMystic  []interface{} `json:"numMystic"`
	Title      interface{}   `json:"title"`
	Region     interface{}   `json:"region"`
	Stages     []int         `json:"stages"`
}
type Variables struct {
	AuctionType string      `json:"auctionType"`
	Criteria    Criteria    `json:"criteria"`
	From        int         `json:"from"`
	Size        int         `json:"size"`
	Sort        string      `json:"sort"`
	Owner       interface{} `json:"owner"`
}

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
// --data '{"data":{"axies":{"total":22,"results":[{"id":"5156442","name":"P4-2","stage":4,"class":"Plant","breedCount":3,"image":"https://storage.googleapis.com/assets.axieinfinity.com/axies/5156442/axie/axie-full-transparent.png","title":"","genes":"0x300000000b04a2220c2010ca0c2511020ca328cc002130ca0cc330cc0c2330c2","battleInfo":{"banned":false,"__typename":"AxieBattleInfo"},"auction":{"currentPrice":"108538657407407407","currentPriceUSD":"423.90","__typename":"Auction"},"stats":{"hp":58,"speed":32,"skill":31,"morale":43,"__typename":"AxieStats"},"parts":[{"id":"eyes-papi","name":"Papi","class":"Plant","type":"Eyes","specialGenes":null,"__typename":"AxiePart"},{"id":"ears-hollow","name":"Hollow","class":"Plant","type":"Ears","specialGenes":null,"__typename":"AxiePart"},{"id":"back-pumpkin","name":"Pumpkin","class":"Plant","type":"Back","specialGenes":null,"__typename":"AxiePart"},{"id":"mouth-serious","name":"Serious","class":"Plant","type":"Mouth","specialGenes":null,"__typename":"AxiePart"},{"id":"horn-little-branch","name":"Little Branch","class":"Beast","type":"Horn","specialGenes":null,"__typename":"AxiePart"},{"id":"tail-carrot","name":"Carrot","class":"Plant","type":"Tail","specialGenes":null,"__typename":"AxiePart"}],"__typename":"Axie"}],"__typename":"Axies"}}}' \

type BattleInfo struct {
	Banned   bool   `json:"banned"`
	Typename string `json:"__typename"`
}
type Auction struct {
	CurrentPrice    string `json:"currentPrice"`
	CurrentPriceUSD string `json:"currentPriceUSD"`
	Typename        string `json:"__typename"`
}
type Stats struct {
	Hp       int    `json:"hp"`
	Speed    int    `json:"speed"`
	Skill    int    `json:"skill"`
	Morale   int    `json:"morale"`
	Typename string `json:"__typename"`
}
type Parts struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Class        string      `json:"class"`
	Type         string      `json:"type"`
	SpecialGenes interface{} `json:"specialGenes"`
	Typename     string      `json:"__typename"`
}
type Results struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Stage      int        `json:"stage"`
	Class      string     `json:"class"`
	BreedCount int        `json:"breedCount"`
	Image      string     `json:"image"`
	Title      string     `json:"title"`
	Genes      string     `json:"genes"`
	BattleInfo BattleInfo `json:"battleInfo"`
	Auction    Auction    `json:"auction"`
	Stats      Stats      `json:"stats"`
	Parts      []Parts    `json:"parts"`
	Typename   string     `json:"__typename"`
}
type Axies struct {
	Total    int       `json:"total"`
	Results  []Results `json:"results"`
	Typename string    `json:"__typename"`
}
type Data struct {
	Axies Axies `json:"axies"`
}

type DataRespone struct {
	Data Data `json:"data"`
}

// IconComponent type
type IconComponentCustom struct {
	Type        linebot.FlexComponentType
	URL         string
	Margin      linebot.FlexComponentMarginType
	Size        linebot.FlexIconSizeType
	AspectRatio linebot.FlexIconAspectRatioType
	OffsetStart string
	OffsetEnd   string
}

// MarshalJSON method of IconComponent
func (c *IconComponentCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type        linebot.FlexComponentType       `json:"type"`
		URL         string                          `json:"url"`
		Margin      linebot.FlexComponentMarginType `json:"margin,omitempty"`
		Size        linebot.FlexIconSizeType        `json:"size,omitempty"`
		AspectRatio linebot.FlexIconAspectRatioType `json:"aspectRatio,omitempty"`
		OffsetStart string                          `json:"offsetStart,omitempty"`
		OffsetEnd   string                          `json:"offsetEnd,omitempty"`
	}{
		Type:        linebot.FlexComponentTypeIcon,
		URL:         c.URL,
		Margin:      c.Margin,
		Size:        c.Size,
		AspectRatio: c.AspectRatio,
		OffsetStart: c.OffsetStart,
		OffsetEnd:   c.OffsetEnd,
	})
}
