package twitter

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/modules"
)

func imageToBase64String(path string) string {
	file, _ := os.Open(path)
	fileData, _ := ioutil.ReadAll(file)
	return base64.StdEncoding.EncodeToString(fileData)
}

// UploadMedia Upload image function
func UploadMedia(api *anaconda.TwitterApi, mediaPath string) anaconda.Media {
	media, err := api.UploadMedia(imageToBase64String(mediaPath))

	if err != nil {
		modules.ErrorMessage("Upload failed")
	}

	return media
}
