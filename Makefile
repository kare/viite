IMPORT_PATH := kkn.fi/viite

.PHONY: test
test:
	go test -v $(IMPORT_PATH)

