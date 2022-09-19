include .env
export

BUILD=./cmd/lake
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
	# Installing Binaries
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/tinylib/msgp@v1.1.6
	# Generating
	( cd $(APP_PATH) && go generate ./... )

# # format

.PHONY: format
format:
	# Installing Binaries
	go install github.com/daixiang0/gci@v0.7.0
	go install github.com/momaek/formattag@v0.0.8
	# Formatting
	(cd $(APP_PATH) && go fmt ./...)
	(cd $(APP_PATH) && find . -name "*.go" -exec formattag -file {} \;)
	(cd $(APP_PATH) && find . -name "*.go" -exec gci write {} \;)

.PHONY: ghz
ghz:
	ghz --insecure \
     --proto api/item_service/service/v1/item.proto \
      --call item_service.service.v1.ItemService/GetItem \
      -d '{"id":"$(ID)"}' \
      -n 100000 -c 4 \
      0.0.0.0:8080
