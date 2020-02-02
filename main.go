package main

import (
	"flag"
	"fmt"
	"log"

	
	"github.com/mtojek/ohlavpn/vpn"
)

func main() {
	var countryCode string
	var limit int
	var geoIP bool
	var verify bool

	flag.StringVar(&countryCode, "c", "us", "country code")
	flag.IntVar(&limit, "l", 5, "proxy server limit")
	
	flag.BoolVar(&verify, "v", false, "verify proxy connectivity")
	flag.Parse()

	vpnClient := vpn.NewClient()
	err := vpnClient.Initialize()
	if err != nil {
		log.Fatalf("Error occurred while initializing VPN API vpnClient: %v", err)
	}

	tunnels, err := vpnClient.FindTunnels(countryCode, limit)
	if err != nil {
		log.Fatalf("Error occurred while finding VPN tunnels: %v", err)
	}

	if len(tunnels.Servers) == 0 {
		log.Fatal("No proxy servers found")
	}

	for _, tunnel := range tunnels.Servers {
		fmt.Print(tunnel.String())

		

		fmt.Println()
	}
}
