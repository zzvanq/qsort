BINARY_NAME=qsort

run:
	go run main.go

build:
	go build -o ${BINARY_NAME} main.go

clean:
	rm ${BINARY_NAME}

.PHONY: run build clean
