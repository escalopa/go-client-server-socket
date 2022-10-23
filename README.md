# go-client-server-socket

This project goal is to create a connections socket between `client` & `server` only using `sockets` 

### Installation

1. Clone the repo with command 
```shell
git clone https://github.com/escalopa/go-client-server-socket.git
```
2. Install go from the [link](https://go.dev/doc/install)(if not already installed)
2. Set GOROOT to path where go is installed, on linux by default it is on `/usr/local/go`
3. Set GOPATH to home directory of the device
4. Open root dir of the project and run the following commands

### Server
Run server side using the following command
```shell
go run ./server/timeserver.go
```

### Client
Run client side using the following command
```shell
go run ./client/timeclient.go
```
