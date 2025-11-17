package handlers

import (
	"encoding/json"
	"net/http"

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
	var blog types.BlogItem
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.blogService.CreateNewBlog(h.db, blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
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
	var blogs []types.Blog
	blogs, _ = h.blogService.QueryAllBlogs(h.db)

	if err := json.NewEncoder(w).Encode(blogs); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) QueryAllHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.blogService.QueryAllBlogs(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	projects, err := h.projectService.QueryAllProjects(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(struct {
		Blogs    []types.Blog    `json:"blogs"`
		Projects []types.Project `json:"projects"`
	}{
		Blogs:    blogs,
		Projects: projects,
	}); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) CreateNewProjectHandler(w http.ResponseWriter, r *http.Request) {
	var project types.ProjectItem
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.projectService.CreateNewProject(h.db, project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) QueryAllProjectHandler(w http.ResponseWriter, r *http.Request) {
	println("All projects:")
	var projects []types.Project
	projects, _ = h.projectService.QueryAllProjects(h.db)

	if err := json.NewEncoder(w).Encode(projects); err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
}

func (h *Handlers) QueryProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["id"]
	println("projectId:", projectId)
	if projectId != "" {
		h.projectService.QueryProject(h.db, projectId)
	}
}
