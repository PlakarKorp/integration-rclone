package main

import (
	"github.com/PlakarKorp/integration-rclone/exporter/rclone"
	"github.com/PlakarKorp/go-kloset-sdk/sdk"
)

func main() {
	if err := sdk.RunExporter(rclone.NewRcloneExporter); err != nil {
		panic(err)
	}
}
