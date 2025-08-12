package services

import (
	"gorm.io/gorm"
	"strconv"

	types "go_webserv/internal/types"
)

type ProjectService struct {

}

func InitProjectService() *ProjectService {
	return &ProjectService{}
}

func (p *ProjectService) CreateNewProject(db *gorm.DB) error {
	return nil
}

func (p *ProjectService) QueryProject(db *gorm.DB, projectId string) (*types.ProjectItem, error) {
	project := types.Project{}
	projectIdInInt, err := strconv.Atoi(projectId)
	if err != nil {
		return nil, err
	}
	db.First(&project, "project_id = ?", projectIdInInt)
	println(project.ProjectName)
	return nil, nil
}

func (p *ProjectService) QueryAllProjects(db *gorm.DB) ([]types.Project, error) {
	var projects []types.Project
	result := db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, project := range projects {
		println(project.ProjectName)
	}

	return projects, nil
}
