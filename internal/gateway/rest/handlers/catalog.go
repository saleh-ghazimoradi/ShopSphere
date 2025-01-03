package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/dto"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/helper"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service/serviceModels"
	"net/http"
	"strconv"
)

type CatalogHandler struct {
	catalogService service.Catalog
	authService    helper.Auth
}

func (c *CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	req := dto.Category{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return badRequestError(ctx, "invalid request body")
	}

	if err = c.catalogService.CreateCategory(context.Background(), &req); err != nil {
		return internalError(ctx, err)
	}

	return successMessage(ctx, "category created successfully", nil)
}

func (c *CatalogHandler) GetCategories(ctx *fiber.Ctx) error {

	cates, err := c.catalogService.GetCategories(context.Background())
	if err != nil {
		return errorMessage(ctx, http.StatusNotFound, err)
	}

	return successMessage(ctx, "categories", cates)
}

func (c *CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return badRequestError(ctx, "invalid category id")
	}

	cat, err := c.catalogService.GetCategory(context.Background(), id)
	if err != nil {
		return errorMessage(ctx, http.StatusNotFound, err)
	}

	return successMessage(ctx, "category", cat)
}

func (c *CatalogHandler) EditCategory(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	req := dto.UpdateCategory{}

	if err := ctx.BodyParser(&req); err != nil {
		return badRequestError(ctx, "invalid update request body")
	}

	updatedCate, err := c.catalogService.EditCategory(context.Background(), id, &req)
	if err != nil {

	}

	return successMessage(ctx, "edit category", updatedCate)
}

func (c *CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	if err := c.catalogService.DeleteCategory(context.Background(), id); err != nil {
		return errorMessage(ctx, http.StatusNotFound, err)
	}
	return successMessage(ctx, "Category deleted successfully", nil)
}

func (c *CatalogHandler) CreateProducts(ctx *fiber.Ctx) error {
	req := dto.Product{}
	if err := ctx.BodyParser(&req); err != nil {
		return badRequestError(ctx, "invalid request body")
	}

	user, _ := c.authService.GetCurrentUser(ctx)

	if err := c.catalogService.CreateProduct(context.Background(), &req, user); err != nil {
		return internalError(ctx, err)
	}

	return successMessage(ctx, "Create Product", nil)
}

func (c *CatalogHandler) GetProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	product, err := c.catalogService.GetProduct(context.Background(), id)
	if err != nil {
		return errorMessage(ctx, http.StatusNotFound, err)
	}
	return successMessage(ctx, "Product", product)
}

func (c *CatalogHandler) GetProducts(ctx *fiber.Ctx) error {
	products, err := c.catalogService.GetProducts(context.Background())
	if err != nil {
		return errorMessage(ctx, http.StatusNotFound, err)
	}
	return successMessage(ctx, "Products", products)
}

func (c *CatalogHandler) EditProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	req := dto.UpdateProduct{}
	if err := ctx.BodyParser(&req); err != nil {
		return badRequestError(ctx, "invalid request body")
	}

	user, _ := c.authService.GetCurrentUser(ctx)

	product, err := c.catalogService.EditProduct(context.Background(), id, &req, user)
	if err != nil {
		return internalError(ctx, err)
	}

	return successMessage(ctx, "EditProduct", product)
}

func (c *CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	req := dto.UpdateStock{}
	if err := ctx.BodyParser(&req); err != nil {
		return badRequestError(ctx, "invalid request body")
	}

	user, _ := c.authService.GetCurrentUser(ctx)

	product := serviceModels.Product{
		ID:     uint(id),
		Stock:  uint(*req.Stock),
		UserId: user.ID,
	}

	updatedProduct, err := c.catalogService.UpdateProductStock(context.Background(), &product)
	if err != nil {
		return internalError(ctx, err)
	}

	return successMessage(ctx, "Update Stock", updatedProduct)
}

func (c *CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	user, _ := c.authService.GetCurrentUser(ctx)
	if err := c.catalogService.DeleteProduct(context.Background(), id, user); err != nil {
		return errorMessage(ctx, http.StatusNotFound, err)
	}

	return successMessage(ctx, "Delete product", nil)
}

func NewCatalogHandler(catalogService service.Catalog, authService helper.Auth) *CatalogHandler {
	return &CatalogHandler{
		catalogService: catalogService,
		authService:    authService,
	}
}
