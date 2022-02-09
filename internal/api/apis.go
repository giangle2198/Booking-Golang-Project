package api

import (
	"booking-center-service/config"
	"booking-center-service/internal/common"
	"booking-center-service/internal/dto"
	"booking-center-service/internal/helper"
	"booking-center-service/internal/usecase"
	"booking-center-service/pb"
	"context"
)

type (
	api struct {
		cfg            *config.Config
		pbConverter    helper.PbConverter
		baseUsecase    usecase.BaseUsecase
		productUseCase usecase.ProductUsecase
	}
)

func NewAPI(cfg *config.Config,
	pbConverter helper.PbConverter,
	baseUsecase usecase.BaseUsecase,
	productUsecase usecase.ProductUsecase,
) pb.APIServer {
	return &api{
		cfg:            cfg,
		pbConverter:    pbConverter,
		baseUsecase:    baseUsecase,
		productUseCase: productUsecase,
	}
}

func (s *api) GetCountries(ctx context.Context, req *pb.GetCountriesRequest) (res *pb.GetCountriesResponse, err error) {
	var (
		resDTO *dto.GetCountriesResponseDTO
	)

	resDTO, err = s.baseUsecase.GetCountries(ctx)
	if err != nil {
		res = &pb.GetCountriesResponse{
			StatusCode:    common.REJECTStatus,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}
		return res, nil
	}
	res = &pb.GetCountriesResponse{}
	s.pbConverter.ToPb(res, resDTO)
	return res, nil
}

func (s *api) GetServices(ctx context.Context, req *pb.GetServicesRequest) (res *pb.GetServicesResponse, err error) {
	var (
		resDTO *dto.GetCountriesResponseDTO
	)

	resDTO, err = s.baseUsecase.GetCountries(ctx)
	if err != nil {
		res = &pb.GetServicesResponse{
			StatusCode:    common.REJECTStatus,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}
		return res, nil
	}
	res = &pb.GetServicesResponse{}
	s.pbConverter.ToPb(res, resDTO)
	return res, nil
}
