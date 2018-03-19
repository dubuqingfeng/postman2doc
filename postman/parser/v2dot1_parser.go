package parser

import (
	"encoding/json"
	"github.com/dubuqingfeng/postman2doc/postman/models"
	"html/template"
	"fmt"
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
	request.PayloadRaw = template.HTML(item.Request.RawBody.Raw)
	request.Responses = this.buildResponse(item)
	request.Headers = item.Request.RawHeaders
	return request
}

func (this CollectionV2dot1Parser) buildResponse(item models.ItemV2dot1) ([]models.Response) {
	requests := []models.Response{}
	for i, item := range item.Response {
		fmt.Print(i)
		request := models.Response{}
		request.Name = item.Name
		request.Body = template.HTML(item.Body)
		request.StatusCode = item.Code
		requests = append(requests, request)
	}

	return requests
}
