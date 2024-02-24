package models

type TemplateResponse struct {
	Data  []Template            `json:"data"`
	Links TemplateResponseLinks `json:"links"`
}

type Template struct {
	Type       string             `json:"type"`
	ID         string             `json:"id"`
	Attributes TemplateAttributes `json:"attributes"`
	Links      TemplateLinks      `json:"links"`
}

type TemplateAttributes struct {
	Name       string `json:"name"`
	EditorType string `json:"editor_type"`
	HTML       string `json:"html"`
	Text       string `json:"text"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

type TemplateLinks struct {
	Self string `json:"self"`
}

type TemplateResponseLinks struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}
