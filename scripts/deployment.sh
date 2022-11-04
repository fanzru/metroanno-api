#!/bin/sh
set -e


echo "================================================================"
echo "===========================DEPLOYMENT==========================="
echo "================================================================"

echo "Update codebase..."
cd ~/project/metroanno-api
git fetch origin main
git reset --hard origin/main

echo "Installing dependencies 🛠"
go mod tidy

echo "Restart pm2 service backend 🔥"
pm2 restart pm2.json

echo "Deploying Backend Application Successfully Yeayyyy ......."