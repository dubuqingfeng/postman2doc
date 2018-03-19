package models

type CollectionV2dot1 struct {
	Info  InfoV2dot1   `json:"info"`
	Items []ItemV2dot1 `json:"item"`
}

type InfoV2dot1 struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schema      string `json:"description"`
}

type ItemV2dot1 struct {
	Name     string           `json:"name"`
	Request  requestV2dot1    `json:"request"`
	Response []responseV2dot1 `json:"response"`
}

type requestV2dot1 struct {
	Method string `json:"method"`
	RawHeaders []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"headers"`
	RawBody struct {
		Mode string `json:"mode"`
		Raw  string `json:"raw"`
	} `json:"body"`
	URL         urlV2dot1 `json:"url"`
	Description string    `json:"description"`
}

type urlV2dot1 struct {
	Raw string `json:"raw"`
}

type responseV2dot1 struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Code   int    `json:"code"`
	Headers []struct {
		Name        string `json:"name"`
		Key         string `json:"key"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"headers"`
	Body string `json:"body"`
}
