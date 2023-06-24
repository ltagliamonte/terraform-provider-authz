#!/bin/bash
set -euo pipefail

local_test_username="local-test"

token=$(curl -sSL 'http://localhost:8080/v1/auth' \
  -H 'Content-Type: application/json' \
  --data-raw '{"username":"admin","password":"changeme"}' \
  --compressed | jq -r '.access_token')

resp=$(curl -sSL 'http://localhost:8080/v1/clients' \
  -H 'Authorization: Bearer '"${token}"'' \
  -H 'Content-Type: application/json' \
  --data-raw '{"name":"'"${local_test_username}"'"}' \
  --compressed)

curl -sSL 'http://localhost:8080/v1/principals/authz-sa-local-test' \
  -X 'PUT' \
  -H 'Authorization: Bearer '"${token}"'' \
  -H 'Content-Type: application/json' \
  --data-raw '{"id":"authz-sa-local-test","roles":["authz-admin"],"attributes":[]}' \
  --compressed > /dev/null

echo AUTHZ_CLIENT_ID=$(echo ${resp} | jq -r '.client_id')
echo AUTHZ_CLIENT_SECRET=$(echo ${resp} | jq -r '.client_secret')
exit 0
