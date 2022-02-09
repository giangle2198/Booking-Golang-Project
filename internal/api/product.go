package api

import (
	"booking-center-service/internal/common"
	"booking-center-service/internal/dto"
	"booking-center-service/pb"
	"context"
)

func (s *api) GetListProduct(ctx context.Context, req *pb.GetListProductRequest) (res *pb.GetListProductResponse, err error) {
	var (
		resDTO *dto.GetListProductResponseDTO
	)

	resDTO, err = s.productUseCase.GetListProduct(ctx, req.Category)
	if err != nil {
		res = &pb.GetListProductResponse{
			StatusCode:    common.REJECTStatus,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}
		return res, nil
	}
	res = &pb.GetListProductResponse{}
	s.pbConverter.ToPb(res, resDTO)
	return res, nil
}

func (s *api) GetProductDetail(ctx context.Context, req *pb.GetProductDetailRequest) (res *pb.GetProductDetailResponse, err error) {
	var (
		resDTO *dto.GetProductDetailResponseDTO
	)

	resDTO, err = s.productUseCase.GetProductDetail(ctx, req.ProductId)
	if err != nil {
		res = &pb.GetProductDetailResponse{
			StatusCode:    common.REJECTStatus,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}
		return res, nil
	}
	res = &pb.GetProductDetailResponse{}
	s.pbConverter.ToPb(res, resDTO)
	return res, nil
}

func (s *api) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (res *pb.CreateProductResponse, err error) {
	var (
		reqDTO *dto.CreateProductRequestDTO = &dto.CreateProductRequestDTO{}
	)

	s.pbConverter.FromPb(reqDTO, req)
	err = s.productUseCase.CreateProduct(ctx, reqDTO)
	if err != nil {
		res = &pb.CreateProductResponse{
			StatusCode:    common.REJECTStatus,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}
		return res, nil
	}
	res = &pb.CreateProductResponse{
		StatusCode: common.ACCEPTStatus,
	}
	return res, nil
}

func (s *api) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (res *pb.UpdateProductResponse, err error) {
	var (
		reqDTO *dto.UpdateProductRequestDTO = &dto.UpdateProductRequestDTO{}
	)

	s.pbConverter.FromPb(reqDTO, req)
	err = s.productUseCase.UpdateProduct(ctx, reqDTO)
	if err != nil {
		res = &pb.UpdateProductResponse{
			StatusCode:    common.REJECTStatus,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}
		return res, nil
	}
	res = &pb.UpdateProductResponse{
		StatusCode: common.ACCEPTStatus,
	}
	return res, nil
}

func (s *api) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (res *pb.DeleteProductResponse, err error) {

	err = s.productUseCase.DeleteProduct(ctx, req.ProductId)
	if err != nil {
		res = &pb.DeleteProductResponse{
			StatusCode:    common.REJECTStatus,
			ReasonCode:    common.ParseError(err).Code(),
			ReasonMessage: common.ParseError(err).Message(),
		}
		return res, nil
	}
	res = &pb.DeleteProductResponse{
		StatusCode: common.ACCEPTStatus,
	}
	return res, nil
}
