## v1.2.5 (8/2/2023)
- Optimized images using `pngquant`, this reduces the package size, memory usage and overall bandwidth usage.
- Update go-chi to v5.
- Windows ARM binaries are now included.

## v1.2.4 (11/3/2021)
- Added homebrew installation. `brew install ravener/tap/img-api`

## v1.2.3 (15/2/2021)
- Useless release to force a release, fucking kill me.

## v1.2.2 (15/2/2021)
- Useless release to force a release, I suck.

## v1.2.1 (15/2/2021)
- Useless release to force a release.

## v1.2.0 (15/2/2021)
- Print a startup message so it doesn't look like the program is hanging to the end user.
- Start shipping binaries with goreleaser

## v1.1.0 (4/9/2020)
- Added `/kaguya` route

## v1.0.3 (01/08/2020)
- Added a Dockerfile
- Changed constant `VERSION` in routes/stats.go

## v1.0.2 (27/5/2020)
- Added `/dominantColor` endpoint.

## v1.0.1 (20/5/2020)
- Documented recommended avatar sizes for each endpoint.
- Added `goroutines` field in `/stats`
- Improved `/picture` and `/bobross` slightly, increased image height. More height will be cut but this way the width is more balanced so the image is less streched and looks better.

## v1.0.0 (19/5/2020)
- Initial release.
