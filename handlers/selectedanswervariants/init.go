package selectedFieldFillvariants

import (
	"context"
	"mdgkb/ankets-server/models"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IService interface {
	GetAll() ([]*models.SelectedFieldFillVariant, error)
	Get(*string) (*models.SelectedFieldFillVariant, error)
	Create(*models.SelectedFieldFillVariant) error
	Update(*models.SelectedFieldFillVariant) error
	Delete(*string) error

	UpsertMany(models.SelectedFieldFillVariants) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.SelectedFieldFillVariant) error
	getAll() ([]*models.SelectedFieldFillVariant, error)
	get(*string) (*models.SelectedFieldFillVariant, error)
	update(*models.SelectedFieldFillVariant) error
	delete(*string) error

	upsertMany(radios models.SelectedFieldFillVariants) error
	deleteMany([]uuid.UUID) error
}

type Handler struct {
	service IService
	helper  *helper.Helper
}
type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
