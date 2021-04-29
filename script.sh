#!/bin/bash

BACKEND_URL="http://localhost:27301"
HEADERS=(-H "Content-Type: application/json")
URL="$BACKEND_URL/api/1.0/signals"

curl "${HEADERS[@]}" -X POST -d  '{
  "zoneId": "KEY_Q",
  "color": "#FFFFFF",
  "effect": "SET_COLOR",
  "pid": "DK5QPID",
  "clientName": "Shell script",
  "message": "Q App version 3 is available. Download it at https://www.daskeyboard.io/get-started/download/",
  "name": "New Q app version available"}' $URL