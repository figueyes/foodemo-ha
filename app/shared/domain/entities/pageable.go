package entities

type Pageable struct {
	Total int         `json:"total"`
	Page  int64       `json:"page"`
	Limit int64       `json:"limit"`
	Data  interface{} `json:"data"`
}
