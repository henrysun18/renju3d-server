#!/bin/bash

# query all PID info with port 8443. Then awk the last column of the row. Then cut out only the numbers part to get the PID
PID_OF_TLS_SERVER=$(sudo netstat -lnp | grep 8443 |  awk 'NF>1{print $NF}' | cut -d/ -f1)
if [ -z $PID_OF_TLS_SERVER ]
then 
  echo "No TLS server PID was found"
else
  echo "Shutting down Renju3D Server with PID $PID_OF_TLS_SERVER"
  sudo kill -9 ${PID_OF_TLS_SERVER}
fi
# exit out of the script
exit 1