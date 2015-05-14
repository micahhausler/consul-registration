package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Check struct {
	Ttl      string `json:"TTL,omitempty"`
	Http     string `json:"HTTP,omitempty"`
	Interval string `json:"Interval,omitempty"`
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
	url := fmt.Sprintf("%s/v1/agent/service/register", consul)
	fmt.Printf("Posting to %s\n", url)

	data, _ := json.Marshal(registration)
	fmt.Printf("    content: \"%s\"\n", string(data))

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("    Response Status:", resp.Status)
	fmt.Println("    Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("    Response Body:", string(body))
}

func MarkServicePass(serviceId, consul, note string) {
	strUrl := fmt.Sprintf("%s/v1/agent/check/pass/service:%s", consul, serviceId)
	//var Url *url.URL

	Url, _ := url.Parse(strUrl)
	parameters := url.Values{}
	parameters.Add("note", note)
	Url.RawQuery = parameters.Encode()
	http.Get(Url.String())
}

func DeregisterService(serviceId, consul string) {
	//url := fmt.Sprintf("%s/v1/agent/check/deregister/service:%s", consul, serviceId)
	url := fmt.Sprintf("%s/v1/agent/service/deregister/%s", consul, serviceId)
	http.Get(url)
}
