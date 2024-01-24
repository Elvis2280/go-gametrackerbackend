# Game Tracker Backend
## Description
This is the backend for the Game Tracker app. It is a RESTFull API built with Go, Gin, and Postgresql. It is hosted on Railway. The frontend is built with React, React Native and is hosted on Netlify. The frontend repo can be found [here]
## ENV Variables
- `DATABASE_URL` - The URL for the Postgresql database
- `JWT_SECRET` - The secret used to sign JWT tokens
- `JWT_EXPIRES` - The time in hours (Int) for JWT tokens to expire
- `JWT_ISSUER` - The issuer of the JWT tokens

## Commands
- `go run main.go` - Runs the server
- `swag init` - Generates swagger documentation



## swagger endpoint
coming soon /docs/index.html