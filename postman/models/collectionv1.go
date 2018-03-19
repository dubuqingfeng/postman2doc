package models

// by https://github.com/aubm/postmanerator/blob/master/postman/collection_v1.go

type CollectionV1 struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Order       []string `json:"order"`
	Folders []struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Order       []string `json:"order"`
	} `json:"folders"`
	Requests []requestV1 `json:"requests"`
}

type requestV1 struct {
	ID            string                 `json:"id"`
	RawHeaders    string                 `json:"headers"`
	URL           string                 `json:"url"`
	PathVariables map[string]interface{} `json:"pathVariables"`
	Method        string                 `json:"method"`
	Data []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"data"`
	DataMode    string       `json:"dataMode"`
	Tests       string       `json:"tests"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	RawModeData string       `json:"rawModeData"`
	Responses   []responseV1 `json:"responses"`
}

type responseV1 struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	ResponseCode struct {
		Code int `json:"code"`
	} `json:"responseCode"`
	Headers []struct {
		Name        string `json:"name"`
		Key         string `json:"key"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"headers"`
	Text string `json:"text"`
	Name string `json:"name"`
}
