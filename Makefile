BINARY_NAME=orchestrator

build:
	go build --ldflags '-extldflags "-Wl,--allow-multiple-definition"' -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

clean:
	go clean
	rm orchestrator
