# Go Client for Das Keyboard Q Service

A primitive Go client for the Das Keyboard Q service. It supports both localhost and Q Cloud.

https://www.daskeyboard.io/get-started/


## How to use this

This creates a popup message and sets the `Q` key to solid red.

```go
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

// Output the response, just for fun.
bytes, err := json.MarshalIndent(response, "", "  ")
fmt.Println(string(bytes))
```

The output would look something like:

```json
{
  "name": "New Q app version available",
  "message": "Q App version 3 is available.",
  "zoneId": "KEY_Q",
  "color": "#FF0000",
  "effect": "SET_COLOR",
  "pid": "Q_MATRIX",
  "clientName": "GoClient",
  "id": -590,
  "userId": 14339,
  "createdAt": 1619815234590,
  "updatedAt": 1619815234590
}
```

