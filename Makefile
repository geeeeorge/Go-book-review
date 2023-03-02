.PHONY: gen-api
gen-api:
	mkdir -p ./gen/api
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	oapi-codegen -generate "server" -package api api/openapi.yaml > ./gen/api/server.gen.go
	oapi-codegen -generate "types" -package api api/openapi.yaml > ./gen/api/types.gen.go

.PHONY: go-fix-lint
go-fix-lint:
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/geeeeorge/Go-book-review"
