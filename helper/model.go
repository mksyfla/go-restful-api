package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryRespose {
	return web.CategoryRespose{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryRespose {
	var categoryResponses []web.CategoryRespose

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
