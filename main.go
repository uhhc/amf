package main

import (
	_ "net/http/pprof"

	"github.com/uhhc/amf/cmd"
)

// @title AMF API
// @version v1
// @description V1 API
// @host 127.0.0.1
// @BasePath /
func main() {
	cmd.Execute()
}
