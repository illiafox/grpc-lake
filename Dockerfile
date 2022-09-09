### Initial stage: download modules
FROM golang:1.19 as modules
# App dependencies
ADD server/go.mod server/go.sum /server/
# ADD server/go.mod go.sum /server/

# gRPC gen dependencies (replaced in go.mod)
COPY gen/go/api/item_service/go.mod /gen/go/api/item_service/go.mod
COPY gen/go/api/item_service/go.sum /gen/go/api/item_service/go.sum

# download packages
RUN cd /server && go mod download


### Intermediate stage: Build the binary
FROM golang:1.19 as builder

# copy all packages
COPY --from=modules /go/pkg /go/pkg

RUN useradd server

RUN mkdir -p /server
COPY server /server

RUN pwd
# grpc gen
COPY gen/go/api /gen/go/api
RUN ls -R

WORKDIR /server

# Build the binary with go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -ldflags '-s -w -extldflags "-static"' \
    -o /bin/server ./app/cmd/lake


### Final stage: Run the binary
# why not scratch? we need bash to connect to the container
FROM busybox:latest

COPY --from=builder /bin/server /server
ENTRYPOINT ["./server"]