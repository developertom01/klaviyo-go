package models

import (
	"time"

	"github.com/jaswdr/faker/v2"
)

type (
	ImageResponse struct {
		Data Image `json:"data"`
	}

	ImageCollectionResponse struct {
		Data  []Image `json:"data"`
		Links Links   `json:"links"`
	}

	Image struct {
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

func mockImage() Image {
	fake := faker.New()

	return Image{
		Type: "images",
		ID:   fake.UUID().V4(),
		Attributes: ImageAttributes{
			Name:      fake.App().Name(),
			ImageUrl:  fake.Internet().URL(),
			Size:      1024,
			Format:    "png",
			Hidden:    "false",
			UpdatedAt: time.Now(),
		},
	}
}

func MockImageResponse() ImageResponse {
	return ImageResponse{
		Data: mockImage(),
	}
}
