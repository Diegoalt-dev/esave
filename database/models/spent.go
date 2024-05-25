package models

type Spent struct {
	ID          uint64  `json:"id" gorm:"primaryKey"`
	Description string  `json:"description"`
	User        string  `json:"user"`
	Value       float64 `json:"value"`
}
