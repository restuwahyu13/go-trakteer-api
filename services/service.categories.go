package services

import (
	"context"

	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/interfaces"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type CategoriesService = interfaces.ICategoriesService
type categoriesService struct {
	repository repositorys.CategoriesRepository
}

func NewCategoriesService(repository repositorys.CategoriesRepository) *categoriesService {
	return &categoriesService{repository: repository}
}

func (s *categoriesService) CreateService(ctx context.Context, body *dtos.DTOCategories) helpers.APIResponse {
	return s.repository.CreateRepository(ctx, body)
}

func (s *categoriesService) GetAllService(ctx context.Context, query *dtos.DTOCategoriesPagination) helpers.APIResponse {
	return s.repository.GetAllRepository(ctx, query)
}

func (s *categoriesService) GetByIdService(ctx context.Context, params *dtos.DTOCategoriesId) helpers.APIResponse {
	return s.repository.GetByIdRepository(ctx, params)
}

func (s *categoriesService) DeleteByIdService(ctx context.Context, params *dtos.DTOCategoriesId) helpers.APIResponse {
	return s.repository.DeleteByIdRepository(ctx, params)
}

func (s *categoriesService) UpdatedByIdService(ctx context.Context, body *dtos.DTOCategories, params *dtos.DTOCategoriesId) helpers.APIResponse {
	return s.repository.UpdatedByIdRepository(ctx, body, params)
}
