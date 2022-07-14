

ifeq (${OS},Windows_NT)
ROOT_PATH = $(shell cygpath -w $(shell pwd))
BIN = ${APP}.exe
else
ROOT_PATH = $(shell pwd)
endif

.PONY:debug
debug:
	go build -gcflags "all=-N -l" -o ./go_build_file_server .
	dlv --listen=:2345  --headless=true --api-version=2 --check-go-version=false --only-same-user=false exec ./go_build_file_server -- -listen :8181


.PONY:build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -gcflags="all=-trimpath=${ROOT_PATH}" -asmflags="all=-trimpath=${ROOT_PATH}" -o dist/linux/file-server .
