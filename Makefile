.PHONY: gen-api
gen-api:
	mkdir -p ./gen/api
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	oapi-codegen -generate "server" -package api api/openapi.yaml > ./gen/api/server.gen.go
	oapi-codegen -generate "types" -package api api/openapi.yaml > ./gen/api/types.gen.go

.PHONY: go-fix-lint
go-fix-lint:
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/geeeeorge/Go-book-review"

.PHONY: __init-db-args
__init-db-args:
ifndef DB_HOST
	$(warning DB_HOST was not set; localhost is used)
	$(eval DB_HOST := localhost)
endif
ifndef DB_PORT
	$(warning DB_PORT was not set; 3306 is used)
	$(eval DB_PORT := 3306)
endif
ifndef DB_USER
	$(warning DB_USER was not set; root is used)
	$(eval DB_USER := root)
endif
ifndef DB_PASS
	$(warning DB_PASS was not set; passw0rd is used)
	$(eval DB_PASS := passw0rd)
endif
ifndef DB_NAME
	$(warning DB_NAME was not set; lime is used)
	$(eval DB_NAME := lime)
endif

.PHONY: db-migrate
db-migrate: __init-db-args
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	migrate -source "file://ddl" -database "mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up
