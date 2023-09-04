package HadesClient

import (
	"kz_bot/internal/models"
)

func (h *Hades) filterGame(in models.MessageHadesClient) {
	h.in = in
	//rename corp
	//h.in.Corporation = "TestCorp"

	if h.in.ChannelType == 0 {
		h.logicAlliance()
	} else if h.in.ChannelType == 1 {
		h.logicWs1()
	} else if h.in.ChannelType == 2 {
		//ws2
	}
}
