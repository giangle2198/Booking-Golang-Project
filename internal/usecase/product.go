package usecase

import (
	"booking-center-service/config"
	"booking-center-service/internal/common"
	"booking-center-service/internal/dto"
	"booking-center-service/internal/repository"
	"booking-libs/log"
	"context"
)

type (
	ProductUsecase interface {
		GetListProduct(ctx context.Context, categories []string) (*dto.GetListProductResponseDTO, error)
		GetProductDetail(ctx context.Context, productID string) (*dto.GetProductDetailResponseDTO, error)
		CreateProduct(ctx context.Context, req *dto.CreateProductRequestDTO) error
		UpdateProduct(ctx context.Context, req *dto.UpdateProductRequestDTO) error
		DeleteProduct(ctx context.Context, productID string) error
	}

	productUsecase struct {
		cfg               *config.Config
		productRepository repository.ProductRepository
	}
)

func NewProductUsecase(cfg *config.Config,
	productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		cfg:               cfg,
		productRepository: productRepository,
	}
}

func (u *productUsecase) GetListProduct(ctx context.Context, categories []string) (res *dto.GetListProductResponseDTO, err error) {

	listProduct, err := u.productRepository.GetListProduct(ctx, categories)
	if err != nil {
		//TODO: log
		log.Logger.Errorf("Error while getting list product detail with err: %v", err)
		return &dto.GetListProductResponseDTO{
			BaseResponseDTO: dto.BaseResponseDTO{
				StatusCode:    common.REJECTStatus,
				ReasonCode:    common.ParseError(err).Code(),
				ReasonMessage: common.ParseError(err).Message(),
			},
		}, err
	}

	res = &dto.GetListProductResponseDTO{
		BaseResponseDTO: dto.BaseResponseDTO{
			StatusCode: common.ACCEPTStatus,
		},
	}

	var products []*dto.ProductDetailDTO
	for _, productItem := range listProduct {
		productDetail := &dto.ProductDetailDTO{
			Id:          productItem.ID.Hex(),
			Title:       productItem.Title,
			Description: productItem.Description,
			Img:         productItem.Img,
			Categories:  productItem.Categories,
			Sizes:       productItem.Sizes,
			Color:       productItem.Color,
			Price:       productItem.Price,
		}
		products = append(products, productDetail)
	}

	res.Products = products
	return res, nil
}
func (u *productUsecase) GetProductDetail(ctx context.Context, productID string) (res *dto.GetProductDetailResponseDTO, err error) {

	productDetail, err := u.productRepository.GetProductDetail(ctx, productID)
	if err != nil {
		//TODO: log
		log.Logger.Errorf("Error while getting product detail with err: %v", err)
		return &dto.GetProductDetailResponseDTO{
			BaseResponseDTO: dto.BaseResponseDTO{
				StatusCode:    common.REJECTStatus,
				ReasonCode:    common.ParseError(err).Code(),
				ReasonMessage: common.ParseError(err).Message(),
			},
		}, err
	}

	res = &dto.GetProductDetailResponseDTO{
		BaseResponseDTO: dto.BaseResponseDTO{
			StatusCode: common.ACCEPTStatus,
		},
	}

	product := &dto.ProductDetailDTO{
		Id:          productDetail.ID.Hex(),
		Title:       productDetail.Title,
		Description: productDetail.Description,
		Img:         productDetail.Img,
		Categories:  productDetail.Categories,
		Sizes:       productDetail.Sizes,
		Color:       productDetail.Color,
		Price:       productDetail.Price,
	}

	res.Product = product
	return res, nil
}
func (u *productUsecase) CreateProduct(ctx context.Context, req *dto.CreateProductRequestDTO) error {

	err := u.productRepository.CreateProduct(ctx, req)
	if err != nil {
		//TODO: log
		log.Logger.Errorf("Error while creating product detail with err: %v", err)
		return err
	}

	return nil
}
func (u *productUsecase) UpdateProduct(ctx context.Context, req *dto.UpdateProductRequestDTO) error {

	err := u.productRepository.UpdateProduct(ctx, req)
	if err != nil {
		//TODO: log
		log.Logger.Errorf("Error while updating product detail with err: %v", err)
		return err
	}

	return nil
}
func (u *productUsecase) DeleteProduct(ctx context.Context, productID string) error {

	err := u.productRepository.DeleteProduct(ctx, productID)
	if err != nil {
		//TODO: log
		log.Logger.Errorf("Error while deleting product detail with err: %v", err)
		return err
	}

	return nil
}
