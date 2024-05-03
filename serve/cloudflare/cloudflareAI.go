package cloudflare

import (
	"fmt"
	"github.com/u2takey/go-utils/json"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	URL    = "https://api.cloudflare.com/client/v4/accounts/%s/ai/run/@cf/openai/whisper"
	METHOD = "POST"
)

// Define la estructura de la palabra
type Word struct {
	Word  string  `json:"word"`
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

// Define la estructura del resultado
type Result struct {
	Text      string `json:"text"`
	WordCount int    `json:"word_count"`
	Words     []Word `json:"words"`
	Vtt       string `json:"vtt"`
}

// Define la estructura principal del JSON
type Response struct {
	Result   Result   `json:"result"`
	Success  bool     `json:"success"`
	Errors   []error  `json:"errors"`
	Messages []string `json:"messages"`
	Key      int
}

func CloudflareAI(key int, path string) Response {
	url := fmt.Sprintf(URL, os.Getenv("CLOUDFLARE_ACCOUNT_ID"))
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return Response{
			Result:   Result{},
			Success:  false,
			Errors:   []error{err},
			Messages: []string{"Error reading file"},
			Key:      key,
		}
	}
	payload := strings.NewReader(string(file))
	client := &http.Client{}
	req, err := http.NewRequest(METHOD, url, payload)
	if err != nil {
		fmt.Println(err)
		return Response{
			Result:   Result{},
			Success:  false,
			Errors:   []error{err},
			Messages: []string{"Error creating request"},
			Key:      key,
		}
	}
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CLOUDFLARE_API_KEY")))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Response{
			Result:   Result{},
			Success:  false,
			Errors:   []error{err},
			Messages: []string{"Error making request"},
			Key:      key,
		}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Response{
			Result:   Result{},
			Success:  false,
			Errors:   []error{err},
			Messages: []string{"Error reading response"},
			Key:      key,
		}
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return Response{
			Result:   Result{},
			Success:  false,
			Errors:   []error{err},
			Messages: []string{"Error parsing response"},
			Key:      key,
		}
	}
	fmt.Println("End of CloudflareAI: ", key)
	return response
}
