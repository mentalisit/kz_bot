package Compendium

import (
	"kz_bot/internal/compendiumCli/bot_api"
	"kz_bot/internal/models"
	"kz_bot/pkg/logger"
	"time"
)

// TechLevels represents tech levels data structure
type TechLevels map[int]models.TechLevel

type Compendium struct {
	Client           bot_api.CompendiumApiClient // Assuming you have CompendiumApiClient defined
	Ident            *models.Identity
	LastRefresh      int64
	LastTokenRefresh int64
	SyncData         *models.SyncData
	Ticker           *time.Ticker
	StorageKey       string
	log              *logger.Logger
}
