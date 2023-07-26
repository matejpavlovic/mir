#!/bin/bash

echo "$1" | python3 scripts/hosts-to-membership.py > membership.json
go run ../cmd/bench params -i bench-config.json -m membership.json -o bench-config.json