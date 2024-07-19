package errors

import (
	"github.com/HarshChakalasiya/krakend-private-repo/fault"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	ErrEnvVarUndefined = "Env not found"
)

func EnvVarNotDefined(varName string, cause error) fault.Fault {
	// data := map[string]any{"var_name": varName}
	return fault.NewBasicFault(ErrEnvVarUndefined)
}

var FaultBundle = faultBundle()

func faultBundle() *i18n.Bundle {

	var faultMgs = []*i18n.Message{
		{
			ID:    ErrEnvVarUndefined,
			One:   "Environment variable {{.var_name}} is not defined",
			Other: "Environment variable {{.var_name}} is not defined",
		},
	}

	b := i18n.NewBundle(language.English)
	b.MustAddMessages(language.English, faultMgs...)
	return b
}
