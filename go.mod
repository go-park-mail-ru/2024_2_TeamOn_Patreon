module github.com/go-park-mail-ru/2024_2_TeamOn_Patreon

go 1.23

require github.com/gorilla/mux v1.8.1 // | >> go get -u github.com/gorilla/mux

require golang.org/x/crypto v0.27.0 // indirect | >> go get golang.org/x/crypto/bcrypt

require github.com/golang-jwt/jwt/v5 v5.2.1 // indirect | >> go get github.com/golang-jwt/jwt/v5

require github.com/stretchr/testify v1.9.0 // indirect  | >> go get github.com/stretchr/testify

require github.com/davecgh/go-spew v1.1.1 // indirect; indirect | >> go get github.com/davecgh/go-spew/spew

require ( // | >> go mod tidy
	github.com/pkg/errors v0.9.1 // indirect 	|| >> go get github.com/pkg/errors
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
