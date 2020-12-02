package services

import (
	"encoding/json"
	"io/ioutil"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"os"
)

//Crawler service
type Crawler interface {
	Branches(string) (*[]branch.Entity, common.CustomError)
}

type crawler struct {
}

//NewCrawler create a new service for crawl
func NewCrawler() Crawler {

	return &crawler{}
}

//Branches crawl branches from institution
func (s *crawler) Branches(baseURL string) (*[]branch.Entity, common.CustomError) {

	//TODO: concat baseURL with resource url

	jsonFile, err := os.Open("./domain/branch/branches.json")

	if err != nil {
		return nil, common.NewInternalServerError("Unable to crawl branches from institution", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	branchJSONData := &branchJSON{}

	jsonUnmarshallErr := json.Unmarshal(byteValue, &branchJSONData)

	if jsonUnmarshallErr != nil {
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	companies := branchJSONData.Data.Brand.Companies[0]

	return &companies.Branches, nil

}

type branchJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				Branches []branch.Entity `json:"branches"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
