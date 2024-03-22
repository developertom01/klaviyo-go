package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/developertom01/klaviyo-go/exceptions"
)

type (
	CatalogItem struct {
		Type         string                    `json:"type"` // catalog-item
		ID           string                    `json:"id"`   //The catalog item ID is a compound ID (string), with format: {integration}:::{catalog}:::{external_id}. Currently, the only supported integration type is $custom, and the only supported catalog is $default.
		Attributes   CatalogItemAttributes     `json:"attributes"`
		Links        DataLinks                 `json:"links"`
		Relationship *CatalogItemRelationships `json:"relationships"`
	}

	CatalogItemCollectionResource struct {
		Data     []CatalogItem         `json:"data"`
		Links    Links                 `json:"links"`
		Included []CatalogItemIncluded `json:"included"`
	}

	CatalogItemResource struct {
		Data CatalogItem `json:"data"`
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

type CatalogItemIncluded map[string]any

func (inc *CatalogItemIncluded) ToCatalogVariant() (*CatalogVariant, bool) {
	byteData, err := json.Marshal(inc)
	if err != nil {
		return nil, false
	}

	var catalogVariant CatalogVariant
	err = json.Unmarshal(byteData, &catalogVariant)
	if err != nil {
		return nil, false
	}

	return &catalogVariant, true
}

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

type CatalogItemSortField string

const (
	CatalogItemSortFieldCreatedASC  CatalogItemSortField = "created"
	CatalogItemSortFieldCreatedDESC CatalogItemSortField = "-created"
)

type CatalogItemIncludedField string

const (
	CatalogItemIncludedFieldVariant CatalogItemIncludedField = "variant"
)

func BuildCatalogItemFieldParams(fields []CatalogItemField) string {
	if fields == nil {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[catalog-item]=%s", strings.Join(formattedFields, ","))
}

type (
	CatalogVariant struct {
		Type          string                       `json:"type"`       // catalog-variant
		ID            string                       `json:"id"`         //The catalog variant ID is a compound ID (string), with format: {integration}:::{catalog}:::{external_id}. Currently, the only supported integration type is $custom, and the only supported catalog is $default.
		Attributes    CatalogVariantAttributes     `json:"attributes"` //CatalogVariantAttributes
		Relationships *CatalogVariantRelationships `json:"relationships,omitempty"`
		Links         DataLinks                    `json:"links"`
	}

	CatalogVariantCollectionResource struct {
		Data  []CatalogVariant `json:"data"`
		Links Links            `json:"links"`
	}

	CatalogVariantResource struct {
		Data CatalogVariant `json:"data"`
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
		InventoryPolicy   *int64         `json:"inventory_policy,omitempty"`
		InventoryQuantity *int64         `json:"inventory_quantity,omitempty"`  //The quantity of the catalog item variant currently in stock.
		Price             *int64         `json:"price,omitempty"`               //This field can be used to set the price on the catalog item variant, which is what gets displayed for the item variant when included in emails. For most price-update use cases, you will also want to update the price on any parent items using the Update Catalog Item Endpoint.
		Url               *string        `json:"url,omitempty"`                 //URL pointing to the location of the catalog item variant on your website.
		ImageFullUrl      *string        `json:"image_full_url,omitempty"`      // URL pointing to the location of a full image of the catalog item variant.
		ImageThumbnailUrl *string        `json:"image_thumbnail_url,omitempty"` //URL pointing to the location of an image thumbnail of the catalog item variant
		Images            []string       `json:"images,omitempty"`              //List of URLs pointing to the locations of images of the catalog item variant.
		CustomMetadata    map[string]any `json:"custom_metadata,omitempty"`     //Flat JSON blob to provide custom metadata about the catalog item variant. May not exceed 100kb.
		Published         *bool          `json:"published,omitempty"`           //Boolean value indicating whether the catalog item variant is published.
		Created           *time.Time     `json:"created,omitempty"`             //Date and time when the catalog item variant was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
		Updated           *time.Time     `json:"updated,omitempty"`             //Date and time when the catalog item variant was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).

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

func BuildCatalogVariantFieldParams(fields []CatalogVariantField) string {
	if fields == nil {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[catalog-variant]=%s", strings.Join(formattedFields, ","))
}

type CatalogVariantSortField string

const (
	CatalogVariantSortFieldCreatedASC  CatalogVariantSortField = "created"
	CatalogVariantSortFieldCreatedDESC CatalogVariantSortField = "-created"
)

type (
	//Resource for bulk catalog creation or update
	CatalogItemBulkJobResource struct {
		Data CatalogItemBulkJob `json:"data"`
	}
	//Collection Resource for bulk catalog creation or update
	CatalogItemBulkJobCollectionResource struct {
		Data  []CatalogItemBulkJob `json:"data"`
		Links Links                `json:"links"`
	}

	//Resource for bulk catalog creation or update. Type is `catalog-item-bulk-create-job` if the job is creation.
	//Type is `catalog-item-bulk-update-job` if job is update
	CatalogItemBulkJob struct {
		Type          string                          `json:"type"` //catalog-item-bulk-create-job or  catalog-item-bulk-update-job
		ID            string                          `json:"id"`   //Unique identifier for retrieving the job. Generated by Klaviyo.
		Attributes    CatalogItemBulkJobAttributes    `json:"attributes"`
		Links         DataLinks                       `json:"links"`
		Relationships CatalogItemBulkJobRelationships `json:"relationships"`
	}

	CatalogItemBulkJobAttributes struct {
		Status         CatalogItemBulkJobStatus `json:"status"`          //Status of the asynchronous job.
		CreatedAt      time.Time                `json:"created_at"`      //The date and time the job was created in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
		TotalCount     int64                    `json:"total_count"`     //The total number of operations to be processed by the job. See completed_count for the job's current progress.
		CompletedCount *int64                   `json:"completed_count"` //The total number of operations that have been completed by the job.
		FailedCount    *int64                   `json:"failed_count"`    //The total number of operations that have failed as part of the job.
		CompletedAt    *time.Time               `json:"completed_at"`    //Date and time the job was completed in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
		Errors         []exceptions.ApiError    `json:"errors"`          //Array of errors encountered during the processing of the job.
		ExpiredAt      *time.Time               `json:"expires_at"`      //Date and time the job expires in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	}

	CatalogItemBulkJobRelationships struct {
		Items RelationshipDataCollection `json:"items"`
	}
)

type CatalogItemBulkJobStatus string

const (
	CatalogItemBulkJobStatusCancelled  CatalogItemBulkJobStatus = "cancelled"
	CatalogItemBulkJobStatusComplete   CatalogItemBulkJobStatus = "complete"
	CatalogItemBulkJobStatusProcessing CatalogItemBulkJobStatus = "processing"
	CatalogItemBulkJobStatusQueued     CatalogItemBulkJobStatus = "queued"
)

type CatalogItemBulkJobField string

const (
	CatalogItemBulkJobFieldStatus         CatalogItemBulkJobField = "status"
	CatalogItemBulkJobFieldCreatedAt      CatalogItemBulkJobField = "created_at"
	CatalogItemBulkJobFieldTotalCount     CatalogItemBulkJobField = "total_count"
	CatalogItemBulkJobFieldCompletedCount CatalogItemBulkJobField = "completed_count"
	CatalogItemBulkJobFieldFailedCount    CatalogItemBulkJobField = "failed_count"
	CatalogItemBulkJobFieldCompletedAt    CatalogItemBulkJobField = "completed_at"
	CatalogItemBulkJobFieldErrors         CatalogItemBulkJobField = "errors"
	CatalogItemBulkJobFieldExpiredAt      CatalogItemBulkJobField = "expired_at"
)

func BuildCatalogItemBulkJobFieldParams(fields []CatalogItemBulkJobField) string {
	if fields == nil {
		return ""
	}

	var formattedFields []string
	for _, field := range fields {
		formattedFields = append(formattedFields, string(field))
	}

	return fmt.Sprintf("fields[catalog-item-bulk-create-job]=%s", strings.Join(formattedFields, ","))
}

type CatalogItemBulkJobIncludeField string

const (
	CatalogItemBulkJobIncludeFieldItem CatalogItemBulkJobIncludeField = "item"
)
