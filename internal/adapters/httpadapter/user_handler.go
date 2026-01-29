package httpadapter

import (
	"database/sql"
	"net/http"

	"crudWithDB/internal/application"
	"crudWithDB/internal/domain"

	"github.com/gin-gonic/gin"
)

type UserHTTPHandler struct {
	service *application.UserService
}

func NewUserHTTPHandler(service *application.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{service: service}
}

func (h *UserHTTPHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/users", h.Create)
	r.GET("/users/:id", h.Get)
	r.PUT("/users/:id", h.Update)
	r.DELETE("/users/:id", h.Delete)
}

func (h *UserHTTPHandler) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	if err := h.service.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *UserHTTPHandler) Get(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHTTPHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	user.ID = id

	if err := h.service.Update(user); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *UserHTTPHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
