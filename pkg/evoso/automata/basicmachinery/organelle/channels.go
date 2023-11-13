package organelle

import (
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

type Channels struct {
	openChannels []chan *environment.Thing
}

func NewChannels() *Channels {
	return &Channels{
		openChannels: make([]chan *environment.Thing, 0),
	}
}

func (channels *Channels) OpenNewChannel() chan *environment.Thing {
	// Create the new channel
	newChannel := make(chan *environment.Thing)

	// Add the new channel to the list of open channels
	channels.openChannels = append(channels.openChannels, newChannel)

	return newChannel
}
