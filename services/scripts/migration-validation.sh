#!/usr/bin/env bash

# this bash script checks if there are incoming migration files which are older than target's branch migration files.
# solution is usually changing the incoming migration file name, so the timestamp is newer than latest migration on `master` branch.

BASE_BRANCH=$1

MIGRATIONS_DIR=gateway/internal/assets/files/migrations/

OLDEST_NEW_MIGRATION_FILE=$(git diff --name-only "origin/${BASE_BRANCH}" --diff-filter=d | grep -m1 "${MIGRATIONS_DIR}")

if [[ -z $OLDEST_NEW_MIGRATION_FILE ]]; then
    echo "no new migrations"
    exit 0
fi

NEWEST_EXISTING_MIGRATION_FILE=$(git ls-tree -r "origin/${BASE_BRANCH}" --name-only | grep "${MIGRATIONS_DIR}" | tail -1)

if [[ -z $NEWEST_EXISTING_MIGRATION_FILE ]]; then
    echo "no existing migrations"
    exit 0
fi

EXISTING_TIMESTAMP=$(basename "$NEWEST_EXISTING_MIGRATION_FILE" | cut -d '_' -f 1)

NEW_TIMESTAMP=$(basename "$OLDEST_NEW_MIGRATION_FILE" | cut -d '_' -f 1)

if [[ $EXISTING_TIMESTAMP -ge $NEW_TIMESTAMP ]]; then
    echo "existing migration timestamp is greater than or equal to incoming migration timestamp. please update your migrations timestamp."
    echo "failing file: ${OLDEST_NEW_MIGRATION_FILE}"
    exit 1
fi

echo "new migration(s) are safe to merge"