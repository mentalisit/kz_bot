package storage

import (
	"fmt"
	"go.uber.org/zap"
	"kz_bot/internal/config"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/mongo"
	"kz_bot/internal/storage/postgres"
	"kz_bot/internal/storage/words"
	"kz_bot/pkg/logger"
)

type Storage struct {
	log   *zap.Logger
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

func NewStorage(log *logger.Logger, cfg *config.ConfigBot) *Storage {

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
	//var h = 0
	//var hades string
	//corp := s.HadesClient.GetAllCorporationHades()
	//for _, client := range corp {
	//	s.CorporationHades[client.Corp] = client
	//	h++
	//	hades = hades + fmt.Sprintf("%s, ", client.Corp)
	//}
	//fmt.Printf("Загружено конфиг хадеса %d : %s\n", h, hades)

	var b = 0
	var bridge string
	bc := s.BridgeConfig.DBReadBridgeConfig()
	for _, configBridge := range bc {
		s.BridgeConfigs[configBridge.NameRelay] = configBridge
		b++
		bridge = bridge + fmt.Sprintf("%s, ", configBridge.HostRelay)
	}
	fmt.Printf("Загружено конфиг мостов %d : %s\n", b, bridge)

	var c = 0
	var rslist string
	rs := s.ConfigRs.ReadConfigRs()
	for _, r := range rs {
		s.CorpConfigRS[r.CorpName] = r
		c++
		rslist = rslist + fmt.Sprintf("%s, ", r.CorpName)
	}
	fmt.Printf("Загружено конфиг RsBot %d : %s\n", c, rslist)
}
func (s *Storage) ReloadDbArray() {
	CorpConfigRS := make(map[string]models.CorporationConfig)
	BridgeConfigs := make(map[string]models.BridgeConfig)
	CorporationHades := make(map[string]models.CorporationHadesClient)

	s.CorpConfigRS = CorpConfigRS
	s.BridgeConfigs = BridgeConfigs
	s.CorporationHades = CorporationHades

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
