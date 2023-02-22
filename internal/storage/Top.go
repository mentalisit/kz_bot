package storage

import "context"

type Top interface {
	TopLevel(ctx context.Context, CorpName, lvlkz string) bool
	TopEventLevel(ctx context.Context, CorpName, lvlkz string, numEvent int) bool
	TopTemp(ctx context.Context) string
	TopTempEvent(ctx context.Context) string
	TopAll(ctx context.Context, CorpName string) bool
	TopAllEvent(ctx context.Context, CorpName string, numberevent int) bool
}
