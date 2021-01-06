package giphyClient

import (
	"net/http"
	"strings"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type GifImage struct {
	Url string `json:"url"`
}

type GifData struct {
	Images map[string]GifImage `json:"images"`
}

type GifResponse struct {
	Data GifData `json:"data"`
}

const apiKeyEnvtVar string = "GIPHY_API_KEY"
const translateApiBase string = "https://api.giphy.com/v1/gifs/translate"

func GetGifUrl(gifInput string) string {
	resp, fetchErr := http.Get(getApiUrl(gifInput))
	if fetchErr != nil {
		fmt.Println("failed to fetch", gifInput)
	}

	gifJson, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		fmt.Println("failed to read response", ioErr)
	}

	var gifResponse GifResponse
	if serializeErr := json.Unmarshal(gifJson, &gifResponse); serializeErr != nil {
		fmt.Println("failed to unmarshal:", serializeErr)
	}

	return gifResponse.Data.Images["fixed_height_small"].Url
}

func getApiKeyParam() string {
	apiKey := os.Getenv(apiKeyEnvtVar)
	return "?api_key=" + apiKey
}

func getTranslateParam(gifInput string) string {
	return "&s=" + gifInput
}

func getApiUrl(gifInput string) string {
	var sb strings.Builder
	sb.WriteString(translateApiBase)
	sb.WriteString(getApiKeyParam())
	sb.WriteString(getTranslateParam(gifInput))

	return sb.String()
}
