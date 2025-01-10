package v1

import (
	"net/http"

	"github.com/Snehashish1609/validator-api/config"
	"github.com/Snehashish1609/validator-api/models"

	"github.com/Snehashish1609/validator-api/internal/common"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type APIHandler struct {
	UserHandler *models.UserHandler
	Config      *config.Config
	// ToDo:
	// can add DB client here (if required)
}

func NewAPIHandler(c *config.Config, uh *models.UserHandler) *APIHandler {
	return &APIHandler{
		UserHandler: uh,
		Config:      c,
	}
}

func (ah *APIHandler) ValidateUser(c *gin.Context) {
	log.Info().
		Msg("ValidateUser called")
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// field validation
	err = ah.UserHandler.Validator.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": common.PayloadValidationFailedMsg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": common.PayloadValidionSuccessMsg})
}
