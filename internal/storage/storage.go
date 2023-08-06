package storage

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/mongo"
	"kz_bot/internal/storage/postgres"
	"kz_bot/internal/storage/words"
)

type Storage struct {
	log   *logrus.Logger
	debug bool
	//	CorpsConfig       *CorpsConfig.Corps
	HadesClient       HadesClient
	BridgeConfig      BridgeConfig
	ConfigRs          ConfigRs
	TimeDeleteMessage TimeDeleteMessage
	Temp              *mongo.DB
	Words             *words.Words
	Subscribe         Subscribe
	Emoji             Emoji
	Count             Count
	Top               Top
	Update            Update
	Timers            Timers
	DbFunc            DbFunc
	//Cache            *postgres.Db
	Event            Event
	CorporationHades map[string]models.CorporationHadesClient
	BridgeConfigs    map[string]models.BridgeConfig
	CorpConfigRS     map[string]models.CorporationConfig
}

func NewStorage(log *logrus.Logger, cfg *config.ConfigBot) *Storage {

	//инициализируем и читаем репозиторий из облока конфига конфигурации
	mongoDB := mongo.InitMongoDB(log)

	//corp := CorpsConfig.NewCorps(log, cfg)

	//подключаю языковой пакет
	w := words.NewWords()

	//инициализируем локальный репозиторий
	local := postgres.NewDb(log, cfg)

	s := &Storage{
		//CorpsConfig:       corp,
		HadesClient:       mongoDB,
		BridgeConfig:      mongoDB,
		TimeDeleteMessage: mongoDB,
		ConfigRs:          mongoDB,
		Temp:              mongoDB,
		Words:             w,
		Subscribe:         local,
		Emoji:             local,
		Count:             local,
		Top:               local,
		Update:            local,
		Timers:            local,
		DbFunc:            local,
		Event:             local,
		CorporationHades:  make(map[string]models.CorporationHadesClient),
		BridgeConfigs:     make(map[string]models.BridgeConfig),
		CorpConfigRS:      make(map[string]models.CorporationConfig),
	}

	go s.loadDbArray()
	return s
}
func (s *Storage) loadDbArray() {
	corp := s.HadesClient.GetAllCorporationHades()
	for _, client := range corp {
		s.CorporationHades[client.Corp] = client
	}

	bc := s.BridgeConfig.DBReadBridgeConfig()
	for _, configBridge := range bc {
		s.BridgeConfigs[configBridge.NameRelay] = configBridge
	}
	rs := s.ConfigRs.ReadConfigRs()
	for _, r := range rs {
		s.CorpConfigRS[r.CorpName] = r
		fmt.Printf("ReadConfigRs %s %s %s %s\n", r.CorpName, r.DsChannel, r.TgChannel, r.WaChannel)
	}
}
func (s *Storage) ReloadDbArray() {
	s.CorpConfigRS = nil
	s.BridgeConfigs = nil
	s.CorporationHades = nil

	corp := s.HadesClient.GetAllCorporationHades()
	for _, client := range corp {
		s.CorporationHades[client.Corp] = client
	}

	bc := s.BridgeConfig.DBReadBridgeConfig()
	for _, configBridge := range bc {
		s.BridgeConfigs[configBridge.NameRelay] = configBridge
	}
	rs := s.ConfigRs.ReadConfigRs()
	for _, r := range rs {
		s.CorpConfigRS[r.CorpName] = r
	}
}
