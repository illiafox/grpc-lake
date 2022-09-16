include .env
export

BUILD=./app/cmd/lake
APP_PATH=server/

.PHONY: run
run:
	 cd $(APP_PATH)$(BUILD) && go run . $(ARGS)


# # docker compose
.PHONY: compose
compose: compose-down
	docker-compose up -d

.PHONY: compose-debug
compose-debug: compose-down
	docker-compose up -d --build

.PHONY: compose-down
compose-down:
	docker-compose down

# # tests
.PHONY: test
test:
	( cd $(APP_PATH) && go test -v ./... )

# # generate
.PHONY: generate
generate:
	( cd $(APP_PATH) && go generate ./... )


# # format

.PHONY: format
format: gci ftag

.PHONY: ftag
ftag:
	(cd $(APP_PATH) && find . -name "*.go" -exec formattag -file {} \;)

.PHONY: gci
gci:
	(cd $(APP_PATH) && find . -name "*.go" -exec gci write {} \;)

.PHONY: ghz
ghz:
	ghz --insecure \
     --proto api/item_service/service/v1/item.proto \
      --call item_service.service.v1.ItemService/GetItem \
      -d '{"id":"$(ID)"}' \
      -n 100000 -c 4 \
      0.0.0.0:8080
