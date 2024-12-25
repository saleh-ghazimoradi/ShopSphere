package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/dto"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/repository"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service/serviceModels"
)

type Catalog interface {
	CreateCategory(ctx context.Context, input *dto.Category) error
	EditCategory(ctx context.Context, id int, input *dto.UpdateCategory) (*serviceModels.Category, error)
	DeleteCategory(ctx context.Context, id int) error
	GetCategories(ctx context.Context) ([]*serviceModels.Category, error)
	GetCategory(ctx context.Context, id int) (*serviceModels.Category, error)
}

type catalogService struct {
	catalogRepository repository.Catalog
}

func (c catalogService) CreateCategory(ctx context.Context, input *dto.Category) error {
	return c.catalogRepository.CreateCategory(ctx, &serviceModels.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})
}

func (c catalogService) EditCategory(ctx context.Context, id int, input *dto.UpdateCategory) (*serviceModels.Category, error) {
	cat, err := c.catalogRepository.FindCategoryById(ctx, id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	if input.Name != nil {
		cat.Name = *input.Name
	}

	if input.ParentId != nil {
		cat.ParentId = *input.ParentId
	}

	if input.ImageUrl != nil {
		cat.ImageUrl = *input.ImageUrl
	}

	if input.DisplayOrder != nil {
		cat.DisplayOrder = *input.DisplayOrder
	}

	updatedCat, err := c.catalogRepository.EditCategory(ctx, cat)
	if err != nil {
		return nil, err
	}

	return updatedCat, nil
}

func (c catalogService) DeleteCategory(ctx context.Context, id int) error {
	if err := c.catalogRepository.DeleteCategory(ctx, id); err != nil {
		return errors.New("category not found")
	}
	return nil
}

func (c catalogService) GetCategories(ctx context.Context) ([]*serviceModels.Category, error) {
	categories, err := c.catalogRepository.FindCategories(ctx)
	if err != nil {
		return nil, errors.New("categories do not exist")
	}
	return categories, nil
}

func (c catalogService) GetCategory(ctx context.Context, id int) (*serviceModels.Category, error) {
	cat, err := c.catalogRepository.FindCategoryById(ctx, id)
	if err != nil {
		return nil, errors.New("category does not exist")
	}
	return cat, nil
}

func NewCatalogService(catalogRepository repository.Catalog) Catalog {
	return &catalogService{
		catalogRepository: catalogRepository,
	}
}
