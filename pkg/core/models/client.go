package models

import "time"

type Client struct {
	SentMessages int
	Address      string
	ConnectedAt  time.Time
}
