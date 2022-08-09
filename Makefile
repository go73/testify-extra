.PHONY: default
default: test

.PHONY: test
test:
	@go test -v -coverprofile coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

.PHONE: fuzz
fuzz: FUZZTIME=3s
fuzz:
	@go test -v -run=FuzzEqualWithEqualInput -fuzz=FuzzEqualWithEqualInput -fuzztime=$(FUZZTIME) github.com/go73/testify-extra/assert
	@go test -v -run=FuzzNotEqualWithEqualInput -fuzz=FuzzNotEqualWithEqualInput -fuzztime=$(FUZZTIME) github.com/go73/testify-extra/assert
	@go test -v -run=FuzzEqualWithEqualInput -fuzz=FuzzEqualWithEqualInput -fuzztime=$(FUZZTIME) github.com/go73/testify-extra/require
	@go test -v -run=FuzzNotEqualWithEqualInput -fuzz=FuzzNotEqualWithEqualInput -fuzztime=$(FUZZTIME) github.com/go73/testify-extra/require

.PHONY: fmt
fmt:
	@gofmt -w $$(find . -name '*.go')

