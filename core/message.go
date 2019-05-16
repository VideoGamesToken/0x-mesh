package core

import (
	peer "github.com/libp2p/go-libp2p-peer"
)

type Message struct {
	// From is the peer ID of the peer who sent the message.
	From peer.ID
	// Data is the underlying data for the message.
	Data []byte
}

// MessageHandler is an interface responsible for validating and storing
// messages as well as selecting messages which are ready to be shared.
type MessageHandler interface {
	// ValidateAndStore validates the given messages and then stores and returns
	// any that are valid. It should only return an error if there was a problem
	// validating the messages. It should not return an error for invalid or
	// duplicate messages.
	ValidateAndStore([]*Message) ([]*Message, error)
	// GetMessagesToShare returns up to max messages to be shared with peers.
	GetMessagesToShare(max int) ([][]byte, error)
}
