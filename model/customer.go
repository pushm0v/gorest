package model

type Customer struct {
	Id      int    `json:"id" gorm:"autoIncrement;primaryKey"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
