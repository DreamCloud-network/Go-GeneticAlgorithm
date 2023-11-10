package place

type PlaceConnection struct {
	origin      *Place
	destination *Place
	weight      float64
}

func NewEmptyConnection() PlaceConnection {
	return PlaceConnection{}
}

// Creates a new connection.
func NewConnection(origin *Place, destination *Place, weight float64) PlaceConnection {
	return PlaceConnection{
		origin:      origin,
		destination: destination,
		weight:      weight,
	}
}

// GetOrigin returns the origin of the connection.
func (c *PlaceConnection) GetOrigin() *Place {
	return c.origin
}

// GetDestination returns the destination of the connection.
func (c *PlaceConnection) GetDestination() *Place {
	return c.destination
}

// GetWeight returns the weight of the connection.
func (c *PlaceConnection) GetWeight() float64 {
	return c.weight
}
