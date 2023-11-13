package organelle

import "github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"

type RequestSignal struct {
	environment.Thing

	Requestor          *environment.Thing
	RequestedThingType environment.ThingType
	Channel            chan *environment.Thing
}

func (organelle *Organelle) newRequestSignal(requestedThingType environment.ThingType, channel chan *environment.Thing) *RequestSignal {

	newRequestMessage := &RequestSignal{
		Thing: *environment.NewThing(environment.Channel, nil),

		Requestor:          &organelle.Thing,
		RequestedThingType: requestedThingType,
		Channel:            channel,
	}

	newRequestMessage.SetObject(newRequestMessage)

	return newRequestMessage
}
