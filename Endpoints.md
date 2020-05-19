# Endpoints
All endpoints are documented here.

The base URL is where you host the server, usually `http://localhost:PORT` where `PORT` is your port or the default `3030` port.

If you exposed it to the internet with `-h 0.0.0.0` then users outside the host can access it via `http://yourserverip:PORT` if you are doing that to e.g share the server with a friend then consider passing a `-s password` flag to add a password and then users will have to request with `Authorization: password` header set. Local requests from `127.0.0.1` still does not need to authenticate.

## GET /ping
Always returns `{"message": "Pong!"}` a quick check to see if the server is alive.

## GET /stats
Returns an object with `{ "version": "version", "uptime": 0, "stats": {...} }`

`version` is the version of the API server.

`uptime` is number of seconds the server has been up for.

`stats` is the full Go object that contains memory usage information. See `go doc runtime.MemStats` for the fields.

Stats is quite big and if you only need uptime you can pass `?noStats=true` to not send the mem stats.

# Image Endpoints
All image endpoints return the image binary directly on `200` success. Any errors are sent with a JSON containing a `message` field.

Replace `{IMAGE_URL}` with an actual url that points to an image smaller than 8 MB.

If there is a `:(sizeXsize)` mentioned then it signals a recommended size that can be asked from Discord to avoid unnecessary resizing. For text input it can signal the max character limit.

- `/religion?avatar={IMAGE_URL}`
- `/beautiful?avatar={IMAGE_URL}`
- `/fear?avatar={IMAGE_URL}`
- `/sacred?avatar={IMAGE_URL}`
- `/painting?avatar={IMAGE_URL}`
- `/color?color={NAME_OR_HEX}` (`#` is automatically stripped off, and name is case insensitive)
- `/delete?avatar={IMAGE_URL}`
- `/garbage?avatar={IMAGE_URL}`
- `/tom?avatar={IMAGE_URL}`
- `/bed?avatar={IMAGE_URL}&target={IMAGE_URL}`
- `/crush?avatar={IMAGE_URL}&target={IMAGE_URL}`
- `/patrick?avatar={IMAGE_URL}`
- `/respect?avatar={IMAGE_URL}`
- `/dipshit?text={TEXT:(max: 33 chars)}`
- `/picture?avatar={IMAGE_URL}` (Note: A little slower compared to the other endpoints)
- `/tweet?text={TEXT:(max: 165 chars)}`
- `/truth?avatar={IMAGE_URL}`
- `/bobross?avatar={IMAGE_URL}`
- `/mask?avatar={IMAGE_URL:(512x512)}` (More feedback required to improve this.)
- `/father?avatar={IMAGE_URL}&text={TEXT:(max: 42 chars)}`
- `/achievement?avatar={IMAGE_URL}&text={TEXT:(max: 21 chars)}`
