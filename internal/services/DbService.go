package services

import (
	"strings"

	"gorm.io/gorm"

	types "go_webserv/internal/types"
)

type DbService struct {
	db *gorm.DB
}

func InitDbService(db *gorm.DB) *DbService {
	return &DbService{
		db: db,
	}
}

func (d *DbService) doInsert(db *gorm.DB, record interface{}) error {
	err := db.Create(record).Error
	if err != nil {
		println("[DbService] Error inserting record:", err)
		return err
	}

	return nil
}

func (d *DbService) doQueryByID(db *gorm.DB, record interface{}, id string) error {
	if err := db.First(record, id).Error; err != nil {
		println("[DbService] Error querying record by ID:", err)

		return err
	}

	return nil
}

func (d *DbService) doQueryAll(db *gorm.DB, posts []interface{}, ) error {
	if err := db.Find(&posts).Error; err != nil {
		println("[DbService] Error querying all records:", err)
	
		return err
	}
	
	return nil
}