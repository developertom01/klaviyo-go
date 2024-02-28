package flows

import (
	"errors"
)

var getFlowsApiCallError = errors.New("Getting account failed")
var serializationError = errors.New("Serializing data failed")
var urlSerializationError = errors.New("Serializing url failed")
