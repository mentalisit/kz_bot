package storage

import (
	"context"
	"kz_bot/internal/models"
)

type Emoji interface {
	EmojiModuleReadUsers(ctx context.Context, name, tip string) models.EmodjiUser
	EmojiUpdate(ctx context.Context, name, tip, slot, emo string) string
	ModuleUpdate(ctx context.Context, name, tip, slot, moduleAndLevel string) string
	WeaponUpdate(ctx context.Context, name, tip, weapon string) string
	EmInsertEmpty(ctx context.Context, tip, name string)
}
