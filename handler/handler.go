package handler

import (
	"context"
	"fmt"
)

type Handler struct {
	timeGen string
}

type SQSEvent struct {
	SqsData string `json:"sqsData" pact:"exampleSQS"`
}

func NewHandler() *Handler {
	h := &Handler{
		timeGen: "random time",
	}

	return h
}

func (h *Handler) Handle(ctx context.Context, event SQSEvent) error {
	if len(event.SqsData) == 0 {
		return fmt.Errorf("no records found")
	}

	// ... actually handling the message using mocked dependencies (the same as in unit-tests)

	return nil
}
