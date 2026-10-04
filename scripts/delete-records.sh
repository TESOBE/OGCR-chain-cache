#!/usr/bin/env bash
# Delete every record of the *_on_chain dynamic entities defined in entities/,
# keeping their definitions and Roles. Lists the records first and asks once
# before deleting; the cacher (make run) writes them again from the chain.
# Name entities to limit it to those.
#
#   scripts/delete-records.sh [ENTITY...]
set -euo pipefail

cd "$(dirname "$0")/.."

if [ $# -gt 0 ]; then
	entities=("$@")
else
	entities=()
	for f in entities/*_on_chain.json; do
		entities+=("$(basename "$f" .json)")
	done
fi

for entity in "${entities[@]}"; do
	go run ./cmd/delete-records "$entity"
	echo
done

read -r -p "Delete the records of ${entities[*]}? [y/N] " answer
case "$answer" in
	y | Y | yes | YES) ;;
	*)
		echo "Nothing was deleted."
		exit 0
		;;
esac

for entity in "${entities[@]}"; do
	go run ./cmd/delete-records -yes "$entity"
done
