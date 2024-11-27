package controller

import (
	tModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/controller/models"
	_ "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/csat/controller/models/mapper"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"net/http"
)

func (h *Handler) CsatTableGet(w http.ResponseWriter, r *http.Request) {
	op := "csat.controller.CsatTableGet"

	ctx := r.Context()

	nDays := r.URL.Query().Get(QueryTimeName)

	stats, err := h.b.GetSTATByTime(ctx, nDays)
	if err != nil {
		logger.StandardResponse(ctx, err.Error(), global.GetCodeError(err), r.Host, op)
		w.WriteHeader(global.GetCodeError(err))
		// отправляем структуру ошибки
		utils.SendModel(&tModels.ModelError{Message: global.GetMsgError(err)}, w, op, ctx)
		return
	}

	tStats := mapper.MapStatsToTables(stats)

	w.WriteHeader(http.StatusOK)
	utils.SendModel(tStats, w, op, ctx)
}
