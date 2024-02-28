package flows

import (
	"github.com/developertom01/klaviyo-go/models"
	"github.com/jaswdr/faker/v2"
)

func mockFlow() models.Flow {
	fake := faker.New()

	name := fake.Blood().Name()
	var status models.FlowsStatus = models.FlowsStatusDraft
	return models.Flow{
		Type: "flow",
		ID:   fake.UUID().V4(),
		Attributes: models.FlowAttributes{
			Name:   &name,
			Status: &status,
			Links: models.DataLinks{
				Self: fake.Internet().URL(),
			},
		},
	}
}

func mockFlowResource() models.FlowResource {
	return models.FlowResource{
		Data: mockFlow(),
	}
}

func mockFlowsCollectionResource(n int) models.FlowCollectionResource {
	data := make([]models.Flow, 0)
	for i := 0; i < n; i++ {
		data = append(data, mockFlow())
	}

	return models.FlowCollectionResource{
		Data:  data,
		Links: models.MockedLinkResponse(),
	}
}

func mockUpdateFlowStatus() UpdateFlowStatusPayload {
	fake := faker.New()

	return UpdateFlowStatusPayload{
		Type: "flow",
		ID:   fake.UUID().V4(),
	}
}
