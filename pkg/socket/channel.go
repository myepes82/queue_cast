package socket

type Channel struct {
	Topic   string
	Clients map[string]string
}

func NewChannel(topic string, clients map[string]string) *Channel {
	return &Channel{
		Topic:   topic,
		Clients: clients,
	}
}

func NewEmptyChannel(topic string) *Channel {
	return &Channel{
		Topic:   topic,
		Clients: map[string]string{},
	}
}

func (ch *Channel) AddClient(id string, client string) {
	ch.Clients[id] = client
}

func (ch *Channel) RemoveClient(id string) {
	delete(ch.Clients, id)
}

func (ch *Channel) GetClient(id string) (string, bool) {
	client, ok := ch.Clients[id]
	return client, ok
}
