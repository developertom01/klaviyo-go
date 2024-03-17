package images

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/developertom01/klaviyo-go/common"
	"github.com/developertom01/klaviyo-go/models"
	"github.com/developertom01/klaviyo-go/options"
	"github.com/stretchr/testify/suite"
)

type ImagesApiTestSuite struct {
	suite.Suite
	api          ImagesApi
	mockedClient *common.MockHTTPClient
}

func (suit *ImagesApiTestSuite) SetupTest() {
	var apiKey = "test-key"
	opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
	session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())
	suit.mockedClient = common.NewMockHTTPClient()
	suit.api = NewImagesApi(session, suit.mockedClient)
}

func (suit *ImagesApiTestSuite) TestUploadImageFromFile() {
	var (
		name   = "image"
		hidden = true
	)

	mockedResponse := models.MockImageResponse()

	multipartFile := strings.NewReader("Some test reader")
	payload := UploadImageFromFilePayload{
		Name:   &name,
		Hidden: &hidden,
	}

	err := common.PrepareMockResponse(http.StatusOK, mockedResponse, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	resp, err := suit.api.UploadImageFromFile(context.Background(), multipartFile, payload)

	suit.Nil(err)
	suit.NotNil(resp)
}

func (suit *ImagesApiTestSuite) TestGetImages() {
	var pageSize = 3

	mockedResponse := models.MockImagesCollectionResponse(pageSize)
	err := common.PrepareMockResponse(http.StatusOK, mockedResponse, suit.mockedClient)
	if err != nil {
		suit.T().Fatal(err)
	}

	resp, err := suit.api.GetImages(context.Background(), "", nil)

	suit.Nil(err)
	suit.NotNil(resp)
	suit.Equal(pageSize, len(resp.Data))
}

func TestImagesApiTestSuite(t *testing.T) {
	suite.Run(t, new(ImagesApiTestSuite))
}
