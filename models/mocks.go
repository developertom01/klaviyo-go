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
