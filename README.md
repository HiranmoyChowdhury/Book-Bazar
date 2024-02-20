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
## Clone the project
First clone the project using this command to your desired directory:
```
git clone https://github.com/HiranmoyChowdhury/Book-Bazar/tree/master
```
## Run using docker
**Building Image:** Run the Following command to build the image from project directory:
```
docker build -t book-bazar .
```
**Running Container:** Run the Following command to run this image in container:
```
docker run -dp 127.0.0.1:8080:8080 book-bazar start
```
Here start is the command that actually instructs the server to start. This server is running with a default port of 8080 You can change this while running the server:
```
docker run -dp 127.0.0.1:<desired port>:<desired port> book-bazar start -d=<desired port>
```
You can register with an **admin name** and **password** which can be useful if you don't want to use the default **admin name** and **password** during the login process:

```
docker run -dp 127.0.0.1:<desired port>:<desired port> book-bazar start -d=<desired port> -u=<user name> -p=<password>
```
## Run on local machine
**Building an executable file:** Run the Following command to build *executable file* from project directory:
```
go build -o book-bazar
```
**Running the executable file:** Run the Following command to run this executable file in your local system:
```
book-bazar start
```
Here start is the command that actually instructs the server to start. This server is running with a default port of 8080 You can change this while running the server:
```
book-bazar start -d=<desired port>
```
You can register with an **admin name** and **password** which can be useful if you don't want to use the default **admin name** and **password** during the login process:

```
book-bazar start -d=<desired port> -u=<user name> -p=<password>
```
# API EndPoints
| Method | API EndPoint | Authentication Type | Description | Curl Command |
| ------------- | ------------- | ------------------ | ----------------------- | ----------------- |
| GET | /api/v1/get-token | Basic-Auth | LogIn with registered user,pass and get bearer token | $ curl -X POST --user  '\<userName\>:\<passWord\>' localhost:\<port no\>/api/v1/get-token |
| POST | /api/v1/book | JWT-Auth | Add a new book and in response get the full info of this book including UUID | $ curl -X POST -H "Authorization: Token \<bearerToken\>" -H "Content-Type:application/json" -d '<bookModelJson>' localhost:\<port no\>/api/v1/book |
| GET | /api/v1/books | JWT-Auth | Get the list of books in response | $ curl -X GET -H "Authorization: Token \<bearerToken\>" localhost:\<port no\>/api/v1/books |
| GET | /api/v1/book/{UUID} | JWT-Auth | Get a certain book with it's UUID | $ curl -X GET -H "Authorization: Token \<bearerToken\>" localhost:\<port no\>/api/v1/book/\<UUID\> |
| PUT | /api/v1/book/{UUID} | JWT-Auth | Update a certain book with it's UUID | $ curl -X PUT -H "Authorization: Bearer \<bearerToken\>" -H "Content-Type:application/json" -d '<bookModelJson>' localhost:\<port no\>/api/v1/book/\<UUID\> |
| DELETE | /api/v1/book/{UUID} | JWT-Auth | Delete a certain book with it's UUIT | $ curl -X DELETE -H "Authorization: Bearer \<bearerToken\>" localhost:\<port no\>/api/v1/book/\<UUID\> |











