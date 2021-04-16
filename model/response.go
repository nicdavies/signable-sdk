package model

type ListResponse struct {
	Http int `json:"http"`
}

type PaginateResponse struct {
	Http   int    `json:"http"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Next   string `json:"next"`
}

type CreateResponse struct {
	Http    int    `json:"http"`
	Message string `json:"message"`
	Href    string `json:"href"`
}

type ReadResponse struct {
	Http int `json:"http"`
}

type UpdateResponse struct {
	Http    int    `json:"http"`
	Message string `json:"message"`
	Href    string `json:"href"`
}

type DeleteResponse struct {
	Http    int    `json:"http"`
	Message string `json:"message"`
}
