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

	client, err := NewClient()
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

func TestSignalDelete(t *testing.T) {
	id := -645

	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	err = client.DeleteSignalByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
}
