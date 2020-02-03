# twilio-microservice
Microservice reference architecture for sending SMS via Twilio API from a container, with the goal of supporting many languages and protocols and security best practice.

Please raise an issue for a feature request or bug



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
```

Alternatively, a Postman collection exists [Twilio-Microservice Postman Collection] (https://github.com/blhagadorn/twilio-microservice/tree/master/postman) with basic endpoints for health check and sending a text message.
