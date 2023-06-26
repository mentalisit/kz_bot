package storage

import "kz_bot/internal/models"

type BridgeConfig interface {
	DBReadBridgeConfig() []models.BridgeConfig
	UpdateBridgeChat(br models.BridgeConfig)
	InsertBridgeChat(br models.BridgeConfig)
}
