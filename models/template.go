package models

type (
	TemplateResponse struct {
		Data  []Template `json:"data"`
		Links Links      `json:"links"`
	}

	Template struct {
		Type       string             `json:"type"` //template
		ID         string             `json:"id"`
		Attributes TemplateAttributes `json:"attributes"`
		Links      DataLinks          `json:"links"`
	}

	TemplateAttributes struct {
		Name string `json:"name"` //The name of the template
		//editor_type has a fixed set of values:
		//SYSTEM_DRAGGABLE: indicates a drag-and-drop editor template
		//SIMPLE: A rich text editor template
		//CODE: A custom HTML template
		//USER_DRAGGABLE: A hybrid template, using custom HTML in the drag-and-drop editor
		EditorType EditorType `json:"editor_type"`
		HTML       string     `json:"html"`    //The rendered HTML of the template
		Text       *string    `json:"text"`    //The template plain_text
		Created    *string    `json:"created"` //The date the template was created in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm)
		Updated    *string    `json:"updated"` //The date the template was updated in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm)
	}
)

// editor_type has a fixed set of values:
// SYSTEM_DRAGGABLE: indicates a drag-and-drop editor template
// SIMPLE: A rich text editor template
// CODE: A custom HTML template
// USER_DRAGGABLE: A hybrid template, using custom HTML in the drag-and-drop editor
type EditorType string

const (
	EditorTypeSystemDraggable = "SYSTEM_DRAGGABLE" //indicates a drag-and-drop editor template
	EditorTypeSimple          = "SIMPLE"           //A rich text editor template
	EditorTypeCode            = "CODE"             //A custom HTML template
	EditorTypeUserDraggable   = "USER_DRAGGABLE"   // A hybrid template, using custom HTML in the drag-and-drop editor
)
