package images

import (
	"context"
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	ImagesApi interface {
		//Upload an image from a file.
		//If you want to import an image from an existing url or a data uri, use the Upload Image From URL endpoint instead.
		UploadImage(ctx context.Context, payload UploadImageFromFilePayload) (*models.ImageResponse, error)
	}

	imageApi struct {
		session    common.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}
)

func NewImagesApi(session common.Session, httpClient common.HTTPClient) ImagesApi {
	var client common.HTTPClient
	if httpClient == nil {
		client = http.DefaultClient
	} else {
		client = httpClient
	}

	return &imageApi{
		session:    session,
		baseApiUrl: common.BASE_URL,
		revision:   common.API_REVISION,
		httpClient: client}
}

func (api *imageApi) UploadImage(ctx context.Context, payload UploadImageFromFilePayload) (*models.ImageResponse, error) {
	panic("")
}
