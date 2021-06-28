.DEFAULT_GOAL := run

PROJ := flabild
SRCS := $(shell find $(CURDIR) -name '*.go')

$(PROJ): $(SRCS)
	go build -o $(PROJ) ./cmd/$(PROJ)/...

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

words_alpha.txt:
	curl -L https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt -o $@

# I guess I can be wittier here with target `./pkg/flabild/choices.go`, but meh...
.PHONY: choices
choices: words_alpha.txt
	rm -f ./pkg/flabild/choices.go
	go build -o generator ./cmd/generator/...
	./generator words_alpha.txt > ./pkg/flabild/choices.go
	go fmt ./pkg/flabild/choices.go
