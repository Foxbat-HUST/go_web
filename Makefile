include .$(pwd)/.env
.PHONY: gen_model
gen_model:
	gen --sqltype=mysql \
   	--connstr "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DB)" \
		--database=$(MYSQL_DB)  \
		--table=$(table) \
		--model=generated \
		--model_naming=$(model) \
		--exclude=schema_migrations \
		--templateDir=./infra/model/internal/gorm/template \
		--mapping=./infra/model/internal/gorm/template/mapping.json \
   	--gorm \
   	--out=./infra/model/internal/gorm/ \
   	--overwrite
	goimports -w ./infra/model/internal/gorm

.PHONY: tool
tool:
	go install github.com/smallnest/gen@latest
	go install golang.org/x/tools/cmd/goimports@latest