.DEFAULT_GOAL := run
.DELETE_ON_ERROR:

SRCS := $(shell find $(CURDIR) -name '*.go')

english.txt:
	curl -L https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt -o $@

plugin/en/en.go: english.txt
	rm -f $@
	go generate ./...

en.so: plugin/en/en.go
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
	rm -rf plugin/*/*
	go clean
