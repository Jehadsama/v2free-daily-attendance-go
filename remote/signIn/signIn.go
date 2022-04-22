package sign

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	utils "v2free-daily-attendance-go/lib/utils"
)

func SignIn() string {
	url := os.Getenv("V2FREE_HOST")
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)

	utils.CheckError(err)

	req.Header.Add("Cookie", "")

	res, err := client.Do(req)

	utils.CheckError(err)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	utils.CheckError(err)

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	utils.CheckError(err)
	var msg string

	switch result["ret"] {
	case 1:
		msg = "sign in successfully"
	default:
		msg = "sign in failed"
	}

	return msg
}
