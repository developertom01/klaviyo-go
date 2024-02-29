package images

import "os"

type UploadImageFromFilePayload struct {
	File   os.File `json:"file"`             //The image file to upload. Supported image formats: jpeg,png,gif. Maximum image size: 5MB.
	Name   *string `json:"name,omitempty"`   //A name for the image. Defaults to the filename if not provided. If the name matches an existing image, a suffix will be added.
	Hidden *bool   `json:"hidden,omitempty"` //If true, this image is not shown in the asset library.
}
