package validate

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
)

func Uuid(uuid string) error {
	if !utils.IsValidUUIDv4(uuid) {
		return global.ErrUuidIsInvalid
	}
	return nil
}
