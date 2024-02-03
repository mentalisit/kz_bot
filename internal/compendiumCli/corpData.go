package compendiumCli

import (
	"errors"
	"fmt"
	"kz_bot/internal/compendiumCli/Compendium"
	"kz_bot/internal/models"
	"kz_bot/pkg/logger"
)

type CompendiumData struct {
	c   *Compendium.Compendium
	log *logger.Logger
}

func GetCompendium(log *logger.Logger, code string, StorageKey string) (*CompendiumData, error) {
	c := Compendium.NewCompendium(log, StorageKey)
	compendiumData := CompendiumData{
		c:   c,
		log: log,
	}

	err := c.Initialize()
	if err != nil {
		log.ErrorErr(err)
		return nil, err
	}
	if c.Ident.Token == "" {
		c.Ident, err = c.CheckConnectCode(code)
		if err != nil {
			log.ErrorErr(err)
			return nil, err
		}

		log.Info(fmt.Sprintf("c.Ident %+v", c.Ident))
		connect, err := c.Connect(c.Ident)
		if err != nil {
			log.ErrorErr(err)
			return nil, err
		}
		c.Ident, err = c.Client.RefreshConnection(connect.Token)
		if err != nil {
			log.ErrorErr(err)
			return nil, err
		}
		c.WriteStorage()
		//log.InfoStruct("connect ", connect)

	}
	//log.Info(fmt.Sprintf("%+v\n", "ok"))
	return &compendiumData, nil
}
func (d *CompendiumData) GetAllRoles() ([]models.CorpRole, error) {
	data, err := d.c.CorpData("")
	if err != nil {
		d.log.ErrorErr(err)
		return nil, err
	}
	return data.Roles, nil
}
func (d *CompendiumData) GetRoleMembers(roleId string) ([]models.CorpMember, error) {
	data, err := d.c.CorpData(roleId)
	if err != nil {
		d.log.ErrorErr(err)
		return nil, err
	}
	return data.Members, nil
}
func (d *CompendiumData) GetMember(roleId, memberName string) (*models.CorpMember, error) {
	data, err := d.c.CorpData(roleId)
	if err != nil {
		d.log.ErrorErr(err)
		return nil, err
	}
	for _, member := range data.Members {
		if memberName == member.Name {
			return &member, err
		}
	}
	return nil, errors.New(fmt.Sprintf("пользователь %s не найден", memberName))
}
func (d *CompendiumData) GetGuild() *models.Guild {
	return &d.c.Ident.Guild
}
func (d *CompendiumData) Shutdown() {
	d.c.Shutdown()
}
