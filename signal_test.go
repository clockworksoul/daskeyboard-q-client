package qclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestSignalCreate(t *testing.T) {
	signal := NewSignalRequest("New Q app version available", "#FF0000", KeyQ).
		WithMessage("Q App version 3 is available.").
		WithEffect(EffectSetColor)

	client, err := New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	response, err := client.CreateSignal(ctx, signal)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))
}

func TestGetShadows(t *testing.T) {
	client, err := New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	response, err := client.GetShadows(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))
}
