package parser

import (
	"github.com/dubuqingfeng/postman2doc/postman/models"
)

type Parser interface {
	Parse(content []byte) (bool, string, models.Collection)
}