.DEFAULT_GOAL := run
.DELETE_ON_ERROR:

SRCS := $(shell find $(CURDIR) -name '*.go')

words_alpha.txt:
	curl -L https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt -o $@

generator:
	go build -o $@ ./cmd/generator/...

plugin/en/en.go: generator
	rm -f $@
	./$< words_alpha.txt > $@
	go fmt $@

en.so: plugin/en/en.go words_alpha.txt
	go build -buildmode=plugin -o en.so $<

flabild: $(SRCS) en.so
	go build -o $@ main.go

.PHONY: run
run: flabild
	./$<

.PHONY: clean
clean:
	rm -f generator
	rm -f flabild
	rm -f *.so
	go clean
