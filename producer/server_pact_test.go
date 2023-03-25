package producer

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider_PactContract(t *testing.T) {
	// Create Pact
	pact := dsl.Pact{}
	go startServer()

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8080",
		//PactURLs:        []string{"../consumer/pacts/consumer-provider.json"},
		BrokerURL: "http://localhost:9292",
		//		Tags:            []string{"master-message-service"},
		ProviderVersion:            "1.0.0",
		Provider:                   "Producer",
		FailIfNoPactsFound:         true,
		PublishVerificationResults: true,
		//PactURLs: []string{"http://localhost:9292/pacts/provider/Provider/consumer/Consumer/latest"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func startServer() {
	server := HandleRequests()
	fmt.Println("Starting Server on port 8080")
	server.ListenAndServe(":8080")
}
