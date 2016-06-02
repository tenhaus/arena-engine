# Botpit

Right now everything runs off of my Google Cloud account, but could run off someone else's account by changing config and downloading a new token to botpit-development-authentication.json

Currently using iam api, pubsub, storage

export BOTPIT_ENV="development"

export GOOGLE_APPLICATION_CREDENTIALS="/Users/chris/Code/go/src/github.com/tenhaus/botpit/botpit-development-authentication.json"

export BOTPIT_CONFIG="/Users/chris/Code/go/src/github.com/tenhaus/botpit/config.json"

### DONE
* Create user
* Delete user
* Create service account
* Delete service account
* Lib for creating pubsub topics and listening for messages
* Config
* Restrict service account permissions to only send/receive on topics

### TODO
* Pull token and send to client
* Client
* Deploy to compute engine (deployment works, but don't know how to start the service yet)
* Rest route and function for /signup
* More 

## Purpose

Botpit is a tournament framework for connecting and controlling player-written bots. It provides a client library and protocol that handles all communication between a game and players.

The application uses Google Cloud Platform throughout. Players sign up by using the command line client, which posts credentials to the rest API hosted on either app engine or compute engine.

The game itself can be anything. The botpit server will provide an api to the game. For example, a number guessing game will be provided with number of players and turn information.

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
