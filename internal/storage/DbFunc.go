package storage

import (
	"context"
	"kz_bot/internal/models"
)

type DbFunc interface {
	ReadAll(ctx context.Context, lvlkz, CorpName string) (users models.Users)
	InsertQueue(ctx context.Context, dsmesid, wamesid, CorpName, name, nameMention, tip, lvlkz, timekz string, tgmesid, numkzN int)
	ElseTrue(ctx context.Context, name string) []models.Sborkz
	DeleteQueue(ctx context.Context, name, lvlkz, CorpName string)
	ReadMesIdDS(ctx context.Context, mesid string) (string, error)
	P30Pl(ctx context.Context, lvlkz, CorpName, name string) int
	UpdateTimedown(ctx context.Context, lvlkz, CorpName, name string)
	Queue(ctx context.Context, corpname string) []string
	OneMinutsTimer(ctx context.Context) []string
	MessageUpdateMin(ctx context.Context, corpname string) ([]string, []int)
	MessageupdateDS(ctx context.Context, dsmesid string, config models.CorporationConfig) models.InMessage
	MessageupdateTG(ctx context.Context, tgmesid int, config models.CorporationConfig) models.InMessage
	NumberQueueLvl(ctx context.Context, lvlkzs, CorpName string) (int, error)
}
