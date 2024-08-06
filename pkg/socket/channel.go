package socket

import "queuecast/pkg/core/models"

type Channel struct {
	Topic   string
	Clients map[string]*models.Client
}

func NewChannel(topic string, clients map[string]*models.Client) *Channel {
	return &Channel{
		Topic:   topic,
		Clients: clients,
	}
}

func NewEmptyChannel(topic string) *Channel {
	return &Channel{
		Topic:   topic,
		Clients: make(map[string]*models.Client),
	}
}

func (ch *Channel) AddClient(id string, client *models.Client) {
	ch.Clients[id] = client
}

func (ch *Channel) RemoveClient(id string) {
	delete(ch.Clients, id)
}

func (ch *Channel) GetClient(id string) (*models.Client, bool) {
	client, ok := ch.Clients[id]
	return client, ok
}
