package consumer

import (
	"anchor_exercise/producer"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestConsumer(t *testing.T) {

	// Create Pact client
	pact := dsl.Pact{
		Consumer: "Consumer",
		Provider: "Producer",
		PactDir:  "./pacts",
	}
	// Shuts down Mock Service when done
	defer pact.Teardown()

	// Pass in your test case as a function to Verify()
	var test = func() error {
		url, _ := url.Parse(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		client := &http.Client{}
		consumer := &Consumer{
			BaseURL:    url,
			httpClient: client,
		}
		_, err := consumer.GetNode(1)
		return err
	}

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given("Node with id 1 exists").
		UponReceiving("A request to get /node/1").
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String("/node/1"),
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(producer.Node{}),
		})

	// Verify
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}

}
