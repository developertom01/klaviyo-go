package images

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	ImagesApi interface {
		//Get all images in an account.
		GetImages(ctx context.Context, filterString string, options *GetImagesOptions) (*models.ImageCollectionResponse, error)
		//Get the image with the given image ID.
		GetImage(ctx context.Context, imageId string, fields []models.ImageField) (*models.ImageResponse, error)
		//Upload an image from a file.
		//If you want to import an image from an existing url or a data uri, use the UploadImageFromUrl instead.
		UploadImageFromFile(ctx context.Context, file io.Reader, payload UploadImageFromFilePayload) (*models.ImageResponse, error)
		//Import an image from a url or data uri.
		//If you want to upload an image from a file, use the Upload Image From File endpoint instead.
		UploadImageFromURL(ctx context.Context, payload UploadImageFromUrlPayload) (*models.ImageResponse, error)
		//Update the image with the given image ID.
		UpdateImage(ctx context.Context, imageId string, payload UpdateImagePayload) (*models.ImageResponse, error)
	}

	imageApi struct {
		session    common.Session
		baseApiUrl string
		revision   string
		httpClient common.HTTPClient
	}

	GetImagesOptions struct {
		PageCursor *string //For more information please visit https://developers.klaviyo.com/en/v2024-02-15/reference/api-overview#pagination
		Sort       *models.ImageSortField
		pageSize   *int //Default: 20. Min: 1. Max: 100.
		Fields     []models.ImageField
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

func buildGetImagesOptionsParams(filter string, opt *GetImagesOptions) string {
	if opt == nil {
		return filter
	}
	var params = []string{filter}

	if opt.Fields != nil {
		params = append(params, models.BuildImageFieldParam(opt.Fields))
	}

	if opt.PageCursor != nil {
		params = append(params, fmt.Sprintf("page[cursor]=%s", *opt.PageCursor))
	}
	if opt.Sort != nil {
		params = append(params, fmt.Sprintf("sort=%s", *opt.Sort))
	}

	if opt.pageSize != nil {
		params = append(params, fmt.Sprintf("page[size]=%d", *opt.pageSize))
	}

	return strings.Join(params, "&")
}

func (api *imageApi) GetImages(ctx context.Context, filterString string, options *GetImagesOptions) (*models.ImageCollectionResponse, error) {
	params := buildGetImagesOptionsParams(filterString, options)
	url := fmt.Sprintf("%s/api/images/?%s", api.baseApiUrl, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(imagesApiCallError, err)
	}

	var imageCollection models.ImageCollectionResponse
	err = json.Unmarshal(byteData, &imageCollection)

	return &imageCollection, err
}

func (api *imageApi) GetImage(ctx context.Context, imageId string, fields []models.ImageField) (*models.ImageResponse, error) {
	params := models.BuildImageFieldParam(fields)
	url := fmt.Sprintf("%s/api/images/%s/?%s", api.baseApiUrl, imageId, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	byteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(imagesApiCallError, err)
	}

	var imageResponse models.ImageResponse
	err = json.Unmarshal(byteData, &imageResponse)

	return &imageResponse, err
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

func (api *imageApi) UploadImageFromURL(ctx context.Context, payload UploadImageFromUrlPayload) (*models.ImageResponse, error) {
	url := fmt.Sprintf("%s/api/images/", api.baseApiUrl)

	reqPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	reqPayloadBuffer := bytes.NewBuffer(reqPayload)
	req, err := http.NewRequest(http.MethodPost, url, reqPayloadBuffer)
	if err != nil {
		return nil, err
	}

	responseByteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(imagesApiCallError, err)
	}

	var imageUploadResponse models.ImageResponse
	err = json.Unmarshal(responseByteData, &imageUploadResponse)

	return &imageUploadResponse, err
}

func (api *imageApi) UpdateImage(ctx context.Context, imageId string, payload UpdateImagePayload) (*models.ImageResponse, error) {
	url := fmt.Sprintf("%s/api/images/%s/", api.baseApiUrl, imageId)

	reqPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	reqPayloadBuffer := bytes.NewBuffer(reqPayload)
	req, err := http.NewRequest(http.MethodPatch, url, reqPayloadBuffer)
	if err != nil {
		return nil, err
	}

	responseByteData, err := common.RetrieveData(api.httpClient, req, api.session, api.revision)
	if err != nil {
		return nil, errors.Join(imagesApiCallError, err)
	}

	var imageUploadResponse models.ImageResponse
	err = json.Unmarshal(responseByteData, &imageUploadResponse)

	return &imageUploadResponse, err
}
