package dto

type AnalysesRequest struct {
	WebUrl string `json:"webUrl"`
}

type AnalysesResponse struct {
	Id int64 `json:"id"`
}

type PageInfo struct {
	Url          string
	Title        string
	Headings     map[string]int
	LinkDetails  LinkDetails
	PageVersion  string
	HasLoginForm bool
}

type LinkDetails struct {
	NoofLinks int
	Active    int
	Internal  int
}
