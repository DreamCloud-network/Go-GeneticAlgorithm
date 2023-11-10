package organelle

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/automata/basicmachinery"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

type Channel struct {
	Organelle

	channelThingType environment.ThingType

	chann chan *environment.Thing
}

func NewChannel(basicUnit *basicmachinery.BasicMachinery, channThingType environment.ThingType) {

	newChannel := &Channel{
		Organelle: *NewOrganelle(environment.Channel, basicUnit),

		channelThingType: channThingType,

		chann: make(chan *environment.Thing),
	}

	newChannel.SetObject(newChannel)

	basicUnit.GetInternalEnvironment().AddThing(&newChannel.Thing)
}

func (chann *Channel) CloseChannel() {
	close(chann.chann)
}

func (chann *Channel) SendToChannel(thing *environment.Thing) bool {
	if chann.channelThingType == thing.GetType() {
		chann.chann <- thing
		return true
	}
	return false
}

func (chann *Channel) Activate() {
	if chann.GetExternalEnvironment() == nil {
		log.Println("organelle.Channel.Activate(): Null external environment")
		return
	}

}
