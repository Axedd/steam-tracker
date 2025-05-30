package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Axedd/steam-tracker.git/internal/db"
	"github.com/gin-gonic/gin"
)

type SteamParamHandler struct {
	Queries *db.Queries
}

func NewSteamParamHandler(q *db.Queries) *SteamParamHandler {
	return &SteamParamHandler{Queries: q}
}

func (h *SteamParamHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/steamParams", h.GetParamsByAppID)        // for global
	r.GET("/steamParams/:appid", h.GetParamsByAppID) // for specific
}

type CleanParamDef struct {
	Key          string          `json:"key"`
	Label        string          `json:"label"`
	Type         string          `json:"type"`
	Options      json.RawMessage `json:"options,omitempty"`
	DefaultValue string          `json:"default_value,omitempty"`
	HelpText     string          `json:"help_text,omitempty"`
	AppID        *int32          `json:"appid,omitempty"`
}

func cleanParam(p db.SteamParamDef) CleanParamDef {
	var opts json.RawMessage
	if p.Options.Valid {
		opts = p.Options.RawMessage
	}
	var defVal string
	if p.DefaultValue.Valid {
		defVal = p.DefaultValue.String
	}
	var help string
	if p.HelpText.Valid {
		help = p.HelpText.String
	}
	var appid *int32
	if p.Appid.Valid {
		val := p.Appid.Int32
		appid = &val
	}

	return CleanParamDef{
		Key:          p.Key,
		Label:        p.Label,
		Type:         p.Type,
		Options:      opts,
		DefaultValue: defVal,
		HelpText:     help,
		AppID:        appid,
	}
}

func cleanParams(params []db.SteamParamDef) []CleanParamDef {
	out := make([]CleanParamDef, len(params))
	for i, p := range params {
		out[i] = cleanParam(p)
	}
	return out
}

// GetParamsByAppID responds with the params for the app.
func (h *SteamParamHandler) GetParamsByAppID(c *gin.Context) {
	idParam := c.Param("appid")

	if idParam == "" {
		params, err := h.Queries.GetGlobalParams(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cleanParams(params))
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appid"})
		return
	}

	params, err := h.Queries.GetParamsByAppID(c.Request.Context(), sql.NullInt32{
		Int32: int32(id),
		Valid: true,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "no parameters found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, cleanParams(params))
}
