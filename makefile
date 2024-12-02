
pre-build:
	go mod tidy
	go mod download
	mkdir -p bin
orm:
	go build -o bin/orm.exe ./cmd/extract-orm
