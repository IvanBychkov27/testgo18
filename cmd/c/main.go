// http://прохоренок.рф/pdf/go/ch15-go-funktsii-dlya-raboty-s-katalogami.html
package main

import (
	"fmt"
	"strings"
)

func main() {
	//ip := "127.0.0.1:1254"
	ip := "0000:0000:0000:0000:0000:FFFF:C0A8:1"

	if IsIPv4(ip) {
		fmt.Println("ip v4")
	}
	if IsIPv6(ip) {
		fmt.Println("ip v6")
	}

}

func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}
