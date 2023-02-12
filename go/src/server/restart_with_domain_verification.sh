#!/bin/bash

cd ~/renju3d-server/go/src/server
go build

# start the server in the background, and output logs to a text file for debug purpose
nohup sudo go run . -renew_ssl=true &> server_log.out &

echo "Began hosting ZeroSSL certificates, click Verify Domain now!"

read -n 1 -p "When done, press any key to continue..."
echo ""
read -n 1 -p "Now, download the certificates as ZIP and copy/replace the contents to ~/renju3d-server/go/src/server/zerossl_certs. When done, press any key to continue restarting the server..."

echo ""
echo "Restarting server..."

cd ~/renju3d-server/go/src/server
bash kill_server.sh
bash start_server.sh

echo "Server restarted!"