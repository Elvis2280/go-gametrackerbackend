# Game Tracker Backend
## Description
This is the backend for the Game Tracker app. It is a RESTFull API built with Go, Gin, and Postgresql. It is hosted on Railway. The frontend is built with React, React Native and is hosted on Netlify. The frontend repo can be found [here]
## ENV Variables
- `DATABASE_URL` - The URL for the Postgresql database
- `JWT_SECRET` - The secret used to sign JWT tokens
- `JWT_EXPIRES` - The time in hours (Int) for JWT tokens to expire
- `JWT_ISSUER` - The issuer of the JWT tokens

## Installation
- Clone the repo
- Add the .env file to the root of the project
- Run docker-compose build
- Run docker-compose up

## Commands
- `docker-compose -f docker-compose.debug.yml up --build` - Runs the server port 8080 in development mode
- `swag init` - Generates swagger documentation

## Enable debug mode
This is setup to run in debug mode using docker compose
so you need to setup your IDE to run external debug in port 2345 localhost
so after running the docker-compose up --build you can run the debug mode in your IDE



## swagger endpoint
coming soon /docs/index.html