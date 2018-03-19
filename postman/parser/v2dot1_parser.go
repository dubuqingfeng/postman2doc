package parser

import (
	"encoding/json"
	"github.com/dubuqingfeng/postman2doc/postman/models"
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
	fmt.Print(content)

	return models.Collection{}
}
