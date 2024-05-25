package services

type SpentDto struct {
	ID          uint64  `json:"id"`
	Description string  `json:"description"`
	User        string  `json:"user"`
	Value       float64 `json:"value"`
}
