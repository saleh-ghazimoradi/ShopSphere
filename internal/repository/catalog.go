package repository

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service/serviceModels"
	"gorm.io/gorm"
)

type Catalog interface {
	CreateCategory(ctx context.Context, category *serviceModels.Category) error
	FindCategories(ctx context.Context) ([]*serviceModels.Category, error)
	FindCategoryById(ctx context.Context, id uint) (*serviceModels.Category, error)
	EditCategory(ctx context.Context, category *serviceModels.Category) (*serviceModels.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
}

type catalogRepository struct {
	db *gorm.DB
}

func (c *catalogRepository) CreateCategory(ctx context.Context, category *serviceModels.Category) error {
	err := c.db.WithContext(ctx).Create(category).Error
	if err != nil {
		return errors.New("create category failed")
	}
	return nil
}

func (c *catalogRepository) FindCategories(ctx context.Context) ([]*serviceModels.Category, error) {
	categories := make([]*serviceModels.Category, 0)
	err := c.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, errors.New("find categories failed")
	}

	return categories, nil
}

func (c *catalogRepository) FindCategoryById(ctx context.Context, id uint) (*serviceModels.Category, error) {
	category := &serviceModels.Category{}
	err := c.db.WithContext(ctx).First(category, id).Error
	if err != nil {
		return nil, errors.New("category does not exist")
	}
	return category, nil
}

func (c *catalogRepository) EditCategory(ctx context.Context, category *serviceModels.Category) (*serviceModels.Category, error) {
	err := c.db.WithContext(ctx).Save(category).Error
	if err != nil {
		return nil, errors.New("edit category failed")
	}
	return category, nil
}

func (c *catalogRepository) DeleteCategory(ctx context.Context, id uint) error {
	if err := c.db.WithContext(ctx).Delete(&serviceModels.Category{}, id).Error; err != nil {
		return errors.New("delete category failed")
	}
	return nil
}

func NewCatalogRepository(db *gorm.DB) Catalog {
	return &catalogRepository{
		db: db,
	}
}
