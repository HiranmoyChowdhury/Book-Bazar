# Book Bazar
An API server that can be used as a backend for an e-bookstore. This server is built using cobra CLI, go-chi/chi, Basic Auth, JWT Auth and RESTful API using go.
## Installation
$ go_version=1.17.1
$ cd ~/Downloads
$ sudo apt-get update
$ sudo apt-get install -y build-essential git curl wget
$ wget https://dl.google.com/go/go${go_version}.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go${go_version}.linux-amd64.tar.gz
$ sudo chown -R $(id -u):$(id -g) /usr/local/go
$ rm go${go_version}.linux-amd64.tar.gz
