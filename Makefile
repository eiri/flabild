.DEFAULT_GOAL := run

PROJ := flabild
SRCS := $(shell find $(CURDIR) -name '*.go')

$(PROJ): $(SRCS)
	go build -o $(PROJ) ./...

.PHONY: run
run: $(PROJ)
	$(CURDIR)/$(PROJ)

.PHONY: test
test:
	go test -v -race ./...

.PHONY: clean
clean:
	go clean
	rm -f $(CURDIR)/$(PROJ)
