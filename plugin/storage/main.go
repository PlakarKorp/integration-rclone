package main

import (
	"fmt"
	"os"

	sdk "github.com/PlakarKorp/go-kloset-sdk"
	"github.com/PlakarKorp/integration-rclone/storage"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Printf("Usage: %s\n", os.Args[0])
		os.Exit(1)
	}

	if err := sdk.RunStorage(storage.NewRcloneStorage); err != nil {
		panic(err)
	}
}
