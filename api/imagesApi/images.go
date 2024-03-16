package images

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	ImagesApi interface {
		//Upload an image from a file.
		//If you want to import an image from an existing url or a data uri, use the Upload Image From URL endpoint instead.
		UploadImageFromFile(ctx context.Context, file io.Reader, payload UploadImageFromFilePayload) (*models.ImageResponse, error)
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

func (api *imageApi) UploadImageFromFile(ctx context.Context, file io.Reader, payload UploadImageFromFilePayload) (*models.ImageResponse, error) {
	url := fmt.Sprintf("%s/api/image-upload/", api.baseApiUrl)

	var imageFieldName = "image"

	requestMeta := make(map[string]string)

	if payload.Name != nil {
		requestMeta["name"] = *payload.Name
	}
	if payload.Hidden != nil {
		if *payload.Hidden {
			requestMeta["hidden"] = "true"
		} else {
			requestMeta["hidden"] = "false"
		}
	}

	multipartOptions := common.MultipartOptions{
		File:          file,
		FileFieldName: *payload.Name,
		FileName:      imageFieldName,
		Meta:          requestMeta,
	}
	requestOptions := common.MultipartRequestOption{
		HttpClient: api.httpClient,
		Session:    api.session,
		Url:        url,
		Revision:   api.revision,
	}

	responseByteData, err := common.MakeMultipartRequest(ctx, requestOptions, multipartOptions)
	if err != nil {
		return nil, errors.Join(imagesApiCallError, err)
	}

	var imageUploadResponse models.ImageResponse
	err = json.Unmarshal(responseByteData, &imageUploadResponse)

	return &imageUploadResponse, err
}
