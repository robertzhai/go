package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func convert( b IPAddr ) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s,".")
}

func(ip IPAddr) String() string {

	//return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1],ip[2],ip[3])
	return convert(ip)
	return string(strconv.Itoa(int(ip[0]))) + "." + string(strconv.Itoa(int(ip[1]))) + "." +
			string(strconv.Itoa(int(ip[2])))+ "." + string(strconv.Itoa(int(ip[3])))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}