#!/bin/bash
source .env
cd "$(dirname "$0")/sql/schema/"
goose postgres $DB_URL up
