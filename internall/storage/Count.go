package storage

import "context"

type Count interface {
	СountName(ctx context.Context, name, lvlkz, corpName string) (int, error)
	CountQueue(ctx context.Context, lvlkz, CorpName string) (int, error)
	CountNumberNameActive1(ctx context.Context, lvlkz, CorpName, name string) (int, error)
	CountNameQueue(ctx context.Context, name string) (countNames int)
	CountNameQueueCorp(ctx context.Context, name, corp string) (countNames int)
}
