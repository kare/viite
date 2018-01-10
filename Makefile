IMPORT_PATH := kkn.fi/viite

GOMETALINTER := $(GOPATH)/bin/gometalinter

.PHONY: test
test:
	go test -v $(IMPORT_PATH)

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter --vendor ./...

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
