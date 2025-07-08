package main

import (
	"github.com/PlakarKorp/integration-rclone/importer/rclone"
	"github.com/PlakarKorp/go-kloset-sdk/sdk"
)

func main() {
	if err := sdk.RunImporter(rclone.NewRcloneImporter); err != nil {
		panic(err)
	}
}
