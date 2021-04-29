package qclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

const BackendURL = "http://localhost:27301/api/1.0/signals"

func TestTest(t *testing.T) {
	ctx := context.Background()

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequestWithContext(ctx, "POST", BackendURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	signal := Signal{
		ZoneId:     "KEY_Q",
		Color:      "#FFFFFF",
		Effect:     "SET_COLOR",
		Pid:        "DK5QPID",
		ClientName: "GoClient",
		IsMuted:    false,
		Message:    "Q App version 3 is available. Download it at https://www.daskeyboard.io/get-started/download/",
		Name:       "New Q app version available",
	}

	json, _ := json.MarshalIndent(signal, "", "  ")
	req.Body = io.NopCloser(bytes.NewReader(json))

	fmt.Println(string(json))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res)

	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
