package storage

import "kz_bot/internal/models"

type Event interface {
	NumActiveEvent(CorpName string) (event1 int)    //номер активного ивента
	NumDeactivEvent(CorpName string) (event0 int)   //номер предыдущего ивента
	UpdateActiveEvent0(CorpName string, event1 int) //отключение активного ивента
	EventStartInsert(CorpName string)               //включение ивента
	CountEventNames(CorpName, name string, numberkz, numEvent int) (countEventNames int)
	CountEventsPoints(CorpName string, numberkz, numberEvent int) int
	UpdatePoints(CorpName string, numberkz, points, event1 int) int
	ReadNamesMessage(CorpName string, numberkz, numberEvent int) (nd, nt models.Names, t models.Sborkz)
	NumberQueueEvents(CorpName string) int
}
