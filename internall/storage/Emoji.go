package storage

import (
	"context"
	"kz_bot/internall/models"
)

type Emoji interface {
	EmReadUsers(ctx context.Context, name, tip string) models.EmodjiUser
	EmUpdateEmodji(ctx context.Context, name, tip, slot, emo string) string
	EmInsertEmpty(ctx context.Context, tip, name string)
}
