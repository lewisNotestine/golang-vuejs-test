package giphyClient

import (
	"net/http"
	"strings"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

// type gifDatum struct {

// }

type gifImage struct {
	// FixedHeightSmall gifDatum `json:"fixed_height_small"`
	Url string `json:"url"`
}

type gifData struct {
	Images map[string]gifImage `json:"images"`
}

type GifJson struct {
	Data gifData `json:"data"`
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

	fmt.Println("gifJson is", gifJson)

	var gifData GifJson
	if serializeErr := json.Unmarshal(gifJson, &gifData); serializeErr != nil {
		fmt.Println("failed to unmarshal:", serializeErr)
	}

	fmt.Println("gifData is", gifData)
	return gifData.Data.Images["fixed_height_small"].Url
}

func getApiKeyParam() string {
	apiKey := os.Getenv(apiKeyEnvtVar)
	fmt.Println("api key is", apiKey)
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

	str := sb.String()
	fmt.Println("url is", str)
	return str
}
