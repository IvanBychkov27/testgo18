package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
)

const (
	defaultWhoisServer = "whois.iana.org"
	defaultWhoisPort   = "43"
	asnPrefix          = "AS"
)

var (
	ErrDomainEmpty = errors.New("whois: domain is empty")
	NotFound       = "not found"
)

func main() {
	listDomains := []string{
		"ecoonline.click",
		"othpro.click",
		"adul.pro",
		"sporweb.click",
		"aduld.click",
		"dappauthorise.com",
		"ovingroup.com",
		"yvmloans.com",
		"dappsupportchain.com",
		"data-informative.com",
		"zstrive.com",
		"skrinelinecompany.com",
		"ya.ru",
		"poker-iv.herokuapp.com",
		"fffggg.ru",
	}

	for idx, domain := range listDomains {
		fmt.Printf("%3d: %26s : ", idx, domain)
		checkWhois(domain)
	}

}

func checkWhois(d string) {
	//startDomain := "poker-iv.herokuapp.com"
	//startDomain := "ya.ru"

	domains := splitDomain(d)
	servers := listServers(d)

	var result string

	for _, domain := range domains {
		for _, server := range servers {
			//fmt.Printf("query: domain %s, server %s \n", domain, server)
			data, err := whois(domain, server)
			if err != nil {
				//fmt.Printf("error: domain %s, server %s, error %s \n", domain, server, err.Error())
				continue
			}
			//fmt.Println(data)
			result = parseWhois(data)
			if result == NotFound {
				continue
			}
			break
		}
	}
	fmt.Println(result)
}

func whois(domain, server string) (result string, err error) {
	domain = strings.Trim(strings.TrimSpace(domain), ".")
	if domain == "" {
		return "", ErrDomainEmpty
	}
	if strings.Contains(domain, "/") {
		domain = strings.Split(domain, "/")[0]
	}

	isASN := IsASN(domain)
	if isASN {
		if !strings.HasPrefix(strings.ToUpper(domain), asnPrefix) {
			domain = asnPrefix + domain
		}
	}

	if !strings.Contains(domain, ".") && !strings.Contains(domain, ":") && !isASN {
		return rawQuery(domain, defaultWhoisServer, defaultWhoisPort)
	}

	port := defaultWhoisPort

	result, err = rawQuery(domain, server, port)
	if err != nil {
		return
	}

	refServer, refPort := getRefServer(result)
	if refServer == "" || refServer == server {
		return
	}

	var data string
	data, err = rawQuery(domain, refServer, refPort)
	if err == nil {
		result += data
	}

	return
}

func rawQuery(domain, server, port string) (string, error) {
	conn, errDial := net.Dial("tcp", net.JoinHostPort(server, port))
	if errDial != nil {
		return "", fmt.Errorf("whois: connect to whois server failed: %w", errDial)
	}
	defer conn.Close()

	_, errWrite := conn.Write([]byte(domain + "\r\n"))
	if errWrite != nil {
		return "", fmt.Errorf("whois: send to whois server failed: %w", errWrite)
	}

	buffer, errReadAll := ioutil.ReadAll(conn)
	if errReadAll != nil {
		return "", fmt.Errorf("whois: read from whois server failed: %w", errReadAll)
	}

	return string(buffer), nil
}

func getRefServer(data string) (string, string) {
	tokens := []string{
		"Registrar WHOIS Server: ",
		"whois: ",
		"ReferralServer: ",
	}

	for _, token := range tokens {
		start := strings.Index(data, token)
		if start != -1 {
			start += len(token)
			end := strings.Index(data[start:], "\n")
			server := strings.TrimSpace(data[start : start+end])
			server = strings.TrimPrefix(server, "whois:")
			server = strings.TrimPrefix(server, "rwhois:")
			server = strings.Trim(server, "/")
			port := defaultWhoisPort
			if strings.Contains(server, ":") {
				v := strings.Split(server, ":")
				server, port = v[0], v[1]
			}
			return server, port
		}
	}

	return "", ""
}

func IsASN(s string) bool {
	s = strings.ToUpper(s)

	s = strings.TrimPrefix(s, asnPrefix)
	_, err := strconv.Atoi(s)

	return err == nil
}

func splitDomain(domain string) []string {
	ds := strings.Split(domain, ".")
	if len(ds) < 2 {
		return nil
	}

	domain2 := ds[len(ds)-2] + "." + ds[len(ds)-1]
	if domain == domain2 {
		return []string{domain}
	}
	return []string{domain, domain2}
}

func listServers(domain string) []string {
	ds := strings.Split(domain, ".")
	serverNet := "whois-servers.net"
	if len(ds) > 0 {
		serverNet = ds[len(ds)-1] + "." + serverNet
	}
	servers := []string{
		"whois.nic.ru",
		serverNet,
		"whois.iana.org",
		"whois.tcinet.ru",
		"whois.godaddy.com",
		"whois.arin.net",
	}
	return servers
}

func parseWhois(data string) string {
	const (
		paidTill                            = "paid-till"
		registrarRegistrationExpirationDate = "Registrar Registration Expiration Date"
		registryExpiryDate                  = "Registry Expiry Date"
	)

	ds := strings.Split(data, "\n")

	for _, d := range ds {
		if strings.Contains(d, paidTill) ||
			strings.Contains(d, registrarRegistrationExpirationDate) ||
			strings.Contains(d, registryExpiryDate) {

			d = strings.TrimSpace(d)
			ds := strings.Split(d, " ")
			if len(ds) == 0 {
				return ""
			}
			d = ds[len(ds)-1]

			return d
		}
	}

	return NotFound
}

/* Ссылка на нужный сервер
"Registrar WHOIS Server: ",
"whois: ",
"ReferralServer: ",
*/
