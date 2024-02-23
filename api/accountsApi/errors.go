package accounts

import (
	"errors"
)

var getAccountApiCallError = errors.New("Getting account failed")
var serializationError = errors.New("Serializing data failed")
var urlSerializationError = errors.New("Serializing url failed")
