.DEFAULT_GOAL := run
.DELETE_ON_ERROR:

SRCS := $(shell find $(CURDIR) -name '*.go')

english.txt:
	curl -L https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt -o $@

russian.txt:
	curl -L https://raw.githubusercontent.com/danakt/russian-words/master/russian.txt -o $@
	iconv -f WINDOWS-1251 -t UTF-8 $@ > $@.tmp
	mv $@.tmp $@

plugins/en/en.go: english.txt
	rm -f $@
	go generate ./...

plugins/ru/ru.go: russian.txt
	rm -f $@
	go generate ./...

libflabild-en.so: plugins/en/en.go
	go build -buildmode=plugin -o $@ $<

libflabild-ru.so: plugins/ru/ru.go
	go build -buildmode=plugin -o $@ $<

flabild: $(SRCS) libflabild-en.so libflabild-ru.so
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
