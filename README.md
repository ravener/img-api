# Image API

This is an API written in Golang for image manipulation commands, primarily targets usage from [Discord](https://discordapp.com) Bots.

It is used in my bot [Miyako](https://github.com/pollen5/miyako) and anyone is free to use this.

There used to be an API with the name "Idiotic API" by York. It was great while it lasted but it died now I tried to reinvent some of the endpoints that API had, this time it's open source and self-hosted so it will always be available to you.

## Setup
Install Golang 1.13+ then clone this repository.
```sh
$ git clone https://github.com/pollen5/img-api
$ cd img-api
# Build the binary
$ go build main.go
# start the server
$ ./main
```
The default port is `3030` but you can pass the flag `-p` to change the port.

All API endpoints are listed in [Endpoints.md](Endpoints.md)

## Local Usage or Shared?
There are 2 ways to use this. The first one is the recommended and easiest way: locally.

The server can be accessed only via the same host, only for your bot. This is the default behaviour.

If you want to expose the server to the internet (e.g to share it with a friend.) then you have to run it with `-h 0.0.0.0` but anyone who knows your server IP and port will be able to use the API so you might consider adding some sort of authorization. Run it with `-s password` where password is a secret authorization that you will only share with users you trust and they will have to add an `Authorization` header with it. Requests from localhost won't be required to authenticate and will work as normal.

## Wrappers
I wrote some API wrappers in multiple languages to help users get started quickly.

- [JavaScript](https://github.com/pollen5/img-api.js)
- [Python](https://github.com/pollen5/img-api.py)
- [Golang](https://github.com/pollen5/img-api-go)

**Unofficial Wrappers:**
- If you made any new wrappers then feel free to add it to this list.

Official wrappers are promised to stay updated with all changes where as unofficial ones depends on the person who made them.

## Contributing
Contributions are welcome, whether it is new image template ideas or code contributions. Feel free to open any issues when in doubt.

Also join [our Discord Server](https://discord.gg/mh7vEck) for contacting me and getting updates.

## License
Released under the [MIT License](LICENSE)

I do not own the assets provided.
