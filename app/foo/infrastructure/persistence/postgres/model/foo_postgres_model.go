package model

type FooPostgresModel struct {
	Id      string `json:"id" gorm:"column:id;type:varchar(60);primary_key"`
	Message string `json:"message" gorm:"column:message;type:varchar(500)"`
}

func (FooPostgresModel) TableName() string {
	return "foos"
}
