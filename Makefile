GOOS = windows
GOARCH = amd64
BIN_DIR = bin64

APP_NAME = Password-Keeper.exe

comp:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./$(BIN_DIR)/$(APP_NAME) ./cmd

run:
	./$(BIN_DIR)/$(APP_NAME)

all:
	go run ./cmd