package services

import (
	"axie-notify/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
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

// PingService provides the function that send the serviceName and services information to match and validate is it online or not.
func PingService(commandMessage string, servicesInfo *models.ServicesInfo, timeOut time.Duration) string {
	serviceInfo, err := FindServiceName(commandMessage, servicesInfo)
	if err != nil {
		return "Sorry, the name did not match to any services in our system."
	}
	if len(serviceInfo.ServiceName) > 0 && len(serviceInfo.Port) > 0 {
		serviceStatus := ping(serviceInfo.ServiceName, serviceInfo.IPAddress, serviceInfo.Port, timeOut)
		message, _ := isServiceOnline(serviceInfo.ServiceName, serviceStatus)
		return message
	}
	return ""
}

func FindServiceName(messageText string, servicesInfo *models.ServicesInfo) (*models.ServiceInfo, error) {
	for _, serviceDetail := range *servicesInfo {
		if strings.Contains(strings.ToLower(messageText), strings.ToLower(serviceDetail.ServiceName)) {
			return &serviceDetail, nil
		}
	}
	return nil, errors.New("the name did not match to any services in our system.")
}

// StartPingAllServices provides the function ping to allservice that we send through input.
func StartPingAllServices(servicesInfo *models.ServicesInfo, timeOut time.Duration) []string {
	var lstServiceDowns []string
	for _, serviceDetail := range *servicesInfo {
		serviceStatus := ping(serviceDetail.ServiceName, serviceDetail.IPAddress, serviceDetail.Port, timeOut)
		if message, isOnline := isServiceOnline(serviceDetail.ServiceName, serviceStatus); !(isOnline) {
			lstServiceDowns = append(lstServiceDowns, message)
		}
	}
	return lstServiceDowns
}

func isServiceOnline(serviceName string, status bool) (string, bool) {
	if status {
		return fmt.Sprintf("%s service is working pretty well.", serviceName), true
	} else {
		return fmt.Sprintf("%s service is down, please contact admin.", serviceName), false
	}

}

func ping(serviceName string, ipAddress string, port string, timeOut time.Duration) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ipAddress, port), timeOut)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// NewBankCoreServiceInfo provides the all service information that using in BankCore project.
func NewBankCoreServiceInfo() *models.ServicesInfo {
	bankCoreServices := models.ServicesInfo{}
	for BankServiceName, BankServicePort := range bankServiceInfo {
		serviceInfo := models.ServiceInfo{
			ServiceName: BankServiceName,
			IPAddress:   bankCoreIPAddress,
			Port:        BankServicePort}
		bankCoreServices = append(bankCoreServices, serviceInfo)
	}
	return &bankCoreServices
}
