package services

import (
	"gorm.io/gorm"
	"strconv"
	"time"

	"fmt"

	types "go_webserv/internal/types"
)

type ProjectService struct {

}

func InitProjectService() *ProjectService {
	return &ProjectService{}
}

func (p *ProjectService) CreateNewProject(db *gorm.DB, newProject types.ProjectItem) error {
	err := db.Create(&types.Project{
		ProjectID: Sha256Hex(fmt.Sprintf("%s-%d", newProject.Name, time.Now().UnixNano()))[:16],
		Name:     newProject.Name,
		ContentPath: newProject.ContentPath,
		CreatedAt:   time.Now(),
		LastModified: time.Now(),
	}).Error

	if err != nil {
		println("Error creating project:", err)
		return err
	}

	return nil
}

func (p *ProjectService) QueryProject(db *gorm.DB, projectId string) (*types.ProjectItem, error) {
	project := types.Project{}
	projectIdInInt, err := strconv.Atoi(projectId)
	if err != nil {
		return nil, err
	}
	db.First(&project, "project_id = ?", projectIdInInt)

	return nil, nil
}

func (p *ProjectService) QueryAllProjects(db *gorm.DB) ([]types.Project, error) {
	var projects []types.Project
	result := db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}
