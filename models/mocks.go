package models

import "github.com/jaswdr/faker"

func mockTemplate() Template {
	fake := faker.New()

	return Template{
		Type: "template",
		ID:   fake.UUID().V4(),
		Attributes: TemplateAttributes{
			Name:       fake.Lorem().Word(),
			EditorType: EditorTypeCode,
			HTML:       "",
		},
	}
}

func MockTemplateResponse() TemplateResponse {
	return TemplateResponse{
		Data: mockTemplate(),
	}
}

func mockTag() Tag {
	fake := faker.New()

	return Tag{
		Type: "tag",
		ID:   fake.UUID().V4(),
		Attributes: TagAttributes{
			Name: fake.Lorem().Word(),
		}}
}

func MockTagsCollectionResponse(n int) TagsCollectionResponse {
	tagsData := make([]Tag, 0)

	for i := 0; i < n; i++ {
		tagsData = append(tagsData, mockTag())
	}

	return TagsCollectionResponse{
		Data:  tagsData,
		Links: MockedLinkResponse(),
	}
}

func MockRelationshipData(rType string) RelationshipData {
	fake := faker.New()

	return RelationshipData{
		Type: rType,
		ID:   fake.UUID().V4(),
	}
}

func MockRelationshipDataCollectionResponse(rType string, n int) RelationshipDataCollection {
	relationships := make([]RelationshipData, 0)
	for i := 0; i < n; i++ {
		relationships = append(relationships, MockRelationshipData(rType))
	}

	return RelationshipDataCollection{
		Data: relationships,
	}
}
