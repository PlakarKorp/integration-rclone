package main

import (
	"github.com/PlakarKorp/integration-rclone/exporter/googlephoto"
	"github.com/PlakarKorp/go-kloset-sdk/sdk"
)

func main() {
	if err := sdk.RunExporter(googlephoto.NewGooglePhotoExporter); err != nil {
		panic(err)
	}
}
