# Book Bazar
An API server that can be used as a backend for an e-bookstore. This server is built using cobra CLI, go-chi/chi, Basic Auth, JWT Auth and RESTful API using go.
## Prerequisites
**First install Go if you don't have Go installed in your environment:**
```
$ go_version=1.21.5
$ cd ~/Downloads
$ sudo apt-get update
$ sudo apt-get install -y build-essential git curl wget
$ wget https://dl.google.com/go/go${go_version}.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go${go_version}.linux-amd64.tar.gz
$ sudo chown -R $(id -u):$(id -g) /usr/local/go
$ rm go${go_version}.linux-amd64.tar.gz
```
Add go to your $PATH variable:
```
$ mkdir $HOME/go
$ nano ~/.bashrc
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
$ source ~/.bashrc
$ go version
```
## Run using docker
**Clone the project:** First clone the project using this command to your desired directory:
```
git clone https://github.com/HiranmoyChowdhury/Book-Bazar/tree/master
```
**Building Image:** Run the Following command to build the image:
```
docker build -t book-server-api .
```
**Running Container:** Run the Following command:
```
docker run -dp 127.0.0.1:3000:3000 book-server-api
```




