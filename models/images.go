package models

import "time"

type (
	ImageResponse struct {
		Data ImageData `json:"data"`
	}

	ImageCollectionResponse struct {
		Data  []ImageData `json:"data"`
		Links Links       `json:"links"`
	}

	ImageData struct {
		Type       string          `json:"type"` //image
		ID         string          `json:"id"`   //The ID of the image
		Attributes ImageAttributes `json:"attributes"`
		Links      DataLinks       `json:"links"`
	}

	ImageAttributes struct {
		Name      string    `json:"name"`
		ImageUrl  string    `json:"image_url"`
		Format    string    `json:"format"`
		Size      int       `json:"size"`
		Hidden    string    `json:"hidden"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

// ---- ImageField

type ImageField string

const (
	ImageFieldName      ImageField = "name"
	ImageFieldImageUrl  ImageField = "image_url"
	ImageFieldFormat    ImageField = "format"
	ImageFieldSize      ImageField = "size"
	ImageFieldHidden    ImageField = "hidden"
	ImageFieldUpdatedAt ImageField = "updated_at"
)

// ---- ImageField

type ImageSortField string

const (
	ImageSortFieldNameASC  ImageSortField = "name"
	ImageSortFieldNameDESC ImageSortField = "-name"

	ImageSortFieldImageUrlASC  ImageSortField = "image_url"
	ImageSortFieldImageUrlDESC ImageSortField = "-image_url"

	ImageSortFieldFormatASC  ImageSortField = "format"
	ImageSortFieldFormatDESC ImageSortField = "-format"

	ImageSortFieldSizeASC  ImageSortField = "size"
	ImageSortFieldSizeDESC ImageSortField = "-size"

	ImageSortFieldHiddenASC  ImageSortField = "hidden"
	ImageSortFieldHiddenDESC ImageSortField = "-hidden"

	ImageSortFieldUpdatedAtASC  ImageSortField = "updated_at"
	ImageSortFieldUpdatedAtDESC ImageSortField = "-updated_at"
)
