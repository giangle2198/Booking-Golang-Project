package usecase

import (
	"booking-center-service/config"
	"booking-center-service/internal/common"
	"booking-center-service/internal/dto"
	"booking-center-service/internal/repository"
	"context"
)

type (
	BaseUsecase interface {
		GetCountries(ctx context.Context) (*dto.GetCountriesResponseDTO, error)
		GetServices(ctx context.Context) (*dto.GetServicesResponse, error)
	}

	baseUsecase struct {
		cfg            *config.Config
		baseRepository repository.BaseRepository
	}
)

func NewBaseUsecase(cfg *config.Config,
	baseRepository repository.BaseRepository) BaseUsecase {
	return &baseUsecase{
		cfg:            cfg,
		baseRepository: baseRepository,
	}
}

func (u *baseUsecase) GetCountries(ctx context.Context) (res *dto.GetCountriesResponseDTO, err error) {

	var listCountry []*dto.ManagementDTO
	managements, err := u.baseRepository.GetManagementsByChannel(ctx, common.ChannelCountry)
	if err != nil {
		return nil, err
	}

	for _, management := range managements {
		listCountry = append(listCountry, &dto.ManagementDTO{
			ID:      management.ID,
			Name:    management.Name,
			Channel: management.Channel,
			Value:   management.Value,
		})
	}

	res = &dto.GetCountriesResponseDTO{
		BaseResponseDTO: dto.BaseResponseDTO{
			StatusCode: common.ACCEPTStatus,
		},
		Details: listCountry,
	}
	return res, nil
}
func (u *baseUsecase) GetServices(ctx context.Context) (res *dto.GetServicesResponse, err error) {
	var listService []*dto.ManagementDTO
	managements, err := u.baseRepository.GetManagementsByChannel(ctx, common.ChannelService)
	if err != nil {
		return nil, err
	}

	for _, management := range managements {
		listService = append(listService, &dto.ManagementDTO{
			ID:      management.ID,
			Name:    management.Name,
			Channel: management.Channel,
			Value:   management.Value,
		})
	}

	res = &dto.GetServicesResponse{
		BaseResponseDTO: dto.BaseResponseDTO{
			StatusCode: common.ACCEPTStatus,
		},
		Details: listService,
	}
	return res, nil
}
