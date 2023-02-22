package storage

import (
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/internal/storage/CorpsConfig"
	"kz_bot/internal/storage/memory"
	"kz_bot/internal/storage/postgres"
	"kz_bot/internal/storage/words"
)

type Storage struct {
	log         *logrus.Logger
	debug       bool
	CorpsConfig *CorpsConfig.Corps
	Words       *words.Words
	Subscribe   Subscribe
	Emoji       Emoji
	Count       Count
	Top         Top
	Update      Update
	Timers      Timers
	DbFunc      DbFunc
	Cache       Cache
	Event       Event
}

func NewStorage(log *logrus.Logger, cfg *config.ConfigBot) *Storage {
	//
	mem := &memory.CorpConfig{}
	//инициализируем и читаем репозиторий из облока конфига конфигурации
	corp := CorpsConfig.NewCorps(log, cfg)
	corp.ReadCorps()

	//подключаю языковой пакет
	w := words.NewWords()

	//инициализируем локальный репозиторий
	local := postgres.NewDb(log, cfg)

	return &Storage{
		CorpsConfig: corp,
		Words:       w,
		Subscribe:   local,
		Emoji:       local,
		Count:       local,
		Top:         local,
		Update:      local,
		Timers:      local,
		DbFunc:      local,
		Event:       local,
		Cache:       mem,
	}
}
