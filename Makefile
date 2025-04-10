MAKEFLAGS += --no-print-directory
CLI_NAME=perlin_noise

build:
ifeq ($(OS),Windows_NT) 
	@$(MAKE) build/batch
else
	@$(MAKE) build/bash
endif

build/all:
ifeq ($(OS),Windows_NT) 
	@$(MAKE) build/all/batch
else
	@$(MAKE) build/all/bash
endif


# PLATFORM SPECIFIC
build/batch:
	set CGO_ENABLED=0&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}.exe ./cmd/main.go \

build/bash:
	CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME} ./cmd/main.go; \

build/all/batch:
	set CGO_ENABLED=0&& set GOOS=linux&& set GOARCH=386&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_linux_386 ./cmd/main.go
	set CGO_ENABLED=0&& set GOOS=linux&& set GOARCH=amd64&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_linux_amd64 ./cmd/main.go
	set CGO_ENABLED=0&& set GOOS=linux&& set GOARCH=arm&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_linux_arm ./cmd/main.go
	set CGO_ENABLED=0&& set GOOS=linux&& set GOARCH=arm64&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_linux_arm64 ./cmd/main.go

	set CGO_ENABLED=0&& set GOOS=darwin&& set GOARCH=amd64&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_darwin_amd64 ./cmd/main.go
	set CGO_ENABLED=0&& set GOOS=darwin&& set GOARCH=arm64&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_darwin_arm64 ./cmd/main.go

	set CGO_ENABLED=0&& set GOOS=windows&& set GOARCH=386&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_windows_386.exe ./cmd/main.go
	set CGO_ENABLED=0&& set GOOS=windows&& set GOARCH=amd64&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_windows_amd64.exe ./cmd/main.go
	set CGO_ENABLED=0&& set GOOS=windows&& set GOARCH=arm&& go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_windows_arm.exe ./cmd/main.go \

build/all/bash:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_linux_386 ./cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_linux_amd64 ./cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_linux_arm ./cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_linux_arm64 ./cmd/main.go

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_darwin_amd64 ./cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_darwin_arm64 ./cmd/main.go

	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_windows_386.exe ./cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_windows_amd64.exe ./cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -ldflags="-s -w" -installsuffix cgo -o ./bin/${CLI_NAME}_windows_arm.exe ./cmd/main.go \