#!/bin/bash

BACKEND_URL="http://localhost:27301"
HEADERS=(-H "Content-Type: application/json")

# Page to query
PAGE_NUMBER="0"

# Number of elements per page
NUMBER_OF_SIGNALS_PER_PAGE="2"

# Sort order can be ASC or DESC
SORT_ORDER="ASC"

URL="$BACKEND_URL/api/1.0/signals?page=$PAGE_NUMBER&size=$NUMBER_OF_SIGNALS_PER_PAGE&sort=createdAt,$SORT_ORDER"

echo $URL

curl -I "${HEADERS[@]}" -X GET $URL
