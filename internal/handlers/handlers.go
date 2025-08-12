package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	services "go_webserv/internal/services"
	"go_webserv/internal/types"
)

type Handlers struct {
	db *gorm.DB
	blogService *services.BlogService
	projectService *services.ProjectService
}

func NewHandlers(db *gorm.DB) *Handlers {
	return &Handlers{
		db: db,
		blogService: services.InitBlogService(),
		projectService: services.InitProjectService(),
	}
}

func (h *Handlers) CreateNewBlogHandler(w http.ResponseWriter, r *http.Request) {
	var blog types.Blog
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blog.CreatedAt = time.Now()
	h.db.Create(&blog)
}

func (h *Handlers) QueryBlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogId := vars["id"]
	println("blogId:", blogId)
	if blogId != "" {
		h.blogService.QueryBlog(h.db, blogId)
	}
}

func (h *Handlers) QueryAllBlogHandler(w http.ResponseWriter, r *http.Request) {
	println("All blogs:")
	h.blogService.QueryAllBlogs(h.db)
}

func (h *Handlers) QueryAllHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.blogService.QueryAllBlogs(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	projects, err := h.projectService.QueryAllProjects(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, blog := range blogs {
		println(blog.Subject)
	}

	for _, project := range projects {
		println(project.ProjectName)
	}
}

func (h *Handlers) CreateNewProjectHandler(w http.ResponseWriter, r *http.Request) {
	var project types.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project.CreatedAt = time.Now()
	h.db.Create(&project)
}

func (h *Handlers) QueryAllProjectHandler(w http.ResponseWriter, r *http.Request) {
	println("All projects:")
	h.projectService.QueryAllProjects(h.db)
}

func (h *Handlers) QueryProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["id"]
	println("projectId:", projectId)
	if projectId != "" {
		h.projectService.QueryProject(h.db, projectId)
	}
}
