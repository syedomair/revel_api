package models

type Client struct {
	Id        int64  `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	ApiKey    string `json:"api_key" gorm:"column:api_key"`
	ApiSecret string `json:"api_secret" gorm:"column:api_secret"`
	active    string `json:"active" gorm:"column:active"`
}
