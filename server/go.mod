module server

go 1.18

replace github.com/illiafox/grpc-lake/gen/go/api/item_service v0.0.0-20220820144355-8a74229d423f => ./../gen/go/api/item_service

require (
	github.com/dchenk/msgp/msgp v0.0.0-20180815194337-5c34657c2259
	github.com/go-redis/redis/v9 v9.0.0-beta.2
	github.com/illiafox/grpc-lake/gen/go/api/item_service v0.0.0-20220820144355-8a74229d423f
	github.com/ilyakaznacheev/cleanenv v1.3.0
	github.com/prometheus/client_golang v1.13.0
	github.com/stretchr/testify v1.8.0
	go.mongodb.org/mongo-driver v1.10.1
	go.uber.org/zap v1.22.0
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220805133916-01dd62135a58 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
