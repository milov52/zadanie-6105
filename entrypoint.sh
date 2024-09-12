#!/bin/bash
set -e

# Step 1: Run migrations
goose -dir ./migrations postgres ${POSTGRES_CONN} up

# Step 2: Start the application
./app