package hades

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"kz_bot/internal/clients"
	"kz_bot/internal/hades/ReservCopyPaste"
	"kz_bot/internal/hades/server"
	"kz_bot/internal/models"
	"kz_bot/internal/storage"
	"os"
	"os/exec"
	"time"
)

type Hades struct {
	cl         *clients.Clients
	storage    *storage.Storage
	toGame     chan models.Message
	toMessager chan models.Message
}

func NewHades(client *clients.Clients, storage *storage.Storage) *Hades {
	h := &Hades{
		cl:         client,
		storage:    storage,
		toGame:     make(chan models.Message, 10),
		toMessager: make(chan models.Message, 10),
	}
	server.NewServer(h.toGame, h.toMessager)

	go h.inbox()
	go ReservCopyPaste.RunReserv()
	return h
}
func (h *Hades) inbox() {
	for {
		select {
		case in := <-h.toMessager:
			h.filterGame(in)
		case in := <-h.cl.Ds.ChanToGame:
			h.filterDs(in)
		case in := <-h.cl.Tg.ChanToGame:
			h.filterTg(in)

		default:
			time.Sleep(500 * time.Millisecond)
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
