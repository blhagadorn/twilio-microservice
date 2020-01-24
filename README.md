# twilio-microservice
Microservice for sending SMS from a container

## Env Files
Replace `env.example` with real `.env` attributes 

## Local Docker Build

Build:
`docker build -t twilio-microservice .` inside of the `go/` or other library

Run:
`docker run -p 8086:8086 twilio-microservice:latest`

## Testing

`curl -X GET localhost:8086/healthz`
