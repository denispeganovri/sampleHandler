package handler

import (
	"context"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"os"
	"testing"
)

var (
	dir, _ = os.Getwd()
	logDir = fmt.Sprintf("%s/log", dir)
	pact   = createPact()
)

type testCase struct {
	desc             string
	req              SQSEvent
	expectedResponse error
}

// Set up the Pact client.
func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer: "Sample Consumer",
		Provider: "Sample Handler",
		LogDir:   logDir,
		PactDir:  "pacts",
		LogLevel: "INFO",
	}
}

func Test_genCon(t *testing.T) {
	eventMessage := SQSEvent{SqsData: "sample data"}

	tc := testCase{
		desc:             "test case description",
		req:              eventMessage,
		expectedResponse: nil,
	}

	msg := pact.AddMessage()
	msg.
		Given("given a SQS message").
		ExpectsToReceive("expected to receive no error").
		WithContent(tc.req).
		AsType(&SQSEvent{})

	err := pact.VerifyMessageConsumer(t, msg, handleWrapper)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}
}

func handleWrapper(m dsl.Message) error {
	h := NewHandler()
	err := h.Handle(context.Background(), *m.Content.(*SQSEvent))

	return err
}
