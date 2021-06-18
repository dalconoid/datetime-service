package model

// NowResponse - successful response for [/time/now]
type NowResponse struct {
	Time float64 `json:"time"`
}

// TimeToStrRequest - expected request for [/time/string]
type TimeToStrRequest struct {
	Time float64 `json:"time"`
}

// TimeToStrResponse - successful response for [/time/string]
type TimeToStrResponse struct {
	Str string `json:"str"`
}

// TimeAddRequest - expected request for [/time/add]
type TimeAddRequest struct {
	Time float64 `json:"time"`
	Delta float64 `json:"delta"`
}

// TimeAddResponse - successful response for [/time/add]
type TimeAddResponse struct {
	Time float64 `json:"time"`
	Error string `json:"delta,omitempty"`
}

// TimeCorrectRequest - expected request for [/time/correct]
type TimeCorrectRequest struct {
	Time float64 `json:"time"`
}