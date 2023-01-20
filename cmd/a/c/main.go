// http://прохоренок.рф/pdf/go/ch15-go-funktsii-dlya-raboty-s-katalogami.html
package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type networkQPS struct {
	limit int64
	count int64
}

func (n *networkQPS) inc() {
	atomic.AddInt64(&n.count, 1)
}

func (n *networkQPS) reset() {
	atomic.StoreInt64(&n.count, 0)
}

func (n *networkQPS) setLimit(limit int64) {
	atomic.StoreInt64(&n.limit, limit)
}

func main() {
	var Mx sync.RWMutex

	m := map[int]*networkQPS{}
	count, ok := m[1]
	if !ok {
		count = &networkQPS{100, 250}
		m[1] = count
	}
	fmt.Println("count:", count)

	//Mx.RLock()
	for _, n := range m {
		Mx.Lock()
		n.count = 0
		Mx.Unlock()
	}
	//Mx.RUnlock()

	fmt.Println("count:", m[1])

}

func strQuote(s string) string {
	return strconv.Quote(s)
}

func strAppendQuote(dst []byte, s string) []byte {
	return strconv.AppendQuote(dst, s)
}

func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}
