package utils

import (
	"fmt"
	"github.com/rclone/rclone/fs/config"
	"os"
)

func WriteRcloneConfigFile(name string, remoteMap map[string]string) (*os.File, error) {
	file, err := createTempConf()
	_, err = fmt.Fprintf(file, "[%s]\n", name)
	if err != nil {
		return nil, err
	}
	for k, v := range remoteMap {
		_, err = fmt.Fprintf(file, "%s = %s\n", k, v)
	}
	return file, nil
}

func createTempConf() (*os.File, error) {
	tempFile, err := os.CreateTemp("", "rclone-*.conf")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary config file: %w", err)
	}
	err = config.SetConfigPath(tempFile.Name())
	if err != nil {
		return nil, err
	}
	return tempFile, nil
}

func DeleteTempConf(name string) {
	err := os.Remove(name)
	if err != nil {
		fmt.Printf("Error removing temporary file: %v\n", err)
	}
}
