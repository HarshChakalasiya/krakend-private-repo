package result

import "github.com/HarshChakalasiya/krakend-private-repo/fault"

type Result struct {
	Data  map[string]any `json:"data"`
	Error fault.Fault    `json:"error"`
}

func NewErr(err fault.Fault) Result {
	return Result{
		Error: err,
	}
}

func NewOk(data map[string]any) Result {
	return Result{
		Data: data,
	}
}
