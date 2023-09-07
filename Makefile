build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/backupmanager cmd/main.go
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/amd64/backupmanager cmd/main.go
