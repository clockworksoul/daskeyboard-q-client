#!/bin/bash

BACKEND_URL="http://localhost:27301"
HEADERS=(-H "Content-Type: application/json")
SIGNAL_ID="-290"
URL="$BACKEND_URL/api/1.0/signals/$SIGNAL_ID"

curl "${HEADERS[@]}" -X DELETE $URL