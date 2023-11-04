#!/bin/zsh

curl -X 'GET' \
  'https://api.warframe.market/v1/items' \
  -H 'accept: application/json' \
  -H 'Language: en' | jq .
