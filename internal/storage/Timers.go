package storage

import (
	"context"
	"kz_bot/internal/models"
)

type Timers interface {
	UpdateMitutsQueue(ctx context.Context, name, CorpName string) models.Sborkz
	TimerInsert(ctx context.Context, dsmesid, dschatid string, tgmesid int, tgchatid int64, timed int)
	TimerDeleteMessage(ctx context.Context) []models.Timer
	MinusMin(ctx context.Context) []models.Sborkz
}
