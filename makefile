all: importer exporter

importer:
	@go build -o rclone-importer importer/main.go

exporter:
	@go build -o rclone-exporter exporter/main.go

clean:
	@rm -f rclone-importer rclone-exporter

.PHONY: all importer exporter clean