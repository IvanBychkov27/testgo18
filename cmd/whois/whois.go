/* https://habr.com/ru/post/165869/

Фактически, все сводится к следующему:
	открыть TCP соединение на порт 43 к нужному whoIS серверу,
	послать запрос в определенном формате, закончить его "\r\n" и
	получить результат в определенном формате.

server:
1. Про большинство whoIS серверов знает whoIS.iana.org.
	Достаточно отправить запрос вида "whoIS -h whoIS.iana.org <домен верхнего уровня>" и получить, среди прочей информации,
	название нужного whoIS сервера. Если IANA вернула лишь адрес сайта,
	whoIS сервер с определенной вероятностью может быть «угадан» как whoIS.<адрес сайта без www>.

2.  С достаточно большой вероятностью название whoIS сервера может иметь вид
    whoIS.nic.<домен верхнего уровня> или whoIS.<домен верхнего уровня>.

3. Многие домены второго уровня имеют свои собственные whoIS серверы.
   В подавляющем большинстве случаев, их названия имеют вид whoIS.<домен второго уровня> или whoIS.nic.<домен второго уровня>.

4. Для отдельных доменов верхнего уровня работает алиас <домен верхнего уровня>.whoIS-servers.net.

*/

package main

import (
	"fmt"
	"net"

	"github.com/likexian/whois"
)

func main() {
	//domain, server := "poker-iv.herokuapp.com", "com.whois-servers.net"
	domain, server := "ya.ru", "ru.whoIS-servers.net"
	//domain, server := "ya.ru", "whois.nic.ru"

	res := whoIS(domain, server)
	fmt.Println(res)
	fmt.Println("--------------------------------------------------")

	result, err := whois.Whois(domain)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Println("Done..")
}

func whoIS(domain, server string) string {
	conn, err := net.Dial("tcp", server+":43")
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))

	buf := make([]byte, 1024)
	res := []byte{}
	for {
		numBytes, err := conn.Read(buf)
		sbuf := buf[0:numBytes]
		res = append(res, sbuf...)
		if err != nil {
			break
		}
	}

	return string(res)
}
