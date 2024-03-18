package handler

import (
	service "skillfactory_project/pkg/sercvice"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv *service.Sercive
}

func NewHandler(srv *service.Sercive) *Handler {
	return &Handler{
		srv: srv,
	}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	posts := router.Group("/posts")
	{
		posts.POST("/creeate", h.CreatePost)
		posts.GET("/getPosts", h.GetAllPosts)
		posts.PUT("/updatePost", h.UpdatePost)
		posts.DELETE("/deletePost", h.DeletePost)
	}
	return router
}
