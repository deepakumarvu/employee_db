package simpledb

import "errors"

var (
	KeyAlreadyPresent    = errors.New("Key already present")
	KeyAbsent            = errors.New("Key absent")
	InvalidLastEvalKeyID = errors.New("Invalid last ID")
)
