package CorpsConfig

//type Corps struct {
//	client postgresqlS.Client
//	log    *logger.Logger
//	debug  bool
//	//db     *db.Repository
//}

//func NewCorps(log *logger.Logger, cfg *config.ConfigBot) *Corps {
//	//создаем клиента инет базы данных
//	client, err := postgresqlS.NewClient(context.Background(), log, 5, cfg)
//	if err != nil {
//		log.Fatalln("Ошибка подключения к облачной ДБ ", err)
//	}
//
//	//инициализируем репо
//	//repo := db.NewRepository(client, log)
//
//	return &Corps{
//		client: client,
//		log:    log,
//		debug:  cfg.IsDebug,
//		//db:     repo,
//	}
//}
