package common

type DeleteRequest struct {
	Id int64
}

type PageRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PageResponse struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}
