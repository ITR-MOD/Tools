default: pre-build orm halfscale clean-rgb

pre-build:
	go mod tidy
	go mod download
	mkdir -p bin
orm: pre-build
	go build -o bin/extract-orm.exe ./cmd/extract-orm

halfscale: pre-build
	go build -o bin/halfscale.exe ./cmd/halfscale

clean-rgb: pre-build
	go build -o bin/rbg/clean-red.exe ./cmd/clean-red
	go build -o bin/rbg/clean-green.exe ./cmd/clean-green
	go build -o bin/rbg/clean-blue.exe ./cmd/clean-blue
	go build -o bin/rbg/just-red.exe ./cmd/just-red
	go build -o bin/rbg/just-green.exe ./cmd/just-green
	go build -o bin/rbg/just-blue.exe ./cmd/just-blue