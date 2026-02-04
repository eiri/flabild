.DEFAULT_GOAL := run
.DELETE_ON_ERROR:

SRCS := $(shell find $(CURDIR) -name '*.go')

english.txt:
	curl -L https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt -o $@

plugins/en/en.go: english.txt
	rm -f $@
	go generate ./...

libflabild-en.so: plugins/en/en.go
	go build -buildmode=plugin -o $@ $<

flabild: $(SRCS) libflabild-en.so
	go build -o $@ ./cmd/$@/...

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run: flabild
	./$<

.PHONY: clean
clean:
	rm -f flabild
	rm -f *.so
	rm -rf plugins/*/*
	go clean
