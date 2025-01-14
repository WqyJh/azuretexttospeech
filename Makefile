PROJECT=azuretexttospeech
DISTDIR := bin
BINARY := azuretexttospeech
REG := jesseward
VERSION := 2.0

all: clean generate vet test build
ci: clean vet test build

.PHONY: check_enumer
check_enumer:
	@if command -v your_command_name > /dev/null; then \
		echo "Command exists"; \
	else \
		go install github.com/dmarkham/enumer@latest; \
	fi

.PHONY: generate
generate: check_enumer
	go generate .

.PHONY: vet
vet:
	go vet .

.PHONY: test
test:
	go test -v -race

.PHONY: cleango
clean:
	@go clean -i ./...
	@rm -f $(DISTDIR)/$(BINARY)
	@rm -rf $(DISTDIR)
	@echo "✓ Cleaned build environment."

mkdirs:
	@mkdir -p $(DISTDIR)
	@echo "✓ Created bin directories"

_build_all:
	@go build -o $(DISTDIR)/$(BINARY) .
	@echo "✓ $(PROJECT) was built and copied to $(DISTDIR)/$(BINARY)"

.PHONY: build
build: clean mkdirs _build_all