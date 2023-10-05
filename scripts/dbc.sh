#!/bin/bash

# Connects to the postgres database
# Usage: ./dbc.sh

# Get the database name from the .env file
DB_USER=$(grep DB_USER .env | cut -d '=' -f2)
DB_PORT=$(grep DB_PORT .env | cut -d '=' -f2)
DB_NAME=$(grep DB_NAME .env | cut -d '=' -f2)

# Connect to PostgreSQL
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME
