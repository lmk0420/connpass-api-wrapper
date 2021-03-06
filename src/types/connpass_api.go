package types

import (
	"time"
)

// RequestParam is Request parameter of connpass-API
type RequestParam struct {
	EventID	int
	Keyword	[]string
	KeywordOR []string
	YM []int
	YMD []int
	NickName []string
	OwnerNickname []string
	SeriesID int
	Start int
	Order int
	Count int
	Format Format
}

// ResultSet is struct of connpass-API responce
type ResultSet struct {
	ResultsReturned int `json:"results_returned"`
	Events []Event		`json:"events"`
	ResultsStart     int `json:"results_start"`
	ResultsAvailable int `json:"results_available"`
}

// Event is connpass-API response Series
type Event struct {
	EventURL      string `json:"event_url"`
	EventTypes     string `json:"event_type"`
	OwnerNickname string `json:"owner_nickname"`
	Series        Series `json:"series"`
	UpdatedAt        time.Time `json:"updated_at"`
	Lat              string    `json:"lat"`
	StartedAt        time.Time `json:"started_at"`
	HashTag          string    `json:"hash_tag"`
	Title            string    `json:"title"`
	EventID          int       `json:"event_id"`
	Lon              string    `json:"lon"`
	Waiting          int       `json:"waiting"`
	Limit            int       `json:"limit"`
	OwnerID          int       `json:"owner_id"`
	OwnerDisplayName string    `json:"owner_display_name"`
	Description      string    `json:"description"`
	Address          string    `json:"address"`
	Catch            string    `json:"catch"`
	Accepted         int       `json:"accepted"`
	EndedAt          time.Time `json:"ended_at"`
	Place            string    `json:"place"`
}

// Series is Series field of the response of connpass-API
type Series struct {
	URL   string `json:"url"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Format is the Series field of the response of connpass-API
type Format string

// connpass-API response format is json only
const (
	JSON Format = "json"
)

// NewRequestParam is constructor of RequestParam
func NewRequestParam(eventid int, keyword, keywordor []string, ym, ymd []int, nickname, ownernickname []string,
					 seriesid, start, order, count int, format string) *RequestParam{
	param := new(RequestParam)
	param.EventID = eventid
	param.Keyword = keyword
	param.KeywordOR = keywordor
	param.YM = ym
	param.YMD = ymd
	param.NickName = nickname
	param.OwnerNickname = ownernickname
	param.SeriesID = seriesid
	param.Start = start
	param.Order = order
	param.Count = count
	param.Format = JSON

	return param
}