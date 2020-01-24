# twilio-microservice
Microservice for sending SMS from a container build with distroless containers

## Env Files
Replace `.env-example` with real `.env` attributes. The `TO_NUMBER` and `FROM_NUMBER` are for testing purposes, but should be in the JSON request instead.
When building, please also put the `.env` files inside of the individual language directories.

## Local Docker Build

Build:
`docker build -t twilio-microservice .` inside of the `go/` directory

Run:
`docker run -p 8087:8087 twilio-microservice:latest`

## Testing

`curl -X GET localhost:8087/healthz`

Or to send a text message:

```
curl -X POST 'localhost:8087/text' \
--header 'Content-Type: application/json' \
--data-raw '{
        "from":"+18888888888",
        "to":"+17777777777",
        "body":"Hello, I'\''m not from a monolith!"
}'
