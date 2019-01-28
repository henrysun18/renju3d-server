# renju3d-server
Server-side code to facilitate client-server networking in my Unity game (Renju3D)

Running on a Google Compute Engine f1-micro instance, so this server's primary purposes are to only be the arbiter for assigning available spots in a room, asserting whose turn it is, and ensuring inactive rooms get reset. This server will also give the board state to spectators that may have joined in the middle of the game.

Old method of networking [Renju3D v1.2 Online](http://henrysun.me/unity/Renju3D_v1.2(online)/index.html):

![alt text](https://github.com/henrysun18/renju3d-server/blob/master/firebase-approach.png?raw=true)


New method of networking:

![alt text](https://github.com/henrysun18/renju3d-server/blob/master/compute-engine-approach.png?raw=true)
