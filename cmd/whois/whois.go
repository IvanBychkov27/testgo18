/* https://habr.com/ru/post/165869/

В нескольких словах, whois (от английского «who is» — «кто такой») – сетевой протокол, базирующийся на протоколе TCP.
Его основное предназначение – получение в текстовом виде регистрационных данных о владельцах IP адресов и доменных имен (главным образом, их контактной информации).
Запись о домене обычно содержит имя и контактную информацию «регистранта» (владельца домена) и «регистратора» (организации, которая домен зарегистрировала),
имена DNS серверов, дату регистрации и дату истечения срока ее действия.
Часто whois используется для проверки, свободно ли доменное имя или уже зарегистрировано.

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


Работу с whois я бы свел к решению трех принципиальных задач:

1. Определение правильного whois сервера;
2. Отправка правильного запроса на сервер;
3. Анализ полученного результата.

http://whois.domaintools.com — лучший whois веб сервис, для образца получаемой информации (как пример что должно быть)

1. Определение правильного whois сервера:
         whois.iana.org
         whois.nic.ru
         <домен>.whois-servers.net
        whois.godaddy.com
        whois.arin.net

	Для каждого домена, который мы хотим найти, у нас будет целый список потенциальных whois серверов с большей или меньшей вероятностью работоспособности.
    Например, для russia.edu.ru у нас будет:
    whois.tcinet.ru (результат IANA для «ru»)
    whois.edu.ru << нужный нам whois сервер
    whois.nic.edu.ru
    whois.nic.ru
    whois.ru
    whois.cctld.ru («вымышленный» сервер на основе адреса сайта регистратора)
    ru.whois-servers.net (по факту, работающий алиас для whois.tcinet.ru)

	Если информации недостаточно можно поискать в тексте Whois Server: с указанным следующим сервисом whois.markmonitor.com (whois может отсутствовать) или
	если "Whois Server:" не найдено ищи "whois." везде и проверяй по нему...

	С IDN доменами (типа «правительство.рф»), к счастью, никаких проблем нет. Переводим «правительство.рф» в «xn--80aealotwbjpid2k.xn--p1ai» и ищем на whois.iana.org домен верхнего уровня «xn--p1ai»:

2. Отправка правильного запроса на сервер:
	В подавляющем большинстве случаев нужно просто открыть соединение на порт 43, отправить строку вида "<домен или IP адрес>\r\n" и прочитать данные, которые вернет сервер.
	Кроме того, следующие whois серверы требуют особый синтаксис:
    whois.arin.org: «n + <IP адрес>\r\n»
    whois.denic.de: "-C UTF-8 -T dn,ace <домен>\r\n"
    whois.dk-hostmaster.dk: "--show-handles <домен>\r\n"
    whois.nic.name: «domain = <домен>\r\n»
    whois серверы, принадлежащие VeriSign: «domain <домен>\r\n»

3. Анализ полученного результата:
	Проверка на валидность результата
	Проверка на повтор полученных данных

*/

package main

import (
	"fmt"
	"github.com/likexian/whois"
	"io/ioutil"
	"net"
	"strings"
)

// poker-iv.herokuapp.com

// whois poker-iv.herokuapp.com | grep 'paid-till\|Registrar Registration Expiration Date\|Registry Expiry Date' -m1

func main() {
	//domain, server := "poker-iv.herokuapp.com", "com.whois-servers.net"
	//domain, server := "poker-iv.herokuapp.com", "whois.iana.org" // ok
	//domain, server := "herokuapp.com", "whois.nic.ru"
	//domain, server := "ya.ru", "ru.whois-servers.net"
	domain, server := "ya.ru", "whois.nic.ru"
	//domain, server := "ya.ru", "whois.iana.org" // в результате см. whois:  whois.tcinet.ru  и  remarks: Registration information: http://www.cctld.ru/en - Это собственно и есть whois сервер и сайт регистратора!
	//domain, server := "ya.ru", "whois.godaddy.com"
	//domain, server := "ya.ru", "whois.arin.net"
	//domain, server := "ya.ru", "whois.tcinet.ru"

	//domain, server := "finonline.click", "whois.iana.org"
	//domain, server := "ecoonline.click", "whois.iana.org"
	//domain, server := "js.cdnspace.io", "whois.iana.org"

	/* list domains:
	   ecoonline.click
	   othpro.click
	   adul.pro
	   sporweb.click
	   aduld.click
	   dappauthorise.com
	   ovingroup.com
	   yvmloans.com
	   dappsupportchain.com
	   data-informative.com
	   zstrive.com
	   skrinelinecompany.com
	*/

	res := whoIS(domain, server)
	fmt.Println(res)
	fmt.Println("--------------------------------------------------")

	result, err := whois.Whois(domain)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Println("res   :", parseWhois(res))
	fmt.Println("result:", parseWhois(result))

	fmt.Println("Done..")
}

func whoIS(domain, server string) string {
	conn, err := net.Dial("tcp", server+":43")
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))

	//buf := make([]byte, 1024)
	//res := []byte{}
	//for {
	//	numBytes, err := conn.Read(buf)
	//	sbuf := buf[0:numBytes]
	//	res = append(res, sbuf...)
	//	if err != nil {
	//		break
	//	}
	//}

	buffer, errReadAll := ioutil.ReadAll(conn)
	if errReadAll != nil {
		fmt.Println("error read all:", errReadAll.Error())
	}

	//return string(res)
	return string(buffer)
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

			//d = strings.TrimSpace(d)
			//ds := strings.Split(d, " ")
			//if len(ds) == 0 {
			//	return ""
			//}
			//d = ds[len(ds)-1]

			return d
		}
	}

	return "not found"
}
