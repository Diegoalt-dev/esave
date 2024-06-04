package models

type Spent struct {
	ID          uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string  `json:"description"`
	User        string  `json:"user"`
	Value       float64 `json:"value"`
}
