package models

type Request struct {
	ApiKey  string `json:"api_key"`
	LongUrl string `json:"long_url"`
}

type Response struct {
	ShortUrl string `json:"short_url"`
}
