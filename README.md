# Image API

This is an API written in Golang for image manipulation commands, primarily targets usage from [Discord](https://discord.com) Bots.

It is used in my bot [Miyako](https://github.com/ravener/miyako) and anyone is free to use this.

There used to be an API with the name "Idiotic API" by York. It was great while it lasted but it died now I tried to reinvent some of the endpoints that API had, this time it's open source and self-hosted so it will always be available to you.

## Install
Pre-built binaries are available for Windows, macOS and Linux in [GitHub Releases](https://github.com/ravener/img-api/releases), the easiest way to get started is to just download a release which includes everything needed to get started in your machine without installing anything else.

Simply extract the archive and run the `img-api` binary (`img-api.exe` on Windows)

On Linux/macOS you can type `./img-api` to run it, on Windows you may just double click the exe or type `img-api` in cmd. (Make sure you are in the correct directory)

If you are on Linux/macOS you may also install via [Homebrew](https://brew.sh) just run: `brew install ravener/tap/img-api`

The API will start in `http://localhost:3030` but the port can be changed via `-p`

If the prebuilt binaries doesn't suit you or you'd like to edit the code then continue for instructions on building, this will require [Golang 1.13+](https://golang.org) installed.

## Setup (No Docker)
Install Golang 1.13+ (`git` must also be installed) then clone this repository.
```sh
$ git clone https://github.com/ravener/img-api
$ cd img-api
# Build the binary
$ go build
# start the server
$ ./img-api
```

# Setup (Docker)
To run with docker you can use
```sh
# Replace target port with the one you want to use on your host (this only exposes it locally)
docker run -d -p <target>:3030 iceemc/img-api:latest
# To expose the server on all interfaces
docker run -d -p 0.0.0.0:<target>:3030 iceemc/img-api:latest
```

# Setup (Pterodactyl)
Just follow the below instructions in order, indicated by the **big numbers**
![image](https://user-images.githubusercontent.com/30955604/116635544-0fe1a280-a92d-11eb-9837-8b03f1dd333f.png)


All API endpoints are listed in [Endpoints.md](Endpoints.md)

## Local Usage or Shared?
There are 2 ways to use this. The first one is the recommended and easiest way: locally.

The server can be accessed only via the same host, only for your bot. This is the default behaviour.

If you want to expose the server to the internet (e.g to share it with a friend.) then you have to run it with `-h 0.0.0.0` but anyone who knows your server IP and port will be able to use the API so you might consider adding some sort of authorization. Run it with `-s password` where password is a secret authorization that you will only share with users you trust and they will have to add an `Authorization` header with it. Requests from localhost won't be required to authenticate and will work as normal.

## Wrappers
I wrote some API wrappers in multiple languages to help users get started quickly.

- [JavaScript](https://github.com/ravener/img-api.js)
- [Python](https://github.com/ravener/img-api.py)
- [Golang](https://github.com/ravener/img-api-go)

**Unofficial Wrappers:**
- If you made any new wrappers then feel free to add it to this list.

Official wrappers are promised to stay updated with all changes where as unofficial ones depends on the person who made them.

## Contributing
Contributions are welcome, whether it is new image template ideas or code contributions. Feel free to open any issues when in doubt.

Also join [`#img-api` in my Discord Server](https://discord.gg/wpE3Nfp) for contacting me and getting updates.

Tip: If you use a web-browser to test the images output in development, pass the `-d` flag to disable browser cache to make testing easier.

## License
Released under the [MIT License](LICENSE)

I do not own the assets provided.
