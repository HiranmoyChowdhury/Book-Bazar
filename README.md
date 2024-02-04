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
**Running Container:** Run the Following command to run this image in container:
```
docker run -dp 127.0.0.1:8080:8080 book-server-api start
```
Here start is the command which actually give the instructions to start the server. This server is running with a default port 8080. You can change it during running the server.
Example: 
```
docker run -dp 127.0.0.1:<desired port>:<desired port> book-server-api start -d=<desired port>
```
You can also register by giving **admin name** and **password** which can be useful if you don't want to use the default **admin name** and **password**.


