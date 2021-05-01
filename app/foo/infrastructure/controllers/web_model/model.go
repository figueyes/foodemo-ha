package web_model

type FooWebModel struct {
	Total int   `json:"total"`
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Data  []DataFooWebModel `json:"data"`
}
type DataFooWebModel struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}