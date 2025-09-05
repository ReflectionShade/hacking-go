package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

func main() {
	targetIP := "192.168.0.1/24"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(targetIP),
		nmap.WithPorts("80", "443"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("error with creating new scanner: %v", err)
	}

	results, warn, err := scanner.Run()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	if warn != nil {
		log.Fatalf("warning: %v\n", warn)
	}

	for _, host := range results.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("IP: %q\n", host.Addresses[0])
		if len(host.Addresses) > 1 {
			fmt.Printf("MAC:%v", host.Addresses[1])
		}
	}
}
