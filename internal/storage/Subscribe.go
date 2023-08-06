package storage

import "context"

type Subscribe interface {
	SubscribePing(ctx context.Context, nameMention, lvlkz, CorpName string, tipPing int, TgChannel string) string
	CheckSubscribe(ctx context.Context, name, lvlkz string, TgChannel string, tipPing int) int
	Subscribe(ctx context.Context, name, nameMention, lvlkz string, tipPing int, TgChannel string)
	Unsubscribe(ctx context.Context, name, lvlkz string, TgChannel string, tipPing int)
}
