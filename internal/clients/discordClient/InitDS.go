package discordClient

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	corpsConfig "kz_bot/internal/clients/corpConfig"
	"kz_bot/internal/dbase/dbaseMysql"
)

type Ds struct {
	d discordgo.Session
	corpsConfig.CorpConfig
	dbase dbaseMysql.Db
	log   *logrus.Logger
}

func (d *Ds) InitDS(TokenD string, db *sql.DB, log *logrus.Logger) {
	d.dbase.Db = db
	d.log = log
	DSBot, err := discordgo.New("Bot " + TokenD)
	if err != nil {
		d.log.Panic("Ошибка запуска дискорда", err)
	}

	DSBot.AddHandler(d.messageHandler)
	DSBot.AddHandler(d.MessageReactionAdd)

	err = DSBot.Open()
	if err != nil {
		d.log.Panic("Ошибка открытия ДС", err)
	}
	d.log.Println("Бот DISCORD запущен!!!")
	d.d = *DSBot
}

func (d *Ds) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	d.logicMixDiscord(m)

}

func (d *Ds) MessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	message, err := d.d.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		d.log.Println("Ошибка чтения реакции в ДС", err)
	}
	if message.Author.ID == s.State.User.ID {
		d.readReactionQueue(r, message)
	}
}
