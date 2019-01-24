# renju3d-server
Server-side code to facilitate client-server networking in my Unity game (Renju3D)

Running on a Google Compute Engine f2-micro instance, so this server's primary purposes are to only be the arbiter for whose turn it is, and pass player messages to each other. This server will also give the board state to spectators that may have joined in the middle of the game.
