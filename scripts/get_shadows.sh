#!/bin/bash

BACKEND_URL="http://localhost:27301"
HEADERS=(-H "Content-Type: application/json")
PID="Q_MATRIX"
URL="$BACKEND_URL/api/1.0/signals/pid/$PID/shadows"

curl "${HEADERS[@]}" -X GET $URL