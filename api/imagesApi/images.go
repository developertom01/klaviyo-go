package images

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
)

type (
	ImagesApi interface {
		//Upload an image from a file.
		//If you want to import an image from an existing url or a data uri, use the Upload Image From URL endpoint instead.
		UploadImage(ctx context.Context, fileByte *os.File, payload UploadImageFromFilePayload) (*models.ImageResponse, error)
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

func (api *imageApi) UploadImage(ctx context.Context, file *os.File, payload UploadImageFromFilePayload) (*models.ImageResponse, error) {

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	fileWriter, err := writer.CreateFormFile("file", os.TempDir())
	if err != nil {
		return nil, err
	}

	// Write the file data to the form file field
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	// Close the multipart writer
	writer.Close()
}
