package discover

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type privateIP struct {
	PrivateIP string `json:"private_ip"`
}

// GetIP uses the go-dns function wich is installed on brain machine
func GetIP(key string, value string) string {

	var output privateIP

	tagValue := make(map[string]string)

	tagValue[key] = value

	data, err := json.Marshal(tagValue)

	if err != nil {
		log.Fatal(err)
	}

	dataString := string(json.RawMessage(data))

	body := strings.NewReader(dataString)

	req, err := http.NewRequest("POST", "http://127.0.0.1:3333/ec2", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &output)
	if err != nil {
		log.Fatal(err)
	}

	return output.PrivateIP

}
