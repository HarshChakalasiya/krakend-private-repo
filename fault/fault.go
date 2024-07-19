package fault

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Fault struct {
	Message     string `json:"message"`
	FaultBundle *i18n.Bundle
}

func NewBasicFault(
	message string,

) Fault {
	return Fault{
		Message: message,
	}
}
