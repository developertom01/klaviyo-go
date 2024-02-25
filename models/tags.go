package models

import (
	"fmt"
	"strings"
)

type (
	TagsCollectionResponse struct {
		Data  []Tag `json:"data"`
		Links Links `json:"links"`
	}

	Tag struct {
		Type          string         `json:"type"` //tag
		ID            string         `json:"id"`   //The Tag ID
		Attributes    TagAttributes  `json:"attributes"`
		Links         DataLinks      `json:"links"`
		Relationships *Relationships `json:"relationships,omitempty"`
	}

	TagAttributes struct {
		Name string `json:"name"` //The Tag name
	}

	TafRelationships struct {
		TagGroup  *Relationships `json:"tag-group,omitempty"`
		Lists     *Relationships `json:"lists,omitempty"`
		Segments  *Relationships `json:"segments,omitempty"`
		Campaigns *Relationships `json:"campaigns,omitempty"`
		Flows     *Relationships `json:"flows,omitempty"`
	}
)

type TagField string

const (
	TagFieldName TagField = "name"
)

func BuildTagFieldParam(fields []TagField) string {
	if len(fields) == 0 {
		return ""
	}

	return strings.ReplaceAll(fmt.Sprintf("fields[tag]=%v", fields), " ", ",")
}
