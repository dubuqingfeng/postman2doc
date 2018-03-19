package parser

import (
	"encoding/json"
	"github.com/dubuqingfeng/postman2doc/postman/models"
)

type CollectionV2dot1Parser struct {
	Parser
}

var V2dot1Parser CollectionV2dot1Parser

func init() {
	V2dot1Parser = CollectionV2dot1Parser{}
}

func (this CollectionV2dot1Parser) Parse(content []byte) (bool, string, models.Collection) {
	collectionV2dot1 := models.CollectionV2dot1{}

	if err := json.Unmarshal(content, &collectionV2dot1); err != nil {
		return false, err.Error(), models.Collection{}
	}
	return true, "success", this.buildCollectionFromSource(collectionV2dot1)
}

func (this CollectionV2dot1Parser) buildCollectionFromSource(content models.CollectionV2dot1) models.Collection {
	collection := models.Collection{
		Name:        content.Info.Name,
		Description: content.Info.Description,
		Requests:    make([]models.Request, 0),
	}

	collection.Requests = this.FindRequests(make([]models.Request, 0), content.Items)

	//fmt.Println(collection)
	return collection
}

// TODO refactor
func (this CollectionV2dot1Parser) FindRequests(requests []models.Request, items []models.ItemV2dot1) []models.Request {
	for _, item := range items {
		req := this.buildRequest(item)
		if req.URL == "" {
			requests = this.FindRequests(requests, item.Items)
			requests = append(requests, req)
		} else {
			requests = append(requests, req)
		}
	}
	return requests
}

func (this CollectionV2dot1Parser) buildRequest(item models.ItemV2dot1) (models.Request) {
	request := models.Request{}
	request.Name = item.Name
	request.Method = item.Request.Method
	request.URL = item.Request.URL.Raw
	request.Description = item.Request.Description
	request.PayloadRaw = item.Request.RawBody.Raw
	request.Headers = item.Request.RawHeaders
	return request
}
