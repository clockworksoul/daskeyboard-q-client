#!/bin/bash

BACKEND_URL="http://localhost:27301"
HEADERS=(-H "Content-Type: application/json")
URL="$BACKEND_URL/api/1.0/signals/shadows"

curl "${HEADERS[@]}" -X GET $URL | jq