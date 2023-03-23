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
		"received data": func(m dsl.Message) (interface{}, error) {
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
		"data exists": func(s dsl.State) error {
			data = &SQSEvent{
				SqsData: "12341234",
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
