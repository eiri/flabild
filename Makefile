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
	rm -f $(CURDIR)/generator
	rm -f $(CURDIR)/$(PROJ)
	rm -f $(CURDIR)/*.so

words_alpha.txt:
	curl -L https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt -o $@

# I guess I can be wittier here with target `./plugin/en/en.go`, but meh...
.PHONY: plugin
plugin: words_alpha.txt
	rm -f ./plugin/en/en.go
	go build -o generator ./cmd/generator/...
	./generator words_alpha.txt > ./plugin/en/en.go
	go fmt ./plugin/en/en.go
	go build -buildmode=plugin -o en.so ./plugin/en/en.go
