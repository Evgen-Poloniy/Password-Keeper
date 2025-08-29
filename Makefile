GOOS = windows
GOARCH = amd64
BIN_DIR = bin64

ifeq ($(GOOS),windows)
    APP_NAME = Password\ keeper.exe
else
    APP_NAME = Password\ keeper
endif

.PHONY: comp run all

comp:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./$(BIN_DIR)/$(APP_NAME) ./cmd

run:
	./$(BIN_DIR)/$(APP_NAME)

all:
	go run ./cmd