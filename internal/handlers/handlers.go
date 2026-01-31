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
	postService *services.PostService
	awsService  *services.AwsService
}

func NewHandlers(db *gorm.DB) *Handlers {
	return &Handlers{
		postService: services.InitPostService(services.InitDbService(db)),
		awsService:  services.InitAwsService(),
	}
}

func (h *Handlers) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var post types.PostItem
	var bId = vars["id"]
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.postService.CreateNewPost(post, bId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) QueryPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	id := vars["id"]
	postType := parsePostType(vars["type"])

	if id != "" {
		postItem, _ := h.postService.QueryPost(id, postType)
		if err := json.NewEncoder(w).Encode(postItem); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handlers) QueryAllPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	postType := parsePostType(vars["type"])

	posts, err := h.postService.QueryAllPosts(postType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) FileDownloadLinkHandler(w http.ResponseWriter, r *http.Request) {
	downloadLinkReq := types.DownloadLinkRequest{}
	err := json.NewDecoder(r.Body).Decode(&downloadLinkReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link, err := h.awsService.GetDownloadLink(downloadLinkReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(link); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) FileUploadLinkHandler(w http.ResponseWriter, r *http.Request) {
	uploadLinkReq := types.UploadLinkRequest{}
	err := json.NewDecoder(r.Body).Decode(&uploadLinkReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link, err := h.awsService.GetUploadLink(uploadLinkReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(link); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func parsePostType(typeStr string) types.PostType {
	postType := types.ALL
	switch typeStr {
	case "blog":
		postType = types.BLOG
	case "project":
		postType = types.PROJECT
	}

	return postType;
}