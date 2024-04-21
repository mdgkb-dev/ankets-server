package auth

import (
	"mdgkb/ankets-server/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RestorePassword(c *gin.Context)
}

type IService interface {
	Login(*models.User) (*models.TokensWithUser, error)
	DoesEmailExist(string) (bool, error)
	DropUUID(item *models.User) error
	UpdatePassword(*models.User) error
}

type Handler struct {
	helper *helper.Helper
}

type Service struct {
	helper *helper.Helper
}

type DoesLoginExist struct {
	DoesLoginExist bool
}

var (
	H *Handler
	S *Service
)

func Init(h *helper.Helper) {
	H = &Handler{helper: h}
	S = &Service{helper: h}
}
