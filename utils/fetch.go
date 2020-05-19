package utils

import (
	"net/http"
	"image"
	_ "image/png"
	_ "image/jpeg"
	"errors"
)

func GetImage(url string) (image.Image, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	// Add some limit to avoid security problems.
	// Since this is mainly for Discord Bots, 8 MB is a sane limit.
	// Since normal users without Nitro can only upload maximum of 8 MB.
	// Although most people would use this for user avatars which is even smaller.
	// Not really a big deal but still.
	if res.ContentLength > 1048 * 1048 * 8 {
		res.Body.Close()
		return nil, errors.New("File cannot be larger than 8 MB")
	}

	img, _, err := image.Decode(res.Body)
	return img, err
}
