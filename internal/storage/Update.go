package storage

import "context"

type Update interface {
	MesidTgUpdate(ctx context.Context, mesidtg int, lvlkz string, corpname string)
	MesidDsUpdate(ctx context.Context, mesidds, lvlkz, corpname string)

	UpdateCompliteRS(ctx context.Context, lvlkz string, dsmesid string, tgmesid int, wamesid string, numberkz int, numberevent int, corpname string)
}
