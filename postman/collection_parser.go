package postman

import (
	"fmt"
	"github.com/dubuqingfeng/postman2doc/postman/parser"
	"github.com/dubuqingfeng/postman2doc/postman/models"
	"errors"
	"os"
	"log"
	"io/ioutil"
)

func getPostmanCollection(filename string) (models.Collection, error) {
	fmt.Print(filename)
	b, err := readCollectionFile(filename)
	if err != nil {
		return models.Collection{}, err
	}
	//fmt.Println(b)

	parsers := [2]parser.Parser{parser.V2dot1Parser, parser.V1Parser}
	for i, item := range parsers {
		result, msg, collection := item.Parse(b)
		if result {
			return collection, nil
		} else {
			fmt.Println(string(i) + msg)
		}
	}

	return models.Collection{}, errors.New("ParseFailed")
}

func readCollectionFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//b, err := ioutil.ReadAll(file)
	//fmt.Print(b)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return b, nil
}
