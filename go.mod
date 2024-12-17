module github.com/go-park-mail-ru/2024_2_TeamOn_Patreon

go 1.23

require github.com/gorilla/mux v1.8.1 // | >> go get -u github.com/gorilla/mux

require golang.org/x/crypto v0.28.0 // indirect | >> go get golang.org/x/crypto/bcrypt

require github.com/golang-jwt/jwt/v5 v5.2.1 // indirect | >> go get github.com/golang-jwt/jwt/v5

require github.com/stretchr/testify v1.9.0 // indirect  | >> go get github.com/stretchr/testify

require github.com/davecgh/go-spew v1.1.1 // indirect; indirect | >> go get github.com/davecgh/go-spew/spew

require ( // | >> go mod tidy
	github.com/pkg/errors v0.9.1 // indirect 	|| >> go get github.com/pkg/errors
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/golang-migrate/migrate/v4 v4.18.1
	github.com/google/uuid v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0
	github.com/jackc/pgx/v5 v5.7.1
	github.com/mailru/easyjson v0.9.0
	github.com/microcosm-cc/bluemonday v1.0.27
	github.com/pborman/uuid v1.2.1
	github.com/prometheus/client_golang v1.20.5
	github.com/sahilm/fuzzy v0.1.1
	github.com/satori/go.uuid v1.2.0
	google.golang.org/grpc v1.64.1
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240513163218-0867130af1f8 // indirect
)
