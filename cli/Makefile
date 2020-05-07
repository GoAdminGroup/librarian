GOCMD = go
GOBUILD = $(GOCMD) build
BINARY_NAME = librarian
VERSION = v0.0.2
LAST_VERSION = v0.0.1

all: build-cli

build-cli:
	GO111MODULE=on $(GOBUILD) -ldflags "-w" -o ./build/mac/$(BINARY_NAME) ./...
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/linux/x86_64/$(BINARY_NAME) ./...
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o ./build/linux/armel/$(BINARY_NAME) ./...
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./build/windows/x86_64/$(BINARY_NAME).exe ./...
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o ./build/windows/i386/$(BINARY_NAME).exe ./...
	rm -rf ./build/linux/armel/librarian_linux_armel_$(LAST_VERSION).zip
	rm -rf ./build/linux/x86_64/librarian_linux_x86_64_$(LAST_VERSION).zip
	rm -rf ./build/windows/x86_64/librarian_windows_x86_64_$(LAST_VERSION).zip
	rm -rf ./build/windows/i386/librarian_windows_i386_$(LAST_VERSION).zip
	rm -rf ./build/mac/librarian_darwin_x86_64_$(LAST_VERSION).zip
	zip -qj ./build/linux/armel/librarian_linux_armel_$(VERSION).zip ./build/linux/armel/$(BINARY_NAME)
	zip -qj ./build/linux/x86_64/librarian_linux_x86_64_$(VERSION).zip ./build/linux/x86_64/$(BINARY_NAME)
	zip -qj ./build/windows/x86_64/librarian_windows_x86_64_$(VERSION).zip ./build/windows/x86_64/$(BINARY_NAME).exe
	zip -qj ./build/windows/i386/librarian_windows_i386_$(VERSION).zip ./build/windows/i386/$(BINARY_NAME).exe
	zip -qj ./build/mac/librarian_darwin_x86_64_$(VERSION).zip ./build/mac/$(BINARY_NAME)
	rm -rf ./build/zip/*
	cp ./build/linux/armel/librarian_linux_armel_$(VERSION).zip ./build/zip/
	cp ./build/linux/x86_64/librarian_linux_x86_64_$(VERSION).zip ./build/zip/
	cp ./build/windows/x86_64/librarian_windows_x86_64_$(VERSION).zip ./build/zip/
	cp ./build/windows/i386/librarian_windows_i386_$(VERSION).zip ./build/zip/
	cp ./build/mac/librarian_darwin_x86_64_$(VERSION).zip ./build/zip/