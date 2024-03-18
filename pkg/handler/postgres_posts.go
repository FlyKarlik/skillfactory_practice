package handler

import (
	"net/http"
	"skillfactory_project/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var post model.Posts
	userId := c.Query("id")
	if err := c.BindJSON(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	authorId, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := h.srv.PostsService.CreatePost(authorId, post)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) GetAllPosts(c *gin.Context) {
	userId := c.Query("id")
	authorId, err := strconv.Atoi(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	posts, err := h.srv.PostsService.GetAllPosts(authorId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	userId := c.Query("id")
	authorId, err := strconv.Atoi(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updates model.UpdatePosts

	if err := c.BindJSON(&updates); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updates.UserId = authorId

	if err := h.srv.PostsService.UpdatePost(updates); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"TaskId": updates.TaskId,
		"Status": "Success",
	})
}

func (h *Handler) DeletePost(c *gin.Context) {
	userId := c.Query("id")
	authorId, err := strconv.Atoi(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var deletePost model.Posts

	if err := c.BindJSON(&deletePost); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	deletePost.Author_id = authorId

	if err := h.srv.PostsService.DeletePost(deletePost); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Task_id": deletePost.Id,
	})
}
