package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// ssh root@159.223.210.61 -L 14082:127.0.0.1:8082

func main() {
	ln, errDial := net.Dial("tcp", "127.0.0.1:14082") // адрес сервиса magic checker
	if errDial != nil {
		fmt.Println("error dial to socket: " + errDial.Error())
		return
	}
	defer ln.Close()

	data := []byte("109.194.11.57:41726\n") // адрес user'a

	errW := ln.SetWriteDeadline(time.Now().Add(time.Second * 1))
	if errW != nil {
		fmt.Println("error set write deadline: ", errW.Error())
		return
	}

	_, errWrite := ln.Write(data)
	if errWrite != nil {
		fmt.Println("error write to socket: ", errWrite.Error())
		return
	}

	var res []byte

	err2 := ln.SetReadDeadline(time.Now().Add(time.Second * 1))
	if err2 != nil {
		fmt.Println("error set read deadline: ", err2.Error())
		return
	}

	for {
		buf := make([]byte, 1024)
		n, errRead := ln.Read(buf)
		if errRead != nil {
			fmt.Println("error read from socket: ", errRead.Error())
			break
		}
		res = append(res, buf[:n]...)
		if n < len(buf) {
			break
		}
	}

	if len(res) == 0 {
		fmt.Println("result = nil")
		return
	}

	fmt.Printf("[Result]:\n%s\n", res)

	v := url.Values{}
	v.Add("cmp", "15de34852038474f27bcc390e13dfea0")
	v.Add("headers[REMOTE_ADDR]", "109.194.11.57")                                                                            // IP адрес usera
	v.Add("headers[REMOTE_PORT]", "45074")                                                                                    // порт user'a
	v.Add("headers[HTTP_USER_AGENT]", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0") // UA user'a
	v.Add("adapi", "2.2")
	v.Add("sv", "11785.3")
	v.Add("HTTP_MC_CACHE", string(res))

	cloudBodyB := v.Encode()

	fmt.Println("[cloudBodyB]:")
	fmt.Println(cloudBodyB)

	req, errReq := http.NewRequest(http.MethodPost, "http://check.magicchecker.com/v2.2/index.php", bytes.NewReader([]byte(cloudBodyB)))
	if errReq != nil {
		fmt.Println("error create request: ", errReq.Error())
		return
	}
	req.Header.Add("adapi", "2.2")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{Timeout: time.Second * 2}
	resp, errDo := client.Do(req)
	if errDo != nil {
		fmt.Println("error do request: ", errDo.Error())
		return
	}
	defer resp.Body.Close()

	body, errB := io.ReadAll(resp.Body)
	if errB != nil {
		fmt.Println("error read response body: ", errB.Error())
		return
	}

	fmt.Printf("Response from cloud:\n%s\n", body)

	blocked := struct {
		Success   int `json:"success"`
		IsBlocked int `json:"isBlocked"`
	}{}

	errU := json.Unmarshal(body, &blocked)
	if errU != nil {
		fmt.Println("error:", errU.Error())
		return
	}

	if blocked.IsBlocked == 1 {
		fmt.Println("is blocked:", true)
	}

	// {
	//  "success": 1,
	//  "isBlocked": 1,  <---- проверяем это поле
	//  "errorMessage": "",
	//  "urlType": "redirect",
	//  "url": "https:\/\/sexu.com\/?w=1",
	//  "send_params": 10
	//}
}

/*

POST /v2.2/index.php HTTP/1.1
adapi: 2.2
Content-Type: application/x-www-form-urlencoded; charset=utf-8
Host: check.magicchecker.com
Connection: close
User-Agent: Paw/3.3.6 (Macintosh; OS X/12.5.0) GCDHTTPRequest
Content-Length: 631
cmp=15de34852038474f27bcc390e13dfea0&headers%5BREMOTE_ADDR%5D=167.88.62.21&headers%5BREMOTE_PORT%5D=58295&headers%5BHTTP_USER_AGENT%5D=Mozilla%2F5.0+%28Macintosh%3B+Intel+Mac+OS+X+10.15%3B+rv%3A103.0%29+Gecko%2F20100101+Firefox%2F103.0&adapi=2.2&sv=11785.3&HTTP_MC_CACHE=xij9WjmUJQ%2FwYka9waXy%2FzR3DhZxmksQerYkh2u9CV%2BB04hHjWx2Jq6YWVUF38lRH4ouXbzg7NTmRGMx%2Fwv6XQArsoe9sVDU4cP4dI89%2FGyPuvA0rf%2Fm0ARbaMO9ZxurInn7xy2ThuQ0xxqDfJtu7S41YDN4XxqpwSUSnX2g1%2FLPUnFwtp%2BNCNYJzs7n8ZjyZplwphh%2FeaWnQag8MfA90jDtEf3hfAU0ViLkNydX%2BRAwLX5Sa%2FQCCEmx6sgCbVG3B8%2FASzqIFDtWmLA7dVWf5RCQF0zpcmUvBGYOjdAplHIzUs7%2Bw4Totj%2BMOp8wHSUkimuwEgqpR2In



cmp : 15de34852038474f27bcc390e13dfea0
headers[REMOTE_ADDR] : 167.88.62.21
headers[REMOTE_PORT] : 58295
headers[HTTP_USER_AGENT] : Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0
adapi : 2.2
sv : 11785.3
HTTP_MC_CACHE : xij9WjmUJQ/wYka9waXy/zR3DhZxmksQerYkh2u9CV+B04hHjWx2Jq6YWVUF38lRH4ouXbzg7NTmRGMx/wv6XQArsoe9sVDU4cP4dI89/GyPuvA0rf/m0ARbaMO9ZxurInn7xy2ThuQ0xxqDfJtu7S41YDN4XxqpwSUSnX2g1/LPUnFwtp+NCNYJzs7n8ZjyZplwphh/eaWnQag8MfA90jDtEf3hfAU0ViLkNydX+RAwLX5Sa/QCCEmx6sgCbVG3B8/ASzqIFDtWmLA7dVWf5RCQF0zpcmUvBGYOjdAplHIzUs7+w4Totj+MOp8wHSUkimuwEgqpR2In

*/
