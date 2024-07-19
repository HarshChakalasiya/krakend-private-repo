package osutil

import (
	"syscall"

	"github.com/HarshChakalasiya/krakend-private-repo/errors"
	"github.com/HarshChakalasiya/krakend-private-repo/result"
)

func GetEnvVar(varName string) result.Result {
	val, found := syscall.Getenv(varName)
	if !found {
		return result.NewErr(errors.EnvVarNotDefined(varName, nil))
	} else {
		return result.NewOk(map[string]any{"message": val})
	}
}
