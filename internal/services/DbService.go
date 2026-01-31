package services

import (
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

func (d *DbService) doInsert(record *types.Post) error {
	err := d.db.Create(record).Error
	if err != nil {
		println("[DbService] Error inserting record:", err)
		return err
	}

	return nil
}

func (d *DbService) doQueryByID(record *types.Post, id string) error {
	if err := d.db.First(record, id).Error; err != nil {
		println("[DbService] Error querying record by ID:", err)

		return err
	}

	return nil
}

func (d *DbService) doQueryAll(posts []types.Post, postType types.PostType) error {
	if err := d.db.Find(&posts).Error; err != nil {
		println("[DbService] Error querying all records:", err)
	
		return err
	}
	
	return nil
}