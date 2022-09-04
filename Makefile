include .env
export

BUILD=./app/cmd/lake

.PHONY: run
run:
	 cd server/$(BUILD) && go run .

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
	( cd server && go test -v ./... )

# # format

.PHONY: format
format: gci ftag

.PHONY: ftag
ftag:
	(cd server && find . -name "*.go" -exec formattag -file {} \;)

.PHONY: gci
gci:
	(cd server && find . -name "*.go" -exec gci write {} \;)