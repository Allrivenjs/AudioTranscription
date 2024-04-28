package cloudflare

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CloudflareAI() {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/@cf/openai/whisper", os.Getenv("CLOUDFLARE_ACCOUNT_ID"))
	method := "POST"
	ospath, _ := os.Getwd()
	path := fmt.Sprintf("%s/testdata/audio.wav", ospath)
	fmt.Println(path)
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	payload := strings.NewReader(string(file))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CLOUDFLARE_API_KEY")))
	req.Header.Add("Cookie", "__cfruid=819a88d8c3fd457e936c6a1ebb2979c286ea136c-1714260613")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
