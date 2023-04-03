## Go CLI Chat app
Go CLI Chat app using the standard library package "net/http" for handling HTTP requests and responses, and "gorilla websocket" package for establishing connections.

### main.go file
here, I set up a simple HTTP server on port 8080 and handles WebSocket connections with the `websocket` from Gorilla. When a WebSocket connection is established, the server reads messages from the client, prints them to the console, and sends them back to the client.

To run the server, open up your terminal and `cd` to the project directory. Then run `go run main.go`

### client.go file
This client uses the gorilla/websocket package to connect to the WebSocket server at ws://localhost:8080. It then enters a loop where it prompts the user to enter a message, sends that message to the server, and waits for a response.

To run this client, `cd` into the `client` directory. Then run `go run client.go`