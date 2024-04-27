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
	PROJECT_ABS_PATH="/Users/achillescres/petprojects/itamconnect" ENV_FILE_PATH="./.env" go run cmd/app/main.go

rapi:
	go run cmd/rawapi/rawapi.go

gr: gen run

atlas:
	atlas schema inspect \
      -u "ent://internal/domain/ent/schema/" \
      --dev-url "sqlite://file?mode=memory&_fk=1" \
      -w

YC_CR_HOST=cr.yandex
YC_CR_ID=crpsfpdb7teskvdbr700
YC_CR_URL=$(YC_CR_HOST)/$(YC_CR_ID)

buildx:
	: "${RELEASE_TAG?Need to set RELEASE_TAG}"
	# Build image
	docker buildx build --platform linux/amd64 --tag achillescres/alumi:$(RELEASE_TAG) --load . -f ./Dockerfile

	# Run image
	echo "Test run(locally), want to know is application at least can start. Normally doesn't(locally) start fully because of isolated db host"
	docker run --platform linux/amd64 --rm -p 7771:7771 --name alumi achillescres/alumi:$(RELEASE_TAG)
