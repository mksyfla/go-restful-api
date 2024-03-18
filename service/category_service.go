package service

import (
	"context"
	"golang-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryRespose
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryRespose
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryRespose
	FindAll(ctx context.Context) []web.CategoryRespose
}
