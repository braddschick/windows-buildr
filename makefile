VERSION := 0.5.0

build: version macos

clean:
	@echo "Cleaning up"
	rm -rf bin

version:
	@echo $(VERSION)
	mkdir -p bin/$(VERSION)

macos:
	@echo "Building for MacOS"
	mkdir -p bin/$(VERSION)/macos/amd64
	mkdir -p bin/$(VERSION)/macos/arm64
	GOOS=darwin GOARCH=amd64 go build -o bin/$(VERSION)/macos/amd64/windows-buildr
	GOOS=darwin GOARCH=arm64 go build -o bin/$(VERSION)/macos/arm64/windows-buildr