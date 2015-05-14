package post

import (
	//"bytes"
	"encoding/json"
	"fmt"
	//"net/http"
)

type Check struct {
	Ttl      int    `json:"TTL,omitempty"`
	Http     string `json:"HTTP,omitempty"`
	Interval int    `json:"Interval,omitempty"`
	Script   string `json:"Script,omitempty"`
}
type Registration struct {
	Id      string   `json:"ID,omitempty"`
	Name    string   `json:"Name"`
	Tags    []string `json:"Tags,omitempty"`
	Address string   `json:"Address"`
	Port    int      `json:"Port,omitempty"`
	Check   *Check   `json:"Check"`
}

func RegisterService(registration *Registration, consul string) {
	url := fmt.Sprintf("http://%s/v1/agent/service/register\n", consul)
	fmt.Printf("Posting to %s\n", url)

	data, _ := json.Marshal(registration)
	fmt.Printf("    content: \"%s\"\n", string(data))

	/*
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("    Response Status:", resp.Status)
	*/

}
