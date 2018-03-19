package models

import "html/template"

// by https://github.com/aubm/postmanerator/blob/master/postman/collection.go

type Collection struct {
	Name        string
	Description string
	Requests    []Request
}

type Request struct {
	ID            string
	Name          string
	Description   string
	Method        string
	URL           string
	PayloadType   string
	PayloadRaw    template.HTML
	PayloadParams []KeyValuePair
	PathVariables []KeyValuePair
	Headers       []KeyValuePair
	Responses     []Response
	Tests         string
}

type Response struct {
	ID         string
	Name       string
	Status     string
	StatusCode int
	Body       template.HTML
	Headers    []KeyValuePair
}

type KeyValuePair struct {
	Name        string
	Key         string
	Value       interface{}
	Description string
}
