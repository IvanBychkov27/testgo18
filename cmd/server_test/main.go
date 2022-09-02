package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (app *Application) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("call ", req.RemoteAddr)

	v := url.Values{}
	v.Add("cmp", "15de34852038474f27bcc390e13dfea0")
	v.Add("headers[REMOTE_ADDR]", "167.88.62.21")
	v.Add("headers[REMOTE_PORT]", "58295")
	v.Add("headers[HTTP_USER_AGENT]", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0")
	v.Add("adapi", "2.2")
	v.Add("sv", "11785.3")
	v.Add("HTTP_MC_CACHE", "xij9WjmUJQ/wYka9waXy/zR3DhZxmksQerYkh2u9CV+B04hHjWx2Jq6YWVUF38lRH4ouXbzg7NTmRGMx/wv6XQArsoe9sVDU4cP4dI89/GyPuvA0rf/m0ARbaMO9ZxurInn7xy2ThuQ0xxqDfJtu7S41YDN4XxqpwSUSnX2g1/LPUnFwtp+NCNYJzs7n8ZjyZplwphh/eaWnQag8MfA90jDtEf3hfAU0ViLkNydX+RAwLX5Sa/QCCEmx6sgCbVG3B8/ASzqIFDtWmLA7dVWf5RCQF0zpcmUvBGYOjdAplHIzUs7+w4Totj+MOp8wHSUkimuwEgqpR2In")

	res := v.Encode()

	_, err := rw.Write([]byte(res))
	if err != nil {
		log.Println("error write ", err.Error())
	}

}

func main() {
	app := New()
	addr := "127.0.0.1:2801"
	log.Printf("start server test - listen: %s", addr)
	err := http.ListenAndServe(addr, app)
	if err != nil {
		log.Printf("error %v", err)
		os.Exit(1)
	}
}
