package images

import (
	"errors"
)

var imagesApiCallError = errors.New("Api Request failed")
var serializationError = errors.New("Serializing data failed")
var urlSerializationError = errors.New("Serializing url failed")
