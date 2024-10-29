// Тестовая функция для генерации jwt токена

package main

import (
	"fmt"

	jwt "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/service/jwt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

func main() {
	user := bModels.User{
		UserID:   "9dc785df-17e3-43b9-9475-cf26ec4ac08b",
		Username: "maxround",
		Role:     "4bfa776c-3048-4291-8479-2a31a07f074f",
	}
	tokenStr, _ := jwt.CreateJWT(user, global.TTL)
	fmt.Println(tokenStr)
}
