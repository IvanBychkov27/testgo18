package main

import (
	"fmt"
	wifi "github.com/mark2b/wpa-connect"
	"time"
)

func main() {
	scanWiFi()
}

// сканирование wi-fi сетей
func scanWiFi() {
	if bssList, err := wifi.ScanManager.Scan(); err == nil {
		for _, bss := range bssList {
			print(bss.SSID, bss.Signal, bss.KeyMgmt)
		}
	} else {
		fmt.Println("error scan wi-fi:", err.Error())
	}

}

// подключекние к wi-fi сети
func connectWiFi(ssid, password string) {
	if conn, err := wifi.ConnectManager.Connect(ssid, password, time.Second*60); err == nil {
		fmt.Println("Connected", conn.NetInterface, conn.SSID, conn.IP4.String(), conn.IP6.String())
	} else {
		fmt.Println(err)
	}
}
