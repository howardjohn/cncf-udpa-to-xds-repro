package main

import (
	udpa "github.com/cncf/xds/go/udpa/type/v1"
	_ "google.golang.org/grpc/xds"
)

func main() {
	_ = udpa.TypedStruct{}
}
