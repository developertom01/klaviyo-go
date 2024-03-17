package models

import (
	"fmt"
	"strings"
	"time"
)

type (
	CatalogItem struct {
		Type         string                    `json:"type"` // catalog-item
		ID           string                    `json:"id"`   //The catalog item ID is a compound ID (string), with format: {integration}:::{catalog}:::{external_id}. Currently, the only supported integration type is $custom, and the only supported catalog is $default.
		Attributes   CatalogItemAttributes     `json:"attributes"`
		Links        DataLinks                 `json:"links"`
		Relationship *CatalogItemRelationships `json:"relationships"`
	}

	CatalogItemRelationships struct {
		Variant Relationships `json:"variant,omitempty"`
	}

	CatalogItemAttributes struct {
		ExternalId        *string         `json:"external_id,omitempty"`         //The ID of the catalog item in an external system.
		Title             *string         `json:"title,omitempty"`               // The title of the catalog item.
		Description       *string         `json:"description,omitempty"`         // The description of the catalog item.
		Price             *int64          `json:"price,omitempty"`               // This field can be used to set the price on the catalog item, which is what gets displayed for the item when included in emails. For most price-update use cases, you will also want to update the price on any child variants, using the Update Catalog Variant Endpoint.
		Url               *string         `json:"url,omitempty"`                 //URL pointing to the location of the catalog item on your website.
		ImageFullUrl      *string         `json:"image_full_url,omitempty"`      // URL pointing to the location of a full image of the catalog item.
		ImageThumbnailUrl *string         `json:"image_thumbnail_url,omitempty"` //URL pointing to the location of an image thumbnail of the catalog item
		Images            []string        `json:"images,omitempty"`              //List of URLs pointing to the locations of images of the catalog item.
		CustomMetadata    *map[string]any `json:"custom_metadata,omitempty"`     //Flat JSON blob to provide custom metadata about the catalog item. May not exceed 100kb.
		Published         *bool           `json:"published,omitempty"`           //Boolean value indicating whether the catalog item is published.
		Created           *time.Time      `json:"created,omitempty"`             //Date and time when the catalog item was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
		Updated           *time.Time      `json:"updated,omitempty"`             //Date and time when the catalog item was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	}
)

type CatalogItemField string

const (
	CatalogItemFieldExternalId        CatalogItemField = "external_id"
	CatalogItemFieldTitle             CatalogItemField = "title"
	CatalogItemFieldDescription       CatalogItemField = "description"
	CatalogItemFieldPrice             CatalogItemField = "price"
	CatalogItemFieldUrl               CatalogItemField = "url"
	CatalogItemFieldImage             CatalogItemField = "image_full_url"
	CatalogItemFieldImageThumbnailUrl CatalogItemField = "image_thumbnail_url"
	CatalogItemFieldImages            CatalogItemField = "images"
	CatalogItemFieldCustomMetadata    CatalogItemField = "custom_metadata"
	CatalogItemFieldPublished         CatalogItemField = "published"
	CatalogItemFieldCreated           CatalogItemField = "created"
	CatalogItemFieldUpdated           CatalogItemField = "updated"
)

func CatalogItemFieldParams(fields []CatalogItemField) string {
	if fields == nil {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[catalog-item]=%v", strings.Join(formattedFields, ","))
}

type (
	CatalogVariant struct {
		Type          string                       `json:"type"`       // catalog-variant
		ID            string                       `json:"id"`         //The catalog variant ID is a compound ID (string), with format: {integration}:::{catalog}:::{external_id}. Currently, the only supported integration type is $custom, and the only supported catalog is $default.
		Attributes    CatalogVariantAttributes     `json:"attributes"` //CatalogVariantAttributes
		Relationships *CatalogVariantRelationships `json:"relationships,omitempty"`
		Links         DataLinks                    `json:"links"`
	}

	CatalogVariantRelationships struct {
		Item Relationships `json:"item"`
	}

	CatalogVariantAttributes struct {
		ExternalId  *string `json:"external_id,omitempty"` //The ID of the catalog item variant in an external system.
		Title       *string `json:"title,omitempty"`       // The title of the catalog item variant.
		Description *string `json:"description,omitempty"` //A description of the catalog item variant.
		Sku         *string `json:"sku,omitempty"`         //The SKU of the catalog item variant.
		//This field controls the visibility of this catalog item variant in product feeds/blocks. This field supports the following values:
		//1: a product will not appear in dynamic product recommendation feeds and blocks if it is out of stock.
		//0 or 2: a product can appear in dynamic product recommendation feeds and blocks regardless of inventory quantity.
		InventoryPolicy   *int64          `json:"inventory_policy,omitempty"`
		InventoryQuantity *int64          `json:"inventory_quantity,omitempty"`  //The quantity of the catalog item variant currently in stock.
		Price             *int64          `json:"price,omitempty"`               //This field can be used to set the price on the catalog item variant, which is what gets displayed for the item variant when included in emails. For most price-update use cases, you will also want to update the price on any parent items using the Update Catalog Item Endpoint.
		Url               *string         `json:"url,omitempty"`                 //URL pointing to the location of the catalog item variant on your website.
		ImageFullUrl      *string         `json:"image_full_url,omitempty"`      // URL pointing to the location of a full image of the catalog item variant.
		ImageThumbnailUrl *string         `json:"image_thumbnail_url,omitempty"` //URL pointing to the location of an image thumbnail of the catalog item variant
		Images            []string        `json:"images,omitempty"`              //List of URLs pointing to the locations of images of the catalog item variant.
		CustomMetadata    *map[string]any `json:"custom_metadata,omitempty"`     //Flat JSON blob to provide custom metadata about the catalog item variant. May not exceed 100kb.
		Published         *bool           `json:"published,omitempty"`           //Boolean value indicating whether the catalog item variant is published.
		Created           *time.Time      `json:"created,omitempty"`             //Date and time when the catalog item variant was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
		Updated           *time.Time      `json:"updated,omitempty"`             //Date and time when the catalog item variant was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).

	}
)

type CatalogVariantField string

const (
	CatalogVariantFieldExternalId        CatalogVariantField = "external_id"
	CatalogVariantFieldTitle             CatalogVariantField = "title"
	CatalogVariantFieldDescription       CatalogVariantField = "description"
	CatalogVariantFieldSku               CatalogVariantField = "sku"
	CatalogVariantFieldInventoryPolicy   CatalogVariantField = "inventory_policy"
	CatalogVariantFieldInventoryQuantity CatalogVariantField = "inventory_quantity"
	CatalogVariantFieldPrice             CatalogVariantField = "price"
	CatalogVariantFieldUrl               CatalogVariantField = "url"
	CatalogVariantFieldImage             CatalogVariantField = "image_full_url"
	CatalogVariantFieldImageThumbnailUrl CatalogVariantField = "image_thumbnail_url"
	CatalogVariantFieldImages            CatalogVariantField = "images"
	CatalogVariantFieldCustomMetadata    CatalogVariantField = "custom_metadata"
	CatalogVariantFieldPublished         CatalogVariantField = "published"
	CatalogVariantFieldCreated           CatalogVariantField = "created"
	CatalogVariantFieldUpdated           CatalogVariantField = "updated"
)

func CatalogVariantFieldParams(fields []CatalogVariantField) string {
	if fields == nil {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[catalog-variant]=%v", strings.Join(formattedFields, ","))
}
