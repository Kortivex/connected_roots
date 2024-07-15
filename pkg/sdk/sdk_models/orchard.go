package sdk_models

import "time"

type OrchardsBody struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Location   string         `json:"location"`
	Size       float64        `json:"size"`
	Soil       string         `json:"soil"`
	Fertilizer string         `json:"fertilizer"`
	Composting string         `json:"composting"`
	ImageURL   string         `json:"image_url"`
	UserID     string         `json:"user_id"`
	User       *UsersBody     `json:"user,omitempty"`
	CropTypeID string         `json:"crop_type_id"`
	CropType   *CropTypesBody `json:"crop_type,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type OrchardsResponse struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Location   string             `json:"location"`
	Size       float64            `json:"size"`
	Soil       string             `json:"soil"`
	Fertilizer string             `json:"fertilizer"`
	Composting string             `json:"composting"`
	ImageURL   string             `json:"image_url"`
	UserID     string             `json:"user_id"`
	User       *UsersResponse     `json:"user,omitempty"`
	CropTypeID string             `json:"crop_type_id"`
	CropType   *CropTypesResponse `json:"crop_type,omitempty"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

func (or *OrchardsResponse) ToOrchardBody() *OrchardsBody {
	return &OrchardsBody{
		ID:         or.ID,
		Name:       or.Name,
		Location:   or.Location,
		Size:       or.Size,
		Soil:       or.Soil,
		Fertilizer: or.Fertilizer,
		Composting: or.Composting,
		ImageURL:   or.ImageURL,
		UserID:     or.UserID,
		User:       or.User.ToUsersBody(),
		CropTypeID: or.CropTypeID,
		CropType:   or.CropType.ToCropTypeBody(),
		CreatedAt:  or.CreatedAt,
		UpdatedAt:  or.UpdatedAt,
	}
}
