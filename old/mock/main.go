package main

import (
	"./pubsub"
	"errors"
	"fmt"
)

// publisher is an interface to allow this package to mock the pubsub
// package support.
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct {
	host string
}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {
	fmt.Println(m.host)
	fmt.Println(key)

	err := errors.New("Bad call")
	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return err
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {
	fmt.Println(m.host)
	fmt.Println(key)
	// ADD YOUR MOCK FOR THE SUBSCRIBE CALL.
	return nil
}

func main() {

	// Create a slice of publisher interface values. Assign
	// the address of a pubsub.PubSub value and the address of
	// a mock value.
	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{"mocking localhost"},
	}

	// Range over the interface value to see how the publisher
	// interface provides the level of decoupling the user needs.
	// The pubsub package did not need to provide the interface type.
	for _, p := range pubs {
		p.Publish("pub key", "pub value")
		p.Subscribe("sub key")
	}
}
