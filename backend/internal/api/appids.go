// internal/api/appids.go
package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/Axedd/steam-tracker.git/internal/db"
	"github.com/gin-gonic/gin"
)

// AppIDHandler holds the dependencies for the appids endpoints.
type AppIDHandler struct {
	Queries *db.Queries
}

// NewAppIDHandler constructs a new AppIDHandler.
func NewAppIDHandler(q *db.Queries) *AppIDHandler {
	return &AppIDHandler{Queries: q}
}

// RegisterRoutes registers all appids routes on the given router.
// Notice we accept gin.IRouter, which both *gin.Engine and *gin.RouterGroup implement.
func (h *AppIDHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/appids", h.ListAppIDs)
	r.GET("/appids/:id", h.GetAppIDByID)
}

// ListAppIDs responds with all entries in the appids table.
func (h *AppIDHandler) ListAppIDs(c *gin.Context) {
	apps, err := h.Queries.ListAppIDs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, apps)
}

// GetAppIDByID responds with a single appid by its numeric ID.
func (h *AppIDHandler) GetAppIDByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appid"})
		return
	}
	app, err := h.Queries.GetAppIDByID(c.Request.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, app)
}
