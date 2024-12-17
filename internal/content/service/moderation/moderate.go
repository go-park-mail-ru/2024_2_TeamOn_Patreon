package moderation

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/sahilm/fuzzy"
	"strings"
)

var blacklist = []string{
	"хуй",
	"пизд",
	"залуп",
	"ебат",
	"бляд",
	"блят",
	"пидорас",
	"ебать",
	"ебан",
	"хуе",
}

func IsOkayText(text string) bool {
	words := strings.Split(text, " ")

	for _, black := range blacklist {
		res := fuzzy.Find(black, words)
		if len(res) > 0 {
			logger.StandardDebugF(context.Background(), "isokaytext", "fuzzy find:=%v", res)
			return false
		}
	}
	return true
}
