# Golang Auth proxy example

This is just an example on how to use the `httputil.ReverseProxy` type provided in the golang standard library.

1. Open a terminal in the root folder and run `go run target/main.go`. This will start the target server. It will only respond to request from `localhost:8081`

2. Open another terminal in the root folder and run the following commands:

```sh
# first export some env variables
export SCHEME="http"
export HOST="localhost:8080"
export AUTH="123456"

export SECRET="superSecret"

# then run proxy server
go run proxy/main.go
```
# reverse proxy but without Auth header.
curl -iL -X GET http://localhost:8081

# with header.
curl -iL -X GET -H "AUTH: 123456" http://localhost:8081

This will start our auth proxy in `localhost:8081` and will redirect traffic to `http://localhost:8080` only if there is an header like `"Auth: 123456"`

3. Open another terminal and start making some requests


