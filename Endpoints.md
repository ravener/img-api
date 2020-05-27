# Endpoints
All endpoints are documented here.

The base URL is where you host the server, usually `http://localhost:PORT` where `PORT` is your port or the default `3030` port.

If you exposed it to the internet with `-h 0.0.0.0` then users outside the host can access it via `http://yourserverip:PORT` if you are doing that to e.g share the server with a friend then consider passing a `-s password` flag to add a password and then users will have to request with `Authorization: password` header set. Local requests from `127.0.0.1` still does not need to authenticate.

## GET /ping
Always returns `{"message": "Pong!"}` a quick check to see if the server is alive.

## GET /stats
Returns an object with `{ "version": "version", "uptime": 0, "stats": {...}, "goroutines": 0 }`

`version` is the version of the API server.

`uptime` is number of seconds the server has been up for.

`stats` is the full Go object that contains memory usage information. See `go doc runtime.MemStats` for the fields.

`goroutines` is the number of active goroutines.

Stats is quite big and if you only need the other fields you can pass `?noStats=true` to not send the mem stats.

# Image Endpoints
All image endpoints return the image binary directly on `200` success. Any errors are sent with a JSON containing a `message` field.

Replace `{IMAGE_URL}` with an actual url that points to an image smaller than 8 MB.

Replace `{TEXT}` with text input respecting the following character limit.

Numbers in parenthesis `()` can indicate a limit. For text it's the max character limit and for images it signals the minimum size that can be requested from Discord directly to avoid unnecessary resizing. The API still has to resize them for most of the endpoints but the smaller the input the faster that process is so those are the minimum sizes you should use.

- `/religion?avatar={IMAGE_URL(512)}`
- `/beautiful?avatar={IMAGE_URL(256)}`
- `/fear?avatar={IMAGE_URL(256)}`
- `/sacred?avatar={IMAGE_URL(512)}`
- `/painting?avatar={IMAGE_URL(512)}`
- `/color?color={NAME_OR_HEX}` (`#` is automatically stripped off, and name is case insensitive)
- `/delete?avatar={IMAGE_URL(256)}`
- `/garbage?avatar={IMAGE_URL(512)}`
- `/tom?avatar={IMAGE_URL(256)}`
- `/bed?avatar={IMAGE_URL(128)}&target={IMAGE_URL(128)}`
- `/crush?avatar={IMAGE_URL(512)}&target={IMAGE_URL(512)}`
- `/patrick?avatar={IMAGE_URL(512)}`
- `/respect?avatar={IMAGE_URL(128)}`
- `/dipshit?text={TEXT(33)}`
- `/picture?avatar={IMAGE_URL(1024)}`
- `/tweet?text={TEXT(165)}`
- `/truth?avatar={IMAGE_URL(256)}`
- `/bobross?avatar={IMAGE_URL(512)}`
- `/mask?avatar={IMAGE_URL(512)}` (More feedback required to improve this.)
- `/father?avatar={IMAGE_URL(256)}&text={TEXT(42)}`
- `/achievement?avatar={IMAGE_URL(64)}&text={TEXT(21)}`
- `/dominantColor?avatar={IMAGE_URL}` Get an image dominant color. Response format: `{ "hex": "#FFFFFF", "rgb": { "r": 255, "g": 255, "b": 255 } }`
