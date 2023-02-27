.PHONY: gen-api
gen-api:
	mkdir -p ./gen/api
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	oapi-codegen --config config/oapi-codegen/server.yaml ./api/openapi.yaml

.PHONY: go-fix-lint
go-fix-lint:
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/nagaseitteam/it-lime-backend"
