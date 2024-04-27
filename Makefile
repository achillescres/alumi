new:
	go run -mod=mod entgo.io/ent/cmd/ent new --target internal/domain/ent/schema $(filter-out $@,$(MAKECMDGOALS))
%:
	@:

#gen:
	#go generate ./internal/domain/ent
gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert,sql/schemaconfig,intercept ./ent/schema

build:
	go build cmd/app/main.go

run:
	PROJECT_ABS_PATH=" ENV_FILE_PATH="./.env" go run cmd/app/main.go

rapi:
	go run cmd/rawapi/rawapi.go

gr: gen run

atlas:
	atlas schema inspect \
      -u "ent://internal/domain/ent/schema/" \
      --dev-url "sqlite://file?mode=memory&_fk=1" \
      -w

swagup:
	swag init --parseDependency -g docs/load/load.go

swagapi: swagup rapi
