package main

import (
	"fmt"

	"github.com/HarshChakalasiya/krakend-private-repo/osutil"
)

func main() {
	result := osutil.GetEnvVar("test")
	// result := osutil.GetEnvVar("GOPRIVATE")
	fmt.Println(result)
}
