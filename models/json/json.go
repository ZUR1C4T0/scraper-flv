package json

type VideoServer struct {
	Server      string `json:"server"`
	Title       string `json:"title"`
	Ads         int    `json:"ads"`
	AllowMobile bool   `json:"allow_mobile"`
	Code        string `json:"code"`
}

type VideoResponse struct {
	Sub []VideoServer `json:"SUB"`
}
