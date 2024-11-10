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
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang-migrate/migrate/v4 v4.18.1
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.7.1
	github.com/microcosm-cc/bluemonday v1.0.27
	github.com/pborman/uuid v1.2.1
)

require (
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)
