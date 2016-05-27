# Botpit

### DONE
* Create user
* Delete user
* Create service account
* Delete service account
* Lib for creating pubsub topics
* Config

### TODO
* Restrict service account permissions to only send/receive on topics
* Pull token and send to client
* Client
* Deploy to compute engine (deployment works, but don't know how to start the service yet)
* Rest route and function for /signup
* More but it's time to go home

## Purpose

Botpit is a tournament framework for connecting and controlling player-written bots. It provides a client library and protocol that handles all communication between a game and players.

The application uses Google Cloud Platform throughout. Players sign up by using the command line client, which posts credentials to the rest API hosted on either app engine or compute engine.

### When the player signs up
* Botpit stores their username, encrypted pass, and email in Cloud Storage
* It then creates a service account, which allows the player bot to join our pubsub topics for message routing.
* Botpit creates a pubsub topic for that user to route all game requests
* Permissions and roles are created to restrict the user from doing anything but talking on the provided topics
* Botpit sends back the user pubsub topic, a general game request topic, and auth token. It's probably a good idea to send back dev versions of all of these too, so the player can develop without messing with their standing

### Bot writing
The user downloads the client library using go get. They'll store their user, pass, and token in config files, sets environment variables to point to them.

Interfacing with the client library is simple. The bot writer is provided with a couple of channels that are endpoints for their topics. The channels will feed their bots GameAction structs with game data. The bot writer can then do whatever they want to write their bot. The can boot up instances of compute engine, train their neural networks, whatever. That's the fun of it.

### Server
The controller will listen for game requests from players and create and manage matches. This is still a little fuzzy, but I have a general idea how to do this. Mostly the controller will start Match and Game go routines and open up channels to them. After that the Match and Game will handle everything. When the match is over it will send the data over the channel and the Controller will store it and do whatever it needs to do to finalize the match. A leaderboard is definitely needed.
