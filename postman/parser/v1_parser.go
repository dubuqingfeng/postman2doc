package parser

import (
	"encoding/json"
	"github.com/dubuqingfeng/postman2doc/postman/models"
)

type CollectionV1Parser struct {
	Parser
}

var V1Parser CollectionV1Parser

func init() {
	V1Parser = CollectionV1Parser{}
}

func (this CollectionV1Parser) Parse(content []byte) (bool, string, models.Collection) {
	collectionV1 := models.CollectionV1{}
	if err := json.Unmarshal(content, &collectionV1); err != nil {
		return false, err.Error(), models.Collection{}
	}
	return true, "success", this.buildCollectionFromSource(collectionV1)
}

func (this CollectionV1Parser) buildCollectionFromSource(content models.CollectionV1) models.Collection {
	return models.Collection{}
}
