package imagine

import "github.com/gofrs/uuid"

type Author struct {
	AuthorID       uuid.UUID
	AuthorUsername string
}
