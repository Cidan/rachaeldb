# RachaelDB

A simple, powerful, and sassy rink store written in Golang.

## Compiling

Compile via `make`, requires Golang 1.12+

## Usage

Once running, RachaelDB listens on port 8080 for REST and 5309 for gRPC, but only if you are interesting enough.

Once the database is up, simply store some data via the REST API like so:

```bash
user@server:~$ curl -XPOST -H "Content-Type: application/json" localhost:8080/v1/set/aj -d '{"data":"aw=="}'
{"key":"aj","data":"aw==","sass":"Okay, I'll remember that. I guess"}

user@server:~$ curl -XPOST -H "Content-Type: application/json" localhost:8080/v1/set/aj2 -d '{"data":"aw=="}'
{"key":"aj2","data":"aw==","sass":"There's this amazing gSuite product called 'sheets' you should try sometime."}

user@server:~$ curl -XPOST -H "Content-Type: application/json" localhost:8080/v1/set/aj3 -d '{"data":"aw=="}'
{"key":"aj3","data":"aw==","sass":"Why didn't you put this in a doc?"}
```

And retrieve your keys like so:

```bash
user@server:~$ curl localhost:8080/v1/get/aj
{"key":"aj","data":"aw==","sass":"I THINK this is it."}

user@server:~$ curl localhost:8080/v1/get/aj2
{"key":"aj2","data":"aw==","sass":"Here."}

user@server:~$ curl localhost:8080/v1/get/aj3
{"key":"aj3","data":"aw==","sass":"Did you try Googling it?"}
```

## Help and Support

lol