package catalog

import "github.com/developertom01/klaviyo-go/models"

type (
	CreateCatalogItemPayload struct {
		Data CreateCatalogItemPayloadData
	}

	CreateCatalogItemPayloadData struct {
		Type          string                                    `json:"type"` //catalog-item
		Attributes    CreateCatalogItemAttributesPayload        `json:"attributes"`
		Relationships CreateCatalogItemDataRelationshipsPayload `json:"relationships"`
	}

	CreateCatalogItemDataRelationshipsPayload struct {
		Categories models.RelationshipsCollectionRequestPayload `json:"categories"`
	}

	CreateCatalogItemAttributesPayload struct {
		ExternalId        string          `json:"external_id"`                   //The ID of the catalog item  in an external system.
		CatalogType       *string         `json:"catalog_type,omitempty"`        //The type of catalog. Currently only "$default" is supported.
		IntegrationType   *string         `json:"integration_type,omitempty"`    //The integration type. Currently only "$custom" is supported.
		Title             string          `json:"title"`                         //The title of the catalog item .
		Description       string          `json:"description"`                   //A description of the catalog item .
		Price             *int64          `json:"price,omitempty"`               //This field can be used to set the price on the catalog item , which is what gets displayed for the item  when included in emails. For most price-update use cases, you will also want to update the price on any parent items using the Update Catalog Item Endpoint.
		Url               int64           `json:"url"`                           //URL pointing to the location of the catalog item  on your website.
		ImageFullUrl      *string         `json:"image_full_url,omitempty"`      // URL pointing to the location of a full image of the catalog item .
		ImageThumbnailUrl *string         `json:"image_thumbnail_url,omitempty"` //URL pointing to the location of an image thumbnail of the catalog item
		Images            []string        `json:"images,omitempty"`              //List of URLs pointing to the locations of images of the catalog item .
		CustomMetadata    *map[string]any `json:"custom_metadata,omitempty"`     //Flat JSON blob to provide custom metadata about the catalog item . May not exceed 100kb.
		Published         *bool           `json:"published,omitempty"`           //Boolean value indicating whether the catalog item  is published.

	}
)

type (
	CreateCatalogItemVariantPayload struct {
		Data CreateCatalogItemVariantPayloadData
	}

	CreateCatalogItemVariantPayloadData struct {
		Type          string                                           `json:"type"` //catalog-item
		Attributes    CreateCatalogItemVariantAttributesPayload        `json:"attributes"`
		Relationships CreateCatalogItemVariantDataRelationshipsPayload `json:"relationships"`
	}

	CreateCatalogItemVariantDataRelationshipsPayload struct {
		Item models.RelationshipsRequestPayload `json:"items"`
	}

	CreateCatalogItemVariantAttributesPayload struct {
		ExternalId      string  `json:"external_id"`                //The ID of the catalog item variant in an external system.
		CatalogType     *string `json:"catalog_type,omitempty"`     //The type of catalog. Currently only "$default" is supported.
		IntegrationType *string `json:"integration_type,omitempty"` //The integration type. Currently only "$custom" is supported.
		Title           string  `json:"title"`                      //The title of the catalog item variant.
		Description     string  `json:"description"`                //A description of the catalog item variant.
		Sku             string  `json:"sku"`                        //The SKU of the catalog item variant
		//This field controls the visibility of this catalog item variant in product feeds/blocks. This field supports the following values:
		//1: a product will not appear in dynamic product recommendation feeds and blocks if it is out of stock.
		//0 or 2: a product can appear in dynamic product recommendation feeds and blocks regardless of inventory quantity.
		//Default: 2
		InventoryPolicy   *int64          `json:"inventory_policy,omitempty"`
		InventoryQuantity int64           `json:"inventory_quantity"`            //The quantity of the catalog item variant currently in stock.
		Price             int64           `json:"price"`                         //This field can be used to set the price on the catalog item variant, which is what gets displayed for the item variant when included in emails. For most price-update use cases, you will also want to update the price on any parent items using the Update Catalog Item Endpoint.
		Url               int64           `json:"url"`                           //URL pointing to the location of the catalog item variant on your website.
		ImageFullUrl      *string         `json:"image_full_url,omitempty"`      // URL pointing to the location of a full image of the catalog item variant.
		ImageThumbnailUrl *string         `json:"image_thumbnail_url,omitempty"` //URL pointing to the location of an image thumbnail of the catalog item variant
		Images            []string        `json:"images,omitempty"`              //List of URLs pointing to the locations of images of the catalog item variant.
		CustomMetadata    *map[string]any `json:"custom_metadata,omitempty"`     //Flat JSON blob to provide custom metadata about the catalog item variant. May not exceed 100kb.
		Published         *bool           `json:"published,omitempty"`           //Boolean value indicating whether the catalog item variant is published.

	}
)

type (
	UpdateCatalogItemPayload struct {
		Data UpdateCatalogItemPayloadData
	}
	UpdateCatalogItemPayloadData struct {
		Type         string                               `json:"type"` // catalog-item
		ID           string                               `json:"id"`   //The catalog item ID is a compound ID (string), with format: {integration}:::{catalog}:::{external_id}. Currently, the only supported integration type is $custom, and the only supported catalog is $default.
		Attributes   UpdateCatalogItemPayloadAttributes   `json:"attributes"`
		Relationship UpdateCatalogItemRelationshipPayload `json:"relationships"`
	}

	UpdateCatalogItemPayloadAttributes struct {
		Title             *string        `json:"title,omitempty"`               //The title of the catalog item.
		Price             *int64         `json:"price,omitempty"`               //This field can be used to set the price on the catalog item, which is what gets displayed for the item when included in emails. For most price-update use cases, you will also want to update the price on any child variants, using the Update Catalog Variant Endpoint.
		Description       *string        `json:"description,omitempty"`         //A description of the catalog item.
		Url               *string        `json:"url,omitempty"`                 //URL pointing to the location of the catalog item on your website.
		ImageFullUrl      *string        `json:"image_full_url,omitempty"`      //URL pointing to the location of a full image of the catalog item.
		ImageThumbnailUrl *string        `json:"image_thumbnail_url,omitempty"` //URL pointing to the location of an image thumbnail of the catalog item
		Images            []string       `json:"images,omitempty"`              //List of URLs pointing to the locations of images of the catalog item.
		CustomMetadata    map[string]any `json:"custom_metadata,omitempty"`     //Flat JSON blob to provide custom metadata about the catalog item. May not exceed 100kb.
		Published         *bool          `json:"published,omitempty"`           //Boolean value indicating whether the catalog item is published.
	}

	UpdateCatalogItemRelationshipPayload struct {
		Categories models.RelationshipDataCollection //ID: A list of catalog category IDs representing the categories the item is in
	}
)
