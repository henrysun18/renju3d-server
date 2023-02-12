
# renju3d-server
Server-side code to facilitate client-server networking in my Unity game (Renju3D)

Running on a Google Compute Engine f1-micro instance, so this server's primary purposes are to only be the arbiter for assigning available spots in a room, asserting whose turn it is, and ensuring inactive rooms get reset. This server will also give the board state to spectators that may have joined in the middle of the game.

Old method of networking [Renju3D v1.2 Online](http://henrysun.me/unity/Renju3D_v1.2(online)/index.html):

![alt text](https://github.com/henrysun18/renju3d-server/blob/master/firebase-approach.png?raw=true)


New method of networking:

![alt text](https://github.com/henrysun18/renju3d-server/blob/master/compute-engine-approach.png?raw=true)

To enable REST API calls: go to Google Cloud Console and add TCP 8080 to the firewall allowlist for the 0.0.0.0/0 IP range (all IPv4 IP addresses)

## Setup
1. Check out the source code inside a Google Compute Engine VM instance, via SSH. 
2. In Google Cloud Console, ensure Firewall for this VM is configured to allow Ingress into port TCP:8443 (e.g. 3 dots menu in VM instances page --> View network details --> Firewall burger menu tab --> Create Firewall rule)
3. Inside the remote terminal, enter `cd ~/renju3d-server/go/src/server && bash start_server.sh`. This will start the server in a background process, which continues running even after closing the SSH session. If you see something like `sudo: go: command not found`, then follow [these steps](https://stackoverflow.com/a/71910152). Note that sudo privileges are needed here for TLS related reasons, otherwise clients would fail to connect.
4. To restart the server, enter `bash kill_server.sh && bash start_server.sh`

## Certificate Renewal
ZeroSSL offers free 90-day certificates to enable HTTPS traffic so that Unity WebGL games can actually have permission to communicate with the Go server. This requires manually renewing every 3 months though, but is easy with the following steps:
1. Click the Renew Certificate button in the "Certificate Expiring in 14 Days" reminder email. Alternatively, go directly to this link: https://app.zerossl.com/certificates
2. Click Renew, and choose HTTP File Upload.
3. Download the txt file and upload to ~/renju3d-server/go/src/server/.well-known/pki-validation (can do this by dropping the file directly there via MobaXTerm, or by simply creating a new txt file and copy-pasting the contents).
4. `cd ~/renju3d-server/go/src/server && bash restart_with_domain_verification.sh`
5. When the script prompts you to Verify Domain, go back to the browser and click the Verify Domain button.
6. Back in the terminal, press any button to proceed, and the script will restart the server.
7. Feel free to close the terminal, as the server has been backgrounded and will continue running.