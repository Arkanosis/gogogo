BIN=gogogo
URI=go.arkanosis.net/gogogo
VERSION=0.0.1

SRC= \
	main.go

TARGET=target/$(BIN)

all: $(TARGET)

$(TARGET): $(SRC)
	mkdir -p "$(dir $@)"
	go build -o $@ -ldflags="-X main.version=$(VERSION)" $(URI)

clean:
	rm -rf target

.PHONY: all
