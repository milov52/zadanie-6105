#!/bin/bash
set -e

# Step 1: Run migrations
goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up

# Step 2: Start the application
./app