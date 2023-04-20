package hades

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"kz_bot/internal/clients"
	"kz_bot/internal/hades/ReservCopyPaste"
	"kz_bot/internal/hades/server"
	"kz_bot/internal/models"
	"kz_bot/internal/storage/CorpsConfig/hades"
	"os"
	"os/exec"
	"time"
)

type Hades struct {
	cl *clients.Clients
}

func NewHades(client *clients.Clients) *Hades {
	NewToMessager := make(chan models.Message, 10)
	server.NewServer(client.ToGame, NewToMessager)
	h := &Hades{}
	h.cl = client
	go h.inbox(NewToMessager)
	go ReservCopyPaste.RunReserv()
	return h
}
func (h *Hades) inbox(toMess chan models.Message) {
	for {
		select {
		case in := <-toMess:
			h.filterGame(in)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func (h *Hades) filterGame(msg models.Message) {
	ok, corp := hades.HadesStorage.AllianceName(msg.Corporation)
	sender := "(🎮)" + msg.Sender
	if ok && msg.Command == "text" {
		if msg.ChannelType == 0 && corp.DsChat != "" {
			if h.ifEditMessage(msg, corp) {
				return
			}
			h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChat, corp.GuildId, msg.Avatar)
		}
		if msg.ChannelType == 1 && corp.DsChatWS1 != "" {
			h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChatWS1, corp.GuildId, msg.Avatar)
		}
		if msg.ChannelType == 2 && corp.DsChatWS2 != "" {
			h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChatWS2, corp.GuildId, msg.Avatar)
		}

		text := "(🎮)" + msg.Sender + ": " + msg.Text
		if msg.ChannelType == 0 && corp.TgChat != 0 {
			if h.ifEditMessage(msg, corp) {
				return
			}
			h.cl.Tg.SendChannel(corp.TgChat, text)
		}
		if msg.ChannelType == 1 && corp.TgChatWS1 != 0 {
			h.cl.Tg.SendChannel(corp.TgChatWS1, text)
		}
	} else if ok && msg.Command != "text" {
		if msg.Command == "ответ ds" {
			mesid := h.cl.Ds.SendWebhookForHades(msg.Text, sender, corp.DsChat, corp.GuildId, msg.Avatar)
			h.cl.Ds.DeleteMesageSecond(corp.DsChat, mesid, 180)
		}
		if msg.Command == "ответ tg" {
			h.cl.Tg.SendChannelDelSecond(corp.TgChat, msg.Text, 180)
		}
	}
}
func RestartedHadesBot() {
	name := "ConsoleClient.exe"
	procs, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	for _, proc := range procs {
		if proc.Executable() == name {
			fmt.Println("Restarting hsBot")
			cmd := exec.Command("taskkill", "/F", "/IM", name)
			_ = cmd.Run()
		}
	}
	if errd := os.Chdir("./hadeschat"); errd != nil {
		fmt.Println(errd)
		return
	}
	cmd := exec.Command("cmd.exe", "/c", "start", name)
	errs := cmd.Start()
	if errs != nil {
		fmt.Println(errs)
		return
	}
}
