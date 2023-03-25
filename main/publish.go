package main

import (
	"log"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func main() {

	p := dsl.Publisher{}
	err := p.Publish(types.PublishRequest{
		PactURLs:        []string{"../consumer/pacts/*.json"},
		PactBroker:      "http://localhost:9292",
		ConsumerVersion: "1.0.0",
		Tags:            []string{"master-message-service"},
	})
	if err != nil {
		log.Fatalln("publish error")
	}
}
