package handler

import (
	"path/filepath"

	"github.com/pact-foundation/pact-go/dsl"
	"testing"
)

var data *SQSEvent

func Test_verifyContract(t *testing.T) {
	pact := createPact()

	// Map test descriptions to message producer (handlers)
	functionMappings := dsl.MessageHandlers{
		//Description (ExpectedToReceive)
		"expected to receive no error": func(m dsl.Message) (interface{}, error) {
			if data != nil {
				return data, nil
			} else {
				return map[string]string{
					"message": "data is not found",
				}, nil
			}
		},
	}

	stateMappings := dsl.StateHandlers{
		//Provider state (Given)
		"given a SQS message": func(s dsl.State) error {
			//Content and Type
			data = &SQSEvent{
				SqsData: "sample SQS data",
			}

			return nil
		},
	}

	// Verify the Provider with local Pact Files
	_, err := pact.VerifyMessageProvider(t, dsl.VerifyMessageRequest{
		PactURLs:        []string{"./pacts/sample_consumer-sample_handler.json"},
		MessageHandlers: functionMappings,
		StateHandlers:   stateMappings,
		PactLogDir:      filepath.ToSlash("./logs"),
		PactLogLevel:    "DEBUG",
	})
	if err != nil {
		return
	}
}
