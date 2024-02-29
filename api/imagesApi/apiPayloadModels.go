package images

type UploadImageFromFilePayload struct {
	Name   *string `json:"name,omitempty"`   //A name for the image. Defaults to the filename if not provided. If the name matches an existing image, a suffix will be added.
	Hidden *bool   `json:"hidden,omitempty"` //If true, this image is not shown in the asset library.
}
