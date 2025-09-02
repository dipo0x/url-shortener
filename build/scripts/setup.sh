#!/bin/bash

# Prompt the user for server and database details
echo "Enter server post (default: localhost:8080):"
read -r PORT
PORT=${PORT:-localhost:8080}

echo "Enter database host (default: golang-url-shortener):"
read -r DATABASE_URL
DATABASE_URL=${DATABASE_URL:-golang-url-shortener}

echo "Enter jwt secret key (default: 1D@Gz3d7P!nK*I#T8rE$F3m9L&vJq%W):"
read -r JWT_SECRET_KEY
JWT_SECRET_KEY=${JWT_SECRET_KEY:-D@Gz3d7P!nK*I#T8rE$F3m9L&vJq%W}

echo "Enter url absolute link (default: http://localhost:8080/url):"
read -r URL_ABSOLUTE_URL
URL_ABSOLUTE_URL=${URL_ABSOLUTE_URL:-http://localhost:8080/url}

echo "Enter redis url (default: localhost:6379):"
read -r REDIS_URL
REDIS_URL=${REDIS_URL:-localhost:6379}

# Create the .env file
cat <<EOF > .env
PORT=$PORT
DATABASE_URL=$DATABASE_URL
JWT_SECRET_KEY=$JWT_SECRET_KEY
URL_ABSOLUTE_URL=$URL_ABSOLUTE_URL
REDIS_URL=$REDIS_URL
EOF

echo ".env file created successfully."