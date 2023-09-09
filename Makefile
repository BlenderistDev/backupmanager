build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/backupmanager cmd/main.go
build-linux-arm:
	GOOS=linux GOARCH=arm go build -o bin/linux/arm/backupmanager cmd/main.go
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/amd64/backupmanager cmd/main.go
build-mac-arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/mac/arm64/backupmanager cmd/main.go

pre-build:
	mkdir -p "bin/linux/amd64" "bin/linux/arm" "bin/mac/amd64" "bin/mac/arm64"

build: pre-build build-linux build-linux-arm build-mac build-mac-arm
