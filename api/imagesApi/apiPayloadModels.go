package images

type UploadImageFromFilePayload struct {
	Name   *string `json:"name,omitempty"`   //A name for the image. Defaults to the filename if not provided. If the name matches an existing image, a suffix will be added.
	Hidden *bool   `json:"hidden,omitempty"` //If true, this image is not shown in the asset library.
}

type (
	UploadImageFromUrlAttributes struct {
		ImportFromUrl string  `json:"import_from_url"`  // An existing image url to import the image from. Alternatively, you may specify a base-64 encoded data-uri (data:image/...). Supported image formats: jpeg,png,gif. Maximum image size: 5MB.
		Name          *string `json:"name,omitempty"`   //A name for the image. Defaults to the filename if not provided. If the name matches an existing image, a suffix will be added.
		Hidden        *bool   `json:"hidden,omitempty"` //If true, this image is not shown in the asset library.
	}

	UploadImageFromUrlPayload struct {
		Type       string                       `json:"type"`
		Attributes UploadImageFromUrlAttributes `json:"attributes"`
	}
)
