package services

import services "esave/services/dtos"

type spentService interface {
	CreateSpent(dto services.SpentDto)
	GetSpentById(id uint64) services.SpentDto
	DeleteSpentById(id uint64)
	UpdateSpent(dto services.SpentDto)
}
