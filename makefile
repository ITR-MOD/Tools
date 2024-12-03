default: pre-build orm halfscale

pre-build:
	go mod tidy
	go mod download
	mkdir -p bin
orm: pre-build
	go build -o bin/extract-orm.exe ./cmd/extract-orm

halfscale: pre-build
	go build -o bin/halfscale.exe ./cmd/halfscale
