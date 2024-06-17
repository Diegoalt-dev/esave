package dao

import (
	"esave/database"
	"esave/database/models"
	services "esave/services/dtos"
	"log"
)

type SpentDaoImpl struct {
}

func (dao SpentDaoImpl) CreateSpent(dto services.SpentDto) {
	log.Printf("Creating spent: %v", dto)
	spent := models.Spent{Description: dto.Description, User: dto.User, Value: dto.Value}
	log.Printf("Spent created: %v", spent)
	db := database.GetDb()
	log.Println("Saving spent...")
	db.Create(&spent)
}

func (dao SpentDaoImpl) GetSpentById(id uint64) services.SpentDto {
	db := database.GetDb()
	var spent models.Spent
	db.First(&spent, id)
	return services.SpentDto{ID: spent.ID, Description: spent.Description, User: spent.User, Value: spent.Value}
}

func (dao SpentDaoImpl) DeleteSpentById(id uint64) {
	db := database.GetDb()
	db.Delete(&models.Spent{}, id)
}

func (dao SpentDaoImpl) UpdateSpent(dto services.SpentDto) {
	db := database.GetDb()
	db.Model(&models.Spent{}).Where("id = ?", dto.ID).Updates(models.Spent{Description: dto.Description, User: dto.User, Value: dto.Value})
}
