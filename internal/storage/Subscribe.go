package storage

import "context"

type Subscribe interface {
	SubscribePing(ctx context.Context, nameMention, lvlkz, CorpName string, tipPing int, TgChannel int64) string
	CheckSubscribe(ctx context.Context, name, lvlkz string, TgChannel int64, tipPing int) int
	Subscribe(ctx context.Context, name, nameMention, lvlkz string, tipPing int, TgChannel int64)
	Unsubscribe(ctx context.Context, name, lvlkz string, TgChannel int64, tipPing int)
}
