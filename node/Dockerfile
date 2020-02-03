FROM node:10.17.0 AS build-env
ADD . /app
WORKDIR /app
RUN npm install

FROM gcr.io/distroless/nodejs
COPY --from=build-env /app /app
WORKDIR /app
EXPOSE 8087
COPY .env .
CMD ["server.js"]
