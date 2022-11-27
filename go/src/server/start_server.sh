#!/bin/bash

cd ~/renju3d-server/go/src/server
go build

# start the server in the background, and output logs to a text file for debug purposes. To do this manually, simply execute `sudo go run .` then press ctrl-z, then `disown -h %1 && bg 1`
nohup sudo go run . &> server_log.out &

echo "Renju3D Server started"
