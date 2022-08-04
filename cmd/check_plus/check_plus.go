package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type UnifyList struct {
	Data []ClientType `json:"data"`
}

type ClientType struct {
	Mac string `json:"mac"`
}

type UserList struct {
	Name   string `json:"name"`
	HWaddr string `json:"hwaddr"`
	Online bool   `json:"online"`
}

func main() {
	text := "Start CheckPlus"
	fmt.Println(text)

	listMac := getWifi()
	fmt.Println("number of devices connected to wifi:", len(listMac.Data))

	users := userList()

	for _, d := range listMac.Data {
		u, ok := users[d.Mac]
		if !ok {
			continue
		}
		if !u.Online {
			u.Online = true
			sendMessage(u.Name + ": +")
		}
	}

	fmt.Println("Done...")
}

func userList() map[string]UserList {
	users := map[string]UserList{
		"04:92:26:7c:f6:39": {Name: "Леха", HWaddr: "04:92:26:7c:f6:39"},
		"6e:17:f5:84:0d:48": {Name: "Женя Ковалев", HWaddr: "6e:17:f5:84:0d:48"},
		"ae:7a:72:b2:da:c9": {Name: "Женя Свентитский", HWaddr: "ae:7a:72:b2:da:c9"},
		"4c:f2:02:a7:29:0b": {Name: "Артем", HWaddr: "4c:f2:02:a7:29:0b"},
		"b4:bf:f6:4c:2d:b7": {Name: "Иван", HWaddr: "b4:bf:f6:4c:2d:b7"},
	}
	return users
}

func getWifi() (result UnifyList) {
	options := cookiejar.Options{}
	jar, err := cookiejar.New(&options)
	if err != nil {
		fmt.Println("error new cookiejar:", err.Error())
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	url := "https://192.168.1.250:8443/api/login"
	admin := []byte(`{"username":"admin", "password":"rtkmythkj["}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(admin))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "/login")

	client := &http.Client{Transport: tr, Jar: jar}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error get wifi:", err.Error())
		return
	}

	if strings.TrimSpace(resp.Status) != "200" {
		fmt.Println("error, status code:", resp.Status)
		return
	}

	urlStat := "https://192.168.1.250:8443/api/s/default/stat/sta/"

	req, err = http.NewRequest("POST", urlStat, bytes.NewBuffer(admin))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "/login")

	client = &http.Client{Transport: tr, Jar: jar}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("error get wifi:", err.Error())
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body (len):", len(body))

	errUnmarshal := json.Unmarshal(body, &result)
	if errUnmarshal != nil {
		fmt.Println("error json unmarshal:", errUnmarshal.Error())
		return
	}

	return result
}

func sendMessage(message string) {
	hookUrl := "https://hooks.slack.com/services/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"

	playload := map[string]interface{}{
		"username": "bot-iv",
		"text":     message,
	}

	bodyString, err := sendPostJson(hookUrl, playload)
	if err != nil {
		fmt.Println("error send message:", err.Error())
		return
	}
	if bodyString == "ok" {
		fmt.Println("send message:", message)
	}
}

func sendPostJson(url string, data interface{}) (bodyString string, err error) {
	info, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(info))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	bodyString = string(body)

	if resp.StatusCode != 200 {
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Headers:", resp.Header)
		fmt.Println("Response Body:", bodyString)

		return bodyString, fmt.Errorf("error %w", err)
	}

	return bodyString, nil
}
