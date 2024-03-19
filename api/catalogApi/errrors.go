package catalog

import "errors"

var catalogApiCallError = errors.New("Catalog api failed")
var serializationError = errors.New("Serializing data failed")
var urlSerializationError = errors.New("Serializing url failed")
