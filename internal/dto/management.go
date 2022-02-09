package dto

type (
	BaseResponseDTO struct {
		StatusCode    string
		ReasonCode    string
		ReasonMessage string
	}

	ManagementDTO struct {
		ID      uint32
		Name    string
		Value   string
		Channel string
	}

	GetCountriesResponseDTO struct {
		BaseResponseDTO
		Details []*ManagementDTO
	}

	GetServicesResponse struct {
		BaseResponseDTO
		Details []*ManagementDTO
	}
)
