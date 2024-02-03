.PHONY: test

test: clean_cache
	go test ./...

clean_cache:
	go clean -testcache
