package dto

type (
	ProductDetailDTO struct {
		Id          string
		Title       string
		Description string
		Img         string
		Categories  []string
		Sizes       []string
		Color       string
		Price       string
	}
	GetListProductRequestDTO  struct{}
	GetListProductResponseDTO struct {
		BaseResponseDTO
		Products []*ProductDetailDTO
	}
	GetProductDetailRequestDTO  struct{}
	GetProductDetailResponseDTO struct {
		BaseResponseDTO
		Product *ProductDetailDTO
	}
	CreateProductRequestDTO struct {
		Title       string
		Description string
		Img         string
		Categories  []string
		Sizes       []string
		Color       string
		Price       string
	}
	CreateProductResponseDTO struct {
		BaseResponseDTO
	}
	UpdateProductRequestDTO struct {
		Id          string
		Title       string
		Description string
		Img         string
		Categories  []string
		Sizes       []string
		Color       string
		Price       string
	}
	UpdateProductResponseDTO struct {
		BaseResponseDTO
	}
	DeleteProductRequestDTO  struct{}
	DeleteProductResponseDTO struct {
		BaseResponseDTO
	}
)
