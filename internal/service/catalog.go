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

	CreateProduct(ctx context.Context, input *dto.Product, user *serviceModels.User) error
	EditProduct(ctx context.Context, id int, input *dto.UpdateProduct, user *serviceModels.User) (*serviceModels.Product, error)
	DeleteProduct(ctx context.Context, id int, user *serviceModels.User) error
	GetProduct(ctx context.Context, id int) (*serviceModels.Product, error)
	GetProducts(ctx context.Context) ([]*serviceModels.Product, error)
	GetSellerProducts(ctx context.Context, id int) ([]*serviceModels.Product, error)
	UpdateProductStock(ctx context.Context, product *serviceModels.Product) (*serviceModels.Product, error)
}

type catalogService struct {
	catalogRepository repository.Catalog
}

func (c *catalogService) CreateCategory(ctx context.Context, input *dto.Category) error {
	return c.catalogRepository.CreateCategory(ctx, &serviceModels.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})
}

func (c *catalogService) EditCategory(ctx context.Context, id int, input *dto.UpdateCategory) (*serviceModels.Category, error) {
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

func (c *catalogService) DeleteCategory(ctx context.Context, id int) error {
	if err := c.catalogRepository.DeleteCategory(ctx, id); err != nil {
		return errors.New("category not found")
	}
	return nil
}

func (c *catalogService) GetCategories(ctx context.Context) ([]*serviceModels.Category, error) {
	categories, err := c.catalogRepository.FindCategories(ctx)
	if err != nil {
		return nil, errors.New("categories do not exist")
	}
	return categories, nil
}

func (c *catalogService) GetCategory(ctx context.Context, id int) (*serviceModels.Category, error) {
	cat, err := c.catalogRepository.FindCategoryById(ctx, id)
	if err != nil {
		return nil, errors.New("category does not exist")
	}
	return cat, nil
}

func (c *catalogService) CreateProduct(ctx context.Context, input *dto.Product, user *serviceModels.User) error {
	err := c.catalogRepository.CreateProduct(ctx, &serviceModels.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
		ImageUrl:    input.ImageUrl,
		UserId:      user.ID,
		Stock:       uint(input.Stock),
	})
	return err
}

func (c *catalogService) EditProduct(ctx context.Context, id int, input *dto.UpdateProduct, user *serviceModels.User) (*serviceModels.Product, error) {
	existProduct, err := c.catalogRepository.FindProductById(ctx, id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	if existProduct.UserId != user.ID {
		return nil, errors.New("user not exist")
	}

	if input.Name != nil {
		existProduct.Name = *input.Name
	}

	if input.Description != nil {
		existProduct.Description = *input.Description
	}

	if input.Price != nil {
		existProduct.Price = *input.Price
	}

	if input.CategoryId != nil {
		existProduct.CategoryId = *input.CategoryId
	}

	updatedProduct, err := c.catalogRepository.EditProduct(ctx, existProduct)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (c *catalogService) DeleteProduct(ctx context.Context, id int, user *serviceModels.User) error {
	existProduct, err := c.catalogRepository.FindProductById(ctx, id)
	if err != nil {
		return errors.New("product does not exist")
	}

	if existProduct.UserId != user.ID {
		return errors.New("you don't have management rights of this product")
	}

	err = c.catalogRepository.DeleteProduct(ctx, id)
	if err != nil {
		return errors.New("product cannot be deleted")
	}
	return nil
}

func (c *catalogService) GetProduct(ctx context.Context, id int) (*serviceModels.Product, error) {
	product, err := c.catalogRepository.FindProductById(ctx, id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}
	return product, nil
}

func (c *catalogService) GetProducts(ctx context.Context) ([]*serviceModels.Product, error) {
	products, err := c.catalogRepository.FindProducts(ctx)
	if err != nil {
		return nil, errors.New("products do not exist")
	}
	return products, nil
}

func (c *catalogService) GetSellerProducts(ctx context.Context, id int) ([]*serviceModels.Product, error) {
	products, err := c.catalogRepository.FindSellerProducts(ctx, id)
	if err != nil {
		return nil, errors.New("products do not exist")
	}
	return products, nil
}

func (c *catalogService) UpdateProductStock(ctx context.Context, product *serviceModels.Product) (*serviceModels.Product, error) {
	p, err := c.catalogRepository.FindProductById(ctx, int(product.ID))
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	if p.UserId != product.UserId {
		return nil, errors.New("product does not own user")
	}

	p.Stock = product.Stock
	editProduct, err := c.catalogRepository.EditProduct(ctx, p)
	if err != nil {
		return nil, err
	}

	return editProduct, nil
}

func NewCatalogService(catalogRepository repository.Catalog) Catalog {
	return &catalogService{
		catalogRepository: catalogRepository,
	}
}
