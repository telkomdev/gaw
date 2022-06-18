.PHONY : test format

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

test:
	$(foreach pkg, $(ALL_PACKAGES),\
	go test -race -v $(pkg);)

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w