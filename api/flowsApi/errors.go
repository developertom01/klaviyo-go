package flows

import (
	"errors"
)

var serializationError = errors.New("Serializing data failed")
var urlSerializationError = errors.New("Serializing url failed")
